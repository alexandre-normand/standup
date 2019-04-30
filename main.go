package standup

import (
	"encoding/json"
	"fmt"
	"github.com/alexandre-normand/slackscot/store"
	"github.com/alexandre-normand/slackscot/store/datastoredb"
	"github.com/lithammer/shortuuid"
	"github.com/nlopes/slack"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	capnp "zombiezen.com/go/capnproto2"
)

const (
	slackTokenEnv = "SLACK_TOKEN"
	projectIDEnv  = "PROJECT_ID"
	groupID       = "fakeUserGroup"
)

var projectID string
var slackToken string

func init() {
	slackToken = os.Getenv(slackTokenEnv)
	projectID = os.Getenv(projectIDEnv)
}

func StartStandup(w http.ResponseWriter, r *http.Request) {
	sc := slack.New(slackToken, slack.OptionDebug(true))
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	// Example body: Body is: token=dNiDsgUrCIHce9UTZsYNHXre&team_id=TFSGU56GK&team_domain=fitbit-sbx&channel_id=CHVDDNJCA&channel_name=standupbot&user_id=UFWF3UYK0&user_name=anormand&command=%2Fstandup&text=&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FTFSGU56GK%2F620708849319%2Fuq3x7g41MBcTBwE8iNHWdmKs&trigger_id=618331356036.536572176563.93d3d026323796e62ffc3393e10fd687"
	params, err := url.ParseQuery(string(body))
	if err != nil {
		log.Printf("Error decoding params: %v", err)
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("Body is: %s", string(body))
	log.Printf("Trigger id is: %s", params["trigger_id"])

	triggerID := params["trigger_id"][0]
	callbackID := shortuuid.New()

	err = sc.OpenDialog(triggerID, slack.Dialog{TriggerID: triggerID, CallbackID: callbackID, Title: "Time for Stand Up", Elements: []slack.DialogElement{slack.DialogInput{Type: "textarea", Label: "What did you do yesterday?", Name: "yesterday"}, slack.DialogInput{Type: "textarea", Label: "What are you doing today?", Name: "today"}, slack.DialogInput{Type: "textarea", Label: "Blockers?", Name: "blockers"}}})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RecordStandup(w http.ResponseWriter, r *http.Request) {
	log.Printf("Slack Token is: %s", slackToken)
	sc := slack.New(slackToken, slack.OptionDebug(true))
	store, err := datastoredb.New("standup", projectID)
	if err != nil {
		log.Printf("Error creating persistence: %v", err)
		http.Error(w, err.Error(), 500)
		return
	}

	payload := r.FormValue("payload")
	if payload == "" {
		http.Error(w, "Empty payload", 400)
		return
	}

	var callback slack.InteractionCallback
	err = json.Unmarshal([]byte(payload), &callback)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = persistStatus(store, callback.User.ID, strings.Split(callback.Submission["yesterday"], "\n"), strings.Split(callback.Submission["today"], "\n"), strings.Split(callback.Submission["blockers"], "\n"))
	if err != nil {
		log.Printf("Error creating persistence: %v", err)
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = sc.PostEphemeral(callback.Channel.ID, callback.User.ID, slack.MsgOptionText(":bow: Thanks for providing your status!", false))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, _, err = sc.PostMessage(callback.Channel.ID, slack.MsgOptionText(fmt.Sprintf("Here's today standup status from <@%s>:\n", callback.User.ID), false))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Report(w http.ResponseWriter, r *http.Request) {
	// sc := slack.New(SLACK_TOKEN, slack.OptionDebug(true))

	// storer, err := datastoredb.New("standup", projectID)
	// if err != nil {
	// 	log.Printf("Error creating persistence: %v", err)
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	// entries, err := storer.ScanSilo(groupID)
	// if err != nil {
	// 	log.Printf("Error scanning: %v", err)
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

}

func persistStatus(storer store.GlobalSiloStringStorer, userID string, yesterday []string, today []string, blockers []string) (err error) {
	_, serialized, err := makeStatus(userID, yesterday, today, blockers)
	if err != nil {
		return err
	}

	err = storer.PutSiloString(groupID, userID, serialized)

	return err
}

func makeStatus(userID string, yesterday []string, today []string, blockers []string) (status Status, serialized string, err error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return status, "", err
	}

	status, err = NewRootStatus(seg)
	if err != nil {
		return status, "", err
	}

	status.SetUserID(userID)
	if len(yesterday) > 0 {
		yesterdayEntries, err := status.NewYesterday(int32(len(yesterday)))
		if err != nil {
			return status, "", err
		}

		for i, v := range yesterday {
			yesterdayEntries.Set(i, v)
		}
	}

	var strBuilder strings.Builder
	err = capnp.NewEncoder(&strBuilder).Encode(msg)
	if err != nil {
		return status, "", err
	}

	return status, strBuilder.String(), nil
}
