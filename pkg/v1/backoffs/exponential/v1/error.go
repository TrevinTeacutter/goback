package exponential

var _ error = basic("")

type basic string

func (c basic) Error() string {
	return string(c)
}
