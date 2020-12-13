package dateutils

import "time"

var dateFormat = "2006-01-02T15:04:05Z"

func GetNowAsString() string {

	now := time.Now()
	return now.Format(dateFormat)

}
