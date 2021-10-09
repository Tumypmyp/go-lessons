package lessons

import (
	"fmt"
	"io"
)

type BufferedWriterCloser struct {
}

func (*BufferedWriterCloser) Write(p []byte) (n int, err error) {
	fmt.Println("writes")
	return 2, nil
}

func (*BufferedWriterCloser) Close() error {
	fmt.Println("closes")
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{}
}

func Lesson11() {
	// Interfaces

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 6; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("heloo world"))
	wc.Close()

	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("conversion failed")
	}

	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Close()
	}

	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("i don't know")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}
type ConsoleWriter struct {
}
type Incrementer interface {
	Increment() int
}
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
