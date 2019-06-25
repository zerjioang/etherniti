package ioproto

import (
	"github.com/zerjioang/etherniti/core/modules/encoding/ioproto/gogoproto"
	"github.com/zerjioang/etherniti/core/modules/encoding/ioproto/json"
	"github.com/zerjioang/etherniti/core/modules/encoding/ioproto/msgpack"
	"github.com/zerjioang/etherniti/core/modules/encoding/ioproto/xml"
	"github.com/zerjioang/etherniti/shared/protocol/io"
)

// return appropiate content type as descriped in HTTP header value Content-Type
func EncodingSelector(contentType string) (io.Serializer, io.ContentTypeMode) {
	switch contentType {
	case io.ContentTypeJson:
		return json.Serialize, io.ModeJson
	case io.ContentTypeXml:
		return xml.Serialize, io.ModeXML
	case io.ContentTypeMsgPack:
		return msgpack.Serialize, io.ModeMsgPack
	case io.ContentTypeProtoBuf:
		return gogoproto.Serialize, io.ModeProtoBuff
	default:
		//return json serializer as default when no one matches
		return json.Serialize, io.ModeJson
	}
}

// return appropiate content type as descriped in HTTP header value Content-Type
func EncodingModeSelector(mode io.ContentTypeMode) (io.Serializer, io.ContentTypeMode) {
	switch mode {
	case io.ModeJson:
		return json.Serialize, io.ModeJson
	case io.ModeXML:
		return xml.Serialize, io.ModeXML
	case io.ModeMsgPack:
		return msgpack.Serialize, io.ModeMsgPack
	case io.ModeProtoBuff:
		return gogoproto.Serialize, io.ModeProtoBuff
	default:
		//return json serializer as default when no one matches
		return json.Serialize, io.ModeJson
	}
}

// this method uses user requested encoding serializer via http headers
// and encodes result as byte array
func GetContentTypedBytes(contentType string, v interface{}) []byte {
	srlzr, _ := EncodingSelector(contentType)
	return GetBytesFromSerializer(srlzr, v)
}

func GetBytesFromSerializer(s io.Serializer, v interface{}) []byte {
	raw, _ := s(v)
	return raw
}

func GetBytesFromMode(mode io.ContentTypeMode, v interface{}) []byte {
	raw, _ := EncodingModeSelector(mode)
	data, _ := raw(v)
	return data
}
