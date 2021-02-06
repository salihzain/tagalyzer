package testdata

type OneTag struct {
	Name     string   // want "field:Name is missing tag:json"
	Age      int      // want "field:Age is missing tag:json"
	Location struct { // want "field:Location is missing tag:json"
		X float32 // want "field:X is missing tag:json"
		Y float32 // want "field:Y is missing tag:json"
	}
}
