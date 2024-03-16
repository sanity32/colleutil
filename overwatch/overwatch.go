package overwatch

import (
	"fmt"
	"time"
)

const (
	DEFAULT_TIMEOUT              time.Duration = time.Second * 60
	DEFAULT_COOLDOWN             time.Duration = time.Second * 2
	DEFAULT_MAX_FAILURES_ALLOWED int           = 3
	OVERWATCH_ENCOUNTERED_ERROR  int           = -1
	OVERWATCH_HAS_NO_RESULTS     int           = -2
	OVERWATCH_IS_EXPIRED         int           = -3
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
	C                  chan Result
	cbCheck            func() int
	startedAt          time.Time
	Expire             time.Duration
	Cooldown           time.Duration
	Failures           int
	MaxFailuresAllowed int
	isFinalized        bool
	stopFlag           bool
	stopFlagReturnCh   chan any
	LoudPrints         bool
}

func New(fn func() int, timeout, cooldown time.Duration) *Overwatch {
	if timeout == 0 {
		timeout = DEFAULT_TIMEOUT
	}
	if cooldown == 0 {
		cooldown = DEFAULT_COOLDOWN
	}
	return &Overwatch{
		cbCheck:            fn,
		C:                  make(chan Result),
		Expire:             timeout,
		Cooldown:           cooldown,
		MaxFailuresAllowed: DEFAULT_MAX_FAILURES_ALLOWED,
		LoudPrints:         false,
	}
}

func (o *Overwatch) finalize(result Result) {
	o.isFinalized = true
	defer func() { recover() }()
	o.C <- result
}

func (o Overwatch) printf(s string, vv ...any) {
	if o.LoudPrints {
		fmt.Printf(s, vv...)
	}
}

func (o *Overwatch) Stop() {
	if o.isFinalized {
		o.printf("unable to stop, instance is already finalized")
		return
	}
	o.stopFlag = true
	o.printf("Placing a stop flag\n")
	o.stopFlagReturnCh = make(chan any, 1)
	<-o.stopFlagReturnCh
	o.printf("Stopped!\n")
}

func (o Overwatch) expirationTime() time.Time {
	return o.startedAt.Add(o.Expire)
}

func (o Overwatch) isExpired() bool {
	return time.Now().After(o.expirationTime())
}

func (o *Overwatch) Start() *Overwatch {
	o.startedAt = time.Now()
	go func() {
		success, r, interrupted := o.poll()
		// if readyToSend := r >= 0; readyToSend {
		// 	o.finalize(Result{Err: !success, Result: r, Interrupted: interrupted})
		// }
		o.finalize(Result{Err: !success, Result: r, Interrupted: interrupted})
	}()

	return o
}

func (o *Overwatch) poll() (success bool, result int, interrupted bool) {
	var i int
	for {
		select {
		case <-o.C:
			o.stopFlag = true
		default:
		}

		i++
		o.printf("Overwatch iteration #%v\n", i)
		if o.stopFlag {
			o.printf("Stopping!\n")
			o.stopFlag = false
			o.stopFlagReturnCh <- nil
			return true, -1, true
		}

		resp := o.cbCheck()
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
		time.Sleep(o.Cooldown)
	}
}
