package main

import "fmt"

const SIZE = 5

type Node struct {
	Data  string
	Left  *Node
	Right *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Hash map[string]*Node

type Cache struct {
	LL   LinkedList
	Hash Hash
}

func NewLinkedList() LinkedList {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head

	return LinkedList{Head: head, Tail: tail, Length: 0}
}

func NewCache() Cache {
	return Cache{
		LL:   NewLinkedList(),
		Hash: Hash{},
	}
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Remove %s Node\n", n.Data)
	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left
	c.LL.Length -= 1

	delete(c.Hash, n.Data)

	return n
}

func (c *Cache) Add(n *Node) {
	head := c.LL.Head
	first_elem := head.Right
	first_elem.Left = n
	n.Right = first_elem
	n.Left = head
	head.Right = n
	head.Left = nil

	c.LL.Length += 1

	if c.LL.Length > SIZE {
		c.Remove(c.LL.Tail.Left)
	}
}

func (c *Cache) Check(str string) {
	node := &Node{}

	fmt.Printf("\nChecking Node %s\n", str)

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Data: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Display() {
	c.LL.Display()
}

func (ll *LinkedList) Display() {
	temp := ll.Head.Right

	fmt.Printf("%d :-> ", ll.Length)
	for i := 0; i < ll.Length; i++ {
		fmt.Printf("%s ", temp.Data)
		temp = temp.Right

		if i < ll.Length-1 {
			fmt.Printf(" <--> ")
		}
	}
	fmt.Printf("\n")
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()

	for _, word := range []string{"hello", "world", "dog", "cat", "hello", "bat", "dog"} {
		cache.Check(word)
		cache.Display()
	}
}
