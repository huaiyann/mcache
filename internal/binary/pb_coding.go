package binary

import "google.golang.org/protobuf/proto"

func Protobuf(v proto.Message) BinaryCoding {
	return &jsoner{value: v}
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
