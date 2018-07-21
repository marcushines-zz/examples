package list

import (
	"fmt"
	"testing"
)

func EveryOther(n *Node) *Node {
	// if you pass in nil nothing to do
	if n == nil {
		return n
	}
	// if there is only one node - just pop itself
	if n.next == n {
		n = nil
		return n
	}
	c := n.next
	r := n.next
	d := true
	// Walk through each node deciding if needs to be removed.
	for c.next != n {
		if d {
			c.next = c.next.next
		} else {
			c = c.next
		}
		d = !d
	}
	c.next = r
	return r
}

func TestEveryOther(t *testing.T) {
	tests := []struct {
		desc string
		in   *Root
		out  *Node
	}{{
		desc: "nil",
		in:   New(),
	}, {
		desc: "single",
		in:   New().Append(&Node{v: 1}),
		out: &Node{
			v: 5,
		},
	}, {
		desc: "double",
		in:   New().Append(&Node{v: 1}).Append(&Node{v: 2}),
		out: &Node{
			v: 5,
		},
	}, {
		desc: "triple",
		in:   New().Append(&Node{v: 1}).Append(&Node{v: 2}).Append(&Node{v: 3}),
	}, {
		desc: "quad",
		in:   New().Append(&Node{v: 1}).Append(&Node{v: 2}).Append(&Node{v: 3}).Append(&Node{v: 4}),
	}, {
		desc: "five",
		in:   New().Append(&Node{v: 1}).Append(&Node{v: 2}).Append(&Node{v: 3}).Append(&Node{v: 4}).Append(&Node{v: 5}),
	}, {
		desc: "six",
		in:   New().Append(&Node{v: 1}).Append(&Node{v: 2}).Append(&Node{v: 3}).Append(&Node{v: 4}).Append(&Node{v: 5}).Append(&Node{v: 6}),
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			fmt.Printf("%s %v\n", tt.desc, tt.in.head.ToString())
			got := EveryOther(tt.in.head)
			fmt.Println("out:", got.ToString())
		})
	}
}

func xTestNew(t *testing.T) {
	r := New()
	fmt.Println(r)
	u := r.Reset()
	fmt.Println(r, u)
	r.Append(&Node{v: 5})
	fmt.Println(r, u)
	u.Append(&Node{v: "asdf"})
	fmt.Println(r, u)
	r.Reset()
	fmt.Println(r, u)
	u.Append(&Node{v: "asdf"}).Append(&Node{v: "def"}).Append(&Node{v: "123"}).Append(&Node{v: "zzz"})
	c := r.head.next
	fmt.Println("head", r.head)
	for c != r.head {
		fmt.Println(c)
		c = c.next
	}
	fmt.Println(r.Get(0))
	fmt.Println(r.Get(1))
	fmt.Println(r.Get(3))
	fmt.Println(r.Get(9))
}
