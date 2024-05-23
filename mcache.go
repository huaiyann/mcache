package mcahce

import (
	"sync"
	"time"

	"github.com/huaiyann/mcache/expire"
)

type MCache struct {
	expires *expire.Expire
	m       *sync.Map
}

func New() *MCache {
	mc := &MCache{
		expires: expire.NewExpire(),
		m:       new(sync.Map),
	}
	go mc.doExpire()
	return mc
}

type cacheItem struct {
	key      string
	value    string
	expireAt time.Time
}

func (mc *MCache) Set(key string, value []byte, ttl time.Duration) {
	item := cacheItem{
		key:   key,
		value: string(value),
	}
	if ttl > 0 {
		item.expireAt = time.Now().Add(ttl)
	} else {
		item.expireAt = time.Now().Add(time.Hour * 24 * 365 * 200)
	}
	mc.m.Store(key, item)
	mc.expires.Add(key, item.expireAt)
}

func (mc *MCache) Get(key string) ([]byte, bool) {
	data, ok := mc.m.Load(key)
	if !ok {
		return nil, false
	}
	item := data.(cacheItem)
	if time.Now().After(item.expireAt) {
		mc.m.CompareAndDelete(key, data)
		return nil, false
	}
	return []byte(item.value), true
}

func (mc *MCache) doExpire() {
	for key := range mc.expires.NeedExpire() {
		data, ok := mc.m.Load(key)
		if !ok {
			continue
		}
		item := data.(cacheItem)
		if time.Now().After(item.expireAt) {
			mc.m.CompareAndDelete(key, data)
		}
	}
}
