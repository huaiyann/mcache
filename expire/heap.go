package expire

import "time"

type ExpireItem struct {
	Key      string
	ExpireAt time.Time
}

// A ExpireQueue implements heap.Interface and holds Items.
type ExpireQueue []*ExpireItem

func (eq ExpireQueue) Len() int { return len(eq) }

func (eq ExpireQueue) Less(i, j int) bool {
	return eq[i].ExpireAt.Before(eq[j].ExpireAt)
}

func (eq ExpireQueue) Swap(i, j int) {
	eq[i], eq[j] = eq[j], eq[i]
}

func (eq *ExpireQueue) Push(x any) {
	item := x.(*ExpireItem)
	*eq = append(*eq, item)
}

func (eq *ExpireQueue) Pop() any {
	old := *eq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*eq = old[0 : n-1]
	return item
}
