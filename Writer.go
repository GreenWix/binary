package binary

type Writer struct {
	n   int
	buf []byte
}

// Данная структура не является io.ByteWriter!!!
//noinspection GoStandardMethods
func (w *Writer) WriteByte(number byte) {
	w.Write(1, []byte{number})
}

func (w *Writer) WriteSignedShort(number int16) {
	_ = w.WriteNumber(int64(number), 2)
}

func (w *Writer) WriteUnsignedShort(number uint16) {
	_ = w.WriteNumber(int64(number), 2)
}

func (w *Writer) WriteSignedInt32(number int32) {
	_ = w.WriteNumber(int64(number), 4)
}

func (w *Writer) WriteUnsignedInt32(number uint32) {
	_ = w.WriteNumber(int64(number), 4)
}

func (w *Writer) WriteSignedLong(number int64) {
	_ = w.WriteNumber(number, 8)
}

func (w *Writer) WriteUnsignedLong(number uint64) {
	_ = w.WriteNumber(int64(number), 8)
}

func (w *Writer) WriteByteArray(n int, p []byte) {
	w.WriteByte(byte(n))
	w.Write(n, p)
}

func (w *Writer) WriteNumber(num int64, n int) error {
	if n < 1 || n > 8 {
		return errNumberBytes
	}

	p := make([]byte, n)

	j := uint(n) - 1
	i := j * 8
	// > 0, а не >= 0, потому что иначе происходит overflow
	for i > 0 {
		p[j] = byte((num >> i) & 0xFF)
		i -= 8
		j--
	}
	p[0] = byte(num & 0xFF)

	w.Write(n, p)

	return nil
}

func (w *Writer) Write(n int, p []byte) {
	w.buf = append(w.buf, p...)
	w.n += n
}

func (w *Writer) init(cap int) {
	w.n = 0
	w.buf = make([]byte, 0, cap)
}

func (w *Writer) Buffer() []byte {
	return w.buf
}
