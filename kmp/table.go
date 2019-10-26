package kmp

type Table struct {
	Content []int
}

func (t *Table) Build(pattern []byte)  {
	var i, j int
	var lenPattern = len(pattern)

	i = 1
	t.Content = make([]int, lenPattern + 1)
	t.Content[0] = -1
	for i < lenPattern {
		if pattern[i] == pattern[j] {
			t.Content[i] = t.Content[j]
		} else {
			t.Content[i] = j
			j = t.Content[j]

			for j >= 0 && pattern[i] != pattern[j] {
				j = t.Content[j]
			}
		}
		i += 1
		j += 1
	}

	t.Content[i] = j
}