package quote

type CoinEntity struct {
	Coin
	ID uint `gorm:"primaryKey" json:"id"`
}
