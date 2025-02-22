package main

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/util/runtime"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	// 获取当前时间
	currentTime := time.Now()
	// 设置时区
	location, _ := time.LoadLocation("Asia/Shanghai")
	// 将时间转换为指定时区
	currentTime = currentTime.In(location)
	// 格式化时间为指定的字符串格式
	formattedTime := currentTime.Format("2006-01-02T15:04:05-07:00")

	// 输出格式化后的时间字符串
	fmt.Println("格式化后的时间字符串:", formattedTime)

	formattedTime = currentTime.Format("20060102-150405")
	fmt.Println("格式化后的时间字符串:", formattedTime)
}

var count = 0

func TestTicker(t *testing.T) {
	background := context.Background()
	err := PollUntilContextTimeout(background, time.Second, 20*time.Second, false, printTime())
	if err != nil {
		fmt.Println(err)
	}
}

func printTime() ConditionWithContextFunc {
	return func(_ context.Context) (done bool, err error) {
		count++
		if count > 12 {
			return true, nil
		} else {
			fmt.Println(time.Now())
			return false, nil
		}
	}
}

func PollUntilContextTimeout(ctx context.Context, interval, timeout time.Duration, immediate bool, condition ConditionWithContextFunc) error {
	deadlineCtx, deadlineCancel := context.WithTimeout(ctx, timeout)
	defer deadlineCancel()
	return loopConditionUntilContext(deadlineCtx, interval, immediate, false, condition)
}

// ConditionWithContextFunc returns true if the condition is satisfied, or an error
// if the loop should be aborted.
//
// The caller passes along a context that can be used by the condition function.
type ConditionWithContextFunc func(context.Context) (done bool, err error)

// 多次循环，每次循环后延迟加倍，直到超时
func loopConditionUntilContext(ctx context.Context, initialInterval time.Duration, immediate, sliding bool, condition ConditionWithContextFunc) error {
	t := time.NewTimer(initialInterval)
	interval := initialInterval
	ratio := time.Duration(2)
	defer t.Stop()

	var timeCh <-chan time.Time
	doneCh := ctx.Done()

	// if immediate is true the condition is
	// guaranteed to be executed at least once,
	// if we haven't requested immediate execution, delay once
	if immediate {
		if ok, err := func() (bool, error) {
			defer runtime.HandleCrash()
			return condition(ctx)
		}(); err != nil || ok {
			return err
		}
	} else {
		timeCh = t.C
		select {
		case <-doneCh:
			return ctx.Err()
		case <-timeCh:
		}
	}

	for {
		// checking ctx.Err() is slightly faster than checking a select
		if err := ctx.Err(); err != nil {
			return err
		}

		if !sliding {
			interval = interval * ratio
			t.Reset(interval)
		}
		if ok, err := func() (bool, error) {
			defer runtime.HandleCrash()
			return condition(ctx)
		}(); err != nil || ok {
			return err
		}
		if sliding {
			interval = interval * ratio
			t.Reset(interval)
		}

		if timeCh == nil {
			timeCh = t.C
		}

		// NOTE: b/c there is no priority selection in golang
		// it is possible for this to race, meaning we could
		// trigger t.C and doneCh, and t.C select falls through.
		// In order to mitigate we re-check doneCh at the beginning
		// of every loop to guarantee at-most one extra execution
		// of condition.
		select {
		case <-doneCh:
			return ctx.Err()
		case <-timeCh:
		}
	}
}
