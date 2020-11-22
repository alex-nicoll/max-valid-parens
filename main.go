package main

import (
	"fmt"
)

func main() {
	parens := "()))((()())()()))"
	fmt.Println(parens)
	fmt.Println("atoms:")
	printCells(parens, findAtoms(parens))
}

/**
 * linkedCells is a doubly-linked list of index pairs
 */
type linkedCells struct {
	prev  *linkedCells
	left  int
	right int
	next  *linkedCells
}

//func maxValidSubstring(parens string) string (
//	return ""
//}

func printCells(parens string, head *linkedCells) {
	for head != nil {
		fmt.Printf("%v %s\n", head.left, parens[head.left:head.right+1])
		head = head.next
	}
}

func findAtoms(parens string) *linkedCells {
	// first cell will be ignored and thrown away later
	head := &linkedCells{}
	o := head
	for i := 0; i < len(parens)-1; i++ {
		// Assume that a paren is one byte long
		left, right := parens[i], parens[i+1]
		if left == "("[0] && right == ")"[0] {
			o.next = &linkedCells{prev: o, left: i, right: i + 1}
			o = o.next
		}
	}
	// throwing away first cell, as promised
	head = head.next
	head.prev = nil
	return head
}
