package promotion

import (
	"context"

	"github.com/simplebeauty/xendit-go"
)

// CreatePromotion creates new promotion.
func CreatePromotion(data *CreatePromotionParams) (*xendit.Promotion, *xendit.Error) {
	return CreatePromotionWithContext(context.Background(), data)
}

// CreatePromotionWithContext creates new promotion with context.
func CreatePromotionWithContext(ctx context.Context, data *CreatePromotionParams) (*xendit.Promotion, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreatePromotionWithContext(ctx, data)
}

// GetPromotions gets promotions.
func GetPromotions(data *GetPromotionsParams) ([]xendit.Promotion, *xendit.Error) {
	return GetPromotionsWithContext(context.Background(), data)
}

// GetPromotionsWithContext gets promotions with context.
func GetPromotionsWithContext(ctx context.Context, data *GetPromotionsParams) ([]xendit.Promotion, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetPromotionsWithContext(ctx, data)
}

// UpdatePromotion updates a promotion.
func UpdatePromotion(data *UpdatePromotionParams) (*xendit.Promotion, *xendit.Error) {
	return UpdatePromotionWithContext(context.Background(), data)
}

// UpdatePromotionWithContext updates a promotion with context.
func UpdatePromotionWithContext(ctx context.Context, data *UpdatePromotionParams) (*xendit.Promotion, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.UpdatePromotionWithContext(ctx, data)
}

// DeletePromotion deletes a promotion.
func DeletePromotion(data *DeletePromotionParams) (*xendit.PromotionDeletion, *xendit.Error) {
	return DeletePromotionWithContext(context.Background(), data)
}

// DeletePromotionWithContext deletes a promotion with context.
func DeletePromotionWithContext(ctx context.Context, data *DeletePromotionParams) (*xendit.PromotionDeletion, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.DeletePromotionWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
