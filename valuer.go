package mcache

import (
	"reflect"
	"time"

	gogo_proto "github.com/gogo/protobuf/proto"
	"github.com/huaiyann/mcache/internal/binary"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type RawValuer[T comparable] struct {
	cache *MCache
}

func (v RawValuer[T]) Set(key string, data T, ttl time.Duration) error {
	item := cacheItem[T]{
		key:   key,
		value: data,
	}
	if ttl > 0 {
		item.expireAt = time.Now().Add(ttl)
	} else {
		item.expireAt = time.Now().Add(time.Hour * 24 * 365 * 200)
	}
	v.cache.m.Store(key, item)
	v.cache.expires.Add(key, item.expireAt)
	return nil
}

func (v RawValuer[T]) Get(key string) (T, bool) {
	var null T
	data, ok := v.cache.m.Load(key)
	if !ok {
		return null, false
	}
	item := data.(cacheItem[T])
	if time.Now().After(item.expireAt) {
		v.cache.m.CompareAndDelete(key, data)
		return null, false
	}
	return item.value, true
}

type JsonValuer[T any] struct {
	valuer RawValuer[string]
}

func (v JsonValuer[T]) Set(key string, data T, ttl time.Duration) error {
	buf, err := binary.Json(data).MarshalBinary()
	if err != nil {
		return errors.Wrapf(err, "MarshalBinary")
	}
	return v.valuer.Set(key, string(buf), ttl)
}

func (v JsonValuer[T]) Get(key string) (T, bool, error) {
	var null T
	data, has := v.valuer.Get(key)
	if !has {
		return null, false, nil
	}

	tt := reflect.TypeOf(null)
	vv := reflect.New(tt)

	err := binary.Json(vv.Interface()).UnmarshalBinary([]byte(data))
	if err != nil {
		return null, false, errors.Wrapf(err, "UnmarshalBinary")
	}
	return vv.Elem().Interface().(T), true, nil
}

type ProtobufValuer[T proto.Message] struct {
	valuer RawValuer[string]
}

func (v ProtobufValuer[T]) Set(key string, data T, ttl time.Duration) error {
	buf, err := binary.Protobuf(data).MarshalBinary()
	if err != nil {
		return errors.Wrapf(err, "MarshalBinary")
	}
	return v.valuer.Set(key, string(buf), ttl)
}

func (v ProtobufValuer[T]) Get(key string) (T, bool, error) {
	var null T
	data, has := v.valuer.Get(key)
	if !has {
		return null, false, nil
	}

	tt := reflect.TypeOf(null).Elem()
	vv := reflect.New(tt)

	target := vv.Interface().(proto.Message)

	err := binary.Protobuf(target).UnmarshalBinary([]byte(data))
	if err != nil {
		return null, false, errors.Wrapf(err, "UnmarshalBinary")
	}
	return vv.Interface().(T), true, nil
}

type ProtobufGogoValuer[T gogo_proto.Message] struct {
	valuer RawValuer[string]
}

func (v ProtobufGogoValuer[T]) Set(key string, data T, ttl time.Duration) error {
	buf, err := binary.ProtobufGogo(data).MarshalBinary()
	if err != nil {
		return errors.Wrapf(err, "MarshalBinary")
	}
	return v.valuer.Set(key, string(buf), ttl)
}

func (v ProtobufGogoValuer[T]) Get(key string) (T, bool, error) {
	var null T
	data, has := v.valuer.Get(key)
	if !has {
		return null, false, nil
	}

	tt := reflect.TypeOf(null).Elem()
	vv := reflect.New(tt)

	target := vv.Interface().(gogo_proto.Message)

	err := binary.ProtobufGogo(target).UnmarshalBinary([]byte(data))
	if err != nil {
		return null, false, errors.Wrapf(err, "UnmarshalBinary")
	}
	return vv.Interface().(T), true, nil
}
