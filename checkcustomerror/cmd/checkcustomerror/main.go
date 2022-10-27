package main

import (
	"github.com/wheatandcat/memoir-static-analytics/checkcustomerror"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(checkcustomerror.Analyzer) }
