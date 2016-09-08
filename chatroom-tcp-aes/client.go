package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"net"
	"os"
)

var nick string = ""
var key = []byte("ncuimexsecret777")

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, addr)
	checkErr(err)
	// 读取提示
	data := make([]byte, 1024)
	conn.Read(data)
	fmt.Println(string(data))
	// 输入昵称
	fmt.Print("输入昵称:")
	fmt.Scanf("%v", &nick)
	fmt.Println("Hello " + nick)
	message := []byte("nick|" + nick)
	messageencrypt, err := AesEncrypt(message, key)
	if err != nil {
	}
	conn.Write(messageencrypt)

	go Handle(conn)

	for {
		someTex := ""
		fmt.Scanf("%v", &someTex)
		message := []byte("say|" + nick + "|" + someTex)
		messageencrypt, err := AesEncrypt(message, key)
		if err != nil {
		}
		fmt.Println(message)
		fmt.Println(messageencrypt)
		conn.Write(messageencrypt)
	}
}

const BufLength = 128

func Handle(conn net.Conn) {
	for {
		data := make([]byte, 1024)
		buf := make([]byte, BufLength)
		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				checkErr(err)
			}
			data = append(data, buf[:n]...)
			if n != BufLength {
				break
			}
		}

		fmt.Println(string(data))
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
