package main

import (
	"encoding/json"
	"fmt"

	"github.com/krkeshav/myJson/myJson"
)

type RandomStruct struct {
	String             string                  `json:"string" encrypt:"true"`
	Integer            int                     `json:"integer"`
	Float              float64                 `json:"float" encrypt:"true"`
	Bool               bool                    `json:"bool"`
	Slice              []string                `json:"slice"`
	Map                map[string]string       `json:"map"`
	IntMap             map[int]string          `json:"int_map"` // handle this case "IntMap":{"<int Value>":"Two"}
	OtherStruct        *OtherStruct            `json:"other_struct"`
	MapStrToStruct     map[string]*OtherStruct `json:"map_str_to_struct"`
	SliceOfStruct      []OtherStruct           `json:"slice_struct"`
	SliceOfPointers    []*OtherStruct          `json:"slice_of_pointers"`
	NoJsonTag          string                  `json:"-"`
	OtherStructPointer *OtherStruct            `json:"other_struct_pointer,omitempty"`
	OtherSliceAgain    []string                `json:"other_slice,omitempty"`
	OtherMap           map[string]string       `json:"other_map,omitempty"`
	RawJson            string                  `json:"rawJson"`
}

type OtherStruct struct {
	OtherStructString string            `json:"other_struct_string" encrypt:"true"`
	OtherInt          int               `json:"other_int"`
	OtherMap          map[string]string `json:"other_map" encrypt:"true"`
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
		SliceOfStruct: []OtherStruct{
			{
				OtherStructString: "other string",
				OtherInt:          3,
				OtherMap: map[string]string{
					"otherKey": "other Value",
				},
			},
			{
				OtherStructString: "other string",
				OtherInt:          3,
				OtherMap: map[string]string{
					"otherKey": "other Value",
				},
			},
		},
		RawJson: `{"string":"encrypted","integer":69,"float":"encrypted","bool":false,"slice":["1","2","3"]}`,
		SliceOfPointers: []*OtherStruct{
			{},
		},
	}

	jsonHelper := myJson.NewJsonData(rd)
	rdJson := jsonHelper.EncodeValue()
	compJson, _ := json.Marshal(rd)
	fmt.Println(rdJson)
	fmt.Println(string(compJson))
}
