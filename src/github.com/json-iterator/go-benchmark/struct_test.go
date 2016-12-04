package go_benchmark

import (
	"testing"
	"encoding/json"
	"github.com/json-iterator/go"
)

type SampleStruct struct {
	field1 int
	field2 int `json:",string"`
}

func Benchmark_struct_by_stardard_lib(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		sample := SampleStruct{}
		json.Unmarshal([]byte(`{"field1": 100, "field2": "102"}`), &sample)
	}
}

func Benchmark_struct_by_jsoniter(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		sample := SampleStruct{}
		jsoniter.Unmarshal([]byte(`{"field1": 100, "field2": "102"}`), &sample)
	}
}
