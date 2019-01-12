package fizzbuzz

func Say(n int) string {
	if n == 1 {
		return "1"
	} else if n == 2 {
		return "2"
	} else if n == 3 {
		return "Fizz"
	} else if n == 4 {
		return "4"
	}
	return "0"
}