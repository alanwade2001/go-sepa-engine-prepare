package service

import (
	"testing"

	"github.com/alanwade2001/go-sepa-iso/pacs_008_001_02"
	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		mapper   *Mapper
		ct10     *pain_001_001_03.CreditTransferTransactionInformation10
		pi3      *pain_001_001_03.PaymentInstructionInformation3
		expected *pacs_008_001_02.CreditTransferTransactionInformation11
	}{
		{
			desc:   "",
			mapper: NewMapper(),
			ct10: &pain_001_001_03.CreditTransferTransactionInformation10{
				PmtId:    &pain_001_001_03.PaymentIdentification1{EndToEndId: "e2e-1"},
				Amt:      &pain_001_001_03.AmountType3Choice{InstdAmt: &pain_001_001_03.ActiveOrHistoricCurrencyAndAmount{CcyAttr: "EUR", Value: 10.01}},
				Cdtr:     &pain_001_001_03.PartyIdentification32{Nm: "c-1"},
				CdtrAgt:  &pain_001_001_03.BranchAndFinancialInstitutionIdentification4{FinInstnId: &pain_001_001_03.FinancialInstitutionIdentification7{BIC: "cbic-1"}},
				CdtrAcct: &pain_001_001_03.CashAccount16{Id: &pain_001_001_03.AccountIdentification4Choice{IBAN: "ciban-1"}},
				RmtInf:   &pain_001_001_03.RemittanceInformation5{Ustrd: []string{"rmtinf-1"}},
			},
			pi3: &pain_001_001_03.PaymentInstructionInformation3{
				PmtInfId:    "pmtinfid-1",
				NbOfTxs:     "1",
				CtrlSum:     100.01,
				ReqdExctnDt: "2025-03-16",
				Dbtr:        &pain_001_001_03.PartyIdentification32{Nm: "dnm-1"},
				DbtrAgt:     &pain_001_001_03.BranchAndFinancialInstitutionIdentification4{FinInstnId: &pain_001_001_03.FinancialInstitutionIdentification7{BIC: "dbic-1"}},
				DbtrAcct:    &pain_001_001_03.CashAccount16{Id: &pain_001_001_03.AccountIdentification4Choice{IBAN: "diban-1"}},
			},
			expected: &pacs_008_001_02.CreditTransferTransactionInformation11{
				PmtId:          &pacs_008_001_02.PaymentIdentification3{EndToEndId: "e2e-1"},
				IntrBkSttlmAmt: &pacs_008_001_02.ActiveCurrencyAndAmount{CcyAttr: "EUR", Value: 10.01},
				IntrBkSttlmDt:  "2025-03-16",
				Cdtr:           &pacs_008_001_02.PartyIdentification32{Nm: "c-1"},
				CdtrAgt:        &pacs_008_001_02.BranchAndFinancialInstitutionIdentification4{FinInstnId: &pacs_008_001_02.FinancialInstitutionIdentification7{BIC: "cbic-1"}},
				CdtrAcct:       &pacs_008_001_02.CashAccount16{Id: &pacs_008_001_02.AccountIdentification4Choice{IBAN: "ciban-1"}},
				RmtInf:         &pacs_008_001_02.RemittanceInformation5{Ustrd: []string{"rmtinf-1"}},
				Dbtr:           &pacs_008_001_02.PartyIdentification32{Nm: "dnm-1"},
				DbtrAgt:        &pacs_008_001_02.BranchAndFinancialInstitutionIdentification4{FinInstnId: &pacs_008_001_02.FinancialInstitutionIdentification7{BIC: "dbic-1"}},
				DbtrAcct:       &pacs_008_001_02.CashAccount16{Id: &pacs_008_001_02.AccountIdentification4Choice{IBAN: "diban-1"}},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual, err := tC.mapper.Map(tC.pi3, tC.ct10); err != nil {
				t.Error(err)
			} else {
				assert.Equal(t, tC.expected, actual, "expected should equal actual")
			}
		})
	}
}
