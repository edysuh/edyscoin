package libedyscoin

import (

)

// TODO sender/receipient secure data type
type Transaction struct {
	sender     string
	receipient string
	amount     int64
}
