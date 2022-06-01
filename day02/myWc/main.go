package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup
var countMap map[string]int

func count(file string, flag string) {
	defer wg.Done()

	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("%s: no such file or directory\n", file)
		return
	}
	scanner := bufio.NewScanner(f)
	switch flag {
	case "w":
		scanner.Split(bufio.ScanWords)
	case "l":
		scanner.Split(bufio.ScanLines)
	case "m":
		scanner.Split(bufio.ScanRunes)
	}
	c := 0
	for scanner.Scan() {
		c++
	}
	countMap[file] = c
}

func printWc(files []string) {
	for _, f := range files {
		fmt.Printf("%d\t%s\n", countMap[f], f)
	}
}

func goStart(files []string, flag string) {
	wg.Add(len(files))
	for i := range files {
		go count(files[i], flag)
	}
	wg.Wait()
	printWc(files)
}

func main() {
	wPtr := flag.Bool("w", false, "count words")
	lPtr := flag.Bool("l", false, "count lines")
	mPtr := flag.Bool("m", false, "count runes")
	flag.Parse()

	if len(os.Args) < 2 {
		log.Fatal("Invalid number of arguments")
	}
	countMap = make(map[string]int)
	if !*wPtr && !*lPtr && !*mPtr {
		goStart(os.Args[1:], "w")
	} else if *wPtr && !*lPtr && !*mPtr {
		goStart(os.Args[2:], "w")
	} else if *lPtr && !*wPtr && !*mPtr {
		goStart(os.Args[2:], "l")
	} else if *mPtr && !*wPtr && !*lPtr {
		goStart(os.Args[2:], "m")
	} else {
		log.Fatal("You can only use one flag -w -l or -m")
	}
}
