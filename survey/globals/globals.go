package global

var LogginStatus chan int

const (
	NewLoggin = iota
	LoggedIn
	GetNewQR
	QRSaved
)
