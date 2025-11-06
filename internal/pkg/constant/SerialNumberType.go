package constant

type SerialNumberTypeStruct struct {
	IN      string
	OUT     string
	options []string
}

var SerialNumberType = func() SerialNumberTypeStruct {
	sn := SerialNumberTypeStruct{
		IN:  "in",
		OUT: "out",
	}

	sn.options = []string{
		sn.IN,
		sn.OUT,
	}

	return sn
}()

func (sn SerialNumberTypeStruct) OPTION() []string {
	return append([]string(nil), sn.options...)
}
