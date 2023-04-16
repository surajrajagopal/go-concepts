package arrslicevardic

func SliceReferenceType(num []int) {
	//sum := 0
	for k, v := range num {
		num[k] = v + 1
	}
	return
}

func AppendConcepts(s []int) {
	num := []int{1, 10, 16}

	//append to slice
	s1 := []int{}
	s1 = append(s1, 1, 2, 3, 4)

	//It is also possible to append one slice to another using the ... operator.
	s1 = append(s1, num...)
}
