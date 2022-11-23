package iv_item

import "fmt"

const (
	ALL         = "*"
	TABLE       = "itm_products"
	COLUMN      = "ItemCode, ItemName, ItemType, UOMID1, UOMID2, UOMID3, UOMID4, AutoConvert, CategoryID, BrandID, PriceGroupID, UseSize, SizeRatioID, ValidSizes, UseColor, ValidColors, UseExpiryDate, Volume, MinQuantity, MaxQuantity, MinOrder, Disabled, OrderMultiple, CanBeSold, CanBePurchased, CanBeUsed, OutsourceToSupplierID, Remarks, CreatedBy, CreateDate, LastModBy, LastMod, FamilyID, SubFamilyID, ItemCounter, PriceCurrencyID, OldPrice, ExtraDiscount, Flag, Label, ExtItemCode"
	VALUE       = ":ItemName, :ItemType, :UOMID1, :UOMID2, :UOMID3, :UOMID4, :AutoConvert, :CategoryID, :BrandID, :PriceGroupID, :UseSize, :SizeRatioID, :ValidSizes, :UseColor, :ValidColors, :UseExpiryDate, :Volume, :MinQuantity, :MaxQuantity, :MinOrder, :Disabled, :OrderMultiple, :CanBeSold, :CanBePurchased, :CanBeUsed, :OutsourceToSupplierID, :Remarks, :CreatedBy, :CreateDate, :LastModBy, :LastMod, :FamilyID, :SubFamilyID, :ItemCounter, :PriceCurrencyID, :OldPrice, :ExtraDiscount, :Flag, :Label, :ExtItemCode"
	COLUMNVALUE = "ItemCode = :ItemCode, ItemName = :ItemName, ItemType = :ItemType, UOMID1 = :UOMID1, UOMID2 = :UOMID2, UOMID3 = :UOMID3, UOMID4 = :UOMID4, AutoConvert = :AutoConvert, CategoryID = :CategoryID, BrandID = :BrandID, PriceGroupID = :PriceGroupID, UseSize = :UseSize, SizeRatioID = :SizeRatioID, ValidSizes = :ValidSizes, UseColor = :UseColor, ValidColors = :ValidColors, UseExpiryDate = :UseExpiryDate, Volume = :Volume, MinQuantity = :MinQuantity, MaxQuantity = :MaxQuantity, MinOrder = :MinOrder, Disabled = :Disabled, OrderMultiple = :OrderMultiple, CanBeSold = :CanBeSold, CanBePurchased = :CanBePurchased, CanBeUsed = :CanBeUsed, OutsourceToSupplierID = :OutsourceToSupplierID, Remarks = :Remarks, CreatedBy = :CreatedBy, CreateDate = :CreateDate, LastModBy = :LastModBy, LastMod = :LastMod, FamilyID = :FamilyID, SubFamilyID = :SubFamilyID, ItemCounter = :ItemCounter, PriceCurrencyID = :PriceCurrencyID, OldPrice = :OldPrice, ExtraDiscount = :ExtraDiscount, Flag = :Flag, Label = :Label, ExtItemCode = :ExtItemCode "
	WHERE       = "ItemID"

	SELECT = "SELECT %s FROM %s WHERE %s = ?"
	CREATE = "INSERT INTO %s (%s) VALUES (%s)"
	UPDATE = "UPDATE %s SET %s WHERE %s = ?"
	LIST   = "SELECT %s FROM %s"
	DELETE = "DELETE FROM %s WHERE %s = ?"
)

func QryFindByID() string {
	return fmt.Sprintf(SELECT, ALL, TABLE, WHERE)
}

func QryList() string {
	return fmt.Sprintf(LIST, ALL, TABLE)
}

func QryCreate() string {
	return fmt.Sprintf(CREATE, TABLE, COLUMN, VALUE)
}

func QryUpdate() string {
	return fmt.Sprintf(UPDATE, TABLE, COLUMNVALUE, WHERE)
}

func QryDelete() string {
	return fmt.Sprintf(DELETE, TABLE, WHERE)
}
