package main

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	s1 := []rune(s)
	p1 := []rune(p)
	firstMatch := (len(s1) != 0 && (p1[0] == s1[0] || p[0] == '.'))
	return firstMatch && isMatch(string(s1[1:]), string(p1[1:]))

}

func main() {

}
