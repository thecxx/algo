package search

import (
	"errors"
)

type Sunday struct {
	table [256]int
	i     bool
}

// New sunday engine.
func NewSunday() *Sunday {
	return &Sunday{}
}

func (s *Sunday) Search(haystack, needle []byte) (int, error) {
	// 1. Validate
	nl := len(needle)
	hl := len(haystack)
	if nl <= 0 || hl <= 0 || hl <= nl {
		return -1, errors.New("invalid length")
	}
	if !s.i {
		s.i = true
		// 2. Fill table
		err := s.fillTable(needle)
		if err != nil {
			return -1, err
		}
	}
	// 3. Search
	for i := 0; i < hl-nl+1; /* i++ */ {
		for j := 0; j < nl; j++ {
			if i+nl >= hl {
				break
			}
			if haystack[i+j] == needle[j] {
				continue
			}
			// Move
			i += s.table[haystack[i+nl]]
			goto LabelNotFound
		}
		return i, nil
	LabelNotFound:
	}
	return -1, nil
}

// Fill a table.
func (s *Sunday) fillTable(needle []byte) error {
	n := len(needle)
	if n <= 0 {
		return errors.New("invalid length")
	}
	for i := 0; i < 256; i++ {
		s.table[i] = n
	}
	for i := 0; i < n; i++ {
		s.table[needle[i]] = n - i
	}
	return nil
}
