package testdata

type OneTag struct {
	Name     string   // want "field:Name is missing tag:json" "field:Name is missing tag:gorm"
	Age      int      // want "field:Age is missing tag:json" "field:Age is missing tag:gorm"
	Location struct { // want "field:Location is missing tag:json" "field:Location is missing tag:gorm"
		X float32 // want "field:X is missing tag:json" "field:X is missing tag:gorm"
		Y float32 // want "field:Y is missing tag:json" "field:Y is missing tag:gorm"
	}
}
