package linkedlist

// LinkableList describes the set of methods of a LinkedList
type LinkableList interface {
	InsertFront(data interface{})
	InsertLast(data interface{})
	InsertAfter(prevNode *Node, data interface{})
	GetLastNode() *Node
	DeleteNodeByKey(key interface{})
	Reverse()
}

// Node represents a node in a LinkedList data structure
type Node struct {
	Data interface{}
	Next *Node
}

// NewNode constructs an instance of linkedlist.Node with the supplied 'data'
func NewNode(data interface{}) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

// LinkedList is a linear data structure storing data and address "next" at each node. Its' Head
// reference (pointer to Node) links to each subsequent Node in the LinkedList
type LinkedList struct {
	Head *Node
}

// New constructs an instance of LinkedList
func New() LinkableList {
	return &LinkedList{}
}

// InsertFront inserts the supplied data in a Node at the front of the LinkedList
func (l *LinkedList) InsertFront(data interface{}) {
	newNode := &Node{Data: data}
	newNode.Next = l.Head
	l.Head = newNode
}

// InsertLast inserts the supplied data in a Node at the last position in the LinkedList
func (l *LinkedList) InsertLast(data interface{}) {
	newNode := &Node{Data: data}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	lastNode := l.GetLastNode()
	lastNode.Next = newNode
}

// GetLastNode iterates the LinkedList until it reaches a nil "next" pointer then returns that node
func (l *LinkedList) GetLastNode() *Node {
	temp := l.Head
	for temp.Next != nil {
		temp = temp.Next
	}
	return temp
}

// InsertAfter inserts data in the Node after the supplied prevNode in the LinkedList
func (l *LinkedList) InsertAfter(prevNode *Node, data interface{}) {
	if prevNode == nil {
		return
	}
	newNode := NewNode(data)
	newNode.Next = prevNode.Next
	prevNode.Next = newNode
}

// DeleteNodeByKey deletes the node in the LinkedList having its' Data equal the supplied 'key'
func (l *LinkedList) DeleteNodeByKey(key interface{}) {
	temp := l.Head
	var prev *Node
	if temp != nil && temp.Data == key {
		l.Head = temp.Next
		return
	}
	for temp != nil && temp.Data != key {
		prev = temp
		temp = temp.Next
	}
	if temp == nil {
		return
	}
	prev.Next = temp.Next
}

// Reverse reverses the order of the Nodes in a LinkedList instance
func (l *LinkedList) Reverse() {
	var prev *Node
	current := l.Head
	var temp *Node
	for current != nil {
		temp = current.Next
		current.Next = prev
		prev = current
		current = temp
	}
	l.Head = prev
}
