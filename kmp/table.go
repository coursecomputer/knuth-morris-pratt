package kmp

type Table struct {
	Content []int
}

func (t *Table) Build(pattern []byte)  {
	var idxPattern, lenMatch int
	var lenPattern = len(pattern)

	idxPattern = 1
	t.Content = make([]int, lenPattern + 1)
	t.Content[0] = -1
	for idxPattern < lenPattern {
		if pattern[idxPattern] == pattern[lenMatch] {
			t.Content[idxPattern] = t.Content[lenMatch]
		} else {
			t.Content[idxPattern] = lenMatch
			lenMatch = t.Content[lenMatch]

			for lenMatch >= 0 && pattern[idxPattern] != pattern[lenMatch] {
				lenMatch = t.Content[lenMatch]
			}
		}
		idxPattern += 1
		lenMatch += 1
	}

	t.Content[idxPattern] = lenMatch
}