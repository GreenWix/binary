package binary

import (
	"bytes"
	"testing"
)

func getWriter() *Writer {
	w := &Writer{}
	clean(w)
	return w
}

func clean(w *Writer) {
	w.n = 0
	w.buf = make([]byte, 0, 16)
}

func TestWriter_WriteNumber(t *testing.T) {
	wrongSize := 9
	normalSize := 4
	checkingNum := int64(0x7FFFFFFF)
	numBytes := []byte{0xFF, 0xFF, 0xFF, 0x7F}

	bs := getWriter()

	err := bs.WriteNumber(checkingNum, wrongSize)
	if err == nil {
		t.Errorf("There must be an error with n = %d", wrongSize)
		clean(bs)
		return
	}

	err = bs.WriteNumber(checkingNum, normalSize)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if !bytes.Equal(numBytes, bs.buf) {
		t.Errorf("Wrong bytes. Expected %s, got %s", string(numBytes), string(bs.buf))
	}
}

func TestWriter_WriteByte(t *testing.T) {
	checkingNum := byte(0xFF)
	numBytes := []byte{0xFF}

	bs := getWriter()

	bs.WriteByte(checkingNum)
	if !bytes.Equal(numBytes, bs.buf) {
		t.Errorf("Wrong bytes. Expected %s, got %s", string(numBytes), string(bs.buf))
	}

}

func TestWriter_WriteSignedShort(t *testing.T) {
	checkingNum := int16(0x7FAA)
	numBytes := []byte{0xAA, 0x7F}

	bs := getWriter()

	bs.WriteSignedShort(checkingNum)
	if !bytes.Equal(numBytes, bs.buf) {
		t.Errorf("Wrong bytes. Expected %s, got %s", string(numBytes), string(bs.buf))
	}
}

func TestWriter_WriteUnsignedShort(t *testing.T) {
	checkingNum := uint16(0xFFAA)
	numBytes := []byte{0xAA, 0xFF}

	bs := getWriter()

	bs.WriteUnsignedShort(checkingNum)
	if !bytes.Equal(numBytes, bs.buf) {
		t.Errorf("Wrong bytes. Expected %s, got %s", string(numBytes), string(bs.buf))
	}
}

func TestWriter_WriteSignedInt32(t *testing.T) {
	checkingNum := int32(0x7FAABBCC)
	numBytes := []byte{0xCC, 0xBB, 0xAA, 0x7F}

	bs := getWriter()

	bs.WriteSignedInt32(checkingNum)
	if !bytes.Equal(numBytes, bs.buf) {
		t.Errorf("Wrong bytes. Expected %s, got %s", string(numBytes), string(bs.buf))
	}
}

func TestWriter_WriteUnsignedInt32(t *testing.T) {
	checkingNum := uint32(0xFFAABBCC)
	numBytes := []byte{0xCC, 0xBB, 0xAA, 0xFF}

	bs := getWriter()

	bs.WriteUnsignedInt32(checkingNum)
	if !bytes.Equal(numBytes, bs.buf) {
		t.Errorf("Wrong bytes. Expected %s, got %s", string(numBytes), string(bs.buf))
	}
}

func TestWriter_WriteSignedLong(t *testing.T) {
	checkingNum := int64(0x7FAABBCCDDEEFF11)
	numBytes := []byte{0x11, 0xFF, 0xEE, 0xDD, 0xCC, 0xBB, 0xAA, 0x7F}

	bs := getWriter()

	bs.WriteSignedLong(checkingNum)
	if !bytes.Equal(numBytes, bs.buf) {
		t.Errorf("Wrong bytes. Expected %s, got %s", string(numBytes), string(bs.buf))
	}
}

func TestWriter_WriteUnsignedLong(t *testing.T) {
	checkingNum := uint64(0xFFAABBCCDDEE1122)
	numBytes := []byte{0x22, 0x11, 0xEE, 0xDD, 0xCC, 0xBB, 0xAA, 0xFF}

	bs := getWriter()

	bs.WriteUnsignedLong(checkingNum)
	if !bytes.Equal(numBytes, bs.buf) {
		t.Errorf("Wrong bytes. Expected %s, got %s", string(numBytes), string(bs.buf))
	}
}

func TestWriter_WriteByteArray(t *testing.T) {
	byteAr := []byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}
	byteLen := len(byteAr)
	expectedAr := append([]byte{byte(byteLen)}, byteAr...)

	bs := getWriter()
	bs.WriteByteArray(byteLen, byteAr)

	if !bytes.Equal(expectedAr, bs.buf) {
		t.Errorf("Wrong bytes. Expected %s, got %s", string(expectedAr), string(bs.buf))
	}
}
