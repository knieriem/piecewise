package piecewise

type WalkFn func(i, dx int, x float64) error

func Walk(valRange, nSeg int, mirrorx bool, f WalkFn) error {
	seg := subdivideRange(valRange, nSeg, !mirrorx)

	for seg.next() {
		i, dx := seg.segment()
		ix := i + 1
		if mirrorx {
			ix = nSeg - ix
		}
		x := float64(ix) / float64(nSeg)
		err := f(i, dx, x)
		if err != nil {
			return err
		}
	}
	return nil
}

type segments struct {
	div      int
	extra    int
	i        int
	n        int
	growLeft bool
}

// subdivideRange subdivides an integer value range by n segments.
// In case the range can not be divided evenly, growLeft decides whether
// the remainder shall be distributed over segments on the left
// side (small values) or on the right (higher values).
func subdivideRange(valRange, n int, growLeft bool) *segments {
	extra := valRange % n
	div := valRange / n
	if div == 0 {
		n = extra
	}
	return &segments{
		extra:    extra,
		div:      div,
		n:        n,
		i:        -1,
		growLeft: growLeft,
	}
}

// next advances to the next segment. It returns true on sucess,
// false if it couldn't move to the next segment.
func (seg *segments) next() bool {
	seg.i++
	return seg.i < seg.n
}

// segment returns the index and the length (dx) of the
// current segment.
func (seg *segments) segment() (i, dx int) {
	dx = seg.div
	i = seg.i
	if seg.growLeft {
		if i < seg.extra {
			dx++
		}
	} else if seg.n-i-1 < seg.extra {
		dx++
	}
	return i, dx
}
