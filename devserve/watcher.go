package main

import (
	"gopkg.in/fsnotify.v1"
	"log"
	"os"
	"path/filepath"
)

func Watch(root string) (chan bool, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	if err := registerRecursively(w, root); err != nil {
		return nil, err
	}

	ch := make(chan bool)
	go func() {
		for {
			select {
			case event := <-w.Events:
				log.Println("event:", event)
				processEvent(w, event.Name, ch)
			case err := <-w.Errors:
				log.Println("error:", err)
			}
		}
	}()

	return ch, nil
}

func registerRecursively(w *fsnotify.Watcher, dir string) error {
	if err := w.Add(dir); err != nil {
		return err
	}

	f, err := os.Open(dir)
	if err != nil {
		return err
	}
	fis, err := f.Readdir(0)
	f.Close()
	if err != nil {
		return err
	}

	for _, fi := range fis {
		path := filepath.Join(dir, fi.Name())
		if fi.IsDir() {
			if err := registerRecursively(w, path); err != nil {
				return err
			}
		} else {
			if err := w.Add(path); err != nil {
				return err
			}
		}
	}
	return nil
}

func processEvent(w *fsnotify.Watcher, path string, ch chan bool) {
	fi, err := os.Stat(path)
	if err != nil {
		log.Println("error:", err)
	} else {
		if fi.IsDir() {
			registerRecursively(w, path)
		} else {
			w.Add(path)
		}
		ch <- true
	}
}
