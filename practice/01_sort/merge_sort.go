package _1_sort

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	m := len(a) / 2
	first := MergeSort(a[:m])
	second := MergeSort(a[m:])
	return merge(first, second)
}

func merge(first []int, second []int) []int {
	var final []int
	i := 0
	j := 0
	for (i < len(first)) && (j < len(second)) {
		if first[i] < second[j] {
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}
	// Copy the rest from first
	for ; i < len(first); i++ {
		final = append(final, first[i])
	}
	// Copy the rest from second
	for ; j < len(second); j++ {
		final = append(final, second[j])
	}
	return final
}
