package simple_encryption

import (
	"fmt"
	"strings"
)

type SimpleEncryption struct {
	extraItem int
	cKey      []string
	cryptKey  []string
}

// 初始化
//
// initialization
//
// extraItem: 额外信息的位置，从0开始
//
// extraItem: extra information position, start from 0
//
// key: 密钥
//
// key: key
//
// cryptKey: 密码本
//
// cryptKey: password book
//
// 返回值1: SimpleEncryption对象
//
// return value 1: SimpleEncryption object
//
// 返回值2: 错误信息
//
// return value 2: error message
func New(extraItem int, key string, cryptKey string) (SimpleEncryption, error) {
	var se SimpleEncryption
	if extraItem < 0 {
		return se, fmt.Errorf("extraItem needs to be greater than 0")
	}
	if key == "" {
		return se, fmt.Errorf("key cannot be empty")
	}
	if cryptKey == "" {
		return se, fmt.Errorf("cryptKey cannot be empty")
	}
	se.extraItem = extraItem
	se.cKey = strings.Split(key, "")
	se.cryptKey = strings.Split(cryptKey, "")
	return se, nil
}

// 初始化(Json格式)
//
// initialization(Json format)
//
// key: 密码本
//
// key: password book
//
//	格式:
//
//	format:
//		{
//			"extraItem": 1,
//			"key": ["asfa908#@%.,?", "bcdeftuvwxyz0123456789"]
//		}
//
// 返回值1: SimpleEncryption对象
//
// return value 1: SimpleEncryption object
//
// 返回值2: 错误信息
//
// return value 2: error message
func NewJson(key string) (SimpleEncryption, error) {
	var (
		se  SimpleEncryption
		err error
	)
	se, err = checkKey(key)
	if err != nil {
		return se, err
	}
	return se, nil
}

// 加密
//
// encryption
//
// extra: 额外信息，暂只支持一个字符
//
// extra: extra information, only support one character
//
// str: 需要加密的字符串
//
// str: string need to encrypt
//
// 返回值1: 加密后的字符串
//
// return value 1: encrypted string
//
// 返回值2: 错误信息
//
// return value 2: error message
func (se SimpleEncryption) encrypt(str string, extraStr string) (string, error) {
	reStr := ""
	j := 0
	strArr := strings.Split(str, "")
	extraStrArr := strings.Split(extraStr, "")
	for i := 0; i < len(strArr); i++ {
		tempStr := ""
		if se.extraItem == i {
			tempStr, j = se._cryption(extraStrArr[0], j)
			reStr += tempStr
		}
		tempStr, j = se._cryption(strArr[i], j)
		reStr += tempStr
	}
	return reStr, nil
}

// 解密
//
// decrypt
//
// str: 需要解密的字符串
//
// str: string need to decrypt
//
// 返回值1: 解密后的字符串
//
// return value 1: decrypted string
//
// 返回值2: 额外信息
//
// return value 2: extra information
//
// 返回值3: 错误信息
//
// return value 3: error message
func (se SimpleEncryption) decrypt(str string) (reStr string, extra string, err error) {
	j := 0
	strArr := strings.Split(str, "")
	for i := 0; i < len(str); i++ {
		tempStr := ""
		tempStr, j = se._cryption(strArr[i], j)
		if se.extraItem == i {
			extra = tempStr
			continue
		}
		reStr += tempStr
	}
	return reStr, extra, err
}

/// 私有方法

func (se SimpleEncryption) _cryption(str string, j int) (string, int) {
	reStr := ""
	s := str
	cryptI := se._findChar(s)
	if cryptI == -1 {
		reStr = s
		return reStr, j
	}
	newByte := 0
	newByte, j = se._enByte(j, cryptI)
	reStr = se.cryptKey[newByte]
	return reStr, j
}
func (se SimpleEncryption) _findChar(str string) int {
	cryptI := -1
	for i, v := range se.cryptKey {
		if str == v {
			cryptI = i
			break
		}
	}
	return cryptI
}
func (se SimpleEncryption) _enByte(ki int, cryptI int) (newItem int, keyItem int) {
	k := se.cKey[ki]
	ki++
	cI := se._findChar(k)
	newI := cI - cryptI
	newI = se._cryptloop(newI)
	return newI, ki
}
func (se SimpleEncryption) _cryptloop(i int) int {
	if i < 0 {
		i = len(se.cryptKey) + i
	}
	if i > len(se.cryptKey) {
		i = i - len(se.cryptKey)
	}
	if i < 0 || i > len(se.cryptKey) {
		i = se._cryptloop(i)
	}
	return i
}
