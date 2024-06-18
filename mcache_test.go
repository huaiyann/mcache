package mcache

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/huaiyann/mcache/internal/pb"
	"google.golang.org/protobuf/proto"
)

func BenchmarkMcSet(b *testing.B) {
	mc := New()
	valuer := Raw[int](mc)
	for i := 0; i < b.N; i++ {
		err := valuer.Set(fmt.Sprintf("%d", i), 1, time.Second*time.Duration(i))
		if err != nil {
			b.Errorf("err is %v", err)
		}
	}
}

func BenchmarkMcGet(b *testing.B) {
	valuer := Raw[int](New())
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("%d", i)
		err := valuer.Set(key, 1, 0)
		if err != nil {
			b.Errorf("err is %v", err)
		}
		_, has := valuer.Get(key)
		if !has {
			b.Errorf("has is %v", has)
		}
	}
}

func testRawValuer[T comparable](t *testing.T, valuer RawValuer[T], key string, value T) {
	err := valuer.Set(key, value, time.Hour)
	if err != nil {
		t.Errorf("testRawValuer, key %s, got err: %v", key, err)
	}
	v, has := valuer.Get(key)
	if !has {
		t.Errorf("testRawValuer, key %s, has is: %v", key, has)
	}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("testRawValuer, key %s, want %v but %v", key, value, v)
	}
}

func TestRawValuer(t *testing.T) {
	cache := New()

	var intVal int = 1
	testRawValuer(t, Raw[int](cache), "test_raw_int", intVal)
	testRawValuer(t, Raw[*int](cache), "test_raw_int_ptr", &intVal)

	var strVal string = "1"
	testRawValuer(t, Raw[string](cache), "test_raw_string", strVal)
	testRawValuer(t, Raw[*string](cache), "test_raw_string_ptr", &strVal)

	type StructVal struct {
		Val1 int
		Val2 string
	}
	var structVal = StructVal{Val1: 123, Val2: "123"}
	testRawValuer(t, Raw[StructVal](cache), "test_raw_struct", structVal)
	testRawValuer(t, Raw[*StructVal](cache), "test_raw_struct_ptr", &structVal)
}

func testJsonValuer[T any](t *testing.T, valuer JsonValuer[T], key string, value T) {
	err := valuer.Set(key, value, time.Hour)
	if err != nil {
		t.Errorf("testJsonValuer, key %s, got err: %v", key, err)
	}
	v, has, err := valuer.Get(key)
	if err != nil {
		t.Errorf("testJsonValuer, key %s, got err: %v", key, err)
	}
	if !has {
		t.Errorf("testJsonValuer, key %s, has is: %v", key, has)
	}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("testJsonValuer, key %s, want %v but %v", key, value, v)
	}
}

func TestJsonValuer(t *testing.T) {
	cache := New()

	var intVal int = 1
	testJsonValuer(t, Json[int](cache), "test_json_int", intVal)
	testJsonValuer(t, Json[*int](cache), "test_json_int_ptr", &intVal)

	var strVal = "123"
	testJsonValuer(t, Json[string](cache), "test_json_string", strVal)
	testJsonValuer(t, Json[*string](cache), "test_json_string_ptr", &strVal)

	var mapVal = map[string]string{"1": "1", "2": "2"}
	testJsonValuer(t, Json[map[string]string](cache), "test_json_map", mapVal)
	testJsonValuer(t, Json[*map[string]string](cache), "test_json_map_ptr", &mapVal)

	var bytesVal = []byte("123")
	testJsonValuer(t, Json[[]byte](cache), "test_json_bytes", bytesVal)
	testJsonValuer(t, Json[*[]byte](cache), "test_json_bytes_ptr", &bytesVal)

	var sliceVal = []int{1, 2, 3}
	testJsonValuer(t, Json[[]int](cache), "test_json_slice", sliceVal)
	testJsonValuer(t, Json[*[]int](cache), "test_json_slice_ptr", &sliceVal)

	type StructVal struct {
		Val1 int
		Val2 string
	}
	var structVal = StructVal{Val1: 123, Val2: "123"}
	testJsonValuer(t, Json[StructVal](cache), "test_json_struct", structVal)
	testJsonValuer(t, Json[*StructVal](cache), "test_json_struct_ptr", &structVal)
}

func testProtobufValuer[T proto.Message](t *testing.T, valuer ProtobufValuer[T], key string, value T) {
	err := valuer.Set(key, value, time.Hour)
	if err != nil {
		t.Errorf("testJsonValuer, key %s, got err: %v", key, err)
	}
	v, has, err := valuer.Get(key)
	if err != nil {
		t.Errorf("testJsonValuer, key %s, got err: %v", key, err)
	}
	if !has {
		t.Errorf("testJsonValuer, key %s, has is: %v", key, has)
	}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("testJsonValuer, key %s, want %v but %v", key, value, v)
	}
}

func TestProtobufValuer(t *testing.T) {
	cache := New()

	var structVal = pb.PbExample{ID: 123, Value: 456}
	testProtobufValuer(t, Protobuf[*pb.PbExample](cache), "test_pb_struct_ptr", &structVal)
}

