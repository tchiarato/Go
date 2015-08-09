package main

import "fmt"

type Operator interface {
	Calculate() int
}

type Sum struct {
	operatorA, operatorB int
}

type Subtract struct {
	operatorA, operatorB int
}

func (sum Sum) Calculate() int {
	return sum.operatorA + sum.operatorB
}

func (sum Sum) String() string {
	return fmt.Sprintf("%d + %d", sum.operatorA, sum.operatorB)
}

func (sub Subtract) Calculate() int {
	return sub.operatorA - sub.operatorB
}

func (sub Subtract) String() string {
	return fmt.Sprintf("%d - %d", sub.operatorA, sub.operatorB)
}

func main() {
	operators := make([]Operator, 4)
	operators[0] = Sum{10, 2}
	operators[1] = Subtract{5, 2}
	operators[2] = Subtract{25, 20}
	operators[3] = Sum{20, 20}

	acumulator := 0

	for _, operator := range operators {
		result := operator.Calculate()
		fmt.Printf("%v = %d\n", operator, result)
		acumulator += result
	}

	fmt.Printf("Result = %d", acumulator)
}
