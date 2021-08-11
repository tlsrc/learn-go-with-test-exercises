package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing should change the value", func(t *testing.T) {
		c := NewCounter()
		c.Inc()
		c.Inc()
		c.Inc()

		assertCounter(t, c, 3)
	})

	t.Run("It works concurrently", func(t *testing.T) {
		c := NewCounter()
		wantedCount := 1000

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, c, wantedCount)

	})
}

func assertCounter(t *testing.T, c *Counter, want int) {
	t.Helper()
	if c.Value() != want {
		t.Errorf("got %d, want %d", c.Value(), want)
	}
}
