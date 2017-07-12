// Code generated by capnpc-go. DO NOT EDIT.

package wireformat

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type MessageType uint16

// MessageType_TypeID is the unique identifier for the type MessageType.
const MessageType_TypeID = 0x9808fa17b68373f4

// Values of MessageType.
const (
	MessageType_errorResponse        MessageType = 0
	MessageType_okResponse           MessageType = 1
	MessageType_dhtPutRequest        MessageType = 2
	MessageType_dhtDelRequest        MessageType = 3
	MessageType_dhtModRequest        MessageType = 4
	MessageType_dhtGetRequest        MessageType = 5
	MessageType_dhtLinkRequest       MessageType = 6
	MessageType_dhtGetLinkRequest    MessageType = 7
	MessageType_dhtDeleteLinkRequest MessageType = 8
	MessageType_gossipRequest        MessageType = 9
	MessageType_validatePutRequest   MessageType = 10
	MessageType_validateLinkRequest  MessageType = 11
	MessageType_validateDelRequst    MessageType = 12
	MessageType_validateModRequest   MessageType = 13
)

// String returns the enum's constant name.
func (c MessageType) String() string {
	switch c {
	case MessageType_errorResponse:
		return "errorResponse"
	case MessageType_okResponse:
		return "okResponse"
	case MessageType_dhtPutRequest:
		return "dhtPutRequest"
	case MessageType_dhtDelRequest:
		return "dhtDelRequest"
	case MessageType_dhtModRequest:
		return "dhtModRequest"
	case MessageType_dhtGetRequest:
		return "dhtGetRequest"
	case MessageType_dhtLinkRequest:
		return "dhtLinkRequest"
	case MessageType_dhtGetLinkRequest:
		return "dhtGetLinkRequest"
	case MessageType_dhtDeleteLinkRequest:
		return "dhtDeleteLinkRequest"
	case MessageType_gossipRequest:
		return "gossipRequest"
	case MessageType_validatePutRequest:
		return "validatePutRequest"
	case MessageType_validateLinkRequest:
		return "validateLinkRequest"
	case MessageType_validateDelRequst:
		return "validateDelRequst"
	case MessageType_validateModRequest:
		return "validateModRequest"

	default:
		return ""
	}
}

// MessageTypeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func MessageTypeFromString(c string) MessageType {
	switch c {
	case "errorResponse":
		return MessageType_errorResponse
	case "okResponse":
		return MessageType_okResponse
	case "dhtPutRequest":
		return MessageType_dhtPutRequest
	case "dhtDelRequest":
		return MessageType_dhtDelRequest
	case "dhtModRequest":
		return MessageType_dhtModRequest
	case "dhtGetRequest":
		return MessageType_dhtGetRequest
	case "dhtLinkRequest":
		return MessageType_dhtLinkRequest
	case "dhtGetLinkRequest":
		return MessageType_dhtGetLinkRequest
	case "dhtDeleteLinkRequest":
		return MessageType_dhtDeleteLinkRequest
	case "gossipRequest":
		return MessageType_gossipRequest
	case "validatePutRequest":
		return MessageType_validatePutRequest
	case "validateLinkRequest":
		return MessageType_validateLinkRequest
	case "validateDelRequst":
		return MessageType_validateDelRequst
	case "validateModRequest":
		return MessageType_validateModRequest

	default:
		return 0
	}
}

type MessageType_List struct{ capnp.List }

func NewMessageType_List(s *capnp.Segment, sz int32) (MessageType_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return MessageType_List{l.List}, err
}

func (l MessageType_List) At(i int) MessageType {
	ul := capnp.UInt16List{List: l.List}
	return MessageType(ul.At(i))
}

func (l MessageType_List) Set(i int, v MessageType) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type Time struct{ capnp.Struct }

// Time_TypeID is the unique identifier for the type Time.
const Time_TypeID = 0xdd5a4545747ef1a4

func NewTime(s *capnp.Segment) (Time, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Time{st}, err
}

func NewRootTime(s *capnp.Segment) (Time, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Time{st}, err
}

func ReadRootTime(msg *capnp.Message) (Time, error) {
	root, err := msg.RootPtr()
	return Time{root.Struct()}, err
}

func (s Time) String() string {
	str, _ := text.Marshal(0xdd5a4545747ef1a4, s.Struct)
	return str
}

