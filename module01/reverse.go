package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var ret string

	for _, w := range word {
		ret = string(w) + ret
	}

	return ret
}
