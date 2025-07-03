package main

import (
	"reflect"
	"testing"
)

func TestBinaryTreeIntTraversal(t *testing.T) {
	intTree := &BinaryTree[int]{}
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, val := range values {
		intTree.Insert(val)
	}

	expectedDFS := []int{20, 30, 40, 50, 60, 70, 80} // In-order traversal for a simple BST
	actualDFS := intTree.DFS()
	if !reflect.DeepEqual(actualDFS, expectedDFS) {
		t.Errorf("DFS traversal failed for int tree. Expected %v, got %v", expectedDFS, actualDFS)
	}

	expectedBFS := []int{50, 30, 70, 20, 40, 60, 80} // Level-order traversal
	actualBFS := intTree.BFS()
	if !reflect.DeepEqual(actualBFS, expectedBFS) {
		t.Errorf("BFS traversal failed for int tree. Expected %v, got %v", expectedBFS, actualBFS)
	}
}

func TestBinaryTreeStringTraversal(t *testing.T) {
	stringTree := &BinaryTree[string]{}
	stringValues := []string{"apple", "banana", "cherry", "date", "fig", "grape"}
	for _, val := range stringValues {
		stringTree.Insert(val)
	}

	// Expected DFS (in-order) for the given insertion order
	// This assumes a simple BST insertion where 'a' < 'b' etc.
	expectedDFS := []string{"apple", "banana", "cherry", "date", "fig", "grape"}
	actualDFS := stringTree.DFS()
	if !reflect.DeepEqual(actualDFS, expectedDFS) {
		t.Errorf("DFS traversal failed for string tree. Expected %v, got %v", expectedDFS, actualDFS)
	}

	// Expected BFS (level-order) for the given insertion order
	expectedBFS := []string{"apple", "banana", "cherry", "date", "fig", "grape"}
	actualBFS := stringTree.BFS()
	if !reflect.DeepEqual(actualBFS, expectedBFS) {
		t.Errorf("BFS traversal failed for string tree. Expected %v, got %v", expectedBFS, actualBFS)
	}
}
