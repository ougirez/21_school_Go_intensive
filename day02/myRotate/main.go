package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func archive(file string, dir string) {
	defer wg.Done()
	info, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	if !strings.HasSuffix(file, ".log") {
		fmt.Printf("%s: it's not a log file\n", file)
		return
	}
	var tarName string
	if dir == "" {
		tarName = strings.TrimSuffix(file, ".log")
	} else {
		tarName = dir + strings.TrimSuffix(info.Name(), ".log")
	}
	tarName = tarName + strconv.FormatInt(info.ModTime().Unix(), 10) + ".tar.gz"
	out, err := os.Create(tarName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	gw := gzip.NewWriter(out)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	header, err := tar.FileInfoHeader(info, info.Name())
	header.Name = file
	if err != nil {
		log.Fatal(err)
	}
	err = tw.WriteHeader(header)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(tw, f)
	if err != nil {
		log.Fatal(err)
	}
}

func archiveAll(files []string, dir string) {
	if dir != "" {
		if !strings.HasSuffix(dir, "/") {
			dir = dir + "/"
		}
		if strings.HasPrefix(dir, "/") {
			dir = strings.TrimPrefix(dir, "/")
		}
		_, err := os.Stat(dir)
		if err != nil {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}
	}

	wg.Add(len(files))
	for _, file := range files {
		go archive(file, dir)
	}
	wg.Wait()
}

func main() {
	dirPtr := flag.String("a", "", "directory for archived files")
	flag.Parse()

	if *dirPtr == "" {
		archiveAll(os.Args[1:], "")
	} else {
		archiveAll(os.Args[3:], *dirPtr)
	}
}
