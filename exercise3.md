# Exercise 3 

We discussed how to encode single-read futures with channels in the lecture. 
- Describe in words how to encode multi-read futures with channels, in particular 
      how to type the channel and what effect your encoding has on termination. 

- Use your encoding in Go and write a small program that uses goroutine that takes 
      two numbers, computes their sum and sends back the results using a multi-read 
      future (according to your encoding).

```go
package main

import "fmt"


func f (fut /* add type here*/, p1 int, p2 int) {
    /* add code here */
}

func main() {
    ch := make( /* add type here */ )
    go f (ch, 1, 2)
	
    //should work for any number of reads, test with 2
    fmt.Println(<−ch)
    fmt.Println(<−ch)
}
```