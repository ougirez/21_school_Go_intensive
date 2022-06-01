package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func findAll(path string, f, d, l bool, extension string) {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range files {
		if files[i].IsDir() {
			if d {
				fmt.Println(path + files[i].Name())
			}
			findAll(path+files[i].Name(), f, d, l, extension)
		} else {
			link, err := os.Readlink(path + files[i].Name())
			if err == nil {
				_, errLink := os.Open(path + link)
				if l && errLink != nil {
					fmt.Printf("%s -> %s\n", path+files[i].Name(), "[broken]")
				} else if l {
					fmt.Printf("%s -> %s\n", path+files[i].Name(), link)
				}
			} else if f && strings.HasSuffix(files[i].Name(), extension) {
				fmt.Println(path + files[i].Name())
			}
		}
	}
}

func main() {
	filesPtr := flag.Bool("f", false, "files")
	dirsPtr := flag.Bool("d", false, "directories")
	linksPtr := flag.Bool("sl", false, "symbol links")
	extPtr := flag.String("ext", "", "files extension")

	if len(os.Args) < 2 {
		log.Fatal("You haven't specify the path")
	}
	if _, err := os.Open(os.Args[len(os.Args)-1]); err != nil {
		log.Fatal("There is no such file or directory")
	}
	flag.Parse()
	if !*filesPtr && !*dirsPtr && !*linksPtr {
		findAll(os.Args[1], true, true, true, *extPtr)
	}

	findAll(os.Args[len(os.Args)-1], *filesPtr, *dirsPtr, *linksPtr, *extPtr)
}
