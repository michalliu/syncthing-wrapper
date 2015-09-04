package demo

import "sort"

func ArrayRemove() {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	//               ^     ^     ^
	// indexes to remove
	removal := []int{3, 1, 5}
	result := arrayRemoveInt(source, removal)
	for _, x := range result {
		print(x, ",")
	}
	print("\n")
}

func typeAssertion() {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	//               ^     ^     ^
	// indexes to remove
	removal := []int{3, 1, 5}
	result := arrayRemove(source, removal)
	// type assertion
	// MUST change to another varname
	result2, ok := result.([]int)
	if ok {
		print(len(result2))
		for _, x := range result2 {
			print(x, ",")
		}
		print("\n")
	}
}

func genericType() {
	var result interface{}
	result = 1
	switch v := result.(type) {
	case []int:
		print("[]int")
	case int:
		print("int")
	default:
		print(v)
	}
	print("\n")
}

func arrayRemove(source interface{}, removal []int) interface{} {
	//GENERIC NOT WORKING...
	return make([]int, 0)
}

// switch type
func typeSwitch() {
	var t interface{}
	t = 2.0
	switch ins := t.(type) {
	case float64:
		print("float64")
	default:
		print(ins)
	}
}

// remove element from array by index
func arrayRemoveInt(source []int, removal []int) []int {
	// result
	var result []int

	// sort the removal indexes
	sort.Sort(sort.IntSlice(removal))

	// loop removal
	for i, v := range removal {
		// once at a time
		v = v - i
		//print("v=", v, "\n")
		// remove on element from array
		// noted once you remove one element
		// since the array.length can not change
		// go will add one element to the end of the array
		result = append(source[:v], source[v+1:]...)
		//for _, x := range result {
		//	print(x, ",")
		//}
		//print("\n")
	}
	// we know how musch data we need, remove the extra data
	result = result[:len(source)-len(removal)]
	return result
}

func ArrayMap() {
	source := make(map[int]string)
	source[2] = ""
	source[1] = ""
	source[10] = ""
	idx := make([]int, 0)
	for k := range source {
		idx = append(idx, k)
	}
	for _, x := range idx {
		print(x, ",")
	}
	print("\n")
}
