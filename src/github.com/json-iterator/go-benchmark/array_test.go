package go_benchmark

import (
	"testing"
	"encoding/json"
	"github.com/json-iterator/go"
)

func Benchmark_array_by_stardard_lib(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		sample := make([]int, 0, 10)
		json.Unmarshal([]byte(`[1,2,3,4,5,6,7,8,9]`), &sample)
	}
}

func Benchmark_array_by_jsoniter(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		sample := make([]int, 0, 10)
		jsoniter.Unmarshal([]byte(`[1,2,3,4,5,6,7,8,9]`), &sample)
	}
}

func Benchmark_array_by_jsoniter_direct(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		sample := make([]uint64, 0, 10)
		iter := jsoniter.ParseString(`[1,2,3,4,5,6,7,8,9]`)
		for iter.ReadArray() {
			sample = append(sample, iter.ReadUint64())
		}
	}
}