func (s Time) UnixSec() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Time) SetUnixSec(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s Time) UnixNsec() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s Time) SetUnixNsec(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

// Time_List is a list of Time.
type Time_List struct{ capnp.List }

// NewTime creates a new list of Time.
func NewTime_List(s *capnp.Segment, sz int32) (Time_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return Time_List{l}, err
}

func (s Time_List) At(i int) Time { return Time{s.List.Struct(i)} }

func (s Time_List) Set(i int, v Time) error { return s.List.SetStruct(i, v.Struct) }

// Time_Promise is a wrapper for a Time promised by a client call.
type Time_Promise struct{ *capnp.Pipeline }

func (p Time_Promise) Struct() (Time, error) {
	s, err := p.Pipeline.Struct()
	return Time{s}, err
}

type Message struct{ capnp.Struct }
type Message_Which uint16

const (
	Message_Which_putReq     Message_Which = 0
	Message_Which_getReq     Message_Which = 1
	Message_Which_delReq     Message_Which = 2
	Message_Which_modReq     Message_Which = 3
	Message_Which_linkReq    Message_Which = 4
	Message_Which_delLinkReq Message_Which = 5
)

func (w Message_Which) String() string {
	const s = "putReqgetReqdelReqmodReqlinkReqdelLinkReq"
	switch w {
	case Message_Which_putReq:
		return s[0:6]
	case Message_Which_getReq:
		return s[6:12]
	case Message_Which_delReq:
		return s[12:18]
	case Message_Which_modReq:
		return s[18:24]
	case Message_Which_linkReq:
		return s[24:31]
	case Message_Which_delLinkReq:
		return s[31:41]

	}
	return "Message_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Message_TypeID is the unique identifier for the type Message.
const Message_TypeID = 0x86899ecd650d61eb

func NewMessage(s *capnp.Segment) (Message, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Message{st}, err
}

func NewRootMessage(s *capnp.Segment) (Message, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Message{st}, err
}

func ReadRootMessage(msg *capnp.Message) (Message, error) {
	root, err := msg.RootPtr()
	return Message{root.Struct()}, err
}

func (s Message) String() string {
	str, _ := text.Marshal(0x86899ecd650d61eb, s.Struct)
	return str
}

func (s Message) Which() Message_Which {
	return Message_Which(s.Struct.Uint16(0))
}
func (s Message) Time() (Time, error) {
	p, err := s.Struct.Ptr(1)
	return Time{Struct: p.Struct()}, err
}

func (s Message) HasTime() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Message) SetTime(v Time) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewTime sets the time field to a newly
// allocated Time struct, preferring placement in s's segment.
func (s Message) NewTime() (Time, error) {
	ss, err := NewTime(s.Struct.Segment())
	if err != nil {
		return Time{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) From() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Message) HasFrom() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Message) FromBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Message) SetFrom(v string) error {
	return s.Struct.SetText(2, v)
}

func (s Message) PutReq() (PutRequest, error) {
	p, err := s.Struct.Ptr(0)
	return PutRequest{Struct: p.Struct()}, err
}

