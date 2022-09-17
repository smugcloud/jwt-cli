package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var jwt string

// ConvertedTimes converts to human readable times
type ConvertedTimes struct {
	Expiration     int64 `json:"exp"`
	Issued         int64 `json:"iat"`
	convExpiration time.Time
	convIssued     time.Time
	json           []byte
	out            *bytes.Buffer
}

func main() {
	validate()
	sp := strings.Split(jwt, ".")
	if len(sp) < 2 {
		fmt.Println("invalid jwt format")
		os.Exit(0)
	}
	var decoded []byte
	var err error
	switch len(sp) {
	case 2:
		decoded, err = base64.RawURLEncoding.DecodeString(sp[0])
	case 3:
		decoded, err = base64.RawURLEncoding.DecodeString(sp[1])
	}
	if err != nil {
		log.Fatalf("error decoding base64 %v", err)
	}
	ct := ConvertedTimes{
		out:  &bytes.Buffer{},
		json: decoded,
	}
	err = json.Unmarshal(decoded, &ct)
	if err != nil {
		log.Fatalf("error unmarshaling %v", err)
	}
	ct.convertTimes()
	ct.writeOut()
}

func (c *ConvertedTimes) writeOut() {
	json.Indent(c.out, c.json, "", "\t")
	c.out.Write([]byte("\n"))
	c.out.Write([]byte(fmt.Sprintf("Expires: %s\n", c.convExpiration)))
	c.out.Write([]byte(fmt.Sprintf("Issued: %s\n", c.convIssued)))
	c.out.WriteTo(os.Stdout)
}
func (c *ConvertedTimes) convertTimes() {
	c.convExpiration = time.Unix(c.Expiration, 0)
	c.convIssued = time.Unix(c.Issued, 0)
}

func validate() {
	if len(os.Args) < 2 {
		fmt.Println("you must provide a jwt")
		myUsage()
	}
	jwt = os.Args[1]
	if jwt == "" {
		fmt.Println("you must provide a jwt")
		flag.Usage()
		os.Exit(0)
	}
}

func myUsage() {
	fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
	fmt.Println("\tjwt-cli <jwt to decode>")
	os.Exit(0)
}
func init() {
	flag.Usage = myUsage
	flag.Parse()
}
