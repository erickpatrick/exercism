package clock

import "fmt"

const (
	hoursInADay    = 24
	minutesInAHour = 60
)

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	hours := (0 + h + (m / minutesInAHour)) % hoursInADay
	minutes := 0 + m%minutesInAHour

	if hours < 0 {
		hours = hoursInADay + hours
	}

	// to make counting easier, we midnight to 24h
	if hours == 0 {
		hours = hoursInADay
	}

	if minutes < 0 {
		// if negative minutes we need go hour backwards
		hours -= 1
		minutes += minutesInAHour
	}

	// reset to 0 when reaching 24h
	if hours == hoursInADay {
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
