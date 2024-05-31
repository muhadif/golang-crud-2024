package error

var (
	ErrProductStock    = "product stock lower than request"
	ErrProductNotFound = "product not found"

	ErrCartIsEmpty = "cart cannot empty"

	ErrCheckoutCartIsEmpty = "need checkout first before making payment"
)
