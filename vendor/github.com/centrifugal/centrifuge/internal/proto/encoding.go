package proto

import "sync"

// Encoding determines connection protocol encoding in use.
type Encoding string

const (
	// EncodingJSON means JSON protocol.
	EncodingJSON Encoding = "json"
	// EncodingProtobuf means protobuf protocol.
	EncodingProtobuf Encoding = "protobuf"
)

// GetPushEncoder ...
func GetPushEncoder(enc Encoding) PushEncoder {
	if enc == EncodingJSON {
		return NewJSONPushEncoder()
	}
	return NewProtobufPushEncoder()
}

var (
	jsonReplyEncoderPool     sync.Pool
	protobufReplyEncoderPool sync.Pool
)

// GetReplyEncoder ...
func GetReplyEncoder(enc Encoding) ReplyEncoder {
	if enc == EncodingJSON {
		e := jsonReplyEncoderPool.Get()
		if e == nil {
			return NewJSONReplyEncoder()
		}
		encoder := e.(ReplyEncoder)
		encoder.Reset()
		return encoder
	}
	e := protobufReplyEncoderPool.Get()
	if e == nil {
		return NewProtobufReplyEncoder()
	}
	encoder := e.(ReplyEncoder)
	encoder.Reset()
	return encoder
}

// PutReplyEncoder ...
func PutReplyEncoder(enc Encoding, e ReplyEncoder) {
	if enc == EncodingJSON {
		jsonReplyEncoderPool.Put(e)
		return
	}
	protobufReplyEncoderPool.Put(e)
}

// GetCommandDecoder ...
func GetCommandDecoder(enc Encoding, data []byte) CommandDecoder {
	if enc == EncodingJSON {
		return NewJSONCommandDecoder(data)
	}
	return NewProtobufCommandDecoder(data)
}

// PutCommandDecoder ...
func PutCommandDecoder(enc Encoding, e CommandDecoder) {
	return
}

// GetResultEncoder ...
func GetResultEncoder(enc Encoding) ResultEncoder {
	if enc == EncodingJSON {
		return NewJSONResultEncoder()
	}
	return NewProtobufResultEncoder()
}

// PutResultEncoder ...
func PutResultEncoder(enc Encoding, e ReplyEncoder) {
	return
}

// GetParamsDecoder ...
func GetParamsDecoder(enc Encoding) ParamsDecoder {
	if enc == EncodingJSON {
		return NewJSONParamsDecoder()
	}
	return NewProtobufParamsDecoder()
}

// PutParamsDecoder ...
func PutParamsDecoder(enc Encoding, e ParamsDecoder) {
	return
}
