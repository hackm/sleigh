package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type TO struct {
	Path   string
	Binary []byte
}

func TestJSON(t *testing.T) {
	// fmt.Println("Hello")
	// file, err := os.Open(`C:\Users\iwate\Downloads\dotnet-sdk-2.0.2-win-gs-x64.exe`)
	// if err != nil {
	// 	t.Fatal("Cannot open file")
	// }
	// defer file.Close()
	// stat, err := os.Stat(`C:\Users\iwate\Downloads\dotnet-sdk-2.0.2-win-gs-x64.exe`)
	// if err != nil {
	// 	t.Fatal("Cannot read stat")
	// }
	// b, err := encodeChecksumIndex(file, stat.Size(), BlockSize)
	// if err != nil {
	// 	t.Fatal("Cannot encode")
	// }
	// b64 := base64.RawURLEncoding.EncodeToString(b)

	bs, _ := json.Marshal(TO{
		Path:   "hoge/fuga/piyo",
		Binary: []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	})

	fmt.Printf("%s", string(bs))
}
