package xbytes

import (
	"bytes"
	"encoding/binary"
)

// Int32ToBytes Convert int32 to byte slices
// byte order: big endian or little endian, like binary.BigEndian or binary.LittleEndian
func Int32ToBytes(n int32, order binary.ByteOrder) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, order, n)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// BytesToInt32 Convert byte slices to int32
// byte order: big endian or little endian, like binary.BigEndian or binary.LittleEndian
func BytesToInt32(b []byte, order binary.ByteOrder) (int32, error) {
	var n int32
	buf := bytes.NewReader(b)
	err := binary.Read(buf, order, &n)
	if err != nil {
		return 0, err
	}
	return n, nil
}
