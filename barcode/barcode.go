package barcode

import (
	bcode "github.com/boombuler/barcode" // Import the barcode package.
	"github.com/boombuler/barcode/ean"   // Import the EAN library from the barcode package for Barcode generation.
	"image/png"                          // Import the PNG standard library for the Logger library.
	"log"                                // Import log, going to implement the Shelves loggers at some point.
	"os"                                 // Import os for the Logger library.
)

// GenBarcode takes a EAN-13 Barcode Number and a variable of os.File type.
// It generates a barcode, and encodes it and then returns the barcode
// with a nil error value, if succesful.
func GenBarcode(barnumber string, f *os.File) (*os.File, error) {
	barcodeenc, err := ean.Encode(barnumber)
	if err != nil {
		log.Fatalf("Error encoding the barcode. Error message: %v\n", err.Error())
		return f, err
	}

	barcode, err := bcode.Scale(barcodeenc, 100, 100)
	if err != nil {
		log.Fatalf("Error scaling the barcode. Error message: %v\n", err.Error())
		return f, err
	}

	err = png.Encode(f, barcode)
	if err != nil {
		log.Fatalf("Error encoding the barcode to file. Error message: %v\n", err.Error())
		return f, err
	}

	log.Printf("Succesfully encoded barcode!")

	return f, nil
}
