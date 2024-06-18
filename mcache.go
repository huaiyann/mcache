package mcache

import (
	"sync"
	"time"

	"github.com/huaiyann/mcache/internal/expire"
	"google.golang.org/protobuf/proto"
)

type cacheItemWithExpire interface {
	GetExipireAt() time.Time
}

func New() *MCache {
	mc := &MCache{
		expires: expire.NewExpire(),
		m:       new(sync.Map),
	}
	go mc.doExpire()
	return mc
}

func Raw[T comparable](cache *MCache) RawValuer[T] {
	return RawValuer[T]{cache: cache}
}

func Json[T any](cache *MCache) JsonValuer[T] {
	return JsonValuer[T]{valuer: Raw[string](cache)}
}

func Protobuf[T proto.Message](cache *MCache) ProtobufValuer[T] {
	return ProtobufValuer[T]{valuer: Raw[string](cache)}
}

type MCache struct {
	expires *expire.Expire
	m       *sync.Map
}

func (mc *MCache) doExpire() {
	for key := range mc.expires.NeedExpire() {
		data, ok := mc.m.Load(key)
		if !ok {
			continue
		}
		item := data.(cacheItemWithExpire)
		if time.Now().After(item.GetExipireAt()) {
			mc.m.CompareAndDelete(key, data)
		}
	}
}

type cacheItem[T comparable] struct {
	key      string
	value    T
	expireAt time.Time
}

func (c cacheItem[T]) GetExipireAt() time.Time {
	return c.expireAt
}
