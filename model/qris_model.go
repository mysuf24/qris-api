package model

import "gorm.io/gorm"

type QRISRequestModel struct {
	gorm.Model
	PartnerReferenceNo string
	AmountValue        string
	AmountCurrency     string
	MerchantID         string
	ValidityPeriod     string
	BillDate           string
	BillDescription    string
	ChannelCode        string
	PhoneNo            string
}
