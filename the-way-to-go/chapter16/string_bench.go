package main

import (
	"bytes"
	"fmt"
	"testing"
)

func main() {
	fmt.Printf(" benchPlus: \n", testing.Benchmark(benchPlus).String())
	fmt.Printf(" benchBuffer: \n", testing.Benchmark(benchBuffer).String())
}

// 导致大量的内存开销和拷贝
func benchPlus(b *testing.B) {
	str := ""
	for i := 0; i < b.N; i++ {
		str += "a"
	}
}

// 最快
func benchBuffer(b *testing.B) {
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		buf.WriteString("a")
	}
}
