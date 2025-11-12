package constant

const (
	SerialNumberIN  = "in"
	SerialNumberOUT = "out"
)

type SerialNumberType struct{}

func (snType SerialNumberType) OptionCodeNames() []string {
	return []string{
		SerialNumberIN,
		SerialNumberOUT,
	}
}
