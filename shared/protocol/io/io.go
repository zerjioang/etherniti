package io

//serialization function header

// this functions convert input go object to byte array
type Serializer func(v interface{}) ([]byte, error)

const (
	ContentTypeJson     = "application/json"
	ContentTypeXml      = "application/xml"
	ContentTypeMsgPack  = "application/x-msgpack"
	ContentTypeProtoBuf = "application/protobuf"
)

type ContentTypeMode uint8

const (
	ModeJson ContentTypeMode = iota
	ModeXML
	ModeMsgPack
	ModeProtoBuff
)
