// Code generated by capnpc-go. DO NOT EDIT.

package standup

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Status struct{ capnp.Struct }

// Status_TypeID is the unique identifier for the type Status.
const Status_TypeID = 0x8e9b98e73dbe7428

func NewStatus(s *capnp.Segment) (Status, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 5})
	return Status{st}, err
}

func NewRootStatus(s *capnp.Segment) (Status, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 5})
	return Status{st}, err
}

func ReadRootStatus(msg *capnp.Message) (Status, error) {
	root, err := msg.RootPtr()
	return Status{root.Struct()}, err
}

func (s Status) String() string {
	str, _ := text.Marshal(0x8e9b98e73dbe7428, s.Struct)
	return str
}

func (s Status) UserID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Status) HasUserID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Status) UserIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Status) SetUserID(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Status) Date() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Status) HasDate() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Status) DateBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Status) SetDate(v string) error {
	return s.Struct.SetText(1, v)
}

func (s Status) Yesterday() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.TextList{List: p.List()}, err
}

func (s Status) HasYesterday() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Status) SetYesterday(v capnp.TextList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewYesterday sets the yesterday field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Status) NewYesterday(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s Status) Today() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(3)
	return capnp.TextList{List: p.List()}, err
}

func (s Status) HasToday() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Status) SetToday(v capnp.TextList) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewToday sets the today field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Status) NewToday(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s Status) Blockers() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(4)
	return capnp.TextList{List: p.List()}, err
}

func (s Status) HasBlockers() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Status) SetBlockers(v capnp.TextList) error {
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewBlockers sets the blockers field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Status) NewBlockers(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

// Status_List is a list of Status.
type Status_List struct{ capnp.List }

// NewStatus creates a new list of Status.
func NewStatus_List(s *capnp.Segment, sz int32) (Status_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 5}, sz)
	return Status_List{l}, err
}

func (s Status_List) At(i int) Status { return Status{s.List.Struct(i)} }

func (s Status_List) Set(i int, v Status) error { return s.List.SetStruct(i, v.Struct) }

func (s Status_List) String() string {
	str, _ := text.MarshalList(0x8e9b98e73dbe7428, s.List)
	return str
}

// Status_Promise is a wrapper for a Status promised by a client call.
type Status_Promise struct{ *capnp.Pipeline }

func (p Status_Promise) Struct() (Status, error) {
	s, err := p.Pipeline.Struct()
	return Status{s}, err
}

const schema_b7eae732787dc175 = "x\xda\\\xc9\xb1J\xc3P\x18\x86\xe1\xef;'\xb1K" +
	"I\xfd\xe1\xec]EPh\xc7\x82\xe8\xe0\xa2S\x7f\xbd" +
	"\x82c\x93IiCr\x02-\xd8A\xa8\xe0\xe4\\\x11" +
	"\xc1\xd9IApS\x1c\x1c\x04\xbd\x8e\x827\x11Q\xe8" +
	"\xd2\xf1}\x9f\xf5\x9b=\xd3\x89\xdf\x08\xa8\x8b\xd7\xea\x8d" +
	"\xf0\xba\xb3\x98\xdf^C\x12\xd6\xd5\xfbt\xdc]\xfc\xbc" +
	" \x8e\x1b\x80|<\xc8w\x03\xe8|\xb6\x89\xad\xba\x0c" +
	"~\x98V\xf96\x07>\x1f\xe6\xbd\xe3\xd0\xf2\xa1*\xfb" +
	"\xa4:\x1b\x01\x11\x01\x99\xf6\x00\x1d[\xea\xccPH\xc7" +
	"\xbfy\xb1\x09\xe8\xb9\xa5^\x19\x8a1\x8e\x06\x90\xcb#" +
	"@g\x96zo(\xd6:Z@\xee\xba\x80\xce-\xf5" +
	"\xd9P\xa2\xc81\x02\xe4\xe9\x10\xd0GK\xfd2\xdc\xad" +
	"\xca\xac8\xd8g\x13\x86M\xb0\x95\xfa\x90-\xa3\x9ed" +
	"e\xc8\x8a\xd4\x83\x13&`\xdf\xf2\x9f\x12\xb0\x1dF\xa9" +
	"_\x9d\xf5\xc9\xd9hp\x9a\x15%\x80\x15\xfa\x0d\x00\x00" +
	"\xff\xff\xdc\xb3:\xd0"

func init() {
	schemas.Register(schema_b7eae732787dc175,
		0x8e9b98e73dbe7428)
}
