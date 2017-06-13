package cha

var table = map[int]byte{
	0x0: '0',
	0x1: '1',
	0x2: '2',
	0x3: '3',
	0x4: '4',
	0x5: '5',
	0x6: '6',
	0x7: '7',
	0x8: '8',
	0x9: '9',
	0xa: 'a',
	0xb: 'b',
	0xc: 'c',
	0xd: 'd',
	0xe: 'e',
	0xf: 'f',
}

func Hash(str []byte) []byte {
	raw := 274777
	hash := []byte{}

	for i := 0; i < len(str); i++ {
		raw = raw*int(str[i]) + 33
	}

	for i := 0; i < 6; i++ {
		rem := raw % 16
		raw = raw / 16
		hash = append(hash, table[rem])
	}

	return hash
}
