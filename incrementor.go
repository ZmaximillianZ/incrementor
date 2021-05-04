package incrementor

// defaultMaximumValue default value for maximumValue
const defaultMaximumValue uint = 99

// I represents contract of Incrementor
type I interface {
	GetNumber() uint
	IncrementNumber()
	SetMaximumValue(maximumValue uint)
}

// Incrementor is used to increment a number
type Incrementor struct {
	number       uint
	maximumValue uint
}

// New creates new struct
func New() Incrementor {
	return Incrementor{0, defaultMaximumValue}
}

// GetNumber returns current value. As default its 0
func (incr *Incrementor) GetNumber() uint {
	return incr.number
}

// IncrementNumber increases the current number by one.
// After each call of this getNumber () method will return one more number.
func (incr *Incrementor) IncrementNumber() {
	if incr.checkIsNumberInLimit() {
		incr.number = 0

		return
	}
	incr.number++
}

// SetMaximumValue Sets the maximum value of the current number.
// The stored number cannot exceed the set maximum value.
// When, when calling incrementNumber (), the current number reaches
// this value, it is zeroed, i.e. getNumber () starts
// return zero again, and again one after the next
// calling incrementNumber (), and so on.
// The default maximum is defaultMaximumValue the maximum int value.
// You cannot be allowed to set a number here less than zero.
func (incr *Incrementor) SetMaximumValue(maximumValue uint) {
	incr.maximumValue = maximumValue
}

// checkIsNumberInLimit checks if the number has reached the limit
func (incr *Incrementor) checkIsNumberInLimit() bool {
	return incr.number >= incr.maximumValue
}
