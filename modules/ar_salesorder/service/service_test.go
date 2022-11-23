package ar_salesorder_test

import (
	domain "STACK-GO/modules/ar_salesorder/domain"
	"STACK-GO/modules/ar_salesorder/intrface/mocks"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	so = domain.Salesorder{
		SOID:             1,
		Class:            0,
		SOType:           "",
		SONumber:         "SO/1",
		SOCounter:        1,
		CustomerID:       1,
		SODate:           time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
		CurrencyID:       1,
		Terms:            1,
		Tax:              "10%",
		AmountTax:        1000,
		Discount:         0,
		AmountDiscount:   0,
		Rounding:         0,
		TotalCharges:     0,
		NetTotal:         10000,
		InternalRemarks:  "Testing",
		Closed:           0,
		Destination:      "Medan",
		CustomerRef:      "",
		Importance:       "",
		Void:             0,
		VoidDateTime:     time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
		VoidBy:           "",
		VoidReason:       "",
		Approved:         0,
		ApprovedDateTime: time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
		ApprovedBy:       "",
		DeliveryTerm:     "",
		CreatedBy:        "Test",
		CreateDate:       time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
		LastModBy:        "",
		LastMod:          time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
		Remarks:          "Testing Aja",
		AdditionalTerms:  "",
		SOWarehouseID:    0,
		DeliveryTerms:    "",
		POCreated:        0,
		FPNo:             "",
		FPDate:           time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
		TaxAddress:       "",
		TaxNPWP:          "",
		TaxDate:          time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
		TaxCustomerName:  "",
		SOExchangeRate:   0,
		IgnoreTax:        0,
		Verified:         0,
		VerifiedBy:       "",
		VerifiedDateTime: time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
		ClosedDate:       time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
		Flag:             "",
		DontSI:           0,
		OpenPrice:        0,
		ClosedBy:         "",
		ClosedDateTime:   time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
		ClosedReason:     "",
		CustomerPODate:   time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
		AddressID:        0,
		ClosedBySystem:   0,
		IsCash:           0,
	}
	sodetail = []domain.SalesorderDetails{
		{
			SOID:              1,
			DetailID:          1,
			UOMLevel:          2,
			ItemID:            2,
			QtyOrdered:        20,
			UnitPrice:         200,
			Discount:          0,
			AmountDiscount:    0,
			LineTotal:         4000,
			Remarks:           "",
			CreatedBy:         "Test",
			CreateDate:        time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
			LastModBy:         "",
			LastMod:           time.Date(2022, time.September, 14, 0, 0, 0, 0, time.Local),
			WarehouseID:       0,
			TaxAmount:         0,
			SpecialOrder:      0,
			LineTotalRounding: 0,
			Availability:      "",
			PriceSet:          0,
			PPh22Set:          0,
			SODFlags:          "",
		},
	}
	sopack = domain.SalesorderPackage{
		Salesorder:        so,
		SalesorderDetails: sodetail,
	}
)

type Services struct {
	mock.Mock
}

func TestFindByID_Success(t *testing.T) {
	serviceMock := mocks.NewServices(t)

	// Prepare data mocks
	serviceMock.On("FindByID", "1").
		Return(
			func(s string) *domain.SalesorderPackage {
				return &sopack
			},
			func(s string) error {
				return nil
			},
		)

	// Go try data mocks
	testdata, err := serviceMock.FindByID("1")

	// fmt.Println(t)
	assert.NoError(t, err)
	assert.Equal(t, testdata, &sopack)
	serviceMock.AssertExpectations(t)
}

func TestFindByID_Fail(t *testing.T) {
	serviceMock := mocks.NewServices(t)

	// Prepare data mocks
	serviceMock.On("FindByID", "1").
		Return(
			func(s string) (*domain.SalesorderPackage, error) {
				return &sopack, nil
			},
		)

	// Go try data mocks
	_, err := serviceMock.FindByID("WrongID")

	// fmt.Println(t)
	assert.Error(t, err)
	// assert.NotEqual(t, testdata, &sopack)
	serviceMock.AssertExpectations(t)
}
