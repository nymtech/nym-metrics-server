package metrics

type Sanitizer interface {
	Sanitize(input string) string
}

type sanitizer struct{}

// func (s sanitizer) Sanitize(input string) string {

// }
