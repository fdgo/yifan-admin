package timex

import "time"

func TimeCurrInt64() (nCurrTimeStamp int64) {
	return time.Now().UTC().Unix()
}

func TimeInt64ToTimeString(nTimeStamp int64) (strTime string) {
	tm := time.Unix(nTimeStamp, 0).UTC()
	return tm.Format("2006-01-02 15:04:05")
}
func TimeBeforeOrLater(curr_time int64, tag string, year_diff int, month_diff time.Month, day_diff, hour_diff, min_diff, sec_diff int) (target_time int64) {
	tTime := time.Unix(curr_time, 0).UTC()
	if tag == "before" {
		return time.Date(tTime.Year()-year_diff, tTime.Month()-month_diff, tTime.Day()-day_diff, tTime.Hour()-hour_diff,
			tTime.Minute()-min_diff, tTime.Second()-sec_diff, 0, tTime.Location()).UTC().Unix()
	} else if tag == "after" {
		return time.Date(tTime.Year()+year_diff, tTime.Month()+month_diff, tTime.Day()+day_diff, tTime.Hour()+hour_diff,
			tTime.Minute()+min_diff, tTime.Second()+sec_diff, 0, tTime.Location()).UTC().Unix()
	} else {
		return 0
	}
}
func TimetToInt64(tTime time.Time) (nTime int64) {
	return time.Date(tTime.Year(), tTime.Month(), tTime.Day(), tTime.Hour(),
		tTime.Minute(), tTime.Second(), 0, tTime.Location()).UTC().Unix()
}
