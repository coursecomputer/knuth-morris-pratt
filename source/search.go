package kmp

func Search(buffer, pattern []byte) (position int) {
	var table = Table{}
	var idxBuffer, idxPattern int
	var lenBuffer = len(buffer)
	var lenPattern = len(pattern)

	position = -1
	table.Build(pattern)
	for idxBuffer < lenBuffer {
		if buffer[idxBuffer] == pattern[idxPattern] {
			idxBuffer += 1
			idxPattern += 1

			if idxPattern == lenPattern {
				position = idxBuffer - idxPattern
				break
			}
		} else {
			idxPattern = table.Content[idxPattern]

			if idxPattern < 0 {
				idxBuffer += 1
				idxPattern += 1
			}
		}
	}

	return position
}
