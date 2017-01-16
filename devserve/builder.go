package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func BuildAll(root string) (map[string][]byte, error) {
	m := map[string][]byte{}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		relPath = "/" + relPath
		switch filepath.Ext(relPath) {
		case ".html", ".png", ".jpg", ".css":
			bs, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			m[relPath] = bs
			return nil
		case ".scss":
			s, err := CompileSass(path)
			if err != nil {
				return err
			}
			
			relPath = strings.TrimSuffix(relPath, ".scss") + ".css"
			m[relPath] = []byte(s)
			return nil
		default:
			println(relPath)
			return nil
		}
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}
