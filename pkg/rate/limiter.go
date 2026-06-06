package rate

import (
	"sync"
	"sync/atomic"
	"time"
)

type Limiter struct {
	mutex    *sync.Mutex
	timer    *time.Timer
	tailing  bool
	interval atomic.Int64 // 秒。整数秒で運用する前提
	f        func()
}

func NewLimiter(interval float64, f func()) *Limiter {
	l := &Limiter{mutex: &sync.Mutex{}, f: f}
	l.interval.Store(int64(interval))
	return l
}

// SetInterval は次回以降の cooldown 周期で使われる間隔（秒）を更新する。
// l.mutex を取らないので、f() コールバック内（再入）からも安全に呼べる。
func (l *Limiter) SetInterval(interval float64) {
	l.interval.Store(int64(interval))
}

func (l *Limiter) Run() {
	defer l.mutex.Unlock()
	l.mutex.Lock()

	// timer がない -> 即実行し timer セット (cool down)
	// timer がある -> tailing を true に
	if l.timer == nil {
		l.f()
		l.timer = time.AfterFunc(time.Duration(l.interval.Load())*time.Second, l.run)
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
		l.f()
		l.timer = time.AfterFunc(time.Duration(l.interval.Load())*time.Second, l.run)
	}
}
