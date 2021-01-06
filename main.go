package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Consonant string

// Vowels
const (
	// short vowels
	a string = "a"
	i string = "i"
	u string = "u"
	r string = "r"

	// long vowels
	_a string = "_a"
	_i string = "_i"
	_u string = "_u"
	_r string = "_r"

	// medium -> strong
	e  string = "e" // ai
	ai string = "ai"
	// medium -> strong
	o  string = "o" // au
	au string = "au"
)

// Consonants
const (
	s string = "s"
	t string = "t"
)

type Node struct {
	Prev   *Node
	Letter string
	Next   *Node
	Vowel  bool
}

type List struct {
	Head *Node
	Tail *Node
}

func (list *List) add(newNode *Node) {
	list.Tail.Next = newNode
	newNode.Prev = list.Tail
	newNode.Next = nil
	list.Tail = newNode
}

func (l *List) print() {
	output := ""
	curr := l.Head
	for curr != nil {
		output = fmt.Sprint(output, curr.Letter)
		curr = curr.Next
	}
	fmt.Println(output)
}

func random(letters []string) string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(letters)
	r := rand.Intn(max-min) + min
	return letters[r]
}

func main() {
	fmt.Println("HELLO LIST")

	var consonants = []string{s, t}
	var vowels = []string{a, i, u, r}

	initialNode := Node{Letter: random(vowels)}

	list := List{
		Head: &initialNode,
		Tail: &initialNode,
	}

	nextNode := Node{Letter: random(consonants)}
	list.add(&nextNode)

	nextNode2 := Node{Letter: random(vowels)}
	list.add(&nextNode2)

	list.print()
}

// EVAL
// type Node struct {
//	letter string
//	vowel bool
// }

// type Letter struct {
//  value     string
// 	vowel     bool
// }

// EVAL const x, y = Letter{"x", false}
