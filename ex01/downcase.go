package downcase

func Downcase(str string) (string, error) {
	b := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		b[i] = c
	}
	return string(b), nil
}
