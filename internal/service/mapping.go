package service

import (
	"github.com/alanwade2001/go-sepa-iso/pacs_008_001_02"
	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"
	"github.com/jinzhu/copier"
)

type Mapper struct {
}

func NewMapper() *Mapper {

	mapper := &Mapper{}

	return mapper
}

func (m *Mapper) Map(pmtInf *pain_001_001_03.PaymentInstructionInformation3, ct10 *pain_001_001_03.CreditTransferTransactionInformation10) (*pacs_008_001_02.CreditTransferTransactionInformation11, error) {

	ct11 := &pacs_008_001_02.CreditTransferTransactionInformation11{}
	options := copier.Option{IgnoreEmpty: true, DeepCopy: true, FieldNameMapping: []copier.FieldNameMapping{
		{SrcType: ct10, DstType: ct11,
			Mapping: map[string]string{}}}}

	if err := copier.CopyWithOption(ct11, pmtInf, options); err != nil {
		return nil, err
	}

	if err := copier.CopyWithOption(ct11, ct10, options); err != nil {
		return nil, err
	}

	ct11.IntrBkSttlmAmt = &pacs_008_001_02.ActiveCurrencyAndAmount{}
	if err := copier.CopyWithOption(ct11.IntrBkSttlmAmt, ct10.Amt.InstdAmt, options); err != nil {
		return nil, err
	}

	ct11.IntrBkSttlmDt = pmtInf.ReqdExctnDt

	return ct11, nil
}
