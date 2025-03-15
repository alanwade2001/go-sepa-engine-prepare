package service

import (
	"log"

	"github.com/alanwade2001/go-sepa-engine-data/model"
	"github.com/alanwade2001/go-sepa-engine-data/repository"
)

type Preparer struct {
	ghRepos  *repository.PaymentGroup
	pmtRepos *repository.Payment
	txRepos  *repository.CreditTransfer
	delivery *Delivery
}

func NewPreparer(ghRepos *repository.PaymentGroup, pmtRepos *repository.Payment, txRepos *repository.CreditTransfer, delivery *Delivery) *Preparer {
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

	// if err = s.delivery.PaymentGroupIngested(mdl); err != nil {
	// 	return err
	// }

	return nil
}
