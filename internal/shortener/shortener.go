package shortener

var (
	counter      = 1000
	allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-."
)

func setCounterForTests(value int) {
	counter = value
}

func ShortenURL(originalURL string) string {
	// Dummy implementation for URL shortening
	charPositions := []int{}
	dividend := counter
	for {
		if dividend < len(allowedChars) {
			charPositions = append(charPositions, dividend)
			break
		}
		remainder := dividend % len(allowedChars)
		charPositions = append(charPositions, remainder)
		dividend = dividend / len(allowedChars)
	}
	characters := ""
	for _, charPosition := range charPositions {
		characters += allowedChars[charPosition : charPosition+1]
	}
	counter++
	return characters
}
