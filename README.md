# Priority Queue

pqueue is a heap priority queue.

Example:
```
package main

import (
	"fmt"

	"github.com/yiqinguo/pqueue"
)

func main() {
	q := pqueue.NewPQueue()
	for i := int64(0); i < 5; i++ {
		q.Push(i, i)
	}

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

}
```
