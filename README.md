test file used: https://github.com/json-iterator/test-data/blob/master/large-file.json

# codec

```
// "github.com/ugorji/go/codec"
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
```

5	 268996576 ns/op	33586776 B/op	 1042781 allocs/op

# standard lib
```
// "encoding/json"
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
```

5	 224364066 ns/op	71466916 B/op	  272474 allocs/op

# json iterator
```
// "github.com/json-iterator/go"
func Benchmark_jsoniter(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		file, _ := os.Open("/tmp/large-file.json")
		iter := jsoniter.Parse(file, 4096)
		count := 0
		for iter.ReadArray() {
			iter.Skip()
			count++
		}
		file.Close()
	}
}
```

10	 110209750 ns/op	    4248 B/op	       5 allocs/op