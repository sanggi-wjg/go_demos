package factory

import (
	"errors"
	"fmt"
)

const (
	EMS_CODE = "EMS"
	SF_CODE  = "SF"
)

type PrinterMethod interface {
	GetOutput(code string) string
}

type EMSPrinter struct{}
type SFPrinter struct{}

func (p *EMSPrinter) GetOutput(packageCode string) string {
	return fmt.Sprintf("<html><body><img src=\"http://host.com/%s.png\"></body></html>", packageCode)
}

func (p *SFPrinter) GetOutput(packageCode string) string {
	return fmt.Sprintf("^BZ^X,Y,100,100,%s^BI^BZ", packageCode)
}

func CreatePrinterMethod(transferCompanyCode string) (PrinterMethod, error) {
	switch transferCompanyCode {
	case EMS_CODE:
		return new(EMSPrinter), nil
	case SF_CODE:
		return new(SFPrinter), nil
	default:
		return nil, errors.New(fmt.Sprintf("%s is not implemented", transferCompanyCode))
	}
}
