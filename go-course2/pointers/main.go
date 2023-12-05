package main

import (
	"fmt"
)

// Modifies pointer to an int (value in a mem address pointed to)
// Pointer semantics
func intDelta(n *int) {
	*n = 210
}

// value semantics
func intDelta2(n int) int {
	return n + 9
}

func sliceDelta(n []int) {
	n[0] = 99
}

func mapDelta(mpd map[string]int, s string) {
	mpd[s] = 45
}

func main() {
	val := 45
	fmt.Println(val)
	// point to val (Mem Address for val)
	p2Val := &val
	fmt.Printf("%v %T", p2Val, p2Val)
	fmt.Printf("%T", p2Val)
	// val := *p
	// value in the adress of val(&val), similar *p
	fmt.Println(*&val)
	fmt.Println(*p2Val)
	fmt.Printf("%T", *p2Val)
	*p2Val = 46 //Update val in mem address p2Val - dereferencing mem address
	fmt.Println(val)

	xy := 200
	fmt.Println(xy)
	intDelta(&xy)
	fmt.Println(xy)

	fmt.Println(intDelta2(xy))

	xi := []int{4, 5, 9, 3}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	mp := make(map[string]int)
	mp["jon"] = 56
	fmt.Println(mp["jon"])
	mapDelta(mp, "jon")
	fmt.Println(mp["jon"])

}
