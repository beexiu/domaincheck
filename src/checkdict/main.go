package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

func printHelp(err int) {
	if err > 0 {
		fmt.Printf("Parameter error.\n\n")
	}
	fmt.Printf("Usage: checkdict tld file\n")
	fmt.Printf("    tld     Domains as com, net\n")
	fmt.Printf("    file    Dictionary file\n")
}

func main() {
	if len(os.Args) != 3 {
		printHelp(1)
		os.Exit(1)
	}

	tlds := os.Args[1]
	dict := os.Args[2]

	tldinfo, err := GetTLD(tlds, "./conf/tld.org.json")
	assert(err)
	if tldinfo.WhoisServer == `` || tldinfo.WhoisServer == "null" {
		fmt.Printf("\"%s\" whois server is empty.\n", tlds)
		os.Exit(1)
	}

	resultFile, _ := os.Create("./data/" + tldinfo.Tld + "_" + time.Now().Format("20060102150405") + "_result.txt")

	fileDict, err := os.Open(dict)
	defer fileDict.Close()

	scanner := bufio.NewScanner(fileDict)
	for scanner.Scan() {
		line := scanner.Text()

		// // 不用多线程，会被Block的
		dm := query(line, tldinfo)
		if dm != "" {
			resultFile.WriteString(dm + "\n")
		}

		time.Sleep(1000 * time.Millisecond)
	}

	if err := scanner.Err(); err != nil {
		assert(err)
	}

	resultFile.Close()
}

func query(line string, tldinfo TLD) string {
	conn, err := net.DialTimeout("tcp", tldinfo.WhoisServer+":43", 30*time.Second)
	if err != nil {
		fmt.Printf("connect error :%s  AAA\n", err.Error())
		return ""
	}
	if conn == nil {
		fmt.Printf("connect error")
		return ""
	}
	defer conn.Close()

	line = strings.Trim(line, " ")
	line = strings.Trim(line, "\n")
	domain := line + "." + tldinfo.Tld

	fmt.Fprintf(conn, domain+"\r\n")

	time.Sleep(time.Second)
	var buf = make([]byte, 65536)
	n, err := conn.Read(buf)

	if err == nil {
		newstr := string(buf[0 : n-1])

		reg := regexp.MustCompile(tldinfo.Patterns.NotRegistered)
		re := reg.FindAllString(newstr, -1)
		if re == nil {
			fmt.Printf(domain + "  has been registed\n")
		} else {
			fmt.Printf(domain + " can be regist!!can be regist!!can be regist!! \n")
			return domain
		}
	} else {
		fmt.Printf(err.Error())
	}

	return ""
}
