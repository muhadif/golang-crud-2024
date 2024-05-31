package error

var (
	ErrProductStock    = "product stock lower than request"
	ErrProductNotFound = "product not found"

	ErrCartIsEmpty  = "cart cannot empty"
	ErrCartNotFOund = "cart not found"

	ErrCheckoutCartIsEmpty = "need checkout first before making payment"
)