func (s Message) HasPutReq() bool {
	if s.Struct.Uint16(0) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetPutReq(v PutRequest) error {
	s.Struct.SetUint16(0, 0)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPutReq sets the putReq field to a newly
// allocated PutRequest struct, preferring placement in s's segment.
func (s Message) NewPutReq() (PutRequest, error) {
	s.Struct.SetUint16(0, 0)
	ss, err := NewPutRequest(s.Struct.Segment())
	if err != nil {
		return PutRequest{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) GetReq() (GetRequest, error) {
	p, err := s.Struct.Ptr(0)
	return GetRequest{Struct: p.Struct()}, err
}

func (s Message) HasGetReq() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetGetReq(v GetRequest) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewGetReq sets the getReq field to a newly
// allocated GetRequest struct, preferring placement in s's segment.
func (s Message) NewGetReq() (GetRequest, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewGetRequest(s.Struct.Segment())
	if err != nil {
		return GetRequest{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) DelReq() (DelRequest, error) {
	p, err := s.Struct.Ptr(0)
	return DelRequest{Struct: p.Struct()}, err
}

func (s Message) HasDelReq() bool {
	if s.Struct.Uint16(0) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetDelReq(v DelRequest) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewDelReq sets the delReq field to a newly
// allocated DelRequest struct, preferring placement in s's segment.
func (s Message) NewDelReq() (DelRequest, error) {
	s.Struct.SetUint16(0, 2)
	ss, err := NewDelRequest(s.Struct.Segment())
	if err != nil {
		return DelRequest{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) ModReq() (ModRequest, error) {
	p, err := s.Struct.Ptr(0)
	return ModRequest{Struct: p.Struct()}, err
}

func (s Message) HasModReq() bool {
	if s.Struct.Uint16(0) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetModReq(v ModRequest) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewModReq sets the modReq field to a newly
// allocated ModRequest struct, preferring placement in s's segment.
func (s Message) NewModReq() (ModRequest, error) {
	s.Struct.SetUint16(0, 3)
	ss, err := NewModRequest(s.Struct.Segment())
	if err != nil {
		return ModRequest{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) LinkReq() (LinkRequest, error) {
	p, err := s.Struct.Ptr(0)
	return LinkRequest{Struct: p.Struct()}, err
}

func (s Message) HasLinkReq() bool {
	if s.Struct.Uint16(0) != 4 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetLinkReq(v LinkRequest) error {
	s.Struct.SetUint16(0, 4)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewLinkReq sets the linkReq field to a newly
// allocated LinkRequest struct, preferring placement in s's segment.
func (s Message) NewLinkReq() (LinkRequest, error) {
	s.Struct.SetUint16(0, 4)
	ss, err := NewLinkRequest(s.Struct.Segment())
	if err != nil {
		return LinkRequest{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) DelLinkReq() (DelLinkRequest, error) {
	p, err := s.Struct.Ptr(0)
	return DelLinkRequest{Struct: p.Struct()}, err
}

func (s Message) HasDelLinkReq() bool {
	if s.Struct.Uint16(0) != 5 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetDelLinkReq(v DelLinkRequest) error {
	s.Struct.SetUint16(0, 5)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewDelLinkReq sets the delLinkReq field to a newly
// allocated DelLinkRequest struct, preferring placement in s's segment.
func (s Message) NewDelLinkReq() (DelLinkRequest, error) {
	s.Struct.SetUint16(0, 5)
	ss, err := NewDelLinkRequest(s.Struct.Segment())
	if err != nil {
		return DelLinkRequest{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Message_List is a list of Message.
type Message_List struct{ capnp.List }

// NewMessage creates a new list of Message.
func NewMessage_List(s *capnp.Segment, sz int32) (Message_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return Message_List{l}, err
}

func (s Message_List) At(i int) Message { return Message{s.List.Struct(i)} }

func (s Message_List) Set(i int, v Message) error { return s.List.SetStruct(i, v.Struct) }

// Message_Promise is a wrapper for a Message promised by a client call.
type Message_Promise struct{ *capnp.Pipeline }

func (p Message_Promise) Struct() (Message, error) {
	s, err := p.Pipeline.Struct()
	return Message{s}, err
}

func (p Message_Promise) Time() Time_Promise {
	return Time_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Message_Promise) PutReq() PutRequest_Promise {
	return PutRequest_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) GetReq() GetRequest_Promise {
	return GetRequest_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) DelReq() DelRequest_Promise {
	return DelRequest_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) ModReq() ModRequest_Promise {
	return ModRequest_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) LinkReq() LinkRequest_Promise {
	return LinkRequest_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) DelLinkReq() DelLinkRequest_Promise {
	return DelLinkRequest_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

const schema_8b6f325f59887a6b = "x\xda\x84S]H\x1cW\x14>\xe7\xde\xd9\x99\xd5\xee" +
	"\xb0{\xd9)\xc5'\xfbPZ\x94\xfehE*\xbel" +
	"[\xba\x08E\x8bwk\x7f\xf4\xa5]\xdc\xab\x1d\xdc\xdd" +
	"\xd9\x9d\x99\xd5\xd6\x87\x8a\x88\xa5RZh\x8bP\x1bL" +
	"HH\x08\xe4%o>\x85<\x85\xf8\x92\x84\x04\x12b" +
	"@C\x02\x86\x184\x10\x13\x03\xc1\xc4\xdcpgv\xe3" +
	"\x9a\x97<\x9d=\xdf\xf7\xdd\xb3g\xbe\xef\xde\xb68\xf9" +
	"\x94\xb4G\x8eD\x01x.\xa2\xcb\xcd\xac).\x1e\x9e" +
	"\xfb\x0dx\x13\xa2\x1c\x9b\xfc}\xf0\x87\x8f\x9d? M" +
	"\x0d\x1d\xa0\xa3Sk\xc4dZ3\x00:>\xd3\x96\x11" +
	"P\xeex3Ko\xedF\xff\x03\xd6D\xf6\xe5\x80\x1d" +
	"\x0dz+&\x9bt%}S\xffNI\x8fo\xff\xea" +
	"\xa7\xd3Ckjt\x9dVMK\xb6\x18\x0f\x93\x9d\x86" +
	"\xfa\xd5nL\xc0#9a\xbbb\xc4q\x0b4\xeb\x7f" +
	"T\x10\x9e\x97\x1d\x15\x1f\x0egK\xc5Rw_\xd8\x01" +
	"\xf4#\xf2\xb7\xa9f\x12)5\x04`W\xba\x01\xf8\x05" +
	"\x8a|\x85\xa0I\x9fK\x0b\x15zM\xa1\x97)\xf2U" +
	"\x82\xa6\xb6'-$\x00\xec\x86B\xafR\xe4\xb7\x08\x9a" +
	"\x91g\xd2B\x0a\xc0\xd6\x14\xbaB\x91\xaf\x134\xf5\xa7" +
	"\xd2B\x0d\x80\xdd\xfe\x1c\x80\xafR\xe4\x1b\x04McW" +
	"Z\x18\x01`w\x86\x00\xf8:E\xfe\x80`\x02-\xd4" +
	"\x01\xd8\xfdV\x00\xbeA\x91\xef\x10dH,4\x00\xd8" +
	"\xb6\x02\xb7(\xf2'\x04S\xa5\x8a\x9f\x11eL\xc8\xd3" +
	"\xdf\xc7~\xec\x19\\X\x04@L\x00\xa6FE\x958" +
	"\xe4\xbc\x17w\xcf,\xfd[#r\"\x1f\x12\xdf\x8e\x7f" +
	"\xb3\xb4\xbcw\xf6X\x8d(8\xb9\x90(Z\xab\xef\x9e" +
	"\x9b\xf9d\xb3JL\xe5\xed\xe2X\xc8\xdc;\xbfu\xe2" +
	"\xd2\xe3\xbf\xe6\xab\x8c\xcc\x89|\xaf\"\x81\x06t\xe6d" +
	"\xe1n\xea\xe6\xf5\xda\x0eq\xdf.\x08L\xec\xc7T\x85" +
	"G\\\xa7\x801 \x18\x03|}.\xf1\x81_JB" +
	"E\xd3\x15\x18}\xd4Uc\xd8\xe2\x10\x00\x12\xb6\xa0:" +
	"\xca\xe6U\xd1\xd8\xdf\xaaD\xd8\x9f\xaa\xe8ln\x12\x00" +
	"\x0d6\xfb\x0f\x00F\xd9\xec)\x00l`\xb3\x8akd" +
	"\xd3\xff\x03\xe0\x1blZ}|\x8cM+\x89\x19\x80R" +
	"\xb8\xae\xe3f\x84\x07\xcd%\xa7\xe8\x09\xe9\x8ce\x84W" +
	"r\x8a@=!s?\xf9\xfd\xcaph.W\x84\xe7" +
	"\xab\xfe\x0beg]\xdf\xa7\\\xac\xeb{\xc4A}\xe8" +
	"WJ\xd4\x0bzm\x0c\x1c\xae\x08\x0fkC\x85\x8f\xa2" +
	"74>\x1eHG\x1d\xcf\xb3K\xfb\xb3\xc6\xb3y;" +
	"\x97\xf5\x05\x06\x1b\x95+\x82\xd6\x83\xd5\xa3\x15\xe3\x804" +
	"X\xb6\\Q\xff\xf2\x12\x0b\x16\xae\x1e\xaf\x85A^\x0d" +
	"\xa3\xb9{\xc0.\x04)D\xa9\x06\x10<\x8f\x16u\x8d" +
	"\xdf\xa1\xc8\xdb\xd4\xe5\xc4\xf0u|\xf0%\x00\x7f\x9f\"" +
	"\xef\"8U)\xda?\x7f-\x861\x02\x04#\x80R" +
	"\xf5_yb\x18TL!\xf6\"\x00\x00\xff\xffM\x88" +
	":\xcb"

func init() {
	schemas.Register(schema_8b6f325f59887a6b,
		0x86899ecd650d61eb,
		0x9808fa17b68373f4,
		0xdd5a4545747ef1a4)
}