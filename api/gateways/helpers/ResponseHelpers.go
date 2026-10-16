package helpers

func NormalizeResponse(input string) string {
	if input != "" {
		if input == "0" || input == "85" {
			return "00"
		}
	}
	return input
}
