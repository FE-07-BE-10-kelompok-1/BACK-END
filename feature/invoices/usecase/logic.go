package usecase

import "bookstore/domain"

type invoiceUsecase struct {
	invoiceData domain.InvoiceData
}

func New(id domain.InvoiceData) domain.InvoiceUsecase {
	return &invoiceUsecase{
		invoiceData: id,
	}
}

func (is *invoiceUsecase) CheckStocks(id []uint) ([]domain.Book, error) {
	data, err := is.invoiceData.CheckStocks(id)
	return data, err
}

func (is *invoiceUsecase) GetUserData(id uint) (domain.User, error) {
	data, err := is.invoiceData.GetUser(id)
	return data, err
}

func (is *invoiceUsecase) UpdateStock(id []uint) error {
	err := is.invoiceData.UpdateStock(id)
	return err
}

func (is *invoiceUsecase) InsertInvoice(data domain.Invoice, id []uint) (domain.Invoice, error) {
	invoice, err := is.invoiceData.Insert(data, id)
	return invoice, err
}

func (is *invoiceUsecase) DeleteCarts(userID uint) error {
	err := is.invoiceData.DeleteCarts(userID)
	return err
}
