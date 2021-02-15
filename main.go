package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Guttural, Platal, Cerebral, Dental, Labia

// Vowels
const (
	// WEAK - short and long
	// short vowels
	a string = "a"
	i string = "i"
	u string = "u"
	r string = "r"

	// long vowels
	a_ string = "a_"
	i_ string = "i_"
	u_ string = "u_"
	r_ string = "r_"

	// MEDIUM  By adding an "a" sound to the front of the weak vowels, we get medium vowels.
	e string = "e" // ai, pronoounced a+i together, 'ay' because vocal formation midway between the two.
	o string = "o" // oh, pronunciating a+u together where they meet in the mouth

	// STRONG By strengthening the "a" sound in a medium vowel, we get strong vowels.
	ai string = "ai" // ah-ee, pronounced as a followed by i
	au string = "au" // ah-uu , pronounced as a followed by u
)

// Consonants
const (
	// Unaspirated, Aspirated, Voiced, Voiced
	// Guttural
	k  string = "k"
	kh string = "kh"
	g  string = "g"
	gh string = "gh"
	ng string = "ng" // n guttural

	// Palatal
	c  string = "c"
	ch string = "ch"
	j  string = "j"
	jh string = "jh"
	np string = "np" // n palatal

	// Cerebral
	t_  string = "t_"
	th_ string = "th_"
	d_  string = "d_"
	dh_ string = "dh_"
	nc_ string = "nc_" // n cerebral

	// Dental
	t  string = "t"
	th string = "th"
	d  string = "d"
	dh string = "dh"
	nd string = "nd" // n dental

	// Labial
	p  string = "p"
	ph string = "ph"
	b  string = "b"
	bh string = "bh"
	m  string = "m" // labial only
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

	var vowels = []string{
		a, i, u, r,
		a_, i_, u_, r_,
		e, o,
		ai, au,
	}
	var consonants = []string{
		k, kh, g, gh, ng,
		c, ch, j, jh, np,
		t_, th_, d_, dh_, nc_,
		t, th, d, dh, nd,
		p, ph, b, bh, m,
	}

	initialNode := Node{Letter: random(vowels)}
	list := List{
		Head: &initialNode,
		Tail: &initialNode,
	}

	nextNode := Node{Letter: random(consonants)}
	list.add(&nextNode)
	nextNode2 := Node{Letter: random(vowels)}
	list.add(&nextNode2)
	nextNode3 := Node{Letter: random(consonants)}
	list.add(&nextNode3)

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
