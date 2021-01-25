package main

import (
	"flag"
	"log"

	"github.com/draftcode/draftcode.github.com/rules/renderer"
)

var (
	inputPath  = flag.String("input_file", "", "Input file path")
	outputPath = flag.String("output_file", "", "Output file path")
)

func main() {
	flag.Parse()
	fm, body, err := renderer.ParseDocument(*inputPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := renderer.RenderPage(*outputPath, fm, body); err != nil {
		log.Fatal(err)
	}
}
