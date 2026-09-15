package gbs

import "testing"

func TestIsChannelOnline(t *testing.T) {
	for _, status := range []string{"ON", "OK", "ONLINE", " online "} {
		if !isChannelOnline(status) {
			t.Errorf("status %q should be online", status)
		}
	}

	for _, status := range []string{"OFF", "OFFILE", "UNKNOWN"} {
		if isChannelOnline(status) {
			t.Errorf("status %q should be offline", status)
		}
	}
}