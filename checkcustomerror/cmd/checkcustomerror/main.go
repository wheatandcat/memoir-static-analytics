package main

import (
	"checkcustomerror"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(checkcustomerror.Analyzer) }
