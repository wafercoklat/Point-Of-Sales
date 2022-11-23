package fi_paymentvoucher

import (
	"time"
)

type Paymentvoucher struct {
	SOID             int       `json:"SOID,omitempty" db:"SOID"`
	Class            int       `json:"Class,omitempty" db:"Class"`
	SOType           string    `json:"SOType,omitempty" db:"SOType"`
	SONumber         string    `json:"SONumber,omitempty" db:"SONumber"`
	SOCounter        int       `json:"SOCounter,omitempty" db:"SOCounter"`
	CustomerID       int       `json:"CustomerID,omitempty" db:"CustomerID"`
	SODate           time.Time `json:"SODate,omitempty" db:"SODate"`
	CurrencyID       int       `json:"CurrencyID,omitempty" db:"CurrencyID"`
	Terms            int       `json:"Terms,omitempty" db:"Terms"`
	Tax              string    `json:"Tax,omitempty" db:"Tax"`
	AmountTax        float64   `json:"AmountTax,omitempty" db:"AmountTax"`
	Discount         float64   `json:"Discount,omitempty" db:"Discount"`
	AmountDiscount   float64   `json:"AmountDiscount,omitempty" db:"AmountDiscount"`
	Rounding         float64   `json:"Rounding,omitempty" db:"Rounding"`
	TotalCharges     float64   `json:"TotalCharges,omitempty" db:"TotalCharges"`
	NetTotal         float64   `json:"NetTotal,omitempty" db:"NetTotal"`
	InternalRemarks  string    `json:"InternalRemarks,omitempty" db:"InternalRemarks"`
	Closed           int       `json:"Closed,omitempty" db:"Closed"`
	Destination      string    `json:"Destination,omitempty" db:"Destination"`
	CustomerRef      string    `json:"CustomerRef,omitempty" db:"CustomerRef"`
	Importance       string    `json:"Importance,omitempty" db:"Importance"`
	Void             int       `json:"Void,omitempty" db:"Void"`
	VoidDateTime     time.Time `json:"VoidDateTime,omitempty" db:"VoidDateTime"`
	VoidBy           string    `json:"VoidBy,omitempty" db:"VoidBy"`
	VoidReason       string    `json:"VoidReason,omitempty" db:"VoidReason"`
	Approved         int       `json:"Approved,omitempty" db:"Approved"`
	ApprovedDateTime time.Time `json:"ApprovedDateTime,omitempty" db:"ApprovedDateTime"`
	ApprovedBy       string    `json:"ApprovedBy,omitempty" db:"ApprovedBy"`
	DeliveryTerm     string    `json:"DeliveryTerm,omitempty" db:"DeliveryTerm"`
	CreatedBy        string    `json:"CreatedBy,omitempty" db:"CreatedBy"`
	CreateDate       time.Time `json:"CreateDate,omitempty" db:"CreateDate"`
	LastModBy        string    `json:"LastModBy,omitempty" db:"LastModBy"`
	LastMod          time.Time `json:"LastMod,omitempty" db:"LastMod"`
	Remarks          string    `json:"Remarks,omitempty" db:"Remarks"`
	AdditionalTerms  string    `json:"AdditionalTerms,omitempty" db:"AdditionalTerms"`
	SOWarehouseID    int       `json:"SOWarehouseID,omitempty" db:"SOWarehouseID"`
	DeliveryTerms    string    `json:"DeliveryTerms,omitempty" db:"DeliveryTerms"`
	POCreated        int       `json:"POCreated,omitempty" db:"POCreated"`
	FPNo             string    `json:"FPNo,omitempty" db:"FPNo"`
	FPDate           time.Time `json:"FPDate,omitempty" db:"FPDate"`
	TaxAddress       string    `json:"TaxAddress,omitempty" db:"TaxAddress"`
	TaxNPWP          string    `json:"TaxNPWP,omitempty" db:"TaxNPWP"`
	TaxDate          time.Time `json:"TaxDate,omitempty" db:"TaxDate"`
	TaxCustomerName  string    `json:"TaxCustomerName,omitempty" db:"TaxCustomerName"`
	SOExchangeRate   float64   `json:"SOExchangeRate,omitempty" db:"SOExchangeRate"`
	IgnoreTax        int       `json:"IgnoreTax,omitempty" db:"IgnoreTax"`
	Verified         int       `json:"Verified,omitempty" db:"Verified"`
	VerifiedBy       string    `json:"VerifiedBy,omitempty" db:"VerifiedBy"`
	VerifiedDateTime time.Time `json:"VerifiedDateTime,omitempty" db:"VerifiedDateTime"`
	ClosedDate       time.Time `json:"ClosedDate,omitempty" db:"ClosedDate"`
	Flag             string    `json:"Flag,omitempty" db:"Flag"`
	DontSI           int       `json:"DontSI,omitempty" db:"DontSI"`
	OpenPrice        int       `json:"OpenPrice,omitempty" db:"OpenPrice"`
	ClosedBy         string    `json:"ClosedBy,omitempty" db:"ClosedBy"`
	ClosedDateTime   time.Time `json:"ClosedDateTime,omitempty" db:"ClosedDateTime"`
	ClosedReason     string    `json:"ClosedReason,omitempty" db:"ClosedReason"`
	CustomerPODate   time.Time `json:"CustomerPODate,omitempty" db:"CustomerPODate"`
	AddressID        int       `json:"AddressID,omitempty" db:"AddressID"`
	ClosedBySystem   int       `json:"ClosedBySystem,omitempty" db:"ClosedBySystem"`
	IsCash           int       `json:"IsCash,omitempty" db:"IsCash"`
}

