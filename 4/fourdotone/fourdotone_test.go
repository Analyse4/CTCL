package fourdotone

import "testing"

func TestRouteBetweenNodeDFS(t *testing.T) {
	eg := GenerateExGraph()

	if !RouteBetweenNodeDFS(2, 5, eg, make([]bool, eg.V())) {
		t.Fatalf("wrong, node-o: 2, node-e: 5\n")
	}

	if RouteBetweenNodeDFS(0, 6, eg, make([]bool, eg.V())) {
		t.Fatalf("wrong, node-o: 0, node-e: 3\n")
	}
}

func TestRouteBetweenNodeBFS(t *testing.T) {
	eg := GenerateExGraph()

	if !RouteBetweenNodeBFS(2, 5, eg) {
		t.Fatalf("wrong, node-o: 2, node-e: 5\n")
	}

	if RouteBetweenNodeBFS(0, 6, eg) {
		t.Fatalf("wrong, node-o: 0, node-e: 3\n")
	}
}
