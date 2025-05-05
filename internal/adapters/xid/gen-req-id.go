package xid

import "github.com/rs/xid"

func (self *xidAdapter) GenReqId() string {
	return xid.New().String()
}
