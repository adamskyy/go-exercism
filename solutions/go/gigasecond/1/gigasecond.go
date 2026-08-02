// Package gigasecond is extremely useful for calculating time in gigaseconds
package gigasecond

import "time"

// AddGigasecond return a value a gigasecond later
func AddGigasecond(t time.Time) time.Time {
    return t.Add(time.Second * 1000 * 1000000)
}
