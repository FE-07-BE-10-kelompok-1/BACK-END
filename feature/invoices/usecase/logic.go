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
