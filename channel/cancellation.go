package channel

import (
	"context"
	"fmt"
	"time"
)

func ShowCancellation() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker stopped")
				return
			default:
				fmt.Println("working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	cancel() // 发出取消信号
}
