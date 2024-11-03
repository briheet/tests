package selectme

import (
	"net/http"
	"time"
)

func Racer(slowUrl, fastUrl string) string {
	durationFastUrl := measureResponseTime(fastUrl)
	durationSlowUrl := measureResponseTime(slowUrl)

	if durationFastUrl < durationSlowUrl {
		return fastUrl
	}

	return slowUrl
}

func measureResponseTime(a string) time.Duration {
	startTimeA := time.Now()
	http.Get(a)
	durationA := time.Since(startTimeA)

	return durationA
}
