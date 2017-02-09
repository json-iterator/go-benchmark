package go_benchmark

import (
	"testing"
	"github.com/buger/jsonparser"
	"github.com/json-iterator/go"
	"encoding/json"
)

func BenchmarkJsonParserLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		count := 0
		jsonparser.ArrayEach(largeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			count++
		}, "topics", "topics")
	}
}

func BenchmarkJsoniterLarge(b *testing.B) {
	iter := jsoniter.ParseBytes(largeFixture)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iter.ResetBytes(largeFixture)
		count := 0
		for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
			if "topics" != field {
				iter.Skip()
				continue
			}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				if "topics" != field {
					iter.Skip()
					continue
				}
				for iter.ReadArray() {
					iter.Skip()
					count++
				}
				break
			}
			break
		}
	}
}

func BenchmarkEncodingJsonLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		payload := &LargePayload{}
		json.Unmarshal(largeFixture, payload)
	}
}