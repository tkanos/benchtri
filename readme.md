# Benchtri

Do a benchmark between Iterative Recursive and TailRecursive simple algorithm.

Everybody knows that iterative is faster than recursive, but how fast it is ?

So I created a simple function that adds the first N integers. (e.g. sum(5) = 1 + 2 + 3 + 4 + 5 = 15).


we got the following numbers :

```  
? goos: linux
? goarch: amd64
? pkg: github.com/tkanos/console/benchtri
√ BenchmarkIterative_10-8             300000000          5.50 ns/op
√ BenchmarkIterative_100-8            30000000         40.9 ns/op
√ BenchmarkIterative_1000-8            5000000        299 ns/op
√ BenchmarkIterative_10000-8            500000       2952 ns/op
√ BenchmarkIterative_100000-8            50000      29981 ns/op
√ BenchmarkIterative_1000000-8            5000     282512 ns/op
√ BenchmarkRecursive_10-8             100000000         21.7 ns/op
√ BenchmarkRecursive_100-8             5000000        265 ns/op
√ BenchmarkRecursive_1000-8             500000       2690 ns/op
√ BenchmarkRecursive_10000-8             50000      27931 ns/op
√ BenchmarkRecursive_100000-8             5000     284461 ns/op
√ BenchmarkRecursive_1000000-8             200    5935023 ns/op
√ BenchmarkTailRecursive_10-8         50000000         22.9 ns/op
√ BenchmarkTailRecursive_100-8         5000000        242 ns/op
√ BenchmarkTailRecursive_1000-8         500000       2393 ns/op
√ BenchmarkTailRecursive_10000-8         50000      26300 ns/op
√ BenchmarkTailRecursive_100000-8         5000     266062 ns/op
√ BenchmarkTailRecursive_1000000-8         200    7657117 ns/op
? PASS
? ok   github.com/tkanos/console/benchtri 29.230s
```

For more information : https://medium.com/@felipedutratine/iterative-recursive-and-tail-recursive-camparition-f92e2bafd7e4
