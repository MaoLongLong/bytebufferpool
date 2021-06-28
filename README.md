# ByteBufferPool

> [bpool](https://github.com/minio/minio/blob/master/internal/bpool/bpool.go)
>
> [bytebufferpool](https://github.com/valyala/bytebufferpool)

## Example

```go
func ExampleByteBuffer() {
	buf := Get()

	buf.WriteString("first line\n")
	buf.Write([]byte("second line\n"))
	buf.B = append(buf.B, "third line\n"...)

	fmt.Printf("bytebuffer contents=%q", buf.B)

	Put(buf)

	//Output:
	//bytebuffer contents="first line\nsecond line\nthird line\n"
}
```

## Benchmark

```bash
$ go test -bench=Pool -benchmem
goos: linux
goarch: amd64
pkg: github.com/maolonglong/bytebufferpool
cpu: Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz
BenchmarkByteBufferPoolBuf-4    20764896                52.71 ns/op            0 B/op          0 allocs/op
BenchmarkBPool-4                 7899157               156.4 ns/op             0 B/op          0 allocs/op
BenchmarkWithoutPool-4           4511701               270.5 ns/op          1472 B/op          3 allocs/op
PASS
ok      github.com/maolonglong/bytebufferpool   4.091s
```
