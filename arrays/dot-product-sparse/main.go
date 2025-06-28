package main

type Sparse struct{}

type VecTuple struct {
	Idx int
	Val int
}

func (s *Sparse) DotProduct(arr1, arr2 []int) int {
	v1 := VecTupleFromSlcInt(arr1)
	v2 := VecTupleFromSlcInt(arr2)
	prod := 0

	for i, j := 0, 0; i < len(v1) && j < len(v2); {
		if v1[i].Idx < v2[j].Idx {
			i++
		} else if v1[i].Idx < v2[j].Idx {
			j++
		} else { // i == j
			prod += v1[i].Val * v2[j].Val
			i++
			j++
		}
	}

	return prod
}

func VecTupleFromSlcInt(arr []int) []VecTuple {
	v1 := []VecTuple{}
	for i, e := range arr {
		if e != 0 {
			v1 = append(v1, VecTuple{
				Idx: i,
				Val: e,
			})
		}
	}

	return v1
}
