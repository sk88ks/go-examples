package main

import "fmt"
import "github.com/EverythingMe/inbloom/go/inbloom"

func main() {
	// create a blank filter - expecting 20 members and an error rate of 1/100
	f, err := inbloom.NewFilter(5000, 0.001)
	if err != nil {
		panic(err)
	}

	// the size of the filter
	fmt.Println(f.Len())

	// insert some values
	f.Add("foo")
	f.Add("bar")
	f.Add("buzz")

	// test for existence of keys
	fmt.Println(f.Contains("foo"))
	fmt.Println(f.Contains("wat"))

	//fmt.Println("unmarshaled data:", f)
	marshaled := f.MarshalBase64()
	fmt.Println("marshaled data:", len(marshaled))
	//unmarshaled, _ := inbloom.UnmarshalBase64(marshaled)
	//fmt.Println("unmarshaled data:", unmarshaled)

	// Output:
	// 24
	// true
	// false
	// marshaled data: oU4AZAAAABQAAAAAAEIAABEAGAQAAgAgAAAwEAAJAAA=
}
