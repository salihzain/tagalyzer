package main

import (
	"github.com/salihzain/tagalyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(tagalyzer.Analyzer)
}
