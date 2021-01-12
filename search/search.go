package search

type Searcher interface {
	// Find the position of the first occurrence of a sub content in `haystack`.
	Search(haystack, needle []byte) (int, error)
}
