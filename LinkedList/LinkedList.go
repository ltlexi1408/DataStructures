package linkedlist

type Node struct {
	word string
	next *Node
}

type linkedlist struct {
	head *Node
}

func (list *linkedlist) add(word string) error {
	node := &Node{
		word: word,
		next: nil}

	if list.head == nil {
		list.head = node
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}

	return nil
}
