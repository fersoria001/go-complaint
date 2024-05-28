package queue

type Link[E any] struct {
	Element *E
	Next    *Link[E]
}

func NewLink[E any](options ...OptionsLinkFunc[E]) *Link[E] {
	l := new(Link[E])
	l.Element = nil
	l.Next = nil
	for _, option := range options {
		option(l)
	}
	return l
}

type OptionsLinkFunc[E any] func(*Link[E]) *Link[E]

func WithValue[E any](value *E) OptionsLinkFunc[E] {
	return func(l *Link[E]) *Link[E] {
		l.Element = value
		return l
	}
}
func WithNext[E any](next *Link[E]) OptionsLinkFunc[E] {
	return func(l *Link[E]) *Link[E] {
		l.Next = next
		return l
	}
}
