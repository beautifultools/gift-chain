package wallet

type Wallet struct {
	public_key  string
	private_key string
	address     string
	amount      int
}

func GetWallet() Wallet {
	return Wallet{}
}
