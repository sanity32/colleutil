package overwatch

import (
	"context"
	"fmt"
	"time"
)

const (
	DEFAULT_MAX_FAILURES_ALLOWED int = 3
	OVERWATCH_ENCOUNTERED_ERROR  int = -1
	OVERWATCH_HAS_NO_RESULTS     int = -2
	OVERWATCH_IS_EXPIRED         int = -3
)

var (
	DefaultLoudPrintValue = false
	DefaultExpirationTime = time.Minute * 3
	DefaultCooldown       = time.Second * 2
)

type Result struct {
	Result      int
	Err         bool
	Interrupted bool
}

func (r Result) SuccessFor(n int) bool {
	return r.Result == n && r.Success()
}

func (r Result) Success() bool {
	return !r.Err &&
		!r.Interrupted &&
		r.Result != OVERWATCH_ENCOUNTERED_ERROR &&
		r.Result != OVERWATCH_HAS_NO_RESULTS &&
		r.Result != OVERWATCH_IS_EXPIRED
}

type Overwatch struct {
	c                  chan Result
	Context            context.Context
	cbCheck            func(ctx context.Context) int
	startedAt          time.Time
	expire             time.Duration
	cooldown           time.Duration
	Failures           int
	MaxFailuresAllowed int
	isFinalized        bool
	stopFlag           bool
	LoudPrints         bool
}

func New(ctx context.Context, fn func(ctx context.Context) int) *Overwatch {
	return &Overwatch{
		cbCheck: fn,
		// c:                  make(chan Result),
		MaxFailuresAllowed: DEFAULT_MAX_FAILURES_ALLOWED,
		LoudPrints:         DefaultLoudPrintValue,
		expire:             DefaultExpirationTime,
		cooldown:           DefaultCooldown,
		Context:            ctx,
	}
}

// Deprecated

func NewOld(fn func(ctx context.Context) int, timeout, cooldown time.Duration) *Overwatch {
	if timeout == 0 {
		timeout = DefaultExpirationTime
	}
	if cooldown == 0 {
		cooldown = DefaultCooldown
	}
	return &Overwatch{
		cbCheck:            fn,
		c:                  make(chan Result),
		expire:             timeout,
		cooldown:           cooldown,
		MaxFailuresAllowed: DEFAULT_MAX_FAILURES_ALLOWED,
		LoudPrints:         false,
	}
}

func (o *Overwatch) C() <-chan Result {
	if o.c == nil {
		o.c = make(chan Result)
	}
	return o.c
}

func (o *Overwatch) WithCtx(ctx context.Context) *Overwatch {
	o.Context = ctx
	return o
}

func (o *Overwatch) Cooldown(v time.Duration) *Overwatch {
	o.cooldown = v
	return o
}

func (o *Overwatch) Expire(v time.Duration) *Overwatch {
	o.expire = v
	return o
}

func (o *Overwatch) finalize(result Result) {
	o.isFinalized = true
	defer func() { recover() }()
	o.c <- result
}

func (o Overwatch) printf(s string, vv ...any) {
	if o.LoudPrints {
		fmt.Printf(s, vv...)
	}
}

// func (o *Overwatch) Stop() {
// 	if o.isFinalized {
// 		o.printf("unable to stop, instance is already finalized")
// 		return
// 	}
// 	o.stopFlag = true
// 	o.printf("Placing a stop flag\n")
// 	o.stopFlagReturnCh = make(chan any, 1)
// 	<-o.stopFlagReturnCh
// 	o.printf("Stopped!\n")
// }

func (o Overwatch) expirationTime() time.Time {
	return o.startedAt.Add(o.expire)
}

func (o Overwatch) isExpired() bool {
	return time.Now().After(o.expirationTime())
}

func (o *Overwatch) Start() *Overwatch {
	o.startedAt = time.Now()
	go func() {
		success, r, interrupted := o.poll()
		o.finalize(Result{Err: !success, Result: r, Interrupted: interrupted})
	}()

	return o
}

func (o *Overwatch) poll() (success bool, result int, interrupted bool) {
	var i int
	ctx := o.Context
	if ctx == nil {
		newCtx, cancel := context.WithCancel(context.Background())
		ctx = newCtx
		defer cancel()
	}
	for {
		select {
		case <-ctx.Done():
			o.printf("stopped by context: %s", ctx.Err())
			o.stopFlag = true
		case <-o.c:
			o.printf("stopped by chan signal")
			o.stopFlag = true
		default:
		}

		i++
		o.printf("Overwatch iteration #%v\n", i)
		if o.stopFlag {
			return true, -1, true
		}

		resp := o.cbCheck(ctx)
		o.printf("resp: %#v\n", resp)
		if resp >= 0 {
			return true, resp, false
		}

		if hasErr := resp == OVERWATCH_ENCOUNTERED_ERROR; hasErr {
			if o.Failures++; o.Failures > o.MaxFailuresAllowed {
				return false, OVERWATCH_ENCOUNTERED_ERROR, false
			}
		}

		if o.isExpired() {
			return true, OVERWATCH_IS_EXPIRED, false
		}
		o.printf("Gonna be expired at %v", o.expirationTime())
		time.Sleep(o.cooldown)
	}
}
