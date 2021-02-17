package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key interface{}
}

type List struct {
	head *Node
	tail *Node
}

func main() {
	link := List{}

	fmt.Println(link)
}