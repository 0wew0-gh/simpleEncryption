package simpleEncryption

import (
	"fmt"
	"strings"
)

type SecretKey struct {
	extraItem int
	cKey      []string
	cryptKey  []string
}

// ==============================
//
//	初始化
//	extraItem: 额外信息的位置，从0开始
//	cryptKey: 密码本
//	key: 密钥
//	返回值1: SecretKey对象
//	返回值2: 错误信息
//
// ==============================
//
//	initialization
//	extraItem: extra information position, start from 0
//	key: key
//	cryptKey: password book
//	return value 1: SecretKey object
//	return value 2: error message
func New(extraItem int, key string, cryptKey string) (*SecretKey, error) {
	var se SecretKey
	if extraItem < 0 {
		return &se, fmt.Errorf("extraItem needs to be greater than 0")
	}
	if key == "" {
		return &se, fmt.Errorf("key cannot be empty")
	}
	if cryptKey == "" {
		return &se, fmt.Errorf("cryptKey cannot be empty")
	}
	se.extraItem = extraItem
	se.cKey = strings.Split(key, "")
	se.cryptKey = strings.Split(cryptKey, "")
	return &se, nil
}

// ==============================
//
//	初始化(Json格式)
//	key: 密码本
//		格式:
//		{
//			"extraItem": 1,
//			"key": ["asfa908#@%.,?", "bcdeftuvwxyz0123456789"]
//		}
//	返回值1: SecretKey对象
//	返回值2: 错误信息
//
// ==============================
//
//	initialization(Json format)
//	key: password book
//	format:
//		{
//			"extraItem": 1,
//			"key": ["asfa908#@%.,?", "bcdeftuvwxyz0123456789"]
//		}
//	return value 1: SecretKey object
//	return value 2: error message
func NewJson(key string) (*SecretKey, error) {
	se, err := checkKey(key)
	if err != nil {
		return se, err
	}
	return se, nil
}

// ==============================
//
//	加密
//	extra: 额外信息，暂只支持一个字符
//	str: 需要加密的字符串
//	返回值1: 加密后的字符串
//	返回值2: 错误信息
//
// ==============================
//
//	encryption
//	extra: extra information, only support one character
//	str: string need to encrypt
//	return value 1: encrypted string
//	return value 2: error message
func (se *SecretKey) Encrypt(str string, extraStr string) (reStr string) {
	j := 0
	strArr := strings.Split(str, "")
	extraStrArr := strings.Split(extraStr, "")
	for i := 0; i < len(strArr); i++ {
		tempStr := ""
		if se.extraItem == i {
			tempStr, j = se.cryption(extraStrArr[0], j)
			reStr += tempStr
		}
		tempStr, j = se.cryption(strArr[i], j)
		reStr += tempStr
	}
	return reStr
}

// ==============================
//
//	解密
//	str: 需要解密的字符串
//	返回值1: 解密后的字符串
//	返回值2: 额外信息
//	返回值3: 错误信息
//
// ==============================
//
//	decrypt
//	str: string need to decrypt
//	return value 1: decrypted string
//	return value 2: extra information
//	return value 3: error message
func (se *SecretKey) Decrypt(str string) (reStr string, extra string, err error) {
	j := 0
	strArr := strings.Split(str, "")
	for i := 0; i < len(str); i++ {
		tempStr := ""
		tempStr, j = se.cryption(strArr[i], j)
		if se.extraItem == i {
			extra = tempStr
			continue
		}
		reStr += tempStr
	}
	return reStr, extra, err
}

/// 私有方法

func (se *SecretKey) cryption(str string, j int) (string, int) {
	reStr := ""
	s := str
	cryptI := se.findChar(s)
	if cryptI == -1 {
		reStr = s
		return reStr, j
	}
	newByte := 0
	newByte, j = se.enByte(j, cryptI)
	reStr = se.cryptKey[newByte]
	return reStr, j
}
func (se *SecretKey) findChar(str string) int {
	cryptI := -1
	for i, v := range se.cryptKey {
		if str == v {
			cryptI = i
			break
		}
	}
	return cryptI
}
func (se *SecretKey) enByte(ki int, cryptI int) (newItem int, keyItem int) {
	k := se.cKey[ki]
	ki++
	cI := se.findChar(k)
	newI := cI - cryptI
	newI = se.cryptloop(newI)
	return newI, ki
}
func (se *SecretKey) cryptloop(i int) int {
	if i < 0 {
		i = len(se.cryptKey) + i
	}
	if i > len(se.cryptKey) {
		i = i - len(se.cryptKey)
	}
	if i < 0 || i > len(se.cryptKey) {
		i = se.cryptloop(i)
	}
	return i
}
