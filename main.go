package main

import (
	"flag"
	"fmt"
	"github.com/thoas/go-funk"
	"strconv"
	"strings"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func main() {
	flag.Parse()
	args := flag.Args()
	numbers := make([]int, len(args))
	for i, arg := range args {
		number, err := strconv.Atoi(arg)
		if err != nil {
			panic("cannot parse argument into number: " + arg)
		}
		numbers[i] = number
	}

	n := newNode(funk.UniqInt(numbers))
	printNode(n, 0, 'M')
}

func newNode(numbers []int) *Node {
	if len(numbers) == 0 {
		return nil
	}

	maxInt := funk.MaxInt(numbers)
	idx := funk.IndexOfInt(numbers, maxInt)
	root := &Node{value: maxInt}
	root.left = newNode(numbers[:idx])
	root.right = newNode(numbers[idx+1:])
	return root
}

func printNode(node *Node, indent int, identifier rune) {
	if node == nil {
		return
	}

	fmt.Printf("%s%c:%v\n", strings.Repeat(" ", indent), identifier, node.value)
	printNode(node.left, indent+2, 'L')
	printNode(node.right, indent+2, 'R')
}
