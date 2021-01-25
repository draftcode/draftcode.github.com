package main

import (
	"flag"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/draftcode/draftcode.github.com/rules/renderer"
)

var (
	inputPaths = flag.String("input_files", "", "Input file path, comma separated")
	outputPath = flag.String("output_dir", "", "Output dir path")

	pageTmpl = template.Must(template.New("webpage").Parse(pageTpl))
)

const (
	timeFormat = "2006-01-02T15:04"
	pageTpl    = `
<!DOCTYPE html>
<html lang=ja>
  <meta charset=utf-8>
  <meta content='width=680' name=viewport>
  <meta content=none name=robots>
  <meta content=never name=referrer>
  <title>New contents - draftcode.osak.jp</title>
  <link href=/style.css rel=stylesheet>
  <ul>
    {{range .}}<li><a href='{{.Path}}'>{{.Title}}</a>
    {{end}}
  </ul>
  <div class=back><a href=/>もどりたい</a></div>
</html>`
)

type entry struct {
	Title string
	Time  time.Time
	Path  string
}

func main() {
	flag.Parse()
	var entries []*entry
	for _, ip := range strings.Split(*inputPaths, ",") {
		fm, body, err := renderer.ParseDocument(ip)
		if err != nil {
			log.Fatal(err)
		}
		base := strings.TrimSuffix(filepath.Base(ip), filepath.Ext(ip)) + ".html"
		op := filepath.Join(*outputPath, base)
		if err := renderer.RenderPage(op, fm, body); err != nil {
			log.Fatal(err)
		}
		t, err := time.Parse(timeFormat, fm["date"].(string))
		if err != nil {
			log.Fatal(err)
		}
		entries = append(entries, &entry{
			Title: t.Format("2006-01-02: ") + fm["title"].(string),
			Time:  t,
			Path:  base,
		})
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Time.After(entries[j].Time)
	})

	of, err := os.Create(filepath.Join(*outputPath, "index.html"))
	if err != nil {
		log.Fatal(err)
	}
	defer of.Close()
	if err := pageTmpl.Execute(of, entries); err != nil {
		log.Fatal(err)
	}

}
