// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"sort"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	// Using interface{} to store both strings and floats
// 	dinosaurs := make(map[string]map[string]interface{})

// 	// Read data1.csv
// 	file1, err := os.Open("data1.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file1.Close()

// 	scanner := bufio.NewScanner(file1)
// 	scanner.Scan() // Skip header
// 	for scanner.Scan() {
// 		parts := strings.Split(scanner.Text(), ",")
// 		name := parts[0]
// 		pedalType := parts[1]
// 		lenLeg, _ := strconv.ParseFloat(parts[2], 64)

// 		dinosaurs[name] = make(map[string]interface{})
// 		dinosaurs[name]["pedal"] = pedalType
// 		dinosaurs[name]["lenLeg"] = lenLeg
// 	}

// 	// Read data2.csv
// 	file2, err := os.Open("data2.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file2.Close()

// 	scanner = bufio.NewScanner(file2)
// 	scanner.Scan() // Skip header
// 	for scanner.Scan() {
// 		parts := strings.Split(scanner.Text(), ",")
// 		name := parts[0]
// 		lenStride, _ := strconv.ParseFloat(parts[1], 64)

// 		if dinosaurs[name] == nil {
// 			dinosaurs[name] = make(map[string]interface{})
// 		}
// 		dinosaurs[name]["lenStride"] = lenStride
// 	}

// 	// Calculate speeds
// 	g := 9.81
// 	var bipedal []string

// 	for name, data := range dinosaurs {
// 		if strings.ToLower(data["pedal"].(string)) == "bipedal" {
// 			feet := data["lenLeg"].(float64)
// 			stride := data["lenStride"].(float64)
// 			speed := ((feet * stride) / 2.0) * g
// 			data["speed"] = speed
// 			bipedal = append(bipedal, name)
// 		}
// 	}

// 	// Sort bipedal dinosaurs by speed
// 	sort.Slice(bipedal, func(i, j int) bool {
// 		return dinosaurs[bipedal[i]]["speed"].(float64) < dinosaurs[bipedal[j]]["speed"].(float64)
// 	})

// 	// Print results
// 	fmt.Println("Dinosaurs in increasing order of speed:")
// 	for _, name := range bipedal {
// 		fmt.Printf("%s: %.2f\n", name, dinosaurs[name]["speed"].(float64))
// 	}
// }
