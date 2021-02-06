package tagalyzer

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Tags allow passing multiple values to the same command line flag
type Tags []string

func (t *Tags) String() string {
	return strings.Join(*t, ",")
}

func (t *Tags) includes(newtag string) bool {
	for _, tag := range *t {
		if tag == newtag {
			return true
		}
	}

	return false
}

func (t *Tags) Set(newtag string) error {
	if !t.includes(newtag) {
		*t = append(*t, newtag)
	}

	return nil
}

const doc = `check for missing tag(s) in your struct fields

For example, when working with a large codebase, one could forget to add the json tag
to one or more struct fields. This could result in a json response with key "Name" when
the client expects key "name".

tagalyzer takes a list of tags to check and reports all struct fields that are
missing those tags.
`

var (
	Analyzer = &analysis.Analyzer{
		Name: "tagalyzer",
		Doc:  doc,
		Run:  run,
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
	}

	tags          Tags
	checkEmbedded bool
)

func init() {
	Analyzer.Flags.Var(&tags, "tag", "a list of struct tags to check, for example: -tag json -tag gorm ...")
	Analyzer.Flags.BoolVar(&checkEmbedded, "checkembedded", false, "include embedded fields in tag check")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// we only care about struct nodes
	filterStructs := []ast.Node{
		(*ast.StructType)(nil),
	}

	inspect.Preorder(filterStructs, func(n ast.Node) {
		s := n.(*ast.StructType)

		for _, f := range s.Fields.List {
			for _, tag := range tags {
				if f.Tag == nil || !strings.Contains(f.Tag.Value, tag) {
					var fieldName string

					// check for embedded fields
					if f.Names == nil || len(f.Names) == 0 {
						if !checkEmbedded {
							continue
						}

						fieldName = fmt.Sprintf("%+v", f.Type)
					} else {
						fieldName = f.Names[0].Name
					}

					pass.Reportf(f.Pos(), fmt.Sprintf("field:%v is missing tag:%v", fieldName, tag))
				}
			}
		}
	})

	return nil, nil
}
