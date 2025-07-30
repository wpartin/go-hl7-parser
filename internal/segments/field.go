package segments

func GetField(fields []string, index int) string {
	if index < len(fields) {
		return fields[index]
	}

	return ""
}
