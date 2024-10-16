package iyzipay

type CheckoutForm struct{}

func (checkoutForm CheckoutForm) Retrieve(request RetrieveCheckoutFormRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/iyzipos/checkoutform/auth/ecom/detail", options)
	return response
}
