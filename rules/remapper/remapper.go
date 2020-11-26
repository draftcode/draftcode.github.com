package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	mappingFile = flag.String("mapping_file", "", "Mapping JSON file")
	outputDir   = flag.String("output_dir", "", "Output directory")
)

func main() {
	flag.Parse()
	configs, err := parseConfigs()
	if err != nil {
		log.Fatal(err)
	}

	written := map[string]bool{}
	for base, dstBase := range configs {
		err = filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			rel, err := filepath.Rel(base, path)
			if err != nil {
				return err
			}
			dst := filepath.Join(*outputDir, dstBase, rel)
			if written[dst] {
				return fmt.Errorf("%q is written by two entries", dst)
			}
			if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
				return err
			}
			return copyFile(dst, path)
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func copyFile(dst, src string) error {
	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()
	dstF, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer dstF.Close()

	_, err = io.Copy(dstF, srcF)
	return err

}

func parseConfigs() (map[string]string, error) {
	bs, err := ioutil.ReadFile(*mappingFile)
	if err != nil {
		return nil, err
	}
	configs := []*PackageMappingConfig{}
	if err := json.Unmarshal(bs, &configs); err != nil {
		return nil, err
	}
	m := map[string]string{}
	for _, pmc := range configs {
		for _, pm := range pmc.Mappings {
			p := filepath.Join(pmc.Package, pm.Src)
			if _, ok := m[p]; ok {
				return nil, fmt.Errorf("path %s has duplicate mappings", p)
			}
			m[p] = pm.Dst
		}
	}
	return m, nil
}

type PackageMappingConfig struct {
	Package  string         `json:"package"`
	Mappings []*PathMapping `json:"mappings"`
}

type PathMapping struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}
