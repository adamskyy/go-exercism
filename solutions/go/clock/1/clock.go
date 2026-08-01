package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
    minutesFromHours := (h * 60) 
    totalMinutes := (minutesFromHours + m) % 1440
    if totalMinutes == 0 {
        h = 0
        m = 0
    } else if totalMinutes > 0 {
        h = totalMinutes / 60
        m = totalMinutes % 60
    } else if totalMinutes < 0 {
        if totalMinutes % 60 == 0 {
            h = 24 + totalMinutes / 60
            m = totalMinutes % 60
        } else {
            h = 24 + totalMinutes / 60 + (-1)
        	m = 60 + totalMinutes % 60
        }
    }
    return Clock{hour: h, minute: m}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute + m)
    
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
