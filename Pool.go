package binary

import "sync"

var readerPool = &sync.Pool{
	New: func() interface{} {
		return &Reader{}
	},
}

var writerPool = &sync.Pool{
	New: func() interface{} {
		return &Writer{}
	},
}

func AcquireReader(buf []byte, n int) (r *Reader) {
	r = readerPool.Get().(*Reader)
	r.init(buf, n)
	return
}

func ReleaseReader(r *Reader) {
	readerPool.Put(r)
}

func AcquireWriter(cap int) (w *Writer) {
	w = writerPool.Get().(*Writer)
	w.init(cap)
	return
}

func ReleaseWriter(w *Writer) {
	writerPool.Put(w)
}