func TestDataModifyRaw(t *testing.T) {
	cache := New()

	var intVal int = 1
	originalIntVal := intVal
	intValuer := Raw[int](cache)
	intPtrValuer := Raw[*int](cache)
	err := intValuer.Set("int_val", intVal, time.Hour)
	if err != nil {
		t.Errorf("int valuer set, error: %v", err)
	}
	err = intPtrValuer.Set("int_ptr_val", &intVal, time.Hour)
	if err != nil {
		t.Errorf("int ptr valuer set, error: %v", err)
	}
	intVal = 2 // modify it
	intGot, _ := intValuer.Get("int_val")
	if intGot != originalIntVal {
		t.Errorf("int valuer want %d but %d", originalIntVal, intGot)
	}
	intPtrGot, _ := intPtrValuer.Get("int_ptr_val")
	if *intPtrGot != intVal {
		t.Errorf("int ptr valuer want %d but %d", intVal, *intPtrGot)
	}
}

func TestDataModifyJson(t *testing.T) {
	cache := New()

	var intVal int = 1
	originalIntVal := intVal
	intValuer := Json[int](cache)
	intPtrValuer := Json[*int](cache)
	err := intValuer.Set("int_val", intVal, time.Hour)
	if err != nil {
		t.Errorf("int valuer set, error: %v", err)
	}
	err = intPtrValuer.Set("int_ptr_val", &intVal, time.Hour)
	if err != nil {
		t.Errorf("int ptr valuer set, error: %v", err)
	}
	intVal = 2 // modify it
	intGot, _, _ := intValuer.Get("int_val")
	if intGot != originalIntVal {
		t.Errorf("int valuer want %d but %d", originalIntVal, intGot)
	}
	intPtrGot, _, _ := intPtrValuer.Get("int_ptr_val")
	if intGot != originalIntVal {
		t.Errorf("int ptr valuer want %d but %d", originalIntVal, *intPtrGot)
	}
}

func TestMCacheNoExpire(t *testing.T) {
	valuer := Raw[string](New())
	err := valuer.Set("key1", "data1", 0)
	if err != nil {
		t.Errorf("err is %v", err)
	}
	data, has := valuer.Get("key1")
	if !has {
		t.Errorf("has is %v", has)
	}
	if data != "data1" {
		t.Errorf("want data1 but %s", data)
	}
}

func TestMCacheMiss(t *testing.T) {
	valuer := Raw[string](New())
	err := valuer.Set("key1", "data1", time.Second*10)
	if err != nil {
		t.Errorf("err is %v", err)
	}
	data, has := valuer.Get("key2")
	if has {
		t.Errorf("has is %v", has)
	}
	if data != "" {
		t.Errorf("want null but %s", data)
	}
}

func TestMCacheRewrite(t *testing.T) {
	mc := New()
	valuer := Raw[string](mc)
	err := valuer.Set("key", "data1", time.Second*10)
	if err != nil {
		t.Errorf("err is %v", err)
	}
	err = valuer.Set("key", "data2", time.Second*10)
	if err != nil {
		t.Errorf("err is %v", err)
	}
	err = valuer.Set("key", "data3", time.Second*10)
	if err != nil {
		t.Errorf("err is %v", err)
	}
	data, has := valuer.Get("key")
	if !has {
		t.Errorf("has is %v", has)
	}
	if data != "data3" {
		t.Errorf("want data3 but %s", data)
	}
}

func TestMCacheRewriteExpire(t *testing.T) {
	mc := New()
	valuer := Raw[string](mc)
	err := valuer.Set("key", "data1", time.Second*10)
	if err != nil {
		t.Errorf("err is %v", err)
	}
	err = valuer.Set("key", "data2", time.Second*10)
	if err != nil {
		t.Errorf("err is %v", err)
	}
	err = valuer.Set("key", "data3", time.Nanosecond)
	if err != nil {
		t.Errorf("err is %v", err)
	}
	data, has := valuer.Get("key")
	if has {
		t.Errorf("has is %v", has)
	}
	if data != "" {
		t.Errorf("want null but %s", data)
	}
}

func TestMCacheTickExpireRaw(t *testing.T) {
	mc := New()

	valuer := Raw[*int](mc)
	var val int = 123
	err := valuer.Set("key", &val, time.Millisecond*100)
	if err != nil {
		t.Errorf("err is %v", err)
	}
	<-time.After(time.Millisecond * 300)
	got, has := valuer.Get("key")
	if has {
		t.Errorf("has is %v", has)
	}
	if got != nil {
		t.Errorf("want nil but %v", got)
	}
}

func TestMCacheTickExpireJson(t *testing.T) {
	mc := New()

	valuer := Json[*int](mc)
	var val int = 123
	err := valuer.Set("key", &val, time.Millisecond*100)
	if err != nil {
		t.Errorf("err is %v", err)
	}
	<-time.After(time.Millisecond * 300)
	got, has, err := valuer.Get("key")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if has {
		t.Errorf("has is %v", has)
	}
	if got != nil {
		t.Errorf("want nil but %v", got)
	}
}

func TestMCacheTickExpireProtobuf(t *testing.T) {
	mc := New()

	valuer := Protobuf[*pb.PbExample](mc)
	val := &pb.PbExample{ID: 123, Value: 456}
	err := valuer.Set("key", val, time.Millisecond*100)
	if err != nil {
		t.Errorf("err is %v", err)
	}
	<-time.After(time.Millisecond * 300)
	got, has, err := valuer.Get("key")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if has {
		t.Errorf("has is %v", has)
	}
	if got != nil {
		t.Errorf("want nil but %v", got)
	}
}
