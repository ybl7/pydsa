package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseCSV(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan() // first line of headers, drop it
	for s.Scan() {
		p := s.Text()
		a := strings.Split(p, ",")
		fmt.Println(a)
	}

}
