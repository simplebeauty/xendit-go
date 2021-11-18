package main

import (
	"fmt"
	"log"
	"time"

	"github.com/simplebeauty/xendit-go/promotion"
)

func promotionTest() {
	startTime := time.Now().Add(time.Hour)
	endTime := startTime.Add(time.Hour)

	created, err := promotion.CreatePromotion(&promotion.CreatePromotionParams{
		ReferenceID: fmt.Sprintf("promotion-%s", time.Now().String()),
		Description: "20% discount applied for all BRI cards",
		BinList: []string{
			"400000",
			"460000",
		},
		DiscountPercent:   20,
		Currency:          "IDR",
		StartTime:         &startTime,
		EndTime:           &endTime,
		ChannelCode:       "BRI",
		MinOriginalAmount: 25000,
		MaxDiscountAmount: 5000,
	})
	if err != nil {
		log.Panic(err)
	}

	_, err = promotion.GetPromotions(&promotion.GetPromotionsParams{
		ReferenceID: created.ReferenceID,
	})
	if err != nil {
		log.Panic(err)
	}

	updated, err := promotion.UpdatePromotion(&promotion.UpdatePromotionParams{
		PromotionID: created.ID,
		Description: "20% discount applied for all BCA cards",
		BinList: []string{
			"400000",
			"460000",
		},
		DiscountPercent: 20,
		Currency:        "IDR",
		ChannelCode:     "BCA",
	})
	if err != nil {
		log.Panic(err)
	}

	_, err = promotion.DeletePromotion(&promotion.DeletePromotionParams{
		PromotionID: updated.ID,
	})
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Promotion's integration tests done!")
}
