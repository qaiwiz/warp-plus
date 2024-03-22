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
// the RTT field of each IPInfo struct. This allows for efficient retrieval
// of the IPInfo struct with the lowest RTT value.
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
// This operation has a time complexity of O(1) as the first item in the sorted
// slice is always the one with the lowest RTT value.
func (q *IPInfQueue) Dequeue() IPInfo {
	// Check if the queue is empty.
	if len(q.items) == 0 {
	
