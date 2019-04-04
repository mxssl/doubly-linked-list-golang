package doublylinkedlist

import "testing"

func TestCase1(t *testing.T) {
	dll := NewDoublyLinkedList()
	node := NewNode(1)

	dll.setHead(node)
	expectingSingleNode(t, dll, node)
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

func expectingSingleNode(t *testing.T, dll *DoublyLinkedList, node *Node) {
	if dll.head != node {
		t.Fail()
	}
	if dll.tail != node {
		t.Fail()
	}
}
