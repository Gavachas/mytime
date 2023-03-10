package mytime

import (
	"testing"
)

func TestGetTime(t *testing.T) {
	mytime, err := GetTime("0.beevik-ntp.pool.ntp.org")
	if err != nil || mytime == "" {
		t.Fatal("error:", err, "time:", mytime)
	}
	mytime, err = GetTime("0.ru.pool.ntp.org")
	if err != nil || mytime == "" {
		t.Fatal("error:", err, "time:", mytime)
	}
	mytime, err2 := GetTime("WRONG.time.apple.com")
	if err2 == nil || mytime != "" {
		t.Fatal("error:", err, "time:", mytime)
	}
}
