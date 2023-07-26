package crc

import (
	"fmt"
	"testing"
)

func TestCrc16(t *testing.T) {
	const (
		poly = 0x1021
		init = 0xFFFF
	)

	var crcTable [256]uint16
	var crcTableHi [256]uint8

	for i := 0; i < 256; i++ {
		crc := uint16(i)
		for j := 0; j < 8; j++ {
			if crc&1 == 1 {
				crc = (crc >> 1) ^ poly
			} else {
				crc >>= 1
			}
		}
		crcTable[i] = crc
		crcTableHi[i] = uint8(crc >> 8)
	}

	for _, v := range crcTable {
		fmt.Printf("0x%X, ", v)
	}
}
