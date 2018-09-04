package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

// TLD 域名服务器信息
type TLD struct {
	Tld         string
	Description string
	WhoisServer string
	Patterns    struct {
		NotRegistered string
		WaitPeriod    string
	}
	WaitPeriod int
}

// GetTLD 获取某个TLD的配置信息
func GetTLD(tld string, cfgPath string) (TLD, error) {
	var TldItem TLD

	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		fmt.Printf("打开TLD文件失败\n")
		return TldItem, errors.New("打开TLD文件失败")
	}

	resp := []TLD{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return TldItem, errors.New("解析json失败")
	}

	//遍历数据
	for _, v := range resp {
		v.Patterns.NotRegistered = strings.Trim(v.Patterns.NotRegistered, "/")
		if v.Tld == tld {
			TldItem = v
		}

	}

	return TldItem, nil
}
