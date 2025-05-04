package xid

import "github.com/rs/xid"

func (self *xidAdapter) GenId() (string, error) {
	return xid.New().String(), nil
}
