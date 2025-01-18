package requests

type InvoiceLineCreateRequest struct {
	Name      string
	Amount    int64
	UnitPrice int64
}
