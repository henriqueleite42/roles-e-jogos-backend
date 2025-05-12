package xid

import (
	"github.com/rs/xid"
)

const TARGET_SESSION_ID_LENGTH = 128

func (self *xidAdapter) GenSessionId() string {
	fullId := xid.New().String() +
		xid.New().String() +
		xid.New().String() +
		xid.New().String() +
		xid.New().String() +
		xid.New().String() +
		xid.New().String()
	return fullId[:TARGET_SESSION_ID_LENGTH]
}
