// Golang 示例
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/twindagger/httpsig.v1"
)

type SigAuth struct {
	KeyID    string
	SecretID string
}

func (auth *SigAuth) Sign(r *http.Request) error {
	headers := []string{"(request-target)", "date"}
	signer, err := httpsig.NewRequestSigner(auth.KeyID, auth.SecretID, "hmac-sha256")
	if err != nil {
		return err
	}
	return signer.SignRequest(r, headers, nil)
}

type Asset struct {
	Hostname  string   `json:"hostname"`
	IP        string   `json:"ip"`
	Platform  string   `json:"platform"`
	Protocols []string `json:"protocols"`
	AdminUser string   `json:"admin_user"`
	Nodes     []string `json:"nodes"`
	IsActive  bool     `json:"is_active"`
}

var TestAsset = &Asset{
	Hostname: "test-114515-another",
	IP:       "10.2.22.253",
	Platform: "Linux",
	Protocols: []string{
		"ssh/22",
	},
	AdminUser: "7f492f30-e68b-41e9-9072-dac70b9eb522",
	Nodes: []string{
		"3740b185-7ab5-4849-ae0f-52f574458a39",
	},
	IsActive: true,
}

func CreateAsset(jms_url string, auth *SigAuth) {
	url := jms_url + "/api/v1/assets/assets/"
	gmt_fmt := "Mon, 02 Jan 2006 15:04:05 GMT"
	buff := new(bytes.Buffer)
	if err := json.NewEncoder(buff).Encode(TestAsset); err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, url, buff)
	// req.Header.Add("User-Agent", "python-requests/2.26.0")
	req.Header.Add("Date", time.Now().UTC().Format(gmt_fmt))
	req.Header.Add("Accept", "application/json")
	// This is immportant, default is www-urlencoded
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-JMS-ORG", "00000000-0000-0000-0000-000000000002")
	if err != nil {
		log.Fatal(err)
	}
	if err := auth.Sign(req); err != nil {
		log.Fatal(err)
	}
	for k, values := range req.Header {
		for _, v := range values {
			log.Printf("header: %s: %s", k, v)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Printf("status code: %d", resp.StatusCode)
	log.Printf("request method: %s", resp.Request.Method)
	io.Copy(os.Stdout, resp.Body)
}

func GetUserInfo(jms_url string, auth *SigAuth) {
	url := jms_url + "/api/v1/users/users/"
	gmt_fmt := "Mon, 02 Jan 2006 15:04:05 GMT"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "python-requests/2.26.0")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Date", time.Now().UTC().Format(gmt_fmt))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-JMS-ORG", "00000000-0000-0000-0000-000000000002")
	if err != nil {
		log.Fatal(err)
	}
	if err := auth.Sign(req); err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func main() {
	jms_url := "http://localhost"
	auth := SigAuth{
		KeyID:    "3E7076FB-D5A4-4EC3-B34A-B50DFA546EF9",
		SecretID: "7F6644D2-F6BE-447D-95A1-DEC718DEEC26",
	}
	// GetUserInfo(jms_url, &auth)
	CreateAsset(jms_url, &auth)
}
