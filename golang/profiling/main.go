package main

import (
	"context"
	_ "embed"
	"errors"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

//go:embed go.mod
var GoModuleFileBytes []byte

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	// create cpu profiling file
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatalf("create cpu profiling fail: %v", err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	log.Print("begin")
	defer log.Print("end")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var n int
	for ; ; n++ {
		if err := demo(ctx); errors.Is(err, context.DeadlineExceeded) {
			break
		} else if err != nil {
			log.Fatalf("run demo fail: %v", err)
		}
	}
	log.Printf("run demo %d times", n)
}

func demo(ctx context.Context) error {
	// sleep 1s
	if err := sleep(ctx, time.Second); err != nil {
		return err
	}

	// sleep 100ms
	if err := sleep(ctx, time.Millisecond*100); err != nil {
		return err
	}

	return nil
}

func sleep(ctx context.Context, d time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(d):
		return nil
	}
}
