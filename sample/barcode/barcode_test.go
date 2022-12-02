package barcode_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/sanggi-wjg/go_demos/sample/barcode"
	"github.com/sanggi-wjg/go_demos/sample/file"
)

func TestCreateBarcode(t *testing.T) {
	testCode := "123456789"
	barcode.CreateBarcode(testCode)

	path := fmt.Sprintf("%s.png", testCode)
	if !file.IsFileExist(path) {
		t.Errorf(fmt.Sprintf("%s does not exist", path))
	}
	os.Remove(path)
}
