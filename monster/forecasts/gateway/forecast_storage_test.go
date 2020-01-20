package gateway

import (
	"testing"
	"time"
)

func Test_get24HoursBefore(t *testing.T) {
	now, before := calculate24HoursBefore()

	dayNow := time.Unix(now, 0)
	dayBefore := time.Unix(before, 0)

	if dayNow.Weekday() == dayBefore.Weekday() {
		t.Fail()
	}
}
