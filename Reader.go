package binary

import (
	"errors"
)

type Reader struct {
	off int
	n   int
	buf []byte
}

func (r *Reader) ReadByte() (num byte, err error) {
	bb, err := r.Read(1)
	if err != nil {
		return
	}
	num = bb[0]
	return
}

func (r *Reader) ReadSignedShort() (num int16, err error) {
	n, err := r.ReadNumber(2)
	num = int16(n)
	return
}

func (r *Reader) ReadUnsignedShort() (num uint16, err error) {
	n, err := r.ReadNumber(2)
	num = uint16(n)
	return
}

func (r *Reader) ReadSignedInt32() (num int32, err error) {
	n, err := r.ReadNumber(4)
	num = int32(n)
	return
}

func (r *Reader) ReadUnsignedInt32() (num uint32, err error) {
	n, err := r.ReadNumber(4)
	num = uint32(n)
	return
}

func (r *Reader) ReadSignedLong() (num int64, err error) {
	num, err = r.ReadNumber(8)
	return
}

func (r *Reader) ReadUnsignedLong() (num uint64, err error) {
	n, err := r.ReadNumber(8)
	num = uint64(n)
	return
}

func (r *Reader) ReadByteArray() (n byte, ar []byte, err error) {
	n, err = r.ReadByte() //todo make varint
	if err != nil {
		return
	}

	ar, err = r.Read(int(n))
	if err != nil {
		r.off-- // Потому что ReadByte(), что вызывается выше, инкрементит offset
	}
	return
}

var errNumberBytes = errors.New("number can have 1/2/4/8 bytes")

func (r *Reader) ReadNumber(n int) (num int64, err error) {
	if n < 1 || n > 8 {
		return 0, errNumberBytes
	}

	p, err := r.Read(n)
	if err != nil {
		return 0, err
	}

	// Костыль, чтобы внутри for'а не делать все время приведение типа у переменной n
	unsignedN := uint(n) * 8

	// Странный момент. Ввел дополнительный j, чтобы лишний раз i не делить на 8
	i := uint(0)
	j := i
	for i < unsignedN {
		num |= int64(p[j]) << i
		i += 8
		j++
	}

	return
}

var errBytesAmount = errors.New("amount of bytes must be positive")
var errNotEnoughBytes = errors.New("not enough bytes in buffer")

func (r *Reader) Read(n int) ([]byte, error) {
	if n < 1 {
		return nil, errBytesAmount
	}

	if r.n < (n + r.off) {
		return nil, errNotEnoughBytes
	}

	r.off += n
	return r.buf[r.off-n : r.off], nil
}

func (r *Reader) RemainingAmount() int { return r.n - r.off }

func (r *Reader) Remaining() []byte { return r.buf[r.off:] }

func (r *Reader) Size() int {
	return r.n
}

func (r *Reader) init(buf []byte, n int) {
	r.n = n
	r.buf = buf
	r.off = 0
}
