package service

import (
	"log"

	"github.com/alanwade2001/go-sepa-engine-data/model"
	"github.com/alanwade2001/go-sepa-engine-data/repository"
	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
)

type Preparer struct {
	ghRepos  *repository.PaymentGroup
	pmtRepos *repository.Payment
	txRepos  *repository.Transaction
	delivery *Delivery
}

func NewPreparer(ghRepos *repository.PaymentGroup, pmtRepos *repository.Payment, txRepos *repository.Transaction, delivery *Delivery) *Preparer {
	initiation := &Preparer{
		ghRepos:  ghRepos,
		pmtRepos: pmtRepos,
		txRepos:  txRepos,
		delivery: delivery,
	}

	return initiation
}

func (s *Preparer) Prepare(mdl *model.PaymentGroup) (err error) {

	log.Printf("prepare: [%v]", mdl)

	if pmts, err := s.pmtRepos.GetPaymentsByPaymentGroupID(mdl.ID); err != nil {
		return err
	} else {
		for _, pmt := range pmts {
			if err = s.PreparePayment(pmt); err != nil {
				return nil
			}
		}
	}

	// if err = s.delivery.PaymentGroupIngested(mdl); err != nil {
	// 	return err
	// }

	return nil
}

func (s *Preparer) PreparePayment(pmt *entity.Payment) error {

	log.Printf("ID=[%d]; PmtInfId=[%s]", pmt.Model.ID, pmt.PmtInfID)

	if txs, err := s.txRepos.GetTransactionsByPaymentID(pmt.Model.ID); err != nil {
		return err
	} else {
		for _, tx := range txs {
			if err = s.PrepareTransaction(tx); err != nil {
				return nil
			}
		}
	}

	return nil
}

func (s *Preparer) PrepareTransactions(txs []*entity.Transaction) error {
	for _, tx := range txs {
		if err := s.PrepareTransaction(tx); err != nil {
			return nil
		}
	}
	return nil
}

func (s *Preparer) PrepareTransaction(tx *entity.Transaction) error {
	log.Printf("ID=[%d]; EndToEndId=[%s]", tx.Model.ID, tx.EndToEndID)
	return nil
}
