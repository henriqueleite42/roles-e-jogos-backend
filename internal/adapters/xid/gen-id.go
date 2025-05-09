package xid

import "github.com/rs/xid"

func (self *xidAdapter) GenId() string {
	return xid.New().String()
}
