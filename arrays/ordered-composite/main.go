package main

import "fmt"

func main() {
	fmt.Println("Expected: true, Got:", IsOrderedComposite("cat", "dog", "catdog"))
	fmt.Println("Expected: true, Got:", IsOrderedComposite("cat", "dog", "cadtog"))
	fmt.Println("Expected: false, Got:", IsOrderedComposite("cat", "dog", "ctdaog"))
	fmt.Println("Expected: false, Got:", IsOrderedComposite("cat", "dog", "cdogta"))
	fmt.Println("Expected: true, Got:", IsOrderedComposite("frog", "dog", "frogdog"))
	fmt.Println("Expected: true, Got:", IsOrderedComposite("frog", "dog", "frdogog"))
	fmt.Println("Expected: false, Got:", IsOrderedComposite("frog", "dog", "fordgog"))
	fmt.Println("Expected: false, Got:", IsOrderedComposite("frog", "dog", "frdoggo"))
}

func IsOrderedComposite(a, b, c string) bool {
	i, j, k := 0, 0, 0

	for k < len(c) {
		if i < len(a) && j < len(b) && c[k] == a[i] && c[k] == b[j] {
			// This is for when we have a char that matches both strings but we don't know which string it should be assinged to
			// It's a recursive call that will handle the rest of the string for us so we return wichever succeeded, if any does
			optA := IsOrderedComposite(a[i+1:], b[j:], c[k+1:])
			optB := IsOrderedComposite(a[i:], b[j+1:], c[k+1:])
			return optA || optB
		} else if i < len(a) && c[k] == a[i] {
			i++
			k++
		} else if j < len(b) && c[k] == b[j] {
			j++
			k++
		} else {
			// This means that the char that we found (a) doesn't belong to a or b OR (b) it's in the wrong order
			return false
		}
	}

	return i == len(a) && j == len(b)
}

// The composite string c is an ordered composite of a and b if it contains all chars from a and b and the order of a and b are respected
// "cat" + "dog" -> "catdog" is ordered
// "cat" + "dog" -> "cadtog" is ordered
// "cat" + "dog" -> "ctdaog" is not ordered
// "cat" + "dog" -> "cdogta" is not ordered
// We are told that a and b are len(n) and that c is always len(2n)

// The idea is to iterate down both strings and the output string
