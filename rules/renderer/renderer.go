package renderer

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var (
	pageTmpl = template.Must(template.New("webpage").Parse(pageTpl))
)

const pageTpl = `
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
{{if .Title}}<div class=back><a href=/>もどりたい</a></div>{{end}}
</html>`

type templateInput struct {
	Title string
	Body  template.HTML
}

// RenderPage renders a page.
func RenderPage(fp string, fm map[string]interface{}, body string) error {
	of, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer of.Close()

	in := &templateInput{
		Body: template.HTML(body),
	}
	if t, ok := fm["title"]; ok {
		in.Title = t.(string)
	}
	if err := pageTmpl.Execute(of, in); err != nil {
		return err
	}
	return nil
}

// ParseDocument takes a file path and parses it.
func ParseDocument(fp string) (map[string]interface{}, string, error) {
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
