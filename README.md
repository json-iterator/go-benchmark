test file used: https://github.com/json-iterator/test-data/blob/master/large-file.json

# codec (skip)

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

# standard lib (skip)
```
// "encoding/json"
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
```

5	 215547514 ns/op	71467118 B/op	  272476 allocs/op

# json iterator (skip)
```
// "github.com/json-iterator/go"
func Benchmark_jsoniter(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		file, _ := os.Open("/tmp/large-file.json")
		iter := jsoniter.Parse(file, 1024)
		for iter.ReadArray() {
			iter.Skip()
		}
		file.Close()
	}
}
```

10	 110209750 ns/op	    4248 B/op	       5 allocs/op

# standard lib (struct)

```
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
```

1000000	      1565 ns/op	     496 B/op	      16 allocs/op

# json iterator (struct)

```
type SampleStruct struct {
	field1 int
	field2 int `json:",string"`
}

func Benchmark_struct_by_jsoniter(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		sample := SampleStruct{}
		jsoniter.Unmarshal([]byte(`{"field1": 100, "field2": "102"}`), &sample)
	}
}
```

3000000	       552 ns/op	     144 B/op	       5 allocs/op

# standard lib (array)

```
func Benchmark_array_by_stardard_lib(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		sample := make([]int, 0, 10)
		json.Unmarshal([]byte(`[1,2,3,4,5,6,7,8,9]`), &sample)
	}
}
```

500000	      2478 ns/op	     408 B/op	      14 allocs/op

# json iterator (struct)

```
func Benchmark_array_by_jsoniter(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		sample := make([]int, 0, 10)
		jsoniter.Unmarshal([]byte(`[1,2,3,4,5,6,7,8,9]`), &sample)
	}
}
```

2000000	       740 ns/op	     224 B/op	       4 allocs/op