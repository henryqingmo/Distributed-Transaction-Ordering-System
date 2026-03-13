package ordering

import (
	manager "cs425_mp1/internal/network"
	"sort"
)

type Queue struct {
	items []*QueueItem
}

type QueueItem struct {
	id          string
	tx          string
	priority    float32
	deliverable bool
	sender      string
}

type isisOrdering struct {
	holdbackQueue Queue
	messageMap    map[string]*QueueItem
}

func (q *Queue) Enqueue(item *QueueItem) {
	q.items = append(q.items, item)
	q.Sort()
}

func (q *Queue) Dequeue() *QueueItem {
	x := q.items[0]
	q.items = q.items[1:]
	return x
}

func (q *Queue) Peek() *QueueItem {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

func NewQueueItem(id string, tx string, priority float32, deliverable bool, sender string) *QueueItem {
	return &QueueItem{
		id:          id,
		tx:          tx,
		priority:    priority,
		deliverable: deliverable,
		sender:      sender,
	}
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

func (q *Queue) Sort() {
	sort.Slice(q.items, func(i, j int) bool {
		if q.items[i].priority == q.items[j].priority {
			return q.items[i].sender < q.items[j].sender
		}
		return q.items[i].priority < q.items[j].priority
	})
}

func (o *isisOrdering) OnReceiveTransaction(msg manager.Message) {
	if msg.Transaction.Kind == manager.Deposit {
	}

	if msg.Transaction.Kind == manager.Transfer {
	}

}

func (o *isisOrdering) OnReceivePropose(msg manager.Message) {

}

func (o *isisOrdering) onReceiveAgree(msg manager.Message) {

}
