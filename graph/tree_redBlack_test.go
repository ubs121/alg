package graph

import "testing"

func TestRedBlackBST(t *testing.T) {
	rbt := &RedBlackBST{}

	// Insert 5 nodes
	rbt.Put("E", "5")
	rbt.Put("A", "1")
	rbt.Put("S", "19")
	rbt.Put("Y", "25")
	rbt.Put("Q", "17")

	// Check node counts
	if rbt.Size(rbt.root) != 5 {
		t.Errorf("Expected size of 5, but got %d", rbt.Size(rbt.root))
	}

	// Check node values
	if rbt.Get("E").Data != "5" {
		t.Errorf("Expected node E to have data 5, but got %s", rbt.Get("E").Data)
	}
	if rbt.Get("A").Data != "1" {
		t.Errorf("Expected node A to have data 1, but got %s", rbt.Get("A").Data)
	}
	if rbt.Get("S").Data != "19" {
		t.Errorf("Expected node S to have data 19, but got %s", rbt.Get("S").Data)
	}
	if rbt.Get("Y").Data != "25" {
		t.Errorf("Expected node Y to have data 25, but got %s", rbt.Get("Y").Data)
	}
	if rbt.Get("Q").Data != "17" {
		t.Errorf("Expected node Q to have data 17, but got %s", rbt.Get("Q").Data)
	}
}
