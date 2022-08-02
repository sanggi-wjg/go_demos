package times_test

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestTime(t *testing.T) {
	now := time.Now()
	t.Log(now) // 2022-08-02 11:14:48.8866312 +0900 KST m=+0.002534901

	nowFormat := time.Now().Format("2006-01-02 15:04:05")
	t.Log(nowFormat) // 2022-08-02 11:17:33 (https://flaviocopes.com/go-date-time-format/?fbclid=IwAR0zPE9XTBDAKnrWza6m6cOGQecch-MWdf1VNdcMGvCPSBxfhEIJcje2j6I)

	utc := time.UTC
	kst, _ := time.LoadLocation("Asia/Seoul")

	whenUtc := time.Date(2020, 1, 1, 12, 10, 20, 0, utc)
	t.Log(whenUtc) // 2020-01-01 12:10:20 +0000 UTC
	whenKst := time.Date(2020, 1, 1, 12, 10, 20, 0, kst)
	t.Log(whenKst) // 2020-01-01 12:10:20 +0900 KST
}

func TestDuration(t *testing.T) {
	when := time.Date(2020, 1, 1, 12, 10, 20, 0, time.UTC)
	t.Log(when) // 2020-01-01 12:10:20 +0000 UTC

	whenAddSec := when.Add(time.Second * 10)
	whenAddHour := when.Add(time.Hour * 10)
	t.Log(whenAddSec)  // 2020-01-01 12:10:30 +0000 UTC
	t.Log(whenAddHour) // 2020-01-01 22:10:20 +0000 UTC

	assert.Equal(t, when.Before(whenAddSec), true)
	assert.Equal(t, when.Before(whenAddHour), true)

	assert.Equal(t, when.After(whenAddSec), false)
	assert.Equal(t, when.After(whenAddHour), false)

	assert.Equal(t, when.Equal(whenAddSec), false)
	assert.Equal(t, when.Equal(whenAddHour), false)
}
