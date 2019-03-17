package main

import "golang.org/x/time/rate"

func main() {
	rr := rate.Limit()

	var aa float64 = 1.0

	r := rate.NewLimiter(0, 3)
}
