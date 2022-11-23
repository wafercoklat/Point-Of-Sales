package dt_bank

import (
	"time"
)

type Bank struct {
	ItemID                int       `json:"ItemID,omitempty" db:"ItemID"`
	ItemCode              string    `json:"ItemCode,omitempty" db:"ItemCode"`
	ItemName              string    `json:"ItemName,omitempty" db:"ItemName"`
	ItemType              int       `json:"ItemType,omitempty" db:"ItemType"`
	UOMID1                int       `json:"UOMID1,omitempty" db:"UOMID1"`
	UOMID2                int       `json:"UOMID2,omitempty" db:"UOMID2"`
	UOMID3                int       `json:"UOMID3,omitempty" db:"UOMID3"`
	UOMID4                int       `json:"UOMID4,omitempty" db:"UOMID4"`
	AutoConvert           int       `json:"AutoConvert,omitempty" db:"AutoConvert"`
	CategoryID            int       `json:"CategoryID,omitempty" db:"CategoryID"`
	BrandID               int       `json:"BrandID,omitempty" db:"BrandID"`
	PriceGroupID          int       `json:"PriceGroupID,omitempty" db:"PriceGroupID"`
	UseSize               int       `json:"UseSize,omitempty" db:"UseSize"`
	SizeRatioID           int       `json:"SizeRatioID,omitempty" db:"SizeRatioID"`
	ValidSizes            string    `json:"ValidSizes,omitempty" db:"ValidSizes"`
	UseColor              int       `json:"UseColor,omitempty" db:"UseColor"`
	ValidColors           string    `json:"ValidColors,omitempty" db:"ValidColors"`
	UseExpiryDate         int       `json:"UseExpiryDate,omitempty" db:"UseExpiryDate"`
	Volume                float64   `json:"Volume,omitempty" db:"Volume"`
	MinQuantity           float64   `json:"MinQuantity,omitempty" db:"MinQuantity"`
	MaxQuantity           float64   `json:"MaxQuantity,omitempty" db:"MaxQuantity"`
	MinOrder              float64   `json:"MinOrder,omitempty" db:"MinOrder"`
	Disabled              int       `json:"Disabled,omitempty" db:"Disabled"`
	OrderMultiple         float64   `json:"OrderMultiple,omitempty" db:"OrderMultiple"`
	CanBeSold             int       `json:"CanBeSold,omitempty" db:"CanBeSold"`
	CanBePurchased        int       `json:"CanBePurchased,omitempty" db:"CanBePurchased"`
	CanBeUsed             int       `json:"CanBeUsed,omitempty" db:"CanBeUsed"`
	OutsourceToSupplierID int       `json:"OutsourceToSupplierID,omitempty" db:"OutsourceToSupplierID"`
	Remarks               string    `json:"Remarks,omitempty" db:"Remarks"`
	CreatedBy             string    `json:"CreatedBy,omitempty" db:"CreatedBy"`
	CreateDate            time.Time `json:"CreateDate,omitempty" db:"CreateDate"`
	LastModBy             string    `json:"LastModBy,omitempty" db:"LastModBy"`
	LastMod               time.Time `json:"LastMod,omitempty" db:"LastMod"`
	FamilyID              int       `json:"FamilyID,omitempty" db:"FamilyID"`
	SubFamilyID           int       `json:"SubFamilyID,omitempty" db:"SubFamilyID"`
	ItemCounter           int       `json:"ItemCounter,omitempty" db:"ItemCounter"`
	PriceCurrencyID       int       `json:"PriceCurrencyID,omitempty" db:"PriceCurrencyID"`
	OldPrice              float64   `json:"OldPrice,omitempty" db:"OldPrice"`
	ExtraDiscount         string    `json:"ExtraDiscount,omitempty" db:"ExtraDiscount"`
	Flag                  string    `json:"Flag,omitempty" db:"Flag"`
	Label                 string    `json:"Label,omitempty" db:"Label"`
	ExtItemCode           string    `json:"ExtItemCode,omitempty" db:"ExtItemCode"`
}

var List []Bank
