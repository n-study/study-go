package main

import "fmt"

type test struct {
	name string
	age  int
}

func main() {
	/* var arr [3]int

	fmt.Println(arr)
	arr[0] = 135
	fmt.Println(arr)
	arr[1] = 452
	fmt.Println(arr)
	arr[2] = 653
	fmt.Println(arr)

	arr2 := [4]int{1, 2}
	fmt.Println(arr2)

	arr3 := [...]int{5, 3, 4, 6, 7}
	fmt.Println(arr3)

	arr4 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr4) */
	/*
		var s []int
		fmt.Println(s)
		fmt.Println(len(s), cap(s))

		s = append(s, 5)
		fmt.Println(s)
		fmt.Println(len(s), cap(s))

		s = append(s, 1, 2, 4, 5)
		fmt.Println(s)
		fmt.Println(len(s), cap(s))

		fmt.Println(s[:3]) */

	// var m map[string]string
	// fmt.Println(m)

	// m["a"] = "b"
	// fmt.Println(m)

	/* 	m := make(map[string]string)

	   	fmt.Println(m)
	   	m["a"] = "b"
	   	m["b"] = "c"
	   	m["c"] = "d"
	   	fmt.Println(m) */

	// m := map[string]string{
	// 	"a": "b",
	// 	"b": "c",
	// 	"c": "d",
	// }
	// fmt.Println(m)

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// fmt.Println(m)
	// delete(m, "a")
	// fmt.Println(m)

	// if _, exist := m["a"]; !exist {
	// 	fmt.Println("no key")
	// }

	a := test{
		name: "홍길동",
		age:  10,
	}
	fmt.Println(a)
	// fmt.Println(newFunction("go"))

	a.hi()

	fmt.Println(a.grow())
	fmt.Println(a)
}

func (t test) hi() {
	fmt.Println(t.name + ", hi")
}

func (t *test) grow() int {
	t.age++
	return t.age
}

func newFunction(a string) int {
	a += ", hi"
	fmt.Println(a)
	return 10
}
