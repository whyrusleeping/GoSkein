package skein

type bufRecycler struct {
	ch chan []byte
}

func NewBufRecycler() *bufRecycler {
	b := new(bufRecycler)
	b.ch = make(chan []byte, 64)
	return b
}

var obuf = NewBufRecycler()

func FreeBuf(b []byte) {
	obuf.ch <- b
}

func GetBuf(size int) []byte {
	select {
	case c := <-obuf.ch:
		return c[:size]
	default:
		return make([]byte, size)
	}
	return nil
}
