package portico

import (
	"context"
	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/builders/rebuilders"
	"github.com/globalpayments/go-sdk/api/entities/enums/inquirytype"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	"github.com/globalpayments/go-sdk/api/paymentmethods"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"testing"
)

func TestPorticoGiftTests(t *testing.T) {
	var card, track *paymentmethods.GiftCard

	config := serviceconfigs.NewPorticoConfig()
	config.SecretApiKey = "skapi_cert_MaePAQBr-1QAqjfckFC8FTbRTT120bVQUlfVOjgCBw"
	config.ServiceUrl = "https://cert.api2.heartlandportico.com"
	config.EnableLogging = true
	err := api.ConfigureService(config, "default")
	if err != nil {
		t.Errorf("Failed to configure service with error: %s", err.Error())
	}

	card = paymentmethods.NewGiftCard()
	card.SetNumber("5022440000000000007")

	track = paymentmethods.NewGiftCard()
	track.SetTrackData("%B5022440000000000098^^391200081613?;5022440000000000098=391200081613?")

	giftCreate(t)
	giftAddAlias(t, card)
	giftAddValue(t, card)
	giftBalanceInquiry(t, card)
	giftSale(t, card)
	giftDeactivate(t, card)
	giftRemoveAlias(t, card)
	giftReplace(t, card, track)
	giftReverse(t, card)
	giftRewards(t, card)
	giftTrackAddAlias(t, track)
	giftTrackAddValue(t, track)
	giftTrackBalanceInquiry(t, track)
	giftTrackSale(t, track)
	giftTrackDeactivate(t, track)
	giftTrackRemoveAlias(t, track)
	giftTrackReplace(t, track, card)
	giftTrackReverse(t, track)
	giftTrackRewards(t, track)
	giftReverseWithTransactionId(t, card)
}

func giftCreate(t *testing.T) {
	ctx := context.Background()
	builder := paymentmethods.CreateGiftCard("2145550199")
	transaction, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil {
		t.Errorf("Gift card creation failed with error: %s", err.Error())
		return
	}
	newCard := transaction.GetGiftCard()
	if newCard == nil {
		t.Errorf("Gift card not found")
		return
	}

	if newCard.GetNumber() == "" || newCard.GetAlias() == "" || newCard.GetPIN() == "" {
		t.Errorf("Gift card creation details missing")
	}
}

func giftAddAlias(t *testing.T, card *paymentmethods.GiftCard) {
	ctx := context.Background()
	builder := card.AddAlias("2145550199")
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Add alias failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for add alias: %s", response.GetResponseCode())
	}
}

func giftAddValue(t *testing.T, card *paymentmethods.GiftCard) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("10")
	builder := card.AddValueWithAmount(val)
	builder.WithCurrency("USD")
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Add value failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for add value: %s", response.GetResponseCode())
	}
}

func giftBalanceInquiry(t *testing.T, card *paymentmethods.GiftCard) {
	ctx := context.Background()
	builder := card.BalanceInquiry(inquirytype.Cash)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Balance inquiry failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for balance inquiry: %s", response.GetResponseCode())
	}
}

func giftSale(t *testing.T, card *paymentmethods.GiftCard) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("10")
	builder := card.ChargeWithAmount(val)
	builder.WithCurrency("USD")
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Sale failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for sale: %s", response.GetResponseCode())
	}
}

func giftDeactivate(t *testing.T, card *paymentmethods.GiftCard) {
	ctx := context.Background()
	builder := card.Deactivate()
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Deactivate failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for deactivate: %s", response.GetResponseCode())
	}
}

func giftRemoveAlias(t *testing.T, card *paymentmethods.GiftCard) {
	ctx := context.Background()
	builder := card.RemoveAlias("2145550199")
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Remove alias failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for remove alias: %s", response.GetResponseCode())
	}
}

func giftReplace(t *testing.T, card *paymentmethods.GiftCard, replacementCard *paymentmethods.GiftCard) {
	ctx := context.Background()
	builder := card.ReplaceWith(replacementCard)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Replace failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for replace: %s", response.GetResponseCode())
	}
}

func giftReverse(t *testing.T, card *paymentmethods.GiftCard) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("10")
	builder := card.ReverseWithAmount(val)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Reverse failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for reverse: %s", response.GetResponseCode())
	}
}

func giftRewards(t *testing.T, card *paymentmethods.GiftCard) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("10")
	builder := card.RewardsWithAmount(val)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Rewards failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for rewards: %s", response.GetResponseCode())
	}
}

func giftTrackAddAlias(t *testing.T, track *paymentmethods.GiftCard) {
	ctx := context.Background()
	builder := track.AddAlias("2145550199")
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Track add alias failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for track add alias: %s", response.GetResponseCode())
	}
}

func giftTrackAddValue(t *testing.T, track *paymentmethods.GiftCard) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("10")
	builder := track.AddValueWithAmount(val)
	builder.WithCurrency("USD")
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Track add value failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for track add value: %s", response.GetResponseCode())
	}
}

func giftTrackBalanceInquiry(t *testing.T, track *paymentmethods.GiftCard) {
	ctx := context.Background()
	builder := track.BalanceInquiry(inquirytype.Cash)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Track balance inquiry failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for track balance inquiry: %s", response.GetResponseCode())
	}
}

func giftTrackSale(t *testing.T, track *paymentmethods.GiftCard) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("10")
	builder := track.ChargeWithAmount(val)
	builder.WithCurrency("USD")
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Track sale failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for track sale: %s", response.GetResponseCode())
	}
}

func giftTrackDeactivate(t *testing.T, track *paymentmethods.GiftCard) {
	ctx := context.Background()
	builder := track.Deactivate()
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Track deactivate failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for track deactivate: %s", response.GetResponseCode())
	}
}

func giftTrackRemoveAlias(t *testing.T, track *paymentmethods.GiftCard) {
	ctx := context.Background()
	builder := track.RemoveAlias("2145550199")
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Track remove alias failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for track remove alias: %s", response.GetResponseCode())
	}
}

func giftTrackReplace(t *testing.T, track *paymentmethods.GiftCard, card *paymentmethods.GiftCard) {
	ctx := context.Background()
	builder := track.ReplaceWith(card)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Track replace failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for track replace: %s", response.GetResponseCode())
	}
}

func giftTrackReverse(t *testing.T, track *paymentmethods.GiftCard) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("10")
	builder := track.ReverseWithAmount(val)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Track reverse failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for track reverse: %s", response.GetResponseCode())
	}
}

func giftTrackRewards(t *testing.T, track *paymentmethods.GiftCard) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("10")
	builder := track.RewardsWithAmount(val)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Track rewards failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for track rewards: %s", response.GetResponseCode())
	}
}

func giftReverseWithTransactionId(t *testing.T, card *paymentmethods.GiftCard) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("10")
	builder := card.ChargeWithAmount(val)
	builder.WithCurrency("USD")
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil || response == nil {
		t.Errorf("Charge for reverse with transaction ID failed with error: %s", err.Error())
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected charge response code for reverse with transaction ID: %s", response.GetResponseCode())
		return
	}

	trans := rebuilders.FromId(response.GetTransactionId(), card.GetPaymentMethodType())
	revTrans := trans.ReverseWithAmount(val)
	reverseResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, revTrans)
	if err != nil || reverseResponse == nil {
		t.Errorf("Reverse with transaction ID failed with error: %s", err.Error())
		return
	}
	if reverseResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected reverse response code with transaction ID: %s", reverseResponse.GetResponseCode())
	}
}
