package model

import "gorm.io/gorm"

type QRISRequest struct {
	gorm.Model
	PartnerReferenceNo string         `json:"partnerReferenceNo" binding:"required"`
	Amount             Amount         `json:"amount" binding:"required"`
	MerchantID         string         `json:"merchantId" binding:"required,len=5"`
	ValidityPeriod     string         `json:"validityPeriod" binding:"required"`
	AdditionalInfo     AdditionalInfo `json:"additionalInfo" binding:"required"`
}

type Amount struct {
	Value    string `json:"value" binding:"required"`
	Currency string `json:"currency" binding:"required,len=3"`
}

type AdditionalInfo struct {
	BillDate        string `json:"billDate" binding:"required"`
	BillDescription string `json:"billDescription" binding:"required"`
	ChannelCode     string `json:"channelCode" binding:"required"`
	PhoneNo         string `json:"phoneNo" binding:"required"`
}
