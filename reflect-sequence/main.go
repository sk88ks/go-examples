package main

import (
	"reflect"
	"sort"

	"github.com/k0kubun/pp"
)

type Sequence struct {
	Elements interface{}
}

func (s Sequence) Len() int {
	rv := reflect.ValueOf(s.Elements)
	if rv.Kind() != reflect.Slice {
		return 0
	}
	return rv.Len()
}

func (s Sequence) Swap(i, j int) {
	rv := reflect.ValueOf(s.Elements)
	if rv.Kind() != reflect.Slice {
		return
	}
	iVal := rv.Index(i).Interface()
	jVal := rv.Index(j).Interface()
	rv.Index(i).Set(reflect.ValueOf(jVal))
	rv.Index(j).Set(reflect.ValueOf(iVal))
}

type SortFloat64 struct {
	Sequence
	key string
}

type SortString struct {
	Sequence
	key string
}

func (s SortFloat64) Less(i, j int) bool {
	rv := reflect.ValueOf(s.Elements)
	if rv.Kind() != reflect.Slice {
		return false
	}
	a := rv.Index(i)
	b := rv.Index(j)
	if a.Kind() != reflect.Struct || b.Kind() != reflect.Struct {
		return false
	}

	aVal := a.FieldByName(s.key)
	bVal := b.FieldByName(s.key)
	if aVal.Kind() != reflect.Float64 || bVal.Kind() != reflect.Float64 {
		return false
	}

	aFloat := aVal.Interface().(float64)
	bFloat := bVal.Interface().(float64)
	return aFloat < bFloat
}

func (s SortString) Less(i, j int) bool {
	rv := reflect.ValueOf(s.Elements)
	if rv.Kind() != reflect.Slice {
		return false
	}
	a := rv.Index(i)
	b := rv.Index(j)
	if a.Kind() != reflect.Struct || b.Kind() != reflect.Struct {
		return false
	}

	aVal := a.FieldByName(s.key)
	bVal := b.FieldByName(s.key)
	if aVal.Kind() != reflect.String || bVal.Kind() != reflect.String {
		return false
	}

	aStr := aVal.Interface().(string)
	bStr := bVal.Interface().(string)
	return aStr < bStr
}

func Loop(data interface{}) {
	rv := reflect.ValueOf(data)
	if rv.Kind() != reflect.Ptr {
		return
	}
	rv = rv.Elem()

	switch rv.Kind() {
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			v := rv.Index(i)
			pp.Println(v)
		}
	}
}

type Data struct {
	ID    string
	Name  string
	Age   int
	Score float64
}

func main() {
	data := []Data{
		{
			ID:    "test000",
			Name:  "name000",
			Age:   1,
			Score: 1.2345,
		},
		{
			ID:    "test001",
			Name:  "name001",
			Age:   2,
			Score: 2.3456,
		},
	}

	seq := Sequence{
		Elements: data,
	}

	sf := SortFloat64{
		Sequence: seq,
		key:      "Score",
	}

	sort.Sort(sort.Reverse(sf))
	pp.Println(seq.Elements)

	ss := SortString{
		Sequence: seq,
		key:      "ID",
	}
	sort.Sort(ss)
	pp.Println(seq.Elements)
}
