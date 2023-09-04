package concurrent

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"testing"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string
type Search func(ctx context.Context, query string) (Result, error)

func fakeSearch(kind string) Search {
	return func(_ context.Context, query string) (Result, error) {
		return Result(fmt.Sprintf("%s result for %q", kind, query)), nil
	}
}

func TestReturnResult(t *testing.T) {
	Google := func(ctx context.Context, query string) ([]Result, error) {
		g, ctx := errgroup.WithContext(ctx)

		searches := []Search{Web, Image, Video}
		results := make([]Result, len(searches))
		for i, search := range searches {
			i, search := i, search // https://golang.org/doc/faq#closures_and_goroutines
			g.Go(func() error {
				result, err := search(ctx, query)
				if err == nil {
					results[i] = result
				}
				return err
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		}
		return results, nil
	}

	results, err := Google(context.Background(), "golang")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, result := range results {
		fmt.Println(result)
	}

	// Output:
	// web result for "golang"
	// image result for "golang"
	// video result for "golang"
}

// 使用errGroup统一返回err
func TestErrGroup(t *testing.T) {
	g, _ := errgroup.WithContext(context.Background())

	// 启动第一个任务
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("Task 1 start")
		time.Sleep(2 * time.Second)
		fmt.Println("Task 1 end")
		return nil
	})

	// 启动第二个任务
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("Task 2 start")
		time.Sleep(3 * time.Second)
		fmt.Println("Task 2 end")
		return errors.New("task 2 error")
	})

	// 启动第三个任务
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("Task 3 start")
		time.Sleep(1 * time.Second)
		fmt.Println("Task 3 end")
		return nil
	})
	i := 10
	g.Go(func() error {
		fmt.Println("Task 4 start")
		fmt.Println(i)
		fmt.Println("Task 4 end")
		return nil
	})

	// 等待所有任务完成，并获取聚合错误
	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("All tasks completed successfully")
	}
}

// 闭包
func TestClosure(t *testing.T) {
	g, _ := errgroup.WithContext(context.Background())
	result := make([]int, 10)
	for i := 0; i < 10; i++ {
		// 重新赋值
		i := i
		g.Go(func() error {
			fmt.Println(i)
			add1 := add1(i)
			result[i] = add1
			return nil
		})
	}
	err := g.Wait()
	fmt.Println(result)
	if err != nil {
		print("error")
	}
}

func add1(i int) int {
	return i + 1
}
