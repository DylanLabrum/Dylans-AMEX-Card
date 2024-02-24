package main
import Golang

func main(){
	 type VirtualCard struct {
		Functionality struct {
			OnlineShopping struct {
				Stores        string
				OnlineServices bool
			}
			PaymentMethods struct {
				PayPal       bool
				AppleWallet  struct {
					ApplePay bool
				}
				SamsungPay    bool
				MobileWallets bool
			}
		}
	}
	
	// CreditCard struct represents credit card details
	type CreditCard struct {
		API        bool
		CardDetails struct {
			CardNumber string
			CVV        string
			Expiry     string
			CardName   string
		}
	}
	
	// Data struct represents the overall structure of the data
	type Data struct {
		VirtualCard VirtualCard
		CreditCard  CreditCard
	}
	
	func main() {
		// Initialize data with JSON-like structure
		data := Data{
			VirtualCard: VirtualCard{
				Functionality: struct {
					OnlineShopping struct {
						Stores        string
						OnlineServices bool
					}
					PaymentMethods struct {
						PayPal       bool
						AppleWallet  struct {
							ApplePay bool
						}
						SamsungPay    bool
						MobileWallets bool
					}
				}{
					OnlineShopping: struct {
						Stores        string
						OnlineServices bool
					}{
						Stores:        "Any",
						OnlineServices: true,
					},
					PaymentMethods: struct {
						PayPal       bool
						AppleWallet  struct {
							ApplePay bool
						}
						SamsungPay    bool
						MobileWallets bool
					}{
						PayPal: true,
						AppleWallet: struct {
							ApplePay bool
						}{
							ApplePay: true,
						},
						SamsungPay:    true,
						MobileWallets: true,
					},
				},
			},
			CreditCard: CreditCard{
				API: true,
				CardDetails: struct {
					CardNumber string
					CVV        string
					Expiry     string
					CardName   string
				}{
					CardNumber: "1143790885000828",
					CVV:        "503",
					Expiry:     "12/27",
					CardName:   "DylanLabrum",
				},
			},
		}
	
		// Accessing data elements
		fmt.Println("Card Number:", data.CreditCard.CardDetails.CardNumber)
		fmt.Println("CVV:", data.CreditCard.CardDetails.CVV)
		fmt.Println("Expiry Date:", data.CreditCard.CardDetails.Expiry)
		fmt.Println("Card Name:", data.CreditCard.CardDetails.CardName)
	}
	