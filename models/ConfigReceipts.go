
package models

var ConfigReceiptsTable = "config__receipt"

type ConfigReceipts struct {
	Printer Printer
	PrinterId uint

	Receipt string
}