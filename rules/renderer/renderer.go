package renderer

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"time"

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

const (
	pageTpl = `
<!DOCTYPE html>
<html lang=ja>
  <meta charset=utf-8>
  <meta content='width=680' name=viewport>
  <meta content=none name=robots>
  <meta content=never name=referrer>
  <title>{{if .Title}}{{.Title}} - {{end}}draftcode.osak.jp</title>
  <link href=/style.css rel=stylesheet>
  <article>
    <header>
      {{if .Time}}<time datetime={{.Time}}>{{.Time}}</time>{{end}}
      {{if .Title}}<h1>{{.Title}}</h1>{{end}}
    </header>
    <div>
      {{.Body}}
    </div>
  </article>
{{if .Title}}<div class=back><a href=/>もどりたい</a></div>{{end}}`
)

type templateInput struct {
	Title string
	Time  string
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
	if date, ok := fm["date"]; ok {
		t, err := time.Parse(time.RFC3339, date.(string))
		if err != nil {
			return err
		}
		in.Time = t.Format("2006-01-02")
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
