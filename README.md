# Serialization
Standup uses [Cap'n Proto](https://capnproto.org) for serialization. To change the schema 
and recompile, one needs Cap'n Proto:

`brew install capnp`

And the [Go compiler](https://github.com/capnproto/go-capnproto2/wiki/Getting-Started):
`go get -u zombiezen.com/go/capnproto2/...`

And recompile:
`capnp compile -I$GOPATH/src/zombiezen.com/go/capnproto2/std standup.capnp -ogo`

# Secrets
Uses `env` variables from the application perspective with dev instructions based on [berglas](https://github.com/GoogleCloudPlatform/berglas/tree/master/examples/cloudfunctions/go) for a simple deployment. 

# gcloud functions deploy
* `gcloud functions deploy startStandup --entry-point SecordStandup --runtime go111 --trigger-http --project $PROJECT_ID --service-account ${SA_EMAIL} --set-env-vars "PROJECT_ID=${PROJECT_ID},SLACK_TOKEN=berglas://${BUCKET_ID}/slacktoken?destination=tempfile"`
* `gcloud functions deploy recordStandup --entry-point RecordStandup --runtime go111 --trigger-http --project $PROJECT_ID --service-account ${SA_EMAIL} --set-env-vars "PROJECT_ID=${PROJECT_ID},SLACK_TOKEN=berglas://${BUCKET_ID}/slacktoken?destination=tempfile"`
