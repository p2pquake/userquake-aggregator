package rate

import (
	"sync"
	"time"
)

type Limiter struct {
	mutex    *sync.Mutex
	timer    *time.Timer
	tailing  bool
	interval float64
	f        func()
}

func NewLimiter(interval float64, f func()) Limiter {
	return Limiter{mutex: &sync.Mutex{}, interval: interval, f: f}
}

func (l *Limiter) Run() {
	defer l.mutex.Unlock()
	l.mutex.Lock()

	// timer がない -> 即実行し timer セット (cool down)
	// timer がある -> tailing を true に
	if l.timer == nil {
		l.timer = time.AfterFunc(time.Duration(l.interval)*time.Second, l.run)
		l.f()
	} else {
		l.tailing = true
	}
}

func (l *Limiter) run() {
	defer l.mutex.Unlock()
	l.mutex.Lock()

	// tailing がない -> おわり
	// tailing がある -> 実行し timer セット (cool down)
	l.timer = nil
	if l.tailing {
		l.tailing = false
		l.timer = time.AfterFunc(time.Duration(l.interval)*time.Second, l.run)
		l.f()
	}
}
