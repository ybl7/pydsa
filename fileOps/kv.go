package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseKV(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		p := s.Text()
		p = strings.TrimSpace(p)
		a := strings.SplitN(p, "=", 2)
		fmt.Println(a)
	}
}
