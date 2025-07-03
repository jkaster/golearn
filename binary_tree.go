package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Node represents a node in the binary tree.
// It uses a type parameter T, allowing it to store any type of value.
type Node[T constraints.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// BinaryTree represents the binary tree structure.
// It also uses a type parameter T, consistent with its Node type.
type BinaryTree[T constraints.Ordered] struct {
	Root *Node[T]
}

// Insert adds a new value to the binary tree.
// This is a simple insertion, not necessarily a balanced tree.
func (bt *BinaryTree[T]) Insert(value T) {
	newNode := &Node[T]{Value: value}
	if bt.Root == nil {
		bt.Root = newNode
		return
	}

	current := bt.Root
	for {
		if value == current.Value { // Handle duplicates (optional: could error or ignore)
			return
		}
		if value < current.Value {
			if current.Left == nil {
				current.Left = newNode
				return	
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				return
			}
			current = current.Right
		}
	}
}

// DFS (Depth-First Search) performs an in-order traversal of the tree.
// It returns a slice of values in the order they are visited.
func (bt *BinaryTree[T]) DFS() []T {
	var result []T
	var traverse func(node *Node[T])
	traverse = func(node *Node[T]) {
		if node == nil {
			return
		}
		traverse(node.Left)
		result = append(result, node.Value)
		traverse(node.Right)
	}
	traverse(bt.Root)
	return result
}

// BFS (Breadth-First Search) performs a level-order traversal of the tree.
// It returns a slice of values in the order they are visited.
func (bt *BinaryTree[T]) BFS() []T {
	var result []T
	if bt.Root == nil {
		return result
	}

	queue := []*Node[T]{bt.Root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // Dequeue

		result = append(result, current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left) // Enqueue left child
		}
		if current.Right != nil {
			queue = append(queue, current.Right) // Enqueue right child
		}
	}
	return result
}

// binaryTreeExample demonstrates the usage of the generic BinaryTree with integer and string types.
// It shows insertion, DFS (Depth-First Search), and BFS (Breadth-First Search) traversals.
func binaryTreeExample() {
	fmt.Println("\n--- Binary Tree Example (int) ---")
	intTree := &BinaryTree[int]{}
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, val := range values {
		intTree.Insert(val)
	}

	fmt.Println("DFS (In-order) Traversal:", intTree.DFS())
	fmt.Println("BFS (Level-order) Traversal:", intTree.BFS())

	fmt.Println("\n--- Binary Tree Example (string) ---")
	stringTree := &BinaryTree[string]{}
	stringValues := []string{"apple", "banana", "cherry", "date", "fig", "grape"}
	for _, val := range stringValues {
		stringTree.Insert(val)
	}

	fmt.Println("DFS (In-order) Traversal:", stringTree.DFS())
	fmt.Println("BFS (Level-order) Traversal:", stringTree.BFS())
}
