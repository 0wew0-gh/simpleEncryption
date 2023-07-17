package simpleEncryption

import (
	"fmt"
	"strings"
)

type SecretKey struct {
	//	额外信息的位置，从0开始
	//
	//	Position of extra information, starting from 0
	extraItem int
	//	密钥
	//
	//	Key
	cKey []string
	//	密码本
	//
	//	Password book
	cryptKey []string
}

// ===============
//
//	初始化
//	extraItem	int		额外信息的位置，从0开始
//	key		string		密钥
//	cryptKey	string		密码本
//
//	返回值1		*SecretKey	SecretKey对象
//	返回值2		error		错误信息
//
// ===============
//
//	initialization
//	extraItem	int		Position of extra information,
//									starting from 0
//	key		string		Key
//	cryptKey	string		Password book
//
//	return 1	*SecretKey	SecretKey object
//	return 2	error		Error message
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

// ===============
//
//	初始化(Json格式)
//	key		string		密码本
//		格式:
//		{
//			"extraItem": 1,
//			"key": ["asfa908#@%.,?", "bcdeftuvwxyz0123456789"]
//		}
//	返回值1		*SecretKey	SecretKey对象
//	返回值2		error 		错误信息
//
// ===============
//
//	initialization(Json format)
//	key		string		Password book
//	format:
//		{
//			"extraItem": 1,
//			"key": ["asfa908#@%.,?", "bcdeftuvwxyz0123456789"]
//		}
//	return 1	*SecretKey	SecretKey object
//	return 2	error		Error message
func NewJson(key string) (*SecretKey, error) {
	se, err := checkKey(key)
	if err != nil {
		return se, err
	}
	return se, nil
}

// ===============
//
//	加密
//	str		string		需要加密的字符串
//	extraStr	string		额外信息，只支持一个字符
//
//	返回值1		string		加密后的字符串
//
// ===============
//
//	encryption
//	str		string		string need to encrypt
//	extraStr	string		Extra information, only one
//									character is supported
//
//	return 1	string		Encrypted string
func (se *SecretKey) Encrypt(str string, extraStr string) (reStr string) {
	j := 0
	strArr := strings.Split(str, "")
	extraStrArr := strings.Split(extraStr, "")
	isAddExtraItem := false
	for i := 0; i < len(strArr); i++ {
		tempStr := ""
		if se.extraItem == i {
			isAddExtraItem = true
			tempStr, j = se.cryption(extraStrArr[0], j)
			reStr += tempStr
		}
		tempStr, j = se.cryption(strArr[i], j)
		reStr += tempStr
	}
	if !isAddExtraItem {
		tempStr := ""
		tempStr, _ = se.cryption(extraStrArr[0], j)
		reStr += tempStr
	}
	return reStr
}

// ===============
//
//	解密
//	str		string		需要解密的字符串
//
//	返回值1		string		解密后的字符串
//	返回值2		string		额外信息
//	返回值3		error		错误信息
//
// ===============
//
//	decrypt
//	str		string		string need to decrypt
//
//	return 1	string		Decrypted string
//	return 2	string		Extra information
//	return 3	error		Error message
func (se *SecretKey) Decrypt(str string) (reStr string, extra string, err error) {
	j := 0
	strArr := strings.Split(str, "")
	isAddExtraItem := false
	for i := 0; i < len(str); i++ {
		tempStr := ""
		tempStr, j = se.cryption(strArr[i], j)
		if se.extraItem == i || (!isAddExtraItem && i == len(str)-1) {
			isAddExtraItem = true
			extra = tempStr
			continue
		}
		reStr += tempStr
	}
	return reStr, extra, err
}

// ===============
//
//	对比密码本并返回对应字符串和下标
//	str		string		需要对比的字符串
//	j		int		密码本下标
//
//	返回值1		string		对比后的字符串
//	返回值2		int		密码本下标
//
// ===============
//
//	Compare the password book and return the corresponding string and subscript
//	str		string		String to compare
//	j		int		Password book subscript
//
//	return 1	string		Compared string
//	return 2	int		Password book subscript
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

// ===============
//
//	查找密码本中的字符
//	str		string		需要查找的字符串
//
//	返回值1		int		密码本下标
//
// ===============
//
//	Find characters in the password book
//	str		string		String to find
//
//	return 1	int		Password book subscript
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

// ===============
//
//	根据密码本和密钥返回密码本下标和密钥下标
//	ki		int		密钥下标
//	cryptI		int		密码本下标
//
//	返回值1		int		加密后的密码本下标
//	返回值2		int		密钥下标
//
// ===============
//
//	According to the password book and key, return the password book subscript
//	ki		int		Key subscript
//	cryptI		int		Password book subscript
//
//	return 1	int		Encrypted password book subscript
//	return 2	int		Key subscript
func (se *SecretKey) enByte(ki int, cryptI int) (newItem int, keyItem int) {
	if ki >= len(se.cKey) {
		ki = 0
	}
	k := se.cKey[ki]
	ki++
	cI := se.findChar(k)
	newI := cI - cryptI
	newI = se.cryptloop(newI)
	return newI, ki
}

// ===============
//
//	循环密码本
//		- 如果为负数则从后往前循环
//		- 如果为正数则从前往后循环
//		- 如果处理后还是超出范围，则继续循环
//	i		int		需要循环的下标
//
//	返回值1		int		循环后的下标
//
// ===============
//
//	Cycle password book
//		- If it is negative, it will cycle from back to front
//		- If it is positive, it will cycle from front to back
//		- If it is still out of range after processing, continue to cycle
//	i		int		Subscript to be cycled
//
//	return 1	int		Subscript after cycling
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
