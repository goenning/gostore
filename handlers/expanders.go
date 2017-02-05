package handlers

import "github.com/leekchan/accounting"

// CurrencyExpander formats currency based on current Accounting settings
func CurrencyExpander(args ...interface{}) string {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	value, _ := args[0].(float32)
	return ac.FormatMoney(value)
}
