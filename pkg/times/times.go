package times

import "time"

func TimetToInt64(tTime time.Time) (nTime int64) {
	return time.Date(tTime.Year(), tTime.Month(), tTime.Day(), tTime.Hour(),
		tTime.Minute(), tTime.Second(), 0, tTime.Location()).Unix()
}
