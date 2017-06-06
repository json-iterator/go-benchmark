package with_10_int_fields

import (
	"testing"
	"github.com/golang/protobuf/proto"
	"github.com/json-iterator/go"
	"encoding/json"
	"bytes"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/mailru/easyjson/jwriter"
)

func Benchmark_protobuf_read(b *testing.B) {
	b.ReportAllocs()
	obj := PbTestObject{31415926, 61415923, 31415269, 53141926, 13145926, 43115926, 31419265, 23141596, 43161592, 112}
	data, _ := proto.Marshal(&obj)
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(data, &obj)
	}
}

func Benchmark_protobuf_write(b *testing.B) {
	b.ReportAllocs()
	obj := PbTestObject{31415926, 61415923, 31415269, 53141926, 13145926, 43115926, 31419265, 23141596, 43161592, 112}
	for i := 0; i < b.N; i++ {
		proto.Marshal(&obj)
	}
}

func Benchmark_jsoniter_read(b *testing.B) {
	b.ReportAllocs()
	//obj := PbTestObject{31415926, 61415923, 31415269, 53141926, 13145926, 43115926, 31419265, 23141596, 43161592, 112}
	obj := PbTestObject{112, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	data, _ := jsoniter.Marshal(&obj)
	iter := jsoniter.NewIterator()
	for i := 0; i < b.N; i++ {
		iter.ResetBytes(data)
		iter.ReadVal(&obj)
	}
}

func Benchmark_jsoniter_write(b *testing.B) {
	b.ReportAllocs()
	obj := PbTestObject{31415926, 61415923, 31415269, 53141926, 13145926, 43115926, 31419265, 23141596, 43161592, 112}
	buf := &bytes.Buffer{}
	stream := jsoniter.NewStream(buf, 4096)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		stream.WriteVal(&obj)
	}
}

func Benchmark_json_read(b *testing.B) {
	b.ReportAllocs()
	obj := PbTestObject{31415926, 61415923, 31415269, 53141926, 13145926, 43115926, 31419265, 23141596, 43161592, 112}
	data, _ := jsoniter.Marshal(&obj)
	buf := &bytes.Buffer{}
	decoder := json.NewDecoder(buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(data)
		decoder.Decode(&obj)
	}
}

func Benchmark_json_write(b *testing.B) {
	b.ReportAllocs()
	obj := PbTestObject{31415926, 61415923, 31415269, 53141926, 13145926, 43115926, 31419265, 23141596, 43161592, 112}
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		encoder.Encode(&obj)
	}
}

func Benchmark_thrift(b *testing.B) {
	obj := ThriftTestObject{31415926, 61415923, 31415269, 53141926, 13145926, 43115926, 31419265, 23141596, 43161592, 112}
	buf := thrift.NewTMemoryBuffer()
	protocolFactory := &thrift.TCompactProtocolFactory{}
	protocol := protocolFactory.GetProtocol(buf)
	obj.Write(protocol)
	data := buf.Bytes()
	b.Run("thrift-write", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf.Reset()
			obj.Write(protocol)
		}
	})
	b.Run("thrift-read", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.Write(data)
			obj.Read(protocol)
		}
	})
}


func Benchmark_easyjson(b *testing.B) {
	obj := PbTestObject{31415926, 61415923, 31415269, 53141926, 13145926, 43115926, 31419265, 23141596, 43161592, 112}
	//obj := PbTestObject{112, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	writer := &jwriter.Writer{}
	obj.MarshalEasyJSON(writer)
	buf := &bytes.Buffer{}
	writer.DumpTo(buf)
	b.Run("easyjson-read", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			err := obj.UnmarshalJSON(buf.Bytes())
			if err != nil {
				b.Error(err)
			}
		}
	})
}