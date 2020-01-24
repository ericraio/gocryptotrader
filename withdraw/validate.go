package withdraw

import (
	"errors"
	"strings"

	"github.com/thrasher-corp/gocryptotrader/currency"
)

// Valid takes interface and passes to asset type to check the request meets requirements to submit
func Valid(request *Request) (err error) {
	if request == nil {
		return ErrRequestCannotBeNil
	}

	var allErrors []string
	if request.Amount <= 0 {
		allErrors = append(allErrors, ErrStrAmountMustBeGreaterThanZero)
	}

	switch request.Type {
	case Fiat:
		if request.Fiat == nil {
			return ErrInvalidRequest
		}
		if (request.Currency != currency.Code{}) {
			if !request.Currency.IsFiatCurrency() {
				allErrors = append(allErrors, ErrStrCurrencyNotFiat)
			}
		}
		allErrors = append(allErrors, validateFiat(request)...)
	case Crypto:
		if request.Crypto == nil {
			return ErrInvalidRequest
		}
		if (request.Currency != currency.Code{}) {
			if !request.Currency.IsCryptocurrency() {
				allErrors = append(allErrors, ErrStrCurrencyNotCrypto)
			}
		}
		allErrors = append(allErrors, validateCrypto(request)...)
	default:
		allErrors = append(allErrors, "invalid request type")
	}
	if len(allErrors) > 0 {
		return errors.New(strings.Join(allErrors, ", "))
	}

	return nil
}

// Valid takes interface and passes to asset type to check the request meets requirements to submit
func validateFiat(request *Request) (err []string) {
	if request.Fiat.BankAccountNumber == "" {
		err = append(err, "Bank Account Number cannot be empty")
	}

	if request.Currency == currency.AUD {
		if request.Fiat.BSB == "" {
			err = append(err, "BSB must be set for AUD transfers")
		}
	} else {
		if request.Fiat.IBAN == "" && request.Fiat.SwiftCode == "" {
			err = append(err, "IBAN or Swift must be set")
		}
	}
	return err
}

// ValidateCrypto checks if Crypto request is valid and meets the minimum requirements to submit a crypto withdrawal request
func validateCrypto(request *Request) (err []string) {
	if request.Crypto.Address == "" {
		err = append(err, "Address cannot be empty")
	}

	if request.Crypto.FeeAmount < 0 {
		err = append(err, "FeeAmount cannot be a negative number")
	}

	return
}