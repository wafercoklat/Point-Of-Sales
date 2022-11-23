package intap_purchaseorder

import "fmt"

var (
	ALL    = "*"
	header = "Header"
	detail = "Detail"
	TABLE  = map[string]string{
		header: "intap_purchaseorder",
		detail: "intap_purchaseorderdetails",
	}
	COLUMN = map[string]string{
		header: "Class, SOType, SONumber, SOCounter, CustomerID, SODate, CurrencyID, Terms, Tax, AmountTax, Discount, AmountDiscount, Rounding, TotalCharges, NetTotal, InternalRemarks, Closed, Destination, CustomerRef, Importance, Void, VoidDateTime, VoidBy, VoidReason, Approved, ApprovedDateTime, ApprovedBy, DeliveryTerm, CreatedBy, CreateDate, LastModBy, LastMod, Remarks, AdditionalTerms, SOWarehouseID, DeliveryTerms, POCreated, FPNo, FPDate, TaxAddress, TaxNPWP, TaxDate, TaxCustomerName, SOExchangeRate, IgnoreTax, Verified, VerifiedBy, VerifiedDateTime, ClosedDate, Flag, DontSI, OpenPrice, ClosedBy, ClosedDateTime, ClosedReason, CustomerPODate, AddressID, ClosedBySystem, IsCash  ",
		detail: "SOID,  UOMLevel, ItemID, QtyOrdered, UnitPrice, Discount, AmountDiscount, LineTotal, Remarks, CreatedBy, CreateDate, LastModBy, LastMod, WarehouseID, TaxAmount, SpecialOrder, LineTotalRounding, Availability, PriceSet, PPh22Set, SODFlags",
	}
	VALUE = map[string]string{
		header: ":Class, :SOType, :SONumber, :SOCounter, :CustomerID, :SODate, :CurrencyID, :Terms, :Tax, :AmountTax, :Discount, :AmountDiscount, :Rounding, :TotalCharges, :NetTotal, :InternalRemarks, :Closed, :Destination, :CustomerRef, :Importance, :Void, :VoidDateTime, :VoidBy, :VoidReason, :Approved, :ApprovedDateTime, :ApprovedBy, :DeliveryTerm, :CreatedBy, :CreateDate, :LastModBy, :LastMod, :Remarks, :AdditionalTerms, :SOWarehouseID, :DeliveryTerms, :POCreated, :FPNo, :FPDate, :TaxAddress, :TaxNPWP, :TaxDate, :TaxCustomerName, :SOExchangeRate, :IgnoreTax, :Verified, :VerifiedBy, :VerifiedDateTime, :ClosedDate, :Flag, :DontSI, :OpenPrice, :ClosedBy, :ClosedDateTime, :ClosedReason, :CustomerPODate, :AddressID, :ClosedBySystem, :IsCash  ",
		detail: ":SOID, :UOMLevel, :ItemID, :QtyOrdered, :UnitPrice, :Discount, :AmountDiscount, :LineTotal, :Remarks, :CreatedBy, :CreateDate, :LastModBy, :LastMod, :WarehouseID, :TaxAmount, :SpecialOrder, :LineTotalRounding, :Availability, :PriceSet, :PPh22Set, :SODFlags",
	}

	COLUMNVALUE = map[string]string{
		header: "Class = :Class, SOType = :SOType, SONumber = :SONumber, SOCounter = :SOCounter, CustomerID = :CustomerID, SODate = :SODate, CurrencyID = :CurrencyID, Terms = :Terms, Tax = :Tax, AmountTax = :AmountTax, Discount = :Discount, AmountDiscount = :AmountDiscount, Rounding = :Rounding, TotalCharges = :TotalCharges, NetTotal = :NetTotal, InternalRemarks = :InternalRemarks, Closed = :Closed, Destination = :Destination, CustomerRef = :CustomerRef, Importance = :Importance, Void = :Void, VoidDateTime = :VoidDateTime, VoidBy = :VoidBy, VoidReason = :VoidReason, Approved = :Approved, ApprovedDateTime = :ApprovedDateTime, ApprovedBy = :ApprovedBy, DeliveryTerm = :DeliveryTerm, CreatedBy = :CreatedBy, CreateDate = :CreateDate, LastModBy = :LastModBy, LastMod = :LastMod, Remarks = :Remarks, AdditionalTerms = :AdditionalTerms, SOWarehouseID = :SOWarehouseID, DeliveryTerms = :DeliveryTerms, POCreated = :POCreated, FPNo = :FPNo, FPDate = :FPDate, TaxAddress = :TaxAddress, TaxNPWP = :TaxNPWP, TaxDate = :TaxDate, TaxCustomerName = :TaxCustomerName, SOExchangeRate = :SOExchangeRate, IgnoreTax = :IgnoreTax, Verified = :Verified, VerifiedBy = :VerifiedBy, VerifiedDateTime = :VerifiedDateTime, ClosedDate = :ClosedDate, Flag = :Flag, DontSI = :DontSI, OpenPrice = :OpenPrice, ClosedBy = :ClosedBy, ClosedDateTime = :ClosedDateTime, ClosedReason = :ClosedReason, CustomerPODate = :CustomerPODate, AddressID = :AddressID, ClosedBySystem = :ClosedBySystem, IsCash = :IsCash  ",
		detail: "SOID = :SOID, UOMLevel = :UOMLevel, ItemID = :ItemID, QtyOrdered = :QtyOrdered, UnitPrice = :UnitPrice, Discount = :Discount, AmountDiscount = :AmountDiscount, LineTotal = :LineTotal, Remarks = :Remarks, CreatedBy = :CreatedBy, CreateDate = :CreateDate, LastModBy = :LastModBy, LastMod = :LastMod, WarehouseID = :WarehouseID, TaxAmount = :TaxAmount, SpecialOrder = :SpecialOrder, LineTotalRounding = :LineTotalRounding, Availability = :Availability, PriceSet = :PriceSet, PPh22Set = :PPh22Set, SODFlags = :SODFlags",
	}
	WHERE = "SOID"

	SELECT = "SELECT %s FROM %s WHERE %s = ?"
	CREATE = "INSERT INTO %s (%s) VALUES (%s)"
	UPDATE = "UPDATE %s SET %s WHERE %s = ?"
	LIST   = "SELECT %s FROM %s"
	DELETE = "DELETE FROM %s WHERE %s = ?"
)

func QryFindByID_Header() string {
	return fmt.Sprintf(SELECT, ALL, TABLE[header], WHERE)
}

func QryFindByID_Detail() string {
	return fmt.Sprintf(SELECT, ALL, TABLE[detail], WHERE)
}

func QryList() string {
	return fmt.Sprintf(LIST, ALL, TABLE[header])
}

func QryCreate_Header() string {
	return fmt.Sprintf(CREATE, TABLE[header], COLUMN[header], VALUE[header])
}

func QryCreate_Detail() string {
	return fmt.Sprintf(CREATE, TABLE[detail], COLUMN[detail], VALUE[detail])
}

func QryUpdate_Header() string {
	return fmt.Sprintf(UPDATE, TABLE[header], COLUMNVALUE[header], WHERE)
}

func QryUpdate_Detail() string {
	return fmt.Sprintf(UPDATE, TABLE[detail], COLUMNVALUE[detail], WHERE)
}

func QryDelete_Header() string {
	return fmt.Sprintf(DELETE, TABLE[header], WHERE)
}

func QryDelete_Detail() string {
	return fmt.Sprintf(DELETE, TABLE[detail], WHERE)
}
