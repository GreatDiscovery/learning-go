package util

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

// 测试批量并行工具，控制并发和处理异常, learn from k8s parallelizer
// 首先，这个工具可以控制goroutine的数量，其次可以控制每次执行的任务的并发程度，把任务分成多个块进行处理
func TestParallel(t *testing.T) {
	//fwk.Parallelizer().Until(ctx, numAllNodes, checkNode, metrics.Filter)
	ParallelizerUtil(context.Background(), 4, 10, func(piece int) {
		fmt.Println(fmt.Sprintf("%v协程，正在处理第%v任务", GetGID(), piece))
	}, WithChunkSize(3))
}

type DoWorkPieceFunc func(piece int)
type Options func(*options)
type options struct {
	chunkSize int
}

func WithChunkSize(c int) func(*options) {
	return func(o *options) {
		o.chunkSize = c
	}
}

// ParallelizeUntil is a framework that allows for parallelize N
// independent pieces of work until done or the context is canceled.
func ParallelizerUtil(ctx context.Context, workers, piece int, doWorkPiece DoWorkPieceFunc, opts ...Options) {
	if piece == 0 {
		return
	}
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}
	chunkSize := o.chunkSize
	if chunkSize < 1 {
		chunkSize = 1
	}
	chunks := ceilDev(piece, chunkSize)
	toProcess := make(chan int, chunks)
	for i := 0; i < chunks; i++ {
		toProcess <- i
	}
	// 用完后及时关掉
	close(toProcess)

	var stop <-chan struct{}
	if ctx != nil {
		stop = ctx.Done()
	}
	if chunks < workers {
		workers = chunks
	}

	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			// fixme
			//defer HandleCrash()
			defer wg.Done()
			for chunk := range toProcess {
				start := chunk * chunkSize
				end := start + chunkSize
				if end > piece {
					end = piece
				}
				for p := start; p < end; p++ {
					select {
					case <-stop:
						return
					default:
						doWorkPiece(p)
					}
				}
			}
		}()
	}
	wg.Wait()
}

func ceilDev(a, b int) int {
	return (a + b - 1) / b
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
