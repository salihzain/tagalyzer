package testdata

type IgnoreField struct {
	Name string // want "field:Name is missing tag:json"
	Age  int    // want "field:Age is missing tag:json"
	ssn  string `tagalyzer:"-"`
}
