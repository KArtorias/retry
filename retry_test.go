package retry

import (
	"testing"
	"time"
)

func TestDoFunc(t *testing.T) {
	err := DoFunc(3, time.Second, time.Millisecond*500, "Test", func() error {
		for i := 0; i < 12; i++ {
			t.Logf("data: %d", i)
			time.Sleep(time.Millisecond * 100)
		}
		return nil
	})

	if err != nil {
		t.Error("retry error:", err)
		return
	}
	t.Logf("retry success")
}
