package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/simplebeauty/xendit-go"
	"github.com/simplebeauty/xendit-go/card"
)

// To run this example, create the token first in https://js.xendit.co/test_tokenize.html
func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	createChargeData := card.CreateChargeParams{
		TokenID:          "5e2e81bbbae82e4d54d76473",
		AuthenticationID: "5e2e8322bae82e4d54d7648f",
		ExternalID:       "cardAuth-" + time.Now().String(),
		Amount:           1000000,
		Capture:          new(bool), // false
	}

	chargeResp, err := card.CreateCharge(&createChargeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created charge: %+v\n", chargeResp)

	getChargeData := card.GetChargeParams{
		ChargeID: chargeResp.ID,
	}

	chargeResp, err = card.GetCharge(&getChargeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved charge: %+v\n", chargeResp)

	captureChargeData := card.CaptureChargeParams{
		ChargeID: chargeResp.ID,
		Amount:   500000,
	}

	chargeResp, err = card.CaptureCharge(&captureChargeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("captured charge: %+v\n", chargeResp)

	createRefundData := card.CreateRefundParams{
		IdempotencyKey: "idempotency-" + time.Now().String(),
		ChargeID:       "5e2e8354d97c174c58bcf664",
		Amount:         10000,
		ExternalID:     "refund-" + time.Now().String(),
	}

	refundResp, err := card.CreateRefund(&createRefundData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("refunded charge: %+v\n", refundResp)

	// To run this part (ReverseAuthorization), the charge must not be captured first.
	reverseAuthorizationData := card.ReverseAuthorizationParams{
		ChargeID:   chargeResp.ID,
		ExternalID: "reverseAuth-" + time.Now().String(),
	}

	reverseAuthorizationResp, err := card.ReverseAuthorization(&reverseAuthorizationData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reversed authorization: %+v\n", reverseAuthorizationResp)
}
