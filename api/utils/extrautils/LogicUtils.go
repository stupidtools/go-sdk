package extrautils

func IfThenElse(logic bool, a string, b string) string {
	if logic {
		return a
	}
	return b
}
