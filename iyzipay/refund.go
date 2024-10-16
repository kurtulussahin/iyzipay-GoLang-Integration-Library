package iyzipay

type Refund struct{}

func (refund Refund) Create(request CreateRefundRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/refund", options)
	return response
}
