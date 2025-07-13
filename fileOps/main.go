package main

import "fmt"

func main() {
	fmt.Println("Parse CSV")
	ParseCSV("./sample_data.csv")
	fmt.Println()
	fmt.Println("Parse space separated file")
	ParseTop("./top_output.txt")
	fmt.Println()
	fmt.Println("Parse JSON")
	ParseJson("./sample_data.json")
	fmt.Println()
	fmt.Println("Parse KV")
	fmt.Println()
	ParseKV("./kv.txt")
	fmt.Println()
}
