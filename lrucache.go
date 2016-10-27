package main

import (
	"fmt"
)

func main() {

	var L int32
	fmt.Scanln(&L)

	keys := make(map[int32]*LruNode)
	lsHead := &LruNode{Pre: nil, Next: nil}
	lsTail := &LruNode{Pre: nil, Next: nil}
	lsHead.Next = lsTail
	lsTail.Pre = lsHead

	var N int32
	fmt.Scan(&N)

	missing := int32(0)

	for i := int32(0); i < N; i++ {

		var xi int32
		fmt.Scan(&xi)

		node, hit := keys[xi]

		if hit {

			node.Next.Pre = node.Pre
			node.Pre.Next = node.Next
		} else {

			node = &LruNode{Addr: xi}
			keys[xi] = node
		}

		node.Pre = lsTail.Pre
		lsTail.Pre.Next = node
		node.Next = lsTail
		lsTail.Pre = node

		if int32(len(keys)) > L {

			unused := lsHead.Next
			lsHead.Next = unused.Next
			unused.Next.Pre = lsHead
			delete(keys, unused.Addr)
		}

		if !hit {
			missing++
		}
	}

	fmt.Println(missing)
}

type LruNode struct {
	Pre  *LruNode
	Next *LruNode
	Addr int32
}
