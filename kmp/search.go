package kmp

func Search(buffer, pattern []byte) (occurrence int) {
	var table = Table{}
	var idxBuffer, lenOccPattern int
	var lenBuffer = len(buffer)
	var lenPattern = len(pattern)

	occurrence = -1
	table.Build(pattern)
	for idxBuffer < lenBuffer {
		if buffer[idxBuffer] == pattern[lenOccPattern] {
			idxBuffer += 1
			lenOccPattern += 1

			if lenOccPattern == lenPattern {
				occurrence = idxBuffer - lenOccPattern
				break
			}
		} else {
			lenOccPattern = table.Content[lenOccPattern]

			if lenOccPattern < 0 {
				idxBuffer += 1
				lenOccPattern += 1
			}
		}
	}

	return occurrence
}