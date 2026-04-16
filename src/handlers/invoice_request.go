package handlers

type InvoiceItemRequest struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type CreateInvoiceRequest struct {
	SenderName      string               `json:"sender_name"`
	SenderAddress   string               `json:"sender_address"`
	ReceiverName    string               `json:"receiver_name"`
	ReceiverAddress string               `json:"receiver_address"`
	Items           []InvoiceItemRequest `json:"items"`
}