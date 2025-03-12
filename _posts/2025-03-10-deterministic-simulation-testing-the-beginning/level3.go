package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type Runner struct {
	rnd   *rand.Rand
	now   func() time.Time
	sleep func(time.Duration)
}

func (r Runner) Start(ctx context.Context) error {
	for {
		runMinute := int(r.rnd.Int31n(60))

		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			now := r.now()
			if now.Minute() != runMinute {
				r.sleep(time.Minute) // check if we should run next minute
				continue
			}

			err := r.runJob()
			if err != nil {
				return err
			}

			// we just ran, don't run until next hour
			r.sleep(time.Hour)
		}
	}
}

func (r Runner) runJob() error {
	if r.rnd.Float32() < 0.000000001 {
		return fmt.Errorf("rarely happening bug occurred!")
	}
	return nil
}

type result struct {
	seed    int64
	elapsed time.Duration
}

func main() {
	seeds := make(chan int64)
	go func() {
		for seed := 0; seed < 1000; seed++ {
			seeds <- int64(seed)
		}
	}()

	runners := runtime.GOMAXPROCS(0)
	results := make(chan result, runners)
	ctx := context.Background()

	for range runners {
		go func() {
			for seed := range seeds {
				ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
				runner := newRunner()
				runner.rnd.Seed(seed)

				t0 := time.Now()
				err := runner.Start(ctx)
				if err != nil && !errors.Is(err, context.DeadlineExceeded) {
					results <- result{
						seed:    seed,
						elapsed: time.Now().Sub(t0),
					}
				}
				cancel()
			}
		}()
	}

	for result := range results {
		fmt.Printf("found bug at seed %v after %v\n", result.seed, result.elapsed)
	}
}

func newRunner() *Runner {
	now := time.Time{}

	return &Runner{
		rnd: rand.New(rand.NewSource(1)),
		now: func() time.Time {
			return now
		},
		sleep: func(d time.Duration) {
			now = now.Add(d)
		},
	}
}
