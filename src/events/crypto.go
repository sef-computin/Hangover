package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	_ "log"
	"os"
)

const (
	KEY_PATH     = "key.k"
	DBCREDS_PATH = "dbcreds.dbc"
)

func Aes256Encode(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

func Aes256Decode(cipherText string, encKey string, iv string) (decryptedString string) {
	bKey := []byte(encKey)
	bIV := []byte(iv)
	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return string(cipherTextDecoded)
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func getDatabaseCreds() (string, error){
	var host, port, user, dbname, password string
	var opts []*string = []*string{&host, &port, &user, &password, &dbname}	

	keys_file, err := os.Open(KEY_PATH)
  if err != nil {
      return "", err
  }
  defer keys_file.Close()

  scanner := bufio.NewScanner(keys_file)
	
	scanner.Scan() 
  if err := scanner.Err(); err != nil {
     return "", err
  }
	key := scanner.Text()
	scanner.Scan()
  if err := scanner.Err(); err != nil {
     return "", err
  }
	iv := scanner.Text()

	dbcreds_file, err := os.Open(DBCREDS_PATH)
	defer dbcreds_file.Close()
	scanner = bufio.NewScanner(dbcreds_file)
	var dbcreds string
	for scanner.Scan(){
		dbcreds = fmt.Sprintf("%s%s", dbcreds, scanner.Text())
	}

	if scanner.Err() != nil{
		return "", scanner.Err()
	}

	db_decoded := Aes256Decode(string(dbcreds), string(key), iv)
	scanner = bufio.NewScanner(bytes.NewBufferString(db_decoded))

	i:=0
	for scanner.Scan() && i < 5{
		*(opts[i]) = scanner.Text()
		i++
	}
	
	if scanner.Err() != nil{
		return "", scanner.Err()
	}

 	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)	
	return dbURL, nil
}
