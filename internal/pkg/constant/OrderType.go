package constant

type OrderTypeStruct struct {
	AVAILABLE    string
	STOCK_OUT    string
	BY_REQUEST   string
	DISCONTINUED string
	BOOKING      string
	options      []string
}

var OrderType = func() OrderTypeStruct {
	ot := OrderTypeStruct{
		AVAILABLE:    "available",
		STOCK_OUT:    "stock_out",
		BY_REQUEST:   "by_request",
		DISCONTINUED: "discontinued",
		BOOKING:      "booking",
	}

	ot.options = []string{
		ot.AVAILABLE,
		ot.STOCK_OUT,
		ot.BY_REQUEST,
		ot.DISCONTINUED,
		ot.BOOKING,
	}

	return ot
}()

func (ot OrderTypeStruct) OPTION() []string {
	return append([]string(nil), ot.options...)
}
