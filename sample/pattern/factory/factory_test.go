package factory_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/sanggi-wjg/go_demos/sample/pattern/factory"
)

func TestCreate(t *testing.T) {
	// given
	transferCompanyCode1 := factory.EMS_CODE
	transferCompanyCode2 := factory.SF_CODE
	transferCompanyCode3 := "What"

	// when
	method1, err1 := factory.CreatePrinterMethod(transferCompanyCode1)
	method2, err2 := factory.CreatePrinterMethod(transferCompanyCode2)
	method3, err3 := factory.CreatePrinterMethod(transferCompanyCode3)

	// then
	assert.Equal(t, method1, factory.EMSPrinter{})
	assert.Equal(t, err1, nil)

	assert.Equal(t, method2, factory.SFPrinter{})
	assert.Equal(t, err2, nil)

	assert.Equal(t, method3, nil)
	assert.NotEqual(t, err3, nil)
}

func TestCreateEMS(t *testing.T) {
	// given
	packageCode := "123"
	transferCompanyCode := factory.EMS_CODE

	// when
	printerMethod, err := factory.CreatePrinterMethod(transferCompanyCode)
	output := printerMethod.GetOutput(packageCode)

	// then
	assert.Equal(t, err, nil)
	assert.Equal(t, output, "<html><body><img src=\"http://host.com/123.png\"></body></html>")
}

func TestCreateSF(t *testing.T) {
	// given
	packageCode := "123"
	transferCompanyCode := factory.SF_CODE

	// when
	printerMethod, err := factory.CreatePrinterMethod(transferCompanyCode)
	output := printerMethod.GetOutput(packageCode)

	// then
	assert.Equal(t, err, nil)
	assert.Equal(t, output, "^BZ^X,Y,100,100,123^BI^BZ")
}
