package enums

//go:generate go run github.com/dmarkham/enumer -type=UserState -sql -values
type UserState uint8

const (
	UserUnverified UserState = iota
	Normal
)
