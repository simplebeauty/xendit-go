package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/simplebeauty/xendit-go"
	"github.com/simplebeauty/xendit-go/cardlesscredit"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	createPaymentData := cardlesscredit.CreatePaymentParams{
		CardlessCreditType: xendit.CardlessCreditTypeEnumKREDIVO,
		ExternalID:         "cardless-credit-" + time.Now().String(),
		Amount:             200000,
		PaymentType:        xendit.PaymentTypeEnum3Months,
		Items: []cardlesscredit.Item{
			{
				ID:       "123",
				Name:     "Laptop Asus Ila",
				Price:    200000,
				Type:     "Laptop",
				URL:      "http://asus-ila.com",
				Quantity: 1,
			},
		},
		CustomerDetails: cardlesscredit.CustomerDetails{
			FirstName: "Michael",
			LastName:  "Belajarrock",
			Email:     "michaelbelajarrock@mail.com",
			Phone:     "08123123123",
		},
		ShippingAddress: cardlesscredit.ShippingAddress{
			FirstName:   "Michael",
			LastName:    "Belajarjazz",
			Address:     "Jalan Teknologi No. 12",
			City:        "Jakarta",
			PostalCode:  "40000",
			Phone:       "08123123123",
			CountryCode: "IDN",
		},
		RedirectURL: "https://google.com",
		CallbackURL: "https://google.com",
	}

	resp, err := cardlesscredit.CreatePayment(&createPaymentData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created payment: %+v\n", resp)
}
