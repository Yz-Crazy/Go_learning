package main

import (
	"fmt"
	"strconv"
)

type option interface {
	computer(numA,numB float64) float64
}

type Add struct {

}

func (add *Add) computer(numA,numB float64) float64{
	return numA+numB
}

type Sub struct {

}

func (sub *Sub) computer(numA,numB float64) float64{
	return numA - numB
}

type Mul struct {

}

func (mul *Mul) computer(numA,numB float64) float64{
	return numA*numB
}

type Div struct {

}
func (div *Div) computer(numA,numB float64) float64{
	return numA/numB
}

func main()  {
	var opt option
	var numA,numB,s string
	var floata,floatb float64
	var err error
	for{

		fmt.Scan(&numA)


		floata,err=strconv.ParseFloat(numA, 64)
		if err ==nil{
			break
		}
		fmt.Println("只能输入数字")
	}
	for  {
		fmt.Scan(&numB)
		floatb, err=strconv.ParseFloat(numB, 64)
		if err ==nil{
			break
		}
		fmt.Println("只能输入数字")
	}

	fmt.Scan(&s)


	switch s {
	case "+":
		opt=new(Add)
	case "-":
		opt=new(Sub)
	case "*":
		opt=new(Mul)
	case "/":
		opt=new(Div)
	default:
		fmt.Println("请输入+-*/")
	}
	fmt.Println(opt.computer(floata,floatb))
}