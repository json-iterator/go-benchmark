package go_benchmark

/*
   Small paylod, http log like structure. Size: 190 bytes
*/
var smallFixture []byte = []byte(`{
    "st": 1,
    "sid": 486,
    "tt": "active",
    "gr": 0,
    "uuid": "de305d54-75b4-431b-adb2-eb6b9e546014",
    "ip": "127.0.0.1",
    "ua": "user_agent",
    "tz": -6,
    "v": 1
}`)

type SmallPayload struct {
	St   int `json:"st"`
	Sid  int `json:"-"`
	Tt   string `json:"-"`
	Gr   int `json:"-"`
	Uuid string `json:"uuid"`
	Ip   string `json:"-"`
	Ua   string `json:"ua"`
	Tz   int `json:"tz"`
	V    int `json:"-"`
}

var arrayFixture []byte = []byte(`["1","2","3"]
`)
