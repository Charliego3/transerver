package enums

//go:generate go run github.com/dmarkham/enumer -type=UserState -trimprefix=User -sql -values -output='userstate_string.go'
type UserState uint8

const (
	UserUnverified UserState = iota
	UserNormal
)
