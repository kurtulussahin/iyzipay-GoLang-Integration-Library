package iyzipay

type SubMerchant struct{}

func (subMerchant SubMerchant) Create(request CreateSubMerchantRequest, options Options) string {
	response := makeRequest(request, "POST", "/onboarding/submerchant", options)
	return response
}

func (subMerchant SubMerchant) Update(request UpdateSubMerchantRequest, options Options) string {
	response := makeRequest(request, "PUT", "/onboarding/submerchant", options)
	return response
}

func (subMerchant SubMerchant) Retrieve(request RetrieveSubMerchantRequest, options Options) string {
	response := makeRequest(request, "POST", "/onboarding/submerchant/detail", options)
	return response
}
