package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
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

func main() {
	ctx := context.Background()
	runner := newRunner()

	var seed int64
	for ; ; seed++ {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		runner.rnd.Seed(seed)

		t0 := time.Now()
		err := runner.Start(ctx)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				log.Printf("seed %v timed out", seed)
				continue
			}
			log.Printf("seed: %v after %v", seed, time.Now().Sub(t0))
		}
		cancel()
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
