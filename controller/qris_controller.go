package controller

import (
	"log"
	"net/http"
	"qris-api/config"
	"qris-api/model"

	"github.com/gin-gonic/gin"
)

var allowedChannel = map[string]string{
	"715": "Linkaja QRIS",
	"711": "ShopeePay QRIS",
	"842": "CIMB QRIS",
}

func CreateQRIS(c *gin.Context) {
	var req model.QRISRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[QRIS] Invalid Request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validasi channel code
	if _, ok := allowedChannel[req.AdditionalInfo.ChannelCode]; !ok {
		log.Printf("[QRIS] Invalid Channel Code: %s", req.AdditionalInfo.ChannelCode)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid channelCode"})
		return
	}

	//simulasi generate qr
	qrString := "QRIS:" + req.MerchantID + ":" + req.PartnerReferenceNo

	log.Printf("[QRIS] Generated QR for Merchant: %s, RefNo: %s", req.MerchantID, req.PartnerReferenceNo)

	c.JSON(http.StatusOK, gin.H{
		"message": "QRIS generated successfully",
		"data": gin.H{
			"qrString": qrString,
			"status":   "generated",
		},
	})

	data := model.QRISRequestModel{
		PartnerReferenceNo: req.PartnerReferenceNo,
		AmountValue:        req.Amount.Value,
		AmountCurrency:     req.Amount.Currency,
		MerchantID:         req.MerchantID,
		ValidityPeriod:     req.ValidityPeriod,
		BillDate:           req.AdditionalInfo.BillDate,
		BillDescription:    req.AdditionalInfo.BillDescription,
		ChannelCode:        req.AdditionalInfo.ChannelCode,
		PhoneNo:            req.AdditionalInfo.PhoneNo,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		log.Printf("[QRIS] Failed to save to DB: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save transaction"})
		return
	}
}
