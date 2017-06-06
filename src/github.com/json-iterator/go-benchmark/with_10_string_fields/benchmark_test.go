package with_10_string_fields

import (
	"testing"
	"github.com/mailru/easyjson/jwriter"
	"bytes"
	"github.com/json-iterator/go"
)

func Benchmark_easyjson(b *testing.B) {
	str := "<a href=\"//itunes.apple.com/us/app/twitter/id409789998?mt=12%5C%22\" rel=\"\\\"nofollow\\\"\">Twitter for Mac</a>"
	obj := PbTestObject{str, str, str, str, str, str, str, str, str, str}
	//obj := PbTestObject{112, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	b.ReportAllocs()
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		writer := &jwriter.Writer{}
		obj.MarshalEasyJSON(writer)
		buf.Reset()
		writer.DumpTo(buf)
	}
}

func Benchmark_jsoniter(b *testing.B) {
	str := "<a href=\"//itunes.apple.com/us/app/twitter/id409789998?mt=12%5C%22\" rel=\"\\\"nofollow\\\"\">Twitter for Mac</a>"
	obj := PbTestObject{str, str, str, str, str, str, str, str, str, str}
	//obj := PbTestObject{112, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	b.ReportAllocs()
	stream := jsoniter.NewStream(nil, 512)
	for i := 0; i < b.N; i++ {
		stream.Reset(nil)
		stream.WriteVal(obj)
		if stream.Error != nil {
			b.Error(stream.Error)
		}
	}
}