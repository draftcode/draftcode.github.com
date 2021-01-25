package main

import (
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

const tpl = `
<!DOCTYPE html>
<html lang=ja>
<meta charset=utf-8>
<meta content='width=680' name=viewport>
<meta content=none name=robots>
<meta content=never name=referrer>
<title>{{if .Title}}{{.Title}} - {{end}}draftcode.osak.jp</title>
<link href=/style.css rel=stylesheet>
<body class=page>
<div class=contents>
{{if .Title}}<h1>{{.Title}}</h1>{{end}}
{{.Body}}
</div>
<div class=back><a href=/>もどりたい</a></div>
</html>`

type templateInput struct {
	Title string
	Body  template.HTML
}

var (
	inputPath  = flag.String("input_file", "", "Input file path")
	outputPath = flag.String("output_file", "", "Output file path")

	tmpl = template.Must(template.New("webpage").Parse(tpl))
)

func main() {
	flag.Parse()
	fm, body, err := parseDocument(*inputPath)
	if err != nil {
		log.Fatal(err)
	}

	of, err := os.Create(*outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer of.Close()

	in := &templateInput{
		Body: template.HTML(body),
	}
	if t, ok := fm["title"]; ok {
		in.Title = t.(string)
	}
	if err := tmpl.Execute(of, in); err != nil {
		log.Fatal(err)
	}
}

func parseDocument(fp string) (map[string]interface{}, string, error) {
	bs, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, "", err
	}
	md := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
			extension.GFM,
			extension.DefinitionList,
			highlighting.Highlighting,
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := md.Convert(bs, &buf, parser.WithContext(context)); err != nil {
		return nil, "", err
	}
	return meta.Get(context), buf.String(), nil
}
