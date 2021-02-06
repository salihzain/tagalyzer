# tagalyzer

Static analyzer to find missing tags in your Golang structs

This Analyzer follows the Golang standardized way of creating static analyzers based on the package (https://pkg.go.dev/golang.org/x/tools/go/analysis) and is fully compatible with the command `go vet`

## Usage

First install the tool on your machine

```
$ go get github.com/salihzain/tagalyzer/cmd/tagalyzer
```

Then run inside your project

```
$ tagalyzer -tag=json -tag=gorm  ./...
```

To include embedded fields

```
$ tagalyzer -checkembedded -tag=json -tag=gorm  ./...
```

**Notes**:

You may use any number of tags

`./...` is a spread operator to pass all the packages in the current directory as command line arguments.

You may also run tagalyzer using `go vet` as follows

```
$ go vet -vettool=$(which tagalyzer) -tag json -tag gorm ./...
```

## Example

Suppose we have the following struct in our project:

```go
type User struct {
 Username string `gorm:"primaryKey"`
 Name string `json:"name"`
 Age int
}
```

then running `tagalyzer -tag json -tag gorm ./...` will output the following:

```
path/to/file.go:25:2: field:Username is missing tag:json
path/to/file.go:26:2: field:Name is missing tag:gorm
path/to/file.go:27:2: field:Age is missing tag:json
path/to/file.go:27:2: field:Age is missing tag:gorm
```
