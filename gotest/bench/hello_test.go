package bench

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
	"time"
)

//go:generate go test -v -bench=Hello2 -benchtime=3s -cpu=8 -benchmem -cpuprofile cpu.out
// go tool pprof  cpu.out
/*
BenchmarkHello2-8       32727718                99.8 ns/op             5 B/op          1 allocs/op

BenchmarkHello2-8：-cpu参数指定，-8表示8个CPU线程执行
32727718：表示总共执行了32727718次
99.8 ns/op：表示每次执行耗时99.8 ns
5 B/op:表示每次执行分配的内存（字节）
1 allocs/op：表示每次执行分配了多少次对象
*/

/*
函数名必须以 Benchmark 开头，后面一般跟待测试的函数名
参数为 b *testing.B。
执行基准测试时，需要添加 -bench 参数。

$ go test -benchmem -bench .
*/

/*
type BenchmarkResult struct {
    N         int           // 迭代次数
    T         time.Duration // 基准测试花费的时间
    Bytes     int64         // 一次迭代处理的字节数
    MemAllocs uint64        // 总的分配内存的次数
    MemBytes  uint64        // 总的分配内存的字节数
}
*/
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello()
	}
}

// 如果在运行前基准测试需要一些耗时的配置，则可以使用 b.ResetTimer() 先重置定时器，例如：
func BenchmarkHello2(b *testing.B) {
	time.Sleep(time.Second * 3) // 耗时操作
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

// 使用 RunParallel 测试并发性能
func BenchmarkParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}