package main

import "fmt"

// 类型转换

// T()

func f3() {
	var i11 int8 = 1

	i12 := int64(i11)             // int8 -> int64
	fmt.Printf("i12: %T \n", i12) // int64

	f11 := 12.34                  // float64
	f12 := int64(f11)             // float64 -> int64
	fmt.Printf("f12: %T \n", f12) // int64

	// bool(1) // int -> bool 🚫
}
