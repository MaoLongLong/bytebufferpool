package bytebufferpool

import "fmt"

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
