package main

import "fmt"

// eval(type letter)
// constant Letter a,i,u,r
// type Letter struct {
// 	vowel     string
// 	consonant string
// }
// ^ or let Letter accept a vowel or consonant...
// type Consonant struct {
// 	before Letter
// 	letter Letter
// 	after  Letter
// }

const (
	a string = "a"
	i string = "i"
	u string = "u"
	r string = "r"

	s string = "s"
)

// TODO where to put beforeLetter afterLetter vowelBefore vowelAfter consonantBefore consonantAfter
// type Consonant struct {
// 	Node
// }
// type Vowel struct {
// 	Node
// }

type Node struct {
	Prev   *Node
	Letter Character
	Next   *Node
	vowel  bool
}

type Character struct {
	string
}

type List struct {
	Head *Node
	Tail *Node
}

func (list *List) add(newNode *Node) {
	newNode.Prev = list.Tail
	newNode.Next = nil
	list.Tail = newNode
}

// TODO
// func (list *List) print() {
// 	while list.Head != nil { // or list.head.next
// 		fmt.Println(list.Head.Letter)
// 		list.shift()
// 	}
// }

func main() {
	fmt.Println("HELLO LIST")

	initialNode := Node{Letter: Character{a}}

	list := List{Head: &initialNode, Tail: &initialNode}

	fmt.Println(list.Head)
	fmt.Println(list.Tail)

	nextNode := Node{Letter: Character{s}}
	list.add(&nextNode)

	fmt.Println(list.Head)
	fmt.Println(list.Tail)

	nextNode2 := Node{Letter: Character{i}}
	list.add(&nextNode2)

	fmt.Println(list.Head)
	fmt.Println(list.Tail) // worked
}

// WORKED
// node1 := Node{Letter: Character{s}}
// node0 := Node{Letter: Character{a}, After: &node1}
// fmt.Println(node0.Letter)
// fmt.Println(node0.After.Letter)

// https://travix.io/type-embedding-in-go-ba40dd4264df
// fb := Football{Ball{Radius: 5, Material: "leather"}}

// WORKED node1 := Vowel{Node{Letter: Character{a}}}
