package sort

type Triangle []int

func New(l []int) Triangle {
	if l == nil {
		return nil
	}
	if len(l) < 3 {
		return Triangle(l)
	}
	flipped := false
	curr := 0
	for i := 1; i < len(l); i++ {
		if (l[curr] >= l[i]) != flipped {
			if flipped {
				return nil
			}
			flipped = true
		}
		curr = i
	}
	return Triangle(l)
}

func (t Triangle) Min() *int {
	if len(t) == 0 {
		return nil
	}
	l := len(t) - 1
	if l == 0 {
		return &t[0]
	}
	if t[0] <= t[l] {
		return &t[0]
	}
	return &t[l]
}

func (t Triangle) Max() *int {
	if len(t) == 0 {
		return nil
	}
	l := len(t) - 1
	if l == 0 {
		return &t[0]
	}
	for i := 0; i < l; i++ {
		if t[i] >= t[i+1] {
			return &t[i]
		}
	}
	return &t[l]
}

func (t Triangle) Sort() []int {
	l := len(t)
	s := make(Triangle, l)
	copy(s, t)
	curr := 0
	for curr != l {
		last := t[l-1]
		if t[curr] > last {
			t = append(t[:curr], append(Triangle{last}, t[curr:l-1]...)...)
		} else {
			curr++
		}
	}
	return t

}
