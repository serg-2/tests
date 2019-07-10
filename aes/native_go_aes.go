package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "errors"
    "fmt"
    "io"
    "log"
)

func main() {
    key := []byte("Mega secret key SUper 123OGON !!") // 32 bytes \ 256 bit
    plaintext := []byte("My message for testing SKM! TESTTESTTEST")
    fmt.Printf("%s\n", plaintext)
    ciphertext, err := encrypt(key, plaintext)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%0x\n", ciphertext)
    result, err := decrypt(key, ciphertext)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n", result)
}

func encrypt(key, text []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    ciphertext := make([]byte, aes.BlockSize+len(text))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }
    cfb := cipher.NewCFBEncrypter(block, iv)
    cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))
    return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    if len(text) < aes.BlockSize {
        return nil, errors.New("ciphertext too short")
    }
    iv := text[:aes.BlockSize]
    text = text[aes.BlockSize:]
    cfb := cipher.NewCFBDecrypter(block, iv)
    cfb.XORKeyStream(text, text)
    return text, nil
}

