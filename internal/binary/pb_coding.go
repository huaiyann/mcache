package binary

import (
	gogo_proto "github.com/gogo/protobuf/proto"
	"google.golang.org/protobuf/proto"
)

func Protobuf(v proto.Message) BinaryCoding {
	return &protobuffer{value: v}
}

type protobuffer struct {
	value proto.Message
}

func (j *protobuffer) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, j.value)
}

func (j *protobuffer) MarshalBinary() ([]byte, error) {
	return proto.Marshal(j.value)
}

func ProtobufGogo(v gogo_proto.Message) BinaryCoding {
	return &gogoProtobuffer{value: v}
}

type gogoProtobuffer struct {
	value gogo_proto.Message
}

func (j *gogoProtobuffer) UnmarshalBinary(data []byte) error {
	return gogo_proto.Unmarshal(data, j.value)
}

func (j *gogoProtobuffer) MarshalBinary() ([]byte, error) {
	return gogo_proto.Marshal(j.value)
}
