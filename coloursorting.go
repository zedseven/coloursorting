package coloursorting

import (
	"math"
)

func rgbToHSV(r, g, b float64) (h, s, v float64) {
	var min float64
	var max float64

	if r < g {
		min = r
	} else {
		min = g
	}
	if min > b {
		min = b
	}

	if r > g {
		max = r
	} else {
		max = g
	}
	if max < b {
		max = b
	}

	v = max
	delta := max - min
	if delta < 0.00001 {
		s = 0
		h = math.NaN()
		return
	}

	if max > 0.0 {
		s = delta / max
	} else {
		s = 0
		h = math.NaN()
		return
	}

	if r >= max {
		h = (g - b) / delta
	} else if g >= max {
		h = 2.0 + (b - r) / delta
	} else {
		h = 4.0 + (r - g) / delta
	}

	h *= 60.0 // in degrees

	if h < 0.0 {
		h += 360.0
	}

	return
}

// StepSort implements a roughly hue-based colour sorting algorithm based on the
// "Step Sorting" section from https://www.alanzucconi.com/2015/09/30/colour-sorting/.
type StepSort [][3]int

// Len simply retrieves the length of the slice.
func (a StepSort) Len() int { return len(a) }
// Swap swaps the elements at i and j.
func (a StepSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
// Less determines whether the element at i is less than the element at j.
func (a StepSort) Less(i, j int) bool {
	is0, is1, is2 := step(a[i][0], a[i][1], a[i][2], 8)
	js0, js1, js2 := step(a[j][0], a[j][1], a[j][2], 8)

	if is0 < js0 {
		return true
	} else if is0 == js0 {
		if is1 < js1 {
			return true
		} else if is1 == js1 {
			return is2 < js2
		}
	}

	return false
}

func step(ir, ig, ib, repetitions int) (int, int, int) {
	r := float64(ir)
	g := float64(ig)
	b := float64(ib)

	lum := math.Sqrt(0.241 * r + 0.691 * g + 0.068 * b)

	h, _, v := rgbToHSV(r, g, b)

	h2 := int(h * float64(repetitions))
	lum2 := int(lum * float64(repetitions))
	v2 := int(v * float64(repetitions))

	return h2, lum2, v2
}