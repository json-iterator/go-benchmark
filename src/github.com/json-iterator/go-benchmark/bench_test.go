package go_benchmark

import (
	"testing"
	"github.com/ugorji/go/codec"
	"github.com/json-iterator/go"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bufio"
)

func Test_codec(t *testing.T) {
	result := []struct{}{}
	file, _ := os.Open("/tmp/large-file.json")
	dec := codec.NewDecoder(file, &codec.JsonHandle{})
	dec.Decode(&result)
	file.Close()
	fmt.Println(len(result))
}

func Benchmark_codec(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		result := []struct{}{}
		file, _ := os.Open("/tmp/large-file.json")
		reader := bufio.NewReader(file)
		dec := codec.NewDecoder(reader, &codec.JsonHandle{})
		dec.Decode(&result)
		file.Close()
	}
}

func Benchmark_stardard_lib(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		file, _ := os.Open("/tmp/large-file.json")
		bytes, _ := ioutil.ReadAll(file)
		file.Close()
		result := []struct{}{}
		json.Unmarshal(bytes, &result)
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