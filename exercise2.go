package IN5170_ex5

// Implement an actor process for calculators as in exercise 1

type OP int

const (
	ADD   OP = 0
	INC   OP = 1
	STORE OP = 2
	DUAL  OP = 3
	SNGL  OP = 4
)

type Msg struct {
	op  OP
	p1  int
	p2  int
	ret chan int
}

type State struct {
	/* define state here*/
}

func loop1(ch chan Msg, state State) {
	/* add code here*/
}

/* add functions here */

func main() {
	// simple test case, write more
	//input := make(chan Msg)
	//go loop1(input, /* your starting state */)
	//res := make(chan int)
	//input <− Msg{STORE, 2, 0, res}
	//input <− Msg{INC, 5, 0, res}
	//fmt.Println(<−res) //should print 7
}
