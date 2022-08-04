package midtranspay

import (
	"bookstore/domain"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func InitConnection(serverKey string) snap.Client {
	var s snap.Client
	s.New(serverKey, midtrans.Sandbox)
	return s
}

func CreateTransactions(s snap.Client, token string, books []domain.Book, userData domain.User) (*snap.Response, error) {
	// 2. Initiate Snap request param
	var total int64
	var booksDetails []midtrans.ItemDetails
	for i := 0; i < len(books); i++ {
		convertID := strconv.Itoa(int(books[i].ID))
		booksDetails = append(booksDetails, midtrans.ItemDetails{
			ID:    convertID,
			Name:  books[i].Title,
			Price: int64(books[i].Price),
			Qty:   1,
		})
		total += int64(books[i].Price)
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  token,
			GrossAmt: int64(total),
		},
		EnabledPayments: []snap.SnapPaymentType{
			"gopay",
			"shopeepay",
			"bca_va",
			"bca_klikbca",
			"bca_klikpay",
			"Indomaret",
			"alfamart",
		},
		Items: &booksDetails,
		CustomerDetail: &midtrans.CustomerDetails{
			FName: userData.Fullname,
			Phone: userData.Phone,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, err := s.CreateTransaction(req)
	if err != nil {
		return snapResp, err
	}
	return snapResp, nil
}
