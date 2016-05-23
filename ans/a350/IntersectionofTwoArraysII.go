package a350

func Intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	for _, num := range nums1 {
		if _, ok := m1[num]; !ok {
			m1[num] = 1
		} else {
			m1[num] += 1
		}
	}

	m2 := make(map[int]int)
	for _, num := range nums2 {
		if _, ok := m2[num]; !ok {
			m2[num] = 1
		} else {
			m2[num] += 1
		}
	}

	if len(m2) < len(m1) {
		tmp := m2
		m2 = m1
		m1 = tmp
	}

	ret := []int{}

	for k := range m1 {
		if _, ok := m2[k]; ok {
			count := 0
			if m2[k] > m1[k] {
				count = m1[k]
			} else {
				count = m2[k]
			}
			for i := 0; i < count; i++ {
				ret = append(ret, k)
			}
		}
	}
	return ret
}
