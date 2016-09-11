package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkErr(err)
	listen, err := net.ListenTCP("tcp", addr)
	checkErr(err)
	fmt.Println("Start server...")
	for {
		conn, err := listen.Accept()
		checkErr(err)
		go Handle(conn) // 每次建立一个连接就放到单独的线程内做处理
	}
}

var key = []byte("ncuimexsecret777")
var AesStatus = false

const BufLength = 128

var users map[string]net.Conn = make(map[string]net.Conn, 10)

func Handle(conn net.Conn) {
	conn.Write([]byte("欢迎加入南昌大学聊天室~"))
	for {
		data := make([]byte, 0) //此处做一个输入缓冲以免数据过长读取到不完整的数据
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

		if AesStatus {
			var err error
			data, err = AesDecrypt(data, key) // Aes解密
			if err != nil {
			}
		}

		fmt.Println(data)
		cmd := strings.Split(string(data), "|")
		fmt.Println("==== New Message ====")
		fmt.Println("IP地址 ", conn.RemoteAddr().String())
		fmt.Println("接收日期 ", time.Now().Format("2006-01-02"))
		fmt.Println("接收时间 ", time.Now().Format("15:04:05"))
		fmt.Println("接收内容 ", cmd)
		fmt.Println("你的服务程序监听端口 ", "1234")
		conn.Write([]byte("已处理"))
		fmt.Println("====     End     ====")

		switch cmd[0] {
		case "nick":
			fmt.Println("注册名称:" + cmd[1])
			users[cmd[1]] = conn
		case "say":
			for k, v := range users {
				if k != cmd[1] {
					fmt.Println("给" + k + "发送消息:" + cmd[2])
					v.Write([]byte(cmd[1] + ":[" + cmd[2] + "]"))
				}
			}
		}
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
