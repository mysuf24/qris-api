package model

type Amount struct {
	Value    string `json:"value" binding:"required"`
	Currency string `json:"currency" binding:"required, len=3"`
}

type AdditionalInfo struct {
	BillDate        string `json:"billDate" binding:"required"`
	BillDescription string `json:"billDescription" binding:"required, max=128"`
	ChannelCode     string `json:"channelCode" binding:"required"`
	PhoneNo         string `json:"phoneNo" binding:"required, max=30"`
}

type QRISRequest struct {
	PartnerReferenceNo string         `json:"partnerReferenceNo" binding:"required,max=32"`
	Amount             Amount         `json:"amount" binding:"required"`
	MerchantID         string         `json:"merchantId" binding:"required,max=5"`
	ValidityPeriod     string         `json:"validityPeriod" binding:"required"`
	AdditionalInfo     AdditionalInfo `json:"additionalInfo" binding:"required"`
}
