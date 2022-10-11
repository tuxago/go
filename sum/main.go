package main

func sum(b []int) int {
	s := 0
	for i := range b {
		s += b[i]
	}
	return s
}
