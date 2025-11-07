package constant

const (
	AVAILABLE    = "available"
	STOCK_OUT    = "stock_out"
	BY_REQUEST   = "by_request"
	DISCONTINUED = "discontinued"
	BOOKING      = "booking"
)

type OrderType struct{}

func (ot OrderType) OptionCodeNames() []string {
	return []string{
		AVAILABLE,
		STOCK_OUT,
		BY_REQUEST,
		DISCONTINUED,
		BOOKING,
	}
}
