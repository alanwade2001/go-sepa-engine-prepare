package service

import (
	"encoding/xml"

	"github.com/alanwade2001/go-sepa-engine-data/model"
	"github.com/alanwade2001/go-sepa-engine-data/repository"
	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/exp/slog"
)

type Preparer struct {
	reposMgr *repository.Manager
	mapper   *Mapper
	delivery *Delivery
}

func NewPreparer(reposMgr *repository.Manager, delivery *Delivery) *Preparer {
	initiation := &Preparer{
		reposMgr: reposMgr,
		delivery: delivery,
		mapper:   NewMapper(),
	}

	return initiation
}

func (s *Preparer) Prepare(mdl *model.PaymentGroup) (err error) {

	slog.Info("prepare: [%v]", mdl)

	if pmts, err := s.reposMgr.Payment.GetPaymentsByPaymentGroupID(mdl.ID); err != nil {
		return err
	} else {
		for _, pmt := range pmts {
			if err = s.PreparePayment(pmt); err != nil {
				return nil
			}
		}
	}

	return nil
}

func (s *Preparer) PreparePayment(pmt *entity.Payment) error {

	slog.Info("ID=[%d]; PmtInfId=[%s]", pmt.Model.ID, pmt.PmtInfID)
	pi3 := &pain_001_001_03.PaymentInstructionInformation3{}

	if err := xml.Unmarshal([]byte(pmt.PmtInf), pi3); err != nil {
		return err
	} else if txs, err := s.reposMgr.Transaction.GetTransactionsByPaymentID(pmt.Model.ID); err != nil {
		return err
	} else {
		if err = s.PrepareTransactions(pi3, txs); err != nil {
			return nil
		}
	}

	pMdl := &model.Payment{}
	pMdl.FromEntity(pmt)

	if err := s.delivery.PaymentPrepared(pMdl); err != nil {
		return err
	}

	return nil
}

func (s *Preparer) PrepareTransactions(pi3 *pain_001_001_03.PaymentInstructionInformation3, txs []*entity.Transaction) error {
	for _, tx := range txs {
		if err := s.PrepareTransaction(pi3, tx); err != nil {
			return nil
		}
	}
	return nil
}

func (s *Preparer) PrepareTransaction(pi3 *pain_001_001_03.PaymentInstructionInformation3, tx *entity.Transaction) error {
	slog.Info("prepare transaction", slog.Uint64("ID", uint64(tx.Model.ID)), slog.String("EndToEndId", tx.EndToEndID))

	ct10Text := tx.CdtTrfTxInf
	ct10 := &pain_001_001_03.CreditTransferTransactionInformation10{}
	if err := xml.Unmarshal([]byte(ct10Text), ct10); err != nil {
		slog.Error("failed to unmarshal ct10", "err", err)
		return err
	}

	if ct11, err := s.mapper.Map(pi3, ct10); err != nil {
		slog.Error("failed to map pi3 & ct10", "err", err)
		return err
	} else if cdtTrfTxInf, err := xml.Marshal(ct11); err != nil {
		slog.Error("failed to marshal ct11", "err", err)
		return err
	} else {
		uid := uuid.NewV4()

		settlement := &entity.Settlement{
			EndToEndID:    tx.EndToEndID,
			CdtTrfTxInf:   string(cdtTrfTxInf),
			TxID:          uid.String(),
			TransactionID: tx.Model.ID,
			Transaction:   tx,
		}

		slog.Info("persisting", "txID", settlement.TxID)

		if _, err := s.reposMgr.Settlement.Perist(settlement); err != nil {
			return err
		}

		slog.Info("persisted", "ID", settlement.Model.ID)
	}

	return nil
}
