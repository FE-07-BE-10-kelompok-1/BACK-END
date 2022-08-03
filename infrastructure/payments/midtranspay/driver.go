package midtranspay

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func InitConnection(serverKey string) snap.Client {
	var s snap.Client
	s.New(serverKey, midtrans.Sandbox)
	return s
}
