<a href="https://github.com/globalpayments" target="_blank">
    <img src="https://developer.globalpay.com/static/media/logo.dab7811d.svg" alt="Global Payments logo" title="Global Payments" align="right" width="225" />
</a>

# Heartland & Global Payments Go SDK

This SDK makes it easy to integrate your Go application with our Card Not Present and Card Present APIs. 

## Solutions


### Card Present (Terminal & POS) Specific

* Minimize PCI Compliance requirements with Semi-Integrated / Out of Scope devices
* Secure End-To-End Encryption

## Requirements

- Go v1.18+
- github.com/shopspring/decimal v1.3.1+

## Installation


To install via [go get](https://pkg.go.dev/cmd/go/internal/get):

```
go get https://github.com/globalpayments/go-sdk
```

## Documentation and Examples

You can find the latest SDK documentation along with code examples and test cards on the [Global Payments](https://developer.realexpayments.com) and [Heartland](https://developer.heartlandpaymentsystems.com/documentation) Developer Hubs.

In addition you can find working examples in the our example code repository.

*Quick Tip*: The included [test suite](https://github.com/globalpayments/go-sdk/tree/master/tests/integration) can be a great source of code samples for using the SDK!

#### Process a Payment over UPA Example

```go
    terminal, err := device.CreditSale(val)
	if err != nil {
		// handle errors
	}
	response, err := api.ExecuteTerminal(terminal)
	if err != nil {
		// handle errors
	}

```

#### Test Card Data

Name        | Number           | Exp Month | Exp Year | CVN
----------- | ---------------- | --------- | -------- | ----
Visa        | 4263970000005262 | 12        | 2025     | 123
MasterCard  | 2223000010005780 | 12        | 2019     | 900
MasterCard  | 5425230000004415 | 12        | 2025     | 123
Discover    | 6011000000000087 | 12        | 2025     | 123
Amex        | 374101000000608  | 12        | 2025     | 1234
JCB         | 3566000000000000 | 12        | 2025     | 123
Diners Club | 36256000000725   | 12        | 2025     | 123

#### Testing Exceptions

During your integration you will want to test for specific issuer responses such as 'Card Declined'. Because our sandbox environments do not actually reach out to issuing banks for authorizations, there are specific transaction amounts and/or card numbers that will trigger gateway and issuing bank responses. Please contact your support representative for a complete listing of values used to simulate transaction AVS/CVV results, declines, errors, and other responses that can be caught in your code. 

## Contributing

All our code is open sourced and we encourage fellow developers to contribute and help improve it!

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Ensure SDK tests are passing
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create new Pull Request

## License

This project is licensed under the GNU General Public License v2.0. Please see [LICENSE.md](LICENSE.md) located at the project's root for more details.