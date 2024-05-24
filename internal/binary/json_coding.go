package binary

import jsoniter "github.com/json-iterator/go"

var _json = jsoniter.ConfigCompatibleWithStandardLibrary

func Json(v any) BinaryCoding {
	return &jsoner{value: v}
}

type jsoner struct {
	value any
}

func (j *jsoner) UnmarshalBinary(data []byte) error {
	return _json.Unmarshal(data, j.value)
}

func (j *jsoner) MarshalBinary() ([]byte, error) {
	return _json.Marshal(j.value)
}
