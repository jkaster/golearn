package collections

import "fmt"

// --- Maps (Dictionaries/Hash Tables) ---
// demonstrateMaps shows the usage of Go's built-in map type, which serves as a dictionary or hash table.
func demonstrateMaps() {
	fmt.Println("\n--- Maps Example ---")

	// Create a map using make
	m := make(map[string]int)

	// Set key/value pairs
	m["apple"] = 10
	m["banana"] = 20

	fmt.Println("Map:", m)

	// Get a value by key
	appleCount := m["apple"]
	fmt.Println("Apple count:", appleCount)

	// Check if a key exists (and get its value)
	val, prs := m["banana"]
	fmt.Println("Banana value:", val, "present:", prs)

	val, prs = m["grape"]
	fmt.Println("Grape value:", val, "present:", prs) // 0, false (zero value for int, false for presence)

	// Delete a key/value pair
	delete(m, "apple")
	fmt.Println("Map after deleting apple:", m)

	// Get length of map
	fmt.Println("Map length:", len(m))

	// Declare and initialize a map in one line
	n := map[string]string{"name": "Alice", "city": "New York"}
	fmt.Println("Another map:", n)
}

// --- Stack (LIFO - Last In, First Out) ---
// A simple stack implementation using a slice.
type Stack[T any] struct {
	data []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

// Pop removes and returns the top element from the stack.
// Returns the zero value of T and false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T // Zero value for type T
		return zero, false
	}
	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index] // Remove the element
	return item, true
}

// Peek returns the top element without removing it.
// Returns the zero value of T and false if the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// demonstrateStack shows a generic stack (LIFO) implementation using a slice.
func demonstrateStack() {
	fmt.Println("\n--- Stack Example ---")
	var s Stack[int] // Create an empty stack of integers

	fmt.Println("Is stack empty?", s.IsEmpty())
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Stack size after pushes:", s.Size())

	val, ok := s.Pop()
	fmt.Printf("Popped: %v, success: %v\n", val, ok)

	val, ok = s.Peek()
	fmt.Printf("Peeked: %v, success: %v\n", val, ok)

	fmt.Println("Is stack empty?", s.IsEmpty())

	s.Pop()
	s.Pop()
	val, ok = s.Pop() // Try to pop from empty stack
	fmt.Printf("Popped from empty: %v, success: %v\n", val, ok)
}

// --- Queue (FIFO - First In, First Out) ---
// A simple queue implementation using a slice.
type Queue[T any] struct {
	data []T
}

// Enqueue adds an element to the back of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.data = append(q.data, item)
}

// Dequeue removes and returns the front element from the queue.
// Returns the zero value of T and false if the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	item := q.data[0]
	q.data = q.data[1:] // Remove the first element
	return item, true
}

// Peek returns the front element without removing it.
// Returns the zero value of T and false if the queue is empty.
func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.data[0], true
}

// IsEmpty checks if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return len(q.data)
}

// demonstrateQueue shows a generic queue (FIFO) implementation using a slice.
func demonstrateQueue() {
	fmt.Println("\n--- Queue Example ---")
	var q Queue[string] // Create an empty queue of strings

	fmt.Println("Is queue empty?", q.IsEmpty())
	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")
	fmt.Println("Queue size after enqueues:", q.Size())

	val, ok := q.Dequeue()
	fmt.Printf("Dequeued: %v, success: %v\n", val, ok)

	val, ok = q.Peek()
	fmt.Printf("Peeked: %v, success: %v\n", val, ok)

	fmt.Println("Is queue empty?", q.IsEmpty())

	q.Dequeue()
	q.Dequeue()
	val, ok = q.Dequeue() // Try to dequeue from empty queue
	fmt.Printf("Dequeued from empty: %v, success: %v\n", val, ok)
}

// Collections demonstrates various collection types: maps (hash tables/dictionaries), and custom implementations of stacks and queues.
func Collections() {
	fmt.Println("\n--- Collections Example ---")
	demonstrateMaps()
	demonstrateStack()
	demonstrateQueue()
}
