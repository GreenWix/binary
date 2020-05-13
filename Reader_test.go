package binary

import (
	"bytes"
	"fmt"
	"testing"
)

func getReader(p []byte) *Reader {
	return &Reader{
		off: 0,
		n:   len(p),
		buf: p,
	}
}

func TestReader_ReadNumber(t *testing.T) {
	wrongSize := 9
	normalSize := 4
	checkingNum := int32(0x7FFFFFFF)
	numBytes := []byte{0xFF, 0xFF, 0xFF, 0x7F}

	bs := getReader(numBytes)

	_, err := bs.ReadNumber(wrongSize)
	if err == nil {
		t.Errorf("There must be an error with n = %d", wrongSize)
		bs.off = 0
		return
	}

	num, err := bs.ReadNumber(normalSize)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if int32(num) != checkingNum {
		t.Errorf("Wrong number. Expected %x, got %x", checkingNum, num)
	}

}

func TestReader_ReadByte(t *testing.T) {
	checkingNum := byte(0xFF)
	numBytes := []byte{0xFF}

	bs := getReader(numBytes)

	num, err := bs.ReadByte()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
		return
	}
	if num != checkingNum {
		t.Errorf("Wrong number. Expected %x, got %x", checkingNum, num)
	}
}

func TestReader_ReadSignedShort(t *testing.T) {
	checkingNum := int16(0x7FAA)
	numBytes := []byte{0xAA, 0x7F}

	bs := getReader(numBytes)

	num, err := bs.ReadSignedShort()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if num != checkingNum {
		t.Errorf("Wrong number. Expected %x, got %x", checkingNum, num)
	}
}

func TestReader_ReadUnsignedShort(t *testing.T) {
	checkingNum := uint16(0xFFAA)
	numBytes := []byte{0xAA, 0xFF}

	bs := getReader(numBytes)

	num, err := bs.ReadUnsignedShort()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if num != checkingNum {
		t.Errorf("Wrong number. Expected %x, got %x", checkingNum, num)
	}
}

func TestReader_ReadSignedInt32(t *testing.T) {
	checkingNum := int32(0x7FAABBCC)
	numBytes := []byte{0xCC, 0xBB, 0xAA, 0x7F}

	bs := getReader(numBytes)

	num, err := bs.ReadSignedInt32()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if num != checkingNum {
		t.Errorf("Wrong number. Expected %x, got %x", checkingNum, num)
	}
}

func TestReader_ReadUnsignedInt32(t *testing.T) {
	checkingNum := uint32(0xFFAABBCC)
	numBytes := []byte{0xCC, 0xBB, 0xAA, 0xFF}

	bs := getReader(numBytes)

	num, err := bs.ReadUnsignedInt32()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if num != checkingNum {
		t.Errorf("Wrong number. Expected %x, got %x", checkingNum, num)
	}
}

func TestReader_ReadSignedLong(t *testing.T) {
	checkingNum := int64(0x7FAABBCCDDEEFF11)
	numBytes := []byte{0x11, 0xFF, 0xEE, 0xDD, 0xCC, 0xBB, 0xAA, 0x7F}

	bs := getReader(numBytes)

	num, err := bs.ReadSignedLong()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if num != checkingNum {
		t.Errorf("Wrong number. Expected %x, got %x", checkingNum, num)
	}
}

func TestReader_ReadUnsignedLong(t *testing.T) {
	checkingNum := uint64(0xFFAABBCCDDEE1122)
	numBytes := []byte{0x22, 0x11, 0xEE, 0xDD, 0xCC, 0xBB, 0xAA, 0xFF}

	bs := getReader(numBytes)

	num, err := bs.ReadUnsignedLong()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if num != checkingNum {
		t.Errorf("Wrong number. Expected %x, got %x", checkingNum, num)
	}
}

func TestReader_ReadByteArray(t *testing.T) {
	byteAr := []byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}
	byteLen := byte(len(byteAr))
	wrongByteLen := byteLen + 1

	bs := getReader([]byte{wrongByteLen})
	bs.n += int(byteLen)
	bs.buf = append(bs.buf, byteAr...)

	n, ar, err := bs.ReadByteArray()
	if err == nil {
		fmt.Println(n, ar)
		t.Errorf("There must be an error with wrong byte array size")
		bs.off = 0

	}

	bs.buf[0] = byteLen
	n, ar, err = bs.ReadByteArray()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if n != byteLen {
		t.Errorf("Wrong size array. Expected %d, got %d", byteLen, n)
	}

	if !bytes.Equal(byteAr, ar) {
		t.Errorf("Wrong byte array. Expected %s, got %s", string(byteAr), string(ar))
	}

}
