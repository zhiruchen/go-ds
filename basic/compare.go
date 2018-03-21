package basic

// Comparer comparable data type
type Comparer interface {
	Equal(v Comparer) bool
	Less(v Comparer) bool
}
