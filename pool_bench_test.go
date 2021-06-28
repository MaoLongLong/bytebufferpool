package bytebufferpool

import (
	"testing"

	"github.com/maolonglong/bpool"
)

var str = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
	`Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
		nisi ut aliquip ex ea commodo consequat.
		Duis aute irure dolor in reprehenderit in voluptate velit esse cillum
		dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident,
		sunt in culpa qui officia deserunt mollit anim id est laborum`,
	"Sed ut perspiciatis",
	"sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt",
	"Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit",
	"laboriosam, nisi ut aliquid ex ea commodi consequatur",
	"Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur",
	"vel illum qui dolorem eum fugiat quo voluptas nulla pariatur",
}

func BenchmarkByteBufferPoolBuf(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := Get()
			for _, s := range str {
				buf.WriteString(s)
			}
			Put(buf)
		}
	})
}

func BenchmarkBPool(b *testing.B) {
	pool := bpool.NewBytePoolCap(500, 0, 64)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := pool.Get()
			for _, s := range str {
				buf = append(buf, s...)
			}
			buf = buf[:0]
			pool.Put(buf)
		}
	})
}

func BenchmarkWithoutPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := make([]byte, 0, 64)
			for _, s := range str {
				buf = append(buf, s...)
			}
		}
	})
}
