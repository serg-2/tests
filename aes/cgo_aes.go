package main
//#cgo LDFLAGS: -lcrypto
//#include <openssl/aes.h>
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

const (
	textConst = "Super secret Mdsffsdfsdfsdsdfsdfsdfkl;;ldks;lfksd;lffdgdfg"
	keyConst  = "Supsdfewpop"
)
func main() {
	text := (*C.uchar)(unsafe.Pointer(C.CString(textConst)))
	encOut := [16]C.uchar{}
	decOut := [16]C.uchar{}
//	encKey := C.AES_KEY{}
//	decKey := C.AES_KEY{}
    var encKey,decKey C.AES_KEY

    key := (*C.uchar)(unsafe.Pointer(C.CString(keyConst)))
	if _, err := C.AES_set_encrypt_key(key, 128, &encKey); err != nil {
		log.Fatalf("couldn't set encrypt key: %v", err)
	}
	encOutPtr := (*C.uchar)(unsafe.Pointer(&encOut))
	if _, err := C.AES_encrypt(text, encOutPtr, &encKey); err != nil {
		log.Fatalf("couldn't ecnrypt text: %v", err)
	}

	if _, err := C.AES_set_decrypt_key(key, 128, &decKey); err != nil {
		log.Fatalf("couldn't set decrypt key: %v", err)
	}

	decOutPtr := (*C.uchar)(unsafe.Pointer(&decOut))

	if _, err := C.AES_decrypt(encOutPtr, decOutPtr, &decKey); err != nil {
		log.Fatalf("couldn't decrypt text: %v", err)
	}

    fmt.Println("Begin ENCRYPTED MESSAGE{")

	for _, c := range encOut {
		fmt.Printf("%c", c)
	}

    fmt.Println("}END ENCRYPTED MESSAGE")

	for _, c := range decOut {
		fmt.Printf("%c", c)
	}
	fmt.Println("====")
}

