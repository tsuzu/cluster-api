package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"sigs.k8s.io/cluster-api/hack/tools/staticchecker/internal/analyzer"
)

func main() {
	singlechecker.Main(analyzer.GoDocPrefix)
}
