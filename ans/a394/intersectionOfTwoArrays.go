package a394

func Intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, num := range nums1 {
		if _, ok := m[num]; !ok {
			m[num] = true
		}
	}

	m2 := make(map[int]bool)
	for _, num2 := range nums2 {
		if _, ok := m[num2]; ok {
			if _, ok2 := m2[num2]; !ok2 {
				m2[num2] = true
			}
		}
	}

	keys := make([]int, 0, len(m2))
	for k := range m2 {
		keys = append(keys, k)
	}
	return keys
}
