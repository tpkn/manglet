package skip

import (
	"testing"
)

var in = "1,2,3,,...,,6,,8,8,9,0"

func TestSkipRows(t *testing.T) {
	l := Rows(in)
	if len(l) != 6 {
		t.Errorf("Should be 6 results instead of %v", l)
	}
}

func BenchmarkSkipRows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rows(in)
	}
}
