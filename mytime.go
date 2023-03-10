package mytime

import (
	"time"

	"github.com/beevik/ntp"
)

func GetTime(host string) (string, error) {
	var mytime string
	if host == "" {
		host = "0.beevik-ntp.pool.ntp.org"
	}
	options := ntp.QueryOptions{
		Timeout: 30 * time.Second,
		TTL:     10,
	}

	response, err := ntp.QueryWithOptions(host, options)
	if err != nil {
		return mytime, err

	}
	err = response.Validate()
	if err != nil {
		return mytime, err
	}
	mytime = time.Now().Add(response.ClockOffset).Format("15:04:05")

	return mytime, err
}
