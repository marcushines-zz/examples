package list

import (
	"fmt"
	"reflect"
)

type Root struct {
	head *Node
	tail *Node
}

func New() *Root {
	return &Root{}
}

func (r Root) Get(i int) *Node {
	if r.head == nil {
		return nil
	}
	if i == 0 {
		return r.head
	}
	curr := r.head.next
	for {
		if i == 1 {
			return curr
		}
		if curr == r.tail {
			return nil
		}
		curr = curr.next
		i--
	}
}

func (r *Root) Append(n *Node) *Root {
	if r.head == nil {
		r.head = n
		r.tail = n
		n.next = n
		return r
	}
	r.tail.next = n
	n.next = r.head
	r.tail = n
	return r
}

func (r *Root) Insert(n *Node) *Root {
	n.next = r.head
	r.head = n
	if r.tail == nil {
		r.tail = n
	}
	return r
}

func (r *Root) Reset() *Root {
	r.head = nil
	r.tail = nil
	return r
}

type Node struct {
	v    interface{}
	next *Node
}

func (n *Node) Equal(e *Node) bool {
	if e == nil {
		return false
	}
	if !reflect.DeepEqual(n.v, e.v) {
		return false
	}
	return true
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Value() interface{} {
	return n.v
}

func (n *Node) ToString() string {
	if n == nil {
		return fmt.Sprintf("%v", n)
	}
	v := fmt.Sprintf("%+v", n)
	c := n.next
	for c != n {
		v += fmt.Sprintf(" %+v", c)
		c = c.next
	}
	return v
}
