package main

import "fmt"

func testIf1() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range a {
		if a[i]%2 == 0 {
			fmt.Printf("%d 是偶数\n", a[i])
		} else {
			fmt.Printf("%d 是奇数\n", a[i])
		}
	}
}

func testIf2() {
	for i := 0; i < 20; i++ {
		if i < 10 {
			fmt.Printf("%d 小于10\n", i)
		} else if i < 20 && i > 10 {
			fmt.Printf("%d 大于10，小于20\n", i)
		} else {
			fmt.Printf("%d\n", i)
		}

	}
}

func testIf3() {
	if num := 10; num%2 == 0 {
		fmt.Printf("%d 是个偶数\n", num)
	} else {
		fmt.Printf("%d 是奇数", num)
	}
}

func testFor01() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func testFor02() {
	for i := 0; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)

	}
}

func testFor03() {
	for i := 0; i <= 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

func testFor04() {
	i := 0
	for {
		i++
		if i >= 6 {
			break
		}
		fmt.Println(i)
	}
}

// 用于多分支
func testSwitch01() {

	for i := 0; i <= 10; i++ {
		switch i {
		case 1: // case 可以多个值
			fmt.Println("a")
		case 2:
			fmt.Println("b")
		case 3:
			fmt.Println("c")
		case 4:
			fmt.Println("d")
		default:
			fmt.Println("e")

		}
	}

}

// case 判断条件
func testSwitch02() {
	for i := 0; i < 20; i++ {
		switch {
		case i <= 10:
			fmt.Printf("%d小于10\n", i)
		case i > 10 && i < 15:
			fmt.Printf("%d大于10，小于15\n", i)
		default:
			fmt.Printf("%d不符合标准\n", i)
		}
	}
}

// fallthrough 如果命中含有fallthrough 的case 那么执行完这一个case之后还会执行它下面一个的case
func testSwitch03() {
	num := 70
	switch {
	case num < 10:
		fmt.Println("小于10")
	case num > 20 && num < 80:
		fmt.Println("num>20 && num <80")
		fallthrough
	case num > 80 && num < 100:
		fmt.Println("num >80 && num <100")
	case num > 100 && num < 200:
		fmt.Println("num>100 && num <200")
	default:
		fmt.Println("全部不中")
	}
}

func testMulti() {
	for j := 1; j <= 9; j++ {
		for k := 1; k <= j; k++ {
			fmt.Printf("%d * %d = %d\t", k, j, j*k)
		}
		fmt.Println()
	}

}

func main() {
	testIf1()
	testIf2()
	testIf3()
	testFor01()
	testFor02()
	testFor03()
	testFor04()
	testSwitch01()
	testSwitch02()
	testSwitch03()
	testMulti()
}
