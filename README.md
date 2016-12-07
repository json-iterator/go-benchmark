* jsonparser: https://github.com/buger/jsonparser
* jsoniter pull-api: https://github.com/json-iterator/go
* jsoniter reflect-api: https://github.com/json-iterator/go/blob/master/jsoniter_reflect.go
* encoding/json: golang standard lib
* easy json: https://github.com/mailru/easyjson

The goal is to prove jsoniter is not slow, not to prove it is the fastest, at least good enough.
My motivation of inventing json iterator is the flexibility to mix high level and low level api.
Performance is the by-product of schema based parsing.

* CPU: i7-6700K @ 4.0G
* Level 1 cache size: 4 x 32 KB 8-way set associative instruction caches
* Level 2 cache size: 4 x 256 KB 4-way set associative caches
* Level 3 cache size: 8 MB 16-way set associative shared cache
* Go: 1.8beta1

# small payload

https://github.com/json-iterator/go-benchmark/blob/master/src/github.com/json-iterator/go-benchmark/benchmark_small_payload_test.go

| jsonparser  | jsoniter pull-api | jsoniter reflect-api | encoding/json | easyjson    |
| ---         | ---               | ---                  | ---           | ---         |
| 599 ns/op   | 515 ns/op         | 684 ns/op            | 2453 ns/op    | 687 ns/op   |
| 64 B/op     | 64 B/op           | 256 B/op             | 864 B/op      | 64 B/op     |
| 2 allocs/op | 2 allocs/op       | 4 allocs/op          | 31 allocs/op  | 2 allocs/op |

pull-api is fast, and reflection-api is not slow either.
encoding/json is not that slow on i7-6700K, but much slower on cpu with smaller cache.

![small](/small.png)

# medium payload

https://github.com/json-iterator/go-benchmark/blob/master/src/github.com/json-iterator/go-benchmark/benchmark_medium_payload_test.go

| jsonparser  | jsoniter pull-api | jsoniter reflect-api | encoding/json | easyjson    |
| ---         | ---               | ---                  | ---           | ---         |
| 5238 ns/op  | 4111 ns/op        | 4708 ns/op           | 24939 ns/op   | 7361 ns/op  |
| 104 B/op    | 104 B/op          | 368 B/op             | 808 B/op      | 248 B/op    |
| 4 allocs/op | 4 allocs/op       | 14 allocs/op         | 18 allocs/op  | 8 allocs/op |

reflect-api even out-performed the hand written parser

![medium](/medium.png)

# large payload

https://github.com/json-iterator/go-benchmark/blob/master/src/github.com/json-iterator/go-benchmark/benchmark_large_payload_test.go

| jsonparser  | jsoniter pull-api | encoding/json |
| ---         | ---               | ---           |
| 38334 ns/op | 38463 ns/op       | 290778 ns/op  |
| 0 B/op      | 0 B/op            | 2128 B/op     |
| 0 allocs/op | 0 allocs/op       | 46 allocs/op  |

This is a pure counting usage. jsonparser is faster

![large](/large.png)

# large file

test file used: https://github.com/json-iterator/test-data/blob/master/large-file.json

| jsonparser     | jsoniter pull-api | encoding/json    |
| ---            | ---               | ---              |
| 42698634 ns/op | 37760014 ns/op    | 235354502 ns/op  |
| 67107104 B/op  | 4248 B/op         | 71467896 B/op    |
| 19 allocs/op   | 5 allocs/op       | 272477 allocs/op |

The difference here is because jsonparser take []byte as input, but jsoniter can take io.Reader as input.

![large-file](/large-file.png)
