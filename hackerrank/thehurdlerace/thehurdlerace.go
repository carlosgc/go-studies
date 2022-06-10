// Package thehurdlerace provides a solution implementation for The Hurdle Race problem.
//
// Statement: A video player plays a game in which the character competes in a hurdle race. Hurdles are of varying
// heights, and the characters have a maximum height they can jump. There is a magic potion they can take that will
// increase their maximum jump height by 1 unit for each dose. How many doses of the potion must the character take
// to be able to jump all of the hurdles. If the character can already clear all of the hurdles, return 0.
//
// A complete description of this problem can be found here: https://www.hackerrank.com/challenges/the-hurdle-race
package thehurdlerace

// HurdleRace returns the numbers of potion doses are needed to jump all of the hurdles.
func HurdleRace(k int32, height []int32) int32 {

	maxH := int32(0)
	n := len(height)

	// Finding the highest height.
	for idx := 0; idx < n; idx++ {

		h := height[idx]
		if h > maxH {
			maxH = h
		}
	}

	d := int32(0)
	// Finding the number of doses needed to jump all hurdles.
	if maxH > k {
		d = maxH - k
	}

	return d
}
