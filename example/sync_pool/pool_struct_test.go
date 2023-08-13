package sync_pool

import (
	"encoding/json"
	"sync"
	"testing"
)

type Student struct {
	Name   string
	Age    int32
	Remark []byte
}

var buf, _ = json.Marshal(Student{
	Name:   "John",
	Age:    18,
	Remark: []byte("good"),
})

func UnMarshal() {
	stu := Student{}
	_ = json.Unmarshal(buf, &stu)
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func UnMarshalWithPool() {
	stu := studentPool.Get().(*Student) // get
	json.Unmarshal(buf, stu)
	studentPool.Put(stu) // put
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UnMarshal()
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UnMarshalWithPool()
	}
}
