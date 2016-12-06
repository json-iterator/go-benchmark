test file used: https://github.com/json-iterator/test-data/blob/master/large-file.json

* jsonparser: https://github.com/buger/jsonparser
* jsoniter pull-api: https://github.com/json-iterator/go
* jsoniter reflect-api: https://github.com/json-iterator/go/blob/master/jsoniter_reflect.go
* encoding/json: golang standard lib

# small payload

| jsonparser  | jsoniter pull-api | jsoniter reflect-api | encoding/json |
| --          | --                | --                   | --            |
| 691 ns/op   | 661 ns/op         | 982 ns/op            | 4740 ns/op	  |
| 64 B/op     | 64 B/op           | 256 B/op             | 864 B/op      |
| 2 allocs/op | 2 allocs/op       | 4 allocs/op          | 31 allocs/op  |

