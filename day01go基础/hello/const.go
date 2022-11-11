package main

const pi = 3.14

const (
	v = "v1.0"
	v2
	v3
	v4
)

const (
	week1 = 1
	week2 = 2
	week3 = 3
)

const (
	n1 = iota // 0
	n2        // 1
	n3
	n4
	n5
)

const (
	z1 = iota // 0
)

// 利用iota声明存储的单位常量
const (
	_  = iota             // 0
	KB = 1 << (10 * iota) // 1<<10 <=> 10000000000
	MB = 1 << (10 * iota) // 1<<20
	GB = 1 << (10 * iota) // 1<<30
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

// 声明中插队
const (
	c1 = iota // 0
	c2        // 1
	c3 = 100  // 插队
	c4 = iota // 3
)

const (
	d1, d2 = iota + 1, iota + 2 // 1,2
	d3, d4 = iota + 1, iota + 2 // 2,3
)


