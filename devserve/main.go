package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

var (
	addr      = flag.String("addr", "0.0.0.0:30000", "listen address")
	rootDir   = flag.String("root", "", "root source directory")
	buildOnly = flag.Bool("build", false, "build only")
)

func main() {
	flag.Parse()
	if *rootDir == "" {
		*rootDir = findRoot()
	}

	m, err := BuildAll(*rootDir)
	if err != nil {
		log.Fatal(err)
	}

	if *buildOnly {
		buildPath := filepath.Join(filepath.Dir(*rootDir), "docs")
		if err := output(buildPath, m); err != nil {
			log.Fatal(err)
		}
		return
	}

	mu := sync.Mutex{}
	ch, err := Watch(*rootDir)
	if err != nil {
		log.Fatal("Cannot watch the root dir: " + err.Error())
	}
	go func() {
		for range ch {
			nextM, err := BuildAll(*rootDir)
			if err != nil {
				log.Println("error:", err)
			} else {
				mu.Lock()
				m = nextM
				mu.Unlock()
			}
		}
	}()

	http.HandleFunc("/_/manifest", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `<style>body { font-family: monospace; }</style>`)
		l := []string{}
		for k, _ := range m {
			l = append(l, k)
		}
		sort.Strings(l)
		for _, k := range l {
			io.WriteString(w, fmt.Sprintf(`<a href="%s">%s</a><br>`, k, k))
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			path = "/index.html"
		}
		if b, ok := m[path]; ok {
			ct := mime.TypeByExtension(filepath.Ext(path))
			if ct == "" {
				ct = http.DetectContentType(b)
			}
			if ct != "" {
				w.Header().Add("Content-Type", ct)
			}
			w.Write(b)
		} else {
			http.NotFound(w, r)
		}
	})
	log.Printf("Serving from " + *addr)
	log.Printf("Manifest: /_/manifest")
	http.ListenAndServe(*addr, nil)
}

// findRoot finds the source root directory from working directory.
func findRoot() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("finding the source root: %s", err)
	}
	origPwd := pwd
	for pwd != "/" {
		src := filepath.Join(pwd, "source")
		if _, err := os.Stat(src); err == nil {
			return src
		}
		pwd = filepath.Dir(pwd)
	}
	log.Fatalf("finding the source root, but not found: %s", origPwd)
	return ""
}

func output(buildPath string, m map[string][]byte) error {
	if _, err := os.Stat(buildPath); err == nil {
		if err := os.RemoveAll(buildPath); err != nil {
			return err
		}
	}
	for k, b := range m {
		p := filepath.Join(buildPath, k)
		pp := filepath.Dir(p)
		if _, err := os.Stat(pp); os.IsNotExist(err) {
			if err := os.MkdirAll(pp, 0755); err != nil {
				return err
			}
		}

		if err := ioutil.WriteFile(p, b, 0600); err != nil {
			return err
		}
	}
	return nil
}
