package main

import (
	"encoding/json"
	"fmt"

	"github.com/krkeshav/myJson/myJson"
)

type RandomStruct struct {
	String         string
	Integer        int
	Float          float64
	Slice          []string
	Map            map[string]string
	IntMap         map[int]string // handle this case "IntMap":{"<int Value>":"Two"}
	OtherStruct    *OtherStruct
	MapStrToStruct map[string]*OtherStruct
}

type OtherStruct struct {
	OtherStructString string
	OtherInt          int
	OtherMap          map[string]string
}

func main() {
	rd := &RandomStruct{
		String:  "sample string",
		Integer: 69,
		Float:   69.420,
		Slice:   []string{"1", "2", "3"},
		Map: map[string]string{
			"Key1": "value1",
		},
		IntMap: map[int]string{
			2: "Two",
		},
		OtherStruct: &OtherStruct{
			OtherStructString: "other string",
			OtherInt:          3,
			OtherMap: map[string]string{
				"otherKey": "other Value",
			},
		},
		MapStrToStruct: map[string]*OtherStruct{
			"mapStructkey1": {
				OtherStructString: "other string",
				OtherInt:          3,
				OtherMap: map[string]string{
					"otherKey": "other Value",
				},
			},
		},
	}

	jsonHelper := myJson.NewJsonData(rd)
	rdJson := jsonHelper.EncodeValue()
	compJson, _ := json.Marshal(rd)
	fmt.Println(rdJson)
	fmt.Println(string(compJson))
}
