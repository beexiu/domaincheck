//usr/local/go/bin/go run $0 $@ $(dirname `realpath $0`); exit
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

// DomainList xml domainlist
type DomainList struct {
	XMLName xml.Name `xml:"domainList"`
	Version string   `xml:"version,attr"`
	Domains []Domain `xml:"domain"`
}

// Domain xml domain
type Domain struct {
	Name    string        `xml:"name,attr"`
	Source  string        `xml:"source"`
	Servers []WhoisServer `xml:"whoisServer"`
	Created string        `xml:"created"`
	Changed string        `xml:"changed"`
	State   string        `xml:"state"`
}

// WhoisServer xml domain
type WhoisServer struct {
	Host             string `xml:"host,attr"`
	Source           string `xml:"source"`
	AvailablePattern string `xml:"availablePattern"`
	QueryFormat      string `xml:"queryFormat"`
}

// TLD 域名服务器信息
type TLD struct {
	Tld         string      `json:"tld"`
	Description string      `json:"description"`
	WhoisServer string      `json:"whoisServer"`
	Patterns    TLDPatterns `json:"patterns"`
	WaitPeriod  int         `json:"waitPeriod"`
}

// TLDPatterns patterns
type TLDPatterns struct {
	NotRegistered string `json:"notRegistered"`
	WaitPeriod    string `json:"waitPeriod"`
}

func main() {
	srcTldURL := "https://raw.githubusercontent.com/whois-server-list/whois-server-list/master/whois-server-list.xml"
	resp, err := http.Get(srcTldURL)
	assert(err)

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	assert(err)

	// file, err := os.Open("/Users/brant/whois-server-list.xml")
	// assert(err)
	// defer file.Close()
	// data, err := ioutil.ReadAll(file)
	// assert(err)

	servers := DomainList{}
	err = xml.Unmarshal(data, &servers)
	assert(err)

	var tlds []TLD
	for _, dm := range servers.Domains {
		if len(dm.Servers) > 0 {
			var tld TLD
			tld.Tld = dm.Name
			tld.WhoisServer = dm.Servers[0].Host
			tld.Patterns.NotRegistered = dm.Servers[0].AvailablePattern
			tld.Patterns.NotRegistered = strings.Replace(tld.Patterns.NotRegistered, "\\Q", "/", -1)
			tld.Patterns.NotRegistered = strings.Replace(tld.Patterns.NotRegistered, "\\E", "/", -1)
			tlds = append(tlds, tld)
		}
	}
	jsonSvrData, err := json.Marshal(tlds)
	assert(err)

	fmt.Printf("%s", jsonSvrData)
}
