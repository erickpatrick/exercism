package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	hours := (0 + h + (m / 60)) % 24
	minutes := 0 + m%60

	if hours < 0 {
		hours = 24 + hours
	}

	if hours == 0 {
		hours = 24
	}

	if minutes < 0 {
		hours -= 1
		minutes += 60
	}

	if hours == 24 {
		hours = 0
	}

	return Clock{
		h: hours,
		m: minutes,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
