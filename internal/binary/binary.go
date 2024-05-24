package binary

import "encoding"

type BinaryCoding interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
