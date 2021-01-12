package search

import (
	"testing"
)

func TestSunday_Search(t *testing.T) {
	s := NewSunday()
	index, err := s.Search([]byte("0123456789"), []byte("91"))
	if err != nil {
		t.Errorf("search failed, error = %s\n", err)
		return
	}
	if index < 0 {
		t.Errorf("not found")
		return
	}
	t.Logf("index = %d\n", index)
}

func BenchmarkSunday_Search(b *testing.B) {
	s := NewSunday()
	for i := 0; i < 100000; i++ {
		s.Search([]byte("0123456789abcdefghijklmnopqrstuvwxyz"), []byte("jklmnopqr"))
	}
}