type PaymentvoucherDetails struct {
	SOID              int       `json:"SOID,omitempty" db:"SOID"`
	DetailID          int       `json:"DetailID,omitempty" db:"DetailID"`
	UOMLevel          int       `json:"UOMLevel,omitempty" db:"UOMLevel"`
	ItemID            int       `json:"ItemID,omitempty" db:"ItemID"`
	QtyOrdered        float64   `json:"QtyOrdered,omitempty" db:"QtyOrdered"`
	UnitPrice         float64   `json:"UnitPrice,omitempty" db:"UnitPrice"`
	Discount          float64   `json:"Discount,omitempty" db:"Discount"`
	AmountDiscount    float64   `json:"AmountDiscount,omitempty" db:"AmountDiscount"`
	LineTotal         float64   `json:"LineTotal,omitempty" db:"LineTotal"`
	Remarks           string    `json:"Remarks,omitempty" db:"Remarks"`
	CreatedBy         string    `json:"CreatedBy,omitempty" db:"CreatedBy"`
	CreateDate        time.Time `json:"CreateDate,omitempty" db:"CreateDate"`
	LastModBy         string    `json:"LastModBy,omitempty" db:"LastModBy"`
	LastMod           time.Time `json:"LastMod,omitempty" db:"LastMod"`
	WarehouseID       int       `json:"WarehouseID,omitempty" db:"WarehouseID"`
	TaxAmount         float64   `json:"TaxAmount,omitempty" db:"TaxAmount"`
	SpecialOrder      int       `json:"SpecialOrder,omitempty" db:"SpecialOrder"`
	LineTotalRounding float64   `json:"LineTotalRounding,omitempty" db:"LineTotalRounding"`
	Availability      string    `json:"Availability,omitempty" db:"Availability"`
	PriceSet          float64   `json:"PriceSet,omitempty" db:"PriceSet"`
	PPh22Set          float64   `json:"PPh22Set,omitempty" db:"PPh22Set"`
	SODFlags          string    `json:"SODFlags,omitempty" db:"SODFlags"`
}

type PaymentvoucherPackage struct {
	Paymentvoucher        Paymentvoucher          `json:"Paymentvoucher,omitempty"`
	PaymentvoucherDetails []PaymentvoucherDetails `json:"PaymentvoucherDetails,omitempty"`
}

var List []Paymentvoucher
