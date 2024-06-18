package mcache

import (
	"reflect"
	"time"

	"github.com/huaiyann/mcache/internal/binary"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type RawValuer[T comparable] struct {
	cache *MCache
}

func (k RawValuer[T]) Set(key string, data T, ttl time.Duration) error {
	item := cacheItem[T]{
		key:   key,
		value: data,
	}
	if ttl > 0 {
		item.expireAt = time.Now().Add(ttl)
	} else {
		item.expireAt = time.Now().Add(time.Hour * 24 * 365 * 200)
	}
	k.cache.m.Store(key, item)
	k.cache.expires.Add(key, item.expireAt)
	return nil
}

func (k RawValuer[T]) Get(key string) (T, bool) {
	var null T
	data, ok := k.cache.m.Load(key)
	if !ok {
		return null, false
	}
	item := data.(cacheItem[T])
	if time.Now().After(item.expireAt) {
		k.cache.m.CompareAndDelete(key, data)
		return null, false
	}
	return item.value, true
}

type JsonValuer[T any] struct {
	keyer RawValuer[string]
}

func (k JsonValuer[T]) Set(key string, data T, ttl time.Duration) error {
	buf, err := binary.Json(data).MarshalBinary()
	if err != nil {
		return errors.Wrapf(err, "MarshalBinary")
	}
	return k.keyer.Set(key, string(buf), ttl)
}

func (k JsonValuer[T]) Get(key string) (T, bool, error) {
	var null T
	data, has := k.keyer.Get(key)
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

func (k ProtobufValuer[T]) Set(key string, data T, ttl time.Duration) error {
	buf, err := binary.Protobuf(data).MarshalBinary()
	if err != nil {
		return errors.Wrapf(err, "MarshalBinary")
	}
	return k.valuer.Set(key, string(buf), ttl)
}

func (k ProtobufValuer[T]) Get(key string) (T, bool, error) {
	var null T
	data, has := k.valuer.Get(key)
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
