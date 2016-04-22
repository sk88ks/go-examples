package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/k0kubun/pp"

	"golang.org/x/net/context"
)

// ContextErrFunc : 処理共通化のためのタスクの定義用
type ContextErrFunc func(ctx context.Context) error

// ContextSerial : ContextErrFunc を直列に実行して、エラーが途中で起こったらその時点でエラーを返す
func ContextSerial(fs ...ContextErrFunc) ContextErrFunc {
	return func(ctx context.Context) error {
		for _, f := range fs {
			if f == nil {
				continue
			}
			if err := f(ctx); err != nil {
				return err
			}
		}
		return nil
	}
}

// ContextParallel : ContextErrFunc を並列に実行して、エラーが途中で起こったらその時点でエラーを返す
func ContextParallel(fs ...ContextErrFunc) ContextErrFunc {
	return func(ctx context.Context) error {
		childCtx, cancelAll := context.WithCancel(ctx)
		defer cancelAll()

		doneCh := make(chan struct{}, len(fs))
		errCh := make(chan error, len(fs))
		recoverCh := make(chan interface{}, len(fs))

		for _, f := range fs {
			go func(_f ContextErrFunc) {
				defer func() {
					r := recover()
					if r != nil {
						recoverCh <- r
					}
				}()

				if _f == nil {
					doneCh <- struct{}{}
					return
				}

				if err := _f(childCtx); err != nil {
					errCh <- err
					return
				}
				doneCh <- struct{}{}
			}(f)
		}

		for i := 0; i < len(fs); i++ {
			select {
			case <-ctx.Done():
				pp.Println("CLOSED")
				return ctx.Err()
			case <-doneCh:
			case err := <-errCh:
				return err
			case r := <-recoverCh:
				panic(r)
			}
		}
		return nil
	}
}

func httpDoTask(req *http.Request, f func(*http.Response, error) error) ContextErrFunc {
	return func(ctx context.Context) error {
		// Run the HTTP request in a goroutine and pass the response to f.
		tr := &http.Transport{}
		client := &http.Client{Transport: tr}
		c := make(chan error, 1)
		go func() { c <- f(client.Do(req)) }()
		select {
		case <-ctx.Done():
			tr.CancelRequest(req)
			<-c // Wait for f to return.
			return ctx.Err()
		case err := <-c:
			return err
		}
	}
}

func DoSomething(ctx context.Context) (int, error) {
	return 1, nil
}

func Stream(done <-chan struct{}, ctx context.Context, out chan int) error {
	for {
		v, err := DoSomething(ctx)
		if err != nil {
			return err
		}
		select {
		case d := <-done:
			fmt.Println("DONE")
			fmt.Println(d)
			fmt.Println(ctx.Err())
			return ctx.Err()
		case out <- v:
		}
	}
}

func main() {
	//	req1, _ := http.NewRequest("GET", "http://google.com", nil)
	//	req2, _ := http.NewRequest("GET", "http://yahoo.com", nil)
	//	req3, _ := http.NewRequest("GET", "http://microsoft.com", nil)
	//
	//	// 全体で5秒でタイムアウト
	//	tc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//	defer cancel()
	//
	//	var res1, res2, res3 *http.Response
	//
	//	if err := ContextSerial(
	//		// 2つ並列に実行してから
	//		ContextParallel(
	//			httpDoTask(req1, func(ares1 *http.Response, err error) error {
	//				res1 = ares1
	//				return err
	//			}),
	//			httpDoTask(req2, func(ares2 *http.Response, err error) error {
	//				res2 = ares2
	//				return err
	//			}),
	//		),
	//		// 3つめを実行
	//		httpDoTask(req3, func(ares3 *http.Response, err error) error {
	//			res3 = ares3
	//			return err
	//		}),
	//	)(tc); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(res1.StatusCode, res2.StatusCode, res3.StatusCode)

	out := make(chan int)
	done := make(chan struct{})
	ctx := context.Background()
	c, _ := context.WithCancel(ctx)

	go Stream(done, c, out)

	go func() {
		time.Sleep(200 * time.Millisecond)
		close(done)
		//cancel()
		fmt.Println("Canceled")
		time.Sleep(200 * time.Millisecond)
		close(out)
	}()

	for i := range out {
		fmt.Println(i)
	}
}
