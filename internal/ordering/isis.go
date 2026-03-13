package ordering

import (
	manager "cs425_mp1/internal/network"
)

type Queue struct {
	items []int
}

type isisOrdering struct {
	queue Queue
}

func (q *Queue) Enqueue(x int) {
	q.items = append(q.items, x)
}

func (q *Queue) Dequeue() int {
	x := q.items[0]
	q.items = q.items[1:]
	return x
}

func (o *isisOrdering) HandleMessage(msg manager.Message) {
	switch msg.Type {
	case manager.TypeTransaction:
		o.OnReceiveTransaction(msg)
	case manager.TypePropose:
		o.OnReceivePropose(msg)
	case manager.TypeAgree:
		o.onReceiveAgree(msg)
	}

}

func (o *isisOrdering) OnReceiveTransaction(msg manager.Message) {

}

func (o *isisOrdering) OnReceivePropose(msg manager.Message) {

}

func (o *isisOrdering) onReceiveAgree(msg manager.Message) {

}
