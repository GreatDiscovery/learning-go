package cronjob

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

// https://pkg.go.dev/github.com/robfig/cron

func TestCronFramework(t *testing.T) {
	c := cron.New()

	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.AddFunc("@every 5s", func() {
		fmt.Println("Every 5 seconds!")
	})
	c.Start()

	time.Sleep(600 * time.Second)
	c.Stop() // Stop the scheduler (does not stop any jobs already running).
}
