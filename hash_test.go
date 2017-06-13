package cha

import (
	"testing"
)

func TestLength(t *testing.T) {
	tc := [][]byte{
		[]byte{},
		[]byte("foo"),
		[]byte("baz"),
		[]byte("0"),
		[]byte{0},
		[]byte{0, 1, 2},
		[]byte("supercalifragalisticexpialidocious"),
		[]byte(`Four days will quickly steep themselves in night;
Four nights will quickly dream away the time;
And then the moon, like to a silver bow
New-bent in heaven, shall behold the night
Of our solemnities.`),
	}

	for _, c := range tc {
		got := Hash(c)
		if len(got) != 6 {
			t.Errorf("Expected hash output to be 6 bytes in length, but got: %s", got)
		}
	}
}
