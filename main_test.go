package main

import "testing"

var Result int

func BenchmarkIterative_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = Iterative(10)
	}
}

func BenchmarkIterative_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = Iterative(100)
	}
}

func BenchmarkIterative_1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = Iterative(1000)
	}
}

func BenchmarkIterative_10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = Iterative(10000)
	}
}

func BenchmarkIterative_100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = Iterative(100000)
	}
}

func BenchmarkIterative_1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = Iterative(1000000)
	}
}

////////////////////////////

func BenchmarkRecursive_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = Recursive(10)
	}
}

func BenchmarkRecursive_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = Recursive(100)
	}
}

func BenchmarkRecursive_1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = Recursive(1000)
	}
}

func BenchmarkRecursive_10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = Recursive(10000)
	}
}

func BenchmarkRecursive_100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = Recursive(100000)
	}
}

func BenchmarkRecursive_1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = Recursive(1000000)
	}
}

//////////////////////////////////
func BenchmarkTailRecursive_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = TailRecursive(10, 0)
	}
}

func BenchmarkTailRecursive_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = TailRecursive(100, 0)
	}
}

func BenchmarkTailRecursive_1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = TailRecursive(1000, 0)
	}
}

func BenchmarkTailRecursive_10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = TailRecursive(10000, 0)
	}
}

func BenchmarkTailRecursive_100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = TailRecursive(100000, 0)
	}
}

func BenchmarkTailRecursive_1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = TailRecursive(1000000, 0)
	}
}
