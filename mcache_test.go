package mcahce

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func BenchmarkMcSet(b *testing.B) {
	mc := New()
	for i := 0; i < b.N; i++ {
		mc.Set(fmt.Sprintf("%d", i), nil, time.Second*time.Duration(i))
	}
}

func BenchmarkMcGet(b *testing.B) {
	mc := New()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("%d", i)
		mc.Set(key, nil, time.Nanosecond)
		mc.Get(key)
	}
}

func TestMCache_Get(t *testing.T) {
	mc := New()
	type args struct {
		key   string
		value []byte
		ttl   time.Duration
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 bool
	}{
		{
			name: "cache hit",
			args: args{
				key:   "cache_hit",
				value: []byte("cache_hit"),
				ttl:   time.Second,
			},
			want:  []byte("cache_hit"),
			want1: true,
		},
		{
			name: "cache expire miss",
			args: args{
				key:   "cache_miss",
				value: []byte("cache_miss"),
				ttl:   time.Nanosecond,
			},
			want:  nil,
			want1: false,
		},
		{
			name: "cache no expire",
			args: args{
				key:   "cache_no_expire",
				value: []byte("cache_no_expire"),
			},
			want:  []byte("cache_no_expire"),
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc.Set(tt.args.key, tt.args.value, tt.args.ttl)
			got, got1 := mc.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCache.Get() got = %s, want %s", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MCache.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMCacheRewrite(t *testing.T) {
	mc := New()
	mc.Set("key", []byte("data1"), time.Millisecond*10)
	mc.Set("key", []byte("data2"), time.Millisecond*10)
	mc.Set("key", []byte("data3"), time.Millisecond*100)
	value, ok := mc.Get("key")
	if !ok {
		t.Errorf("ok is %v", ok)
	}
	if !reflect.DeepEqual(value, []byte("data3")) {
		t.Errorf("value want %s but %s", "data3", value)
	}
}

func TestMCacheRewriteExpire(t *testing.T) {
	mc := New()
	mc.Set("key", []byte("data1"), time.Second*10)
	mc.Set("key", []byte("data2"), time.Second*10)
	mc.Set("key", []byte("data3"), time.Nanosecond)
	_, ok := mc.Get("key")
	if ok {
		t.Errorf("ok is %v", ok)
	}
}
