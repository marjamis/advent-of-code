package structures

// MatrixPathQueueCoordinates stores the location in the matrix and it's parent in the path
type MatrixPathQueueCoordinates struct {
	Row    int
	Col    int
	Parent *MatrixPathQueueCoordinates
}

// MatrixPathQueue is a queue that forms the path taken from one set of coordinates to another
type MatrixPathQueue struct {
	data []MatrixPathQueueCoordinates
}

// Enqueue adds a new set of coordinates to the queue
func (q *MatrixPathQueue) Enqueue(s MatrixPathQueueCoordinates) {
	q.data = append(q.data, s)
}

// Dequeue remove the head of the queue
func (q *MatrixPathQueue) Dequeue() MatrixPathQueueCoordinates {
	dequeued := q.data[0]
	q.data = q.data[1:]

	return dequeued
}

// IsEmpty returns if the queue is empty
func (q *MatrixPathQueue) IsEmpty() bool {
	return len(q.data) == 0
}

// Size returns the length of the queue
func (q *MatrixPathQueue) Size() int {
	return len(q.data)
}

// NewQueue creates a new MatrixPathQueue
func NewQueue() *MatrixPathQueue {
	return &MatrixPathQueue{
		data: make([]MatrixPathQueueCoordinates, 0),
	}
}
