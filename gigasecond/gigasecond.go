//Package gigasecond contains time function
package gigasecond

import "time"

// AddGigasecond takes a date and returns the addition of 1 Gigasecond to that date.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)

}
