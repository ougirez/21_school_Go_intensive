package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
)



func main() {
	var fn1, fn2 string

	if len(os.Args) != 5 {
		log.Fatal("Invalid number of arguments")
	}
	for i := 1; i < len(os.Args); i += 2 {
		if os.Args[i] == "--old" {
			fn1 = os.Args[i+1]
		} else if os.Args[i] == "--new" {
			fn2 = os.Args[i+1]
		} else {
			log.Fatal("Use these flags: --old, --new")
		}
	}
	f1, err := os.Open(fn1)
	if err != nil {
        log.Fatal(err)
    }
	var FS1 = make(map[string]bool)
	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		FS1[scanner.Text()] = true
	}
    f1.Close()

	f2, err := os.Open(fn2)
	if err != nil {
        log.Fatal(err)
    }
	scanner = bufio.NewScanner(f2)
	for scanner.Scan() {
		if _, ok := FS1[scanner.Text()]; ok == true {
			delete(FS1, scanner.Text())
		} else {
			fmt.Printf("ADDED %s\n", scanner.Text())
		}
	}
	for key := range FS1 {
		fmt.Printf("REMOVED %s\n", key)
	}
    f2.Close()
}