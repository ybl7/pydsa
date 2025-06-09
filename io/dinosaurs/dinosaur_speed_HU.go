package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// {
	// 	"dinoName": {
	// 		"pedal": "string",
	// 		"lenFeet": float,
	// 		"lenStride": float,
	// 		"speed": float
	// 	}
	// }
	dinos := make(map[string]map[string]interface{})

	// Open first file
	f1, err := os.Open("data1.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	scanner := bufio.NewScanner(f1)
	scanner.Scan() // First line is just headers

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		name := parts[0]
		pedal := parts[1]
		lenFeet, err := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
		if err != nil {
			log.Fatal(err)
		}

		dinos[name] = map[string]interface{}{
			"pedal":   pedal,
			"lenFeet": lenFeet,
		}
	}

	// Open second file
	os.Open("data2.csv")

	// Open first file
	f2, err := os.Open("data2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	scanner2 := bufio.NewScanner(f2)
	scanner2.Scan() // First line is just headers

	g := 9.81
	for scanner2.Scan() {
		parts := strings.Split(scanner2.Text(), ",")
		name := parts[0]
		lenStride, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		if err != nil {
			log.Fatal(err)
		}

		dinos[name]["lenStride"] = lenStride
		// Need to cast the interfaces into concrete types to operate on
		dinos[name]["speed"] = (dinos[name]["lenFeet"].(float64) * dinos[name]["lenStride"].(float64) / 2.0) * g
	}

	var bipedals []string
	for dino, dinoMap := range dinos {
		if strings.ToLower(dinoMap["pedal"].(string)) == "bipedal" {
			bipedals = append(bipedals, dino)
		}
	}
	// This says we will sort slice bipedals, an when we get to element i and j in the slice, if the function returns false, then i is less than j
	sort.Slice(bipedals, func(i, j int) bool {
		return dinos[bipedals[i]]["speed"].(float64) < dinos[bipedals[j]]["speed"].(float64)
	})

	fmt.Println("sorted bipedals by speed ascending: ", bipedals)
}

// Have to admit you learn a lot about casting in Go when you do the probelm in Go, the python solution is half as long, owwing to the Go error handlinng, but also the file reading and scanning being more controlled I think
// Took about 20 mins after already having seen the AI solution and with generous use of google to find the right packages and methods
