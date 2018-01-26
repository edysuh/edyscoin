package libedyscoin

import (

)

// TODO sender/recipient secure data type
type Transaction struct {
	Sender     string
	Recipient  string
	Amount     int64
}
