package simple_encryption

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func checkKey(key string) (SimpleEncryption, error) {
	var (
		se  SimpleEncryption
		err error
	)
	tempJson := make(map[string]interface{}, 0)
	if se.extraItem < 0 || se.cKey == nil || (se.cKey != nil && len(se.cKey) == 0) || se.cryptKey == nil || (se.cryptKey != nil && len(se.cryptKey) == 0) {
		err = json.Unmarshal([]byte(key), &tempJson)
		if err != nil {
			return se, err
		}
	}
	if se.extraItem < 0 || se.cKey == nil || (se.cKey != nil && len(se.cKey) == 0) {
		temp := tempJson["extraItem"]
		switch tempType := temp.(type) {
		case float64:
			se.extraItem = int(tempType)
		case int:
			se.extraItem = tempType
		case int64:
			se.extraItem = int(tempType)
		case string:
			se.extraItem, err = strconv.Atoi(tempType)
			if err != nil {
				return se, err
			}
		default:
			return se, fmt.Errorf("extraItem is not int")
		}
	}
	if se.cryptKey == nil || (se.cryptKey != nil && len(se.cryptKey) == 0) || se.cKey == nil || (se.cKey != nil && len(se.cKey) == 0) {
		temp := tempJson["key"]
		switch tempType := temp.(type) {
		case []interface{}:
			for i := 0; i < len(tempType); i++ {
				switch tempTypeChild := tempType[i].(type) {
				case string:
					if i == 0 {
						se.cKey = strings.Split(tempTypeChild, "")
					} else {
						se.cryptKey = strings.Split(tempTypeChild, "")
					}
				default:
					return se, fmt.Errorf("key.key is not true")
				}
			}
		default:
			return se, fmt.Errorf("key.key is not true")
		}
	}
	if se.extraItem < 0 || se.cKey == nil || (se.cKey != nil && len(se.cKey) == 0) || se.cryptKey == nil || (se.cryptKey != nil && len(se.cryptKey) == 0) {
		return se, fmt.Errorf("key is not true")
	}
	return se, nil
}
