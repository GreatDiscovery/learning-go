package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStopwatch(t *testing.T) {
	stopwatch := NewNamedStopwatch()
	_ = stopwatch.AddMany([]string{
		"task1",
		"task2",
		"task3",
	})
	stopwatch.Start("task1")
	time.Sleep(time.Second)
	stopwatch.Stop("task1")

	stopwatch.Start("task2")
	time.Sleep(2 * time.Second)
	stopwatch.Stop("task2")

	stopwatch.Start("task3")
	time.Sleep(3 * time.Second)
	stopwatch.Stop("task3")

	assert.Equal(t, 1, int(stopwatch.ElapsedSeconds("task1")))
	assert.Equal(t, 2, int(stopwatch.ElapsedSeconds("task2")))
	assert.Equal(t, 3, int(stopwatch.ElapsedSeconds("task3")))

}
