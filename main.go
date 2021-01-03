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

// TODO where to put beforeLetter afterLetter vowelBefore vowelAfter consonantBefore consonantAfter
type Consonant struct {
	Node
}

type Vowel struct {
	Node
}

type Node struct {
	Before *Node
	Letter Character
	// Vowel Character? probably not
	// Consonant Character?
	After *Node
	vowel bool
}

type Character struct {
	string
}

const (
	a string = "a"
	i string = "i"
	u string = "u"
	r string = "r"

	s string = "s"
)

// https://travix.io/type-embedding-in-go-ba40dd4264df
// fb := Football{Ball{Radius: 5, Material: "leather"}}

func main() {
	fmt.Print("HELLO WORLD\n")
	//linked list
	node1 := Node{Letter: Character{s}}
	// vowel := Vowel{node1} // do I even need this Vowel wrap?
	node0 := Consonant{Node{Letter: Character{a}, After: &node1}}
	fmt.Println(node0.Letter)
	fmt.Println(node0.After.Letter)
}

// WORKED node1 := Vowel{Node{Letter: Character{a}}}
