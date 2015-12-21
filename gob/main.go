package main // GobEncode converts data to a bobs byte slice by gob
import (
	"bytes"
	"encoding/gob"
	"reflect"

	"github.com/k0kubun/pp"
)

// Data is data
type Data struct {
	ID      string
	Number  int
	Score   float64
	SubData []Data
}

func conv(data []byte, rec interface{}) error {
	refValue := reflect.ValueOf(rec)
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.DecodeValue(refValue)
}

func GobEncode(data interface{}) []byte {
	refValue := reflect.ValueOf(data)
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.EncodeValue(refValue)
	return buf.Bytes()
}

// GobDecode converts a gobs byte slice to a given receiver
// You should give a gobs formatted byte slice as the arg
func GobDecode(data []byte, rec interface{}) error {
	refValue := reflect.ValueOf(rec)
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.DecodeValue(refValue)
}

func main() {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(Data{
		"test_001",
		0,
		1.02,
		[]Data{
			{
				"sub_1_1",
				0,
				1.1,
				nil,
			},
			{
				"sub_1_2",
				1,
				1.2,
				nil,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	//dec := gob.NewDecoder(&buf)
	//err = dec.Decode(&rec)
	res := Data{}
	err = conv(buf.Bytes(), &res)
	if err != nil {
		panic(err)
	}
	pp.Println(res)

	arr := []int{0, 1, 2, 3, 4, 5}
	byts := GobEncode(arr)

	rec := []int{}
	err = GobDecode(byts, &rec)
	pp.Println(err)
	pp.Println(rec)
}
