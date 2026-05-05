func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}

	m1 := make(map[byte]byte) 
	m2 := make(map[byte]byte) 

	for i := 0; i < len(s); i++ {
		c1 := s[i]
		c2 := t[i]

		if val, ok := m1[c1]; ok {
			if val != c2 {
				return false
			}
		} else {
			m1[c1] = c2
		}

		if val, ok := m2[c2]; ok {
			if val != c1 {
				return false
			}
		} else {
			m2[c2] = c1
		}
	}

	return true
}