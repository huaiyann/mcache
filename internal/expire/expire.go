package expire

import (
	"container/heap"
	"sync"
	"time"
)

type Expire struct {
	lock   *sync.Mutex
	output chan string
	queue  *ExpireQueue
}

func NewExpire() *Expire {
	e := &Expire{
		lock:   new(sync.Mutex),
		output: make(chan string, 128),
		queue:  new(ExpireQueue),
	}
	go e.run()
	return e
}

func (e Expire) Add(key string, expireAt time.Time) {
	e.lock.Lock()
	defer e.lock.Unlock()
	heap.Push(e.queue, &ExpireItem{Key: key, ExpireAt: expireAt})
}

func (e Expire) NeedExpire() <-chan string {
	return e.output
}

func (e Expire) run() {
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()
	for {
		e.lock.Lock()
		if e.queue.Len() == 0 {
			e.lock.Unlock()
			<-ticker.C
			continue
		}
		data := heap.Pop(e.queue).(*ExpireItem)
		e.lock.Unlock()

		now := time.Now()
		if now.Before(data.ExpireAt) {
			// 没到过期时间，稍微等一下，然后塞回去下次重试
			<-ticker.C
			e.lock.Lock()
			heap.Push(e.queue, data)
			e.lock.Unlock()
		} else {
			// 到过期时间，正常输出
			e.output <- data.Key
		}
	}
}
