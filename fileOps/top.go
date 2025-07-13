package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseTop(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for i := 0; i < 7; i++ {
		s.Scan() // drop header
	}
	for s.Scan() {
		line := s.Text()
		line = strings.TrimSpace(line)
		o := strings.Fields(line)
		fmt.Println(o)
	}
}
