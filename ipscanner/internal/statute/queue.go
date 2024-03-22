package statute

import (
	"sort"
	"time"
)

// IPInfQueue struct represents a priority queue of IPInfo structs,
// where the priority is determined by the RTT field.
type IPInfQueue struct {
	items []IPInfo // Slice of IPInfo structs.
}

// Enqueue adds an item to the IPInfQueue and sorts the queue based on
// the RTT field of each IPInfo struct.
func (q *IPInfQueue) Enqueue(item IPInfo) {
	// Add the item to the end of the queue.
	q.items = append(q.items, item)

	// Sort the queue based on the RTT field.
	sort.Slice(q.items, func(i, j int) bool {
		return q.items[i].RTT < q.items[j].RTT
	})
}

// Dequeue removes and returns the IPInfo struct with the lowest RTT value
// from the IPInfQueue. If the queue is empty, an empty IPInfo struct is returned.
func (q *IPInfQueue) Dequeue() IPInfo {
	// Check if the queue is empty.
	if len(q.items) == 0 {
		// Return an empty IPInfo struct.
		return IPInfo{}
	}

	// Remove and store the item with the lowest RTT.
	item := q.items[0]
	q.items = q.items[1:]

	// Set the CreatedAt field of the item to the current time.
	item.CreatedAt = time.Now()

	// Return the item.
	return item
}

// Size returns the number of items currently in the IPInfQueue.
func (q *IPInfQueue) Size() int {
	// Return the length of the items slice.
	return len(q.items)
}

