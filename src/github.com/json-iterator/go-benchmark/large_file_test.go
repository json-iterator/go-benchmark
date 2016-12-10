package go_benchmark

import (
	"testing"
	"github.com/json-iterator/go"
	"os"
	"encoding/json"
	"github.com/buger/jsonparser"
	"io/ioutil"
)

func Test_jsonparser_skip(t *testing.T) {
	file, _ := os.Open("/tmp/large-file.json")
	bytes, _ := ioutil.ReadAll(file)
	file.Close()
	total := 0
	jsonparser.ArrayEach(bytes, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		total++
	})
	if total != 11351 {
		t.Fatal(total)
	}
}

func Test_jsoniter_skip(t *testing.T) {
	for i := 0; i < 100; i++ {
		file, _ := os.Open("/tmp/large-file.json")
		iter := jsoniter.Parse(file, 4096)
		total := 0
		for iter.ReadArray() {
			iter.Skip()
			total++
		}
		file.Close()
		if total != 11351 {
			t.Fatal(total)
		}
	}
}

func Benchmark_jsonparser(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		file, _ := os.Open("/tmp/large-file.json")
		bytes, _ := ioutil.ReadAll(file)
		file.Close()
		total := 0
		jsonparser.ArrayEach(bytes, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			total++
		})
	}
}

func Benchmark_stardard_lib(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		file, _ := os.Open("/tmp/large-file.json")
		result := []struct{}{}
		decoder := json.NewDecoder(file)
		decoder.Decode(&result)
		file.Close()
	}
}

func Benchmark_jsoniter(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		file, _ := os.Open("/tmp/large-file.json")
		iter := jsoniter.Parse(file, 4096)
		for iter.ReadArray() {
			iter.Skip()
		}
		file.Close()
	}
}