package main

import (
	"fmt"
    "maps"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// The builtin `delete` removes key/value pairs from
	// a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// To remove *all* key/value pairs from a map, use
	// the `clear` builtin.
	clear(m)
	fmt.Println("map:", m)

	//在 Go 语言中，从映射（map）中检索值可以同时返回两个结果：值本身和一个布尔值，
	//后者表明键（key）是否真的存在于映射中。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in
	// the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

    	// The `maps` package contains a number of useful
	// utility functions for maps.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
