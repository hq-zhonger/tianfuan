package CloudInspection

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

type CloudInspection struct {
	ThreatBook
	VirusTotal
	FilePath string
	Length   int64
	Result   chan [][]string
}

type ThreatBook struct {
	Api         string
	Sha256      string
	SandBoxType []string
}

type VirusTotal struct {
	Api    string
	Id     string
	Sha256 string
}

func (c *CloudInspection) Run() {
	c.VirusTotal.Api = "764220630042a1105de38580275e6767624bea32272b23dce31e377a19d09d51"
	c.ThreatBook.Api = "4f9366d9e169434194634b09540d736d94603da7f2b94996bd4d10adc929ced1"
	c.Result = make(chan [][]string, 3)
	defer close(c.Result)

	stat, err := os.Stat(c.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	c.Length = stat.Size()
	if c.Length <= 33554432 && c.Length != 0 {
		go c.RunVirusTotal()
		c.RunThreatBook()
		return
	} else {
		if c.Length <= 52428800 && c.Length != 0 {
			c.RunThreatBook()
			return
		} else {
			if c.Length <= 104857600 && c.Length != 0 {
				return
			} else {
				fmt.Println("超过最大长度限制100M")
				return
			}
		}
	}
}

// RunVirusTotal 32M
func (c *CloudInspection) RunVirusTotal() {
	c.VirusTotalUploadFile()
	c.VirusTotalSearchSha256()
	select {
	case <-time.NewTicker(time.Second * 60).C:
		c.VirusTotalSearchMultiengines()
		return
	}
}

// RunThreatBook 50M
func (c *CloudInspection) RunThreatBook() {
	c.ThreatBookUploadFile()
	select {
	case <-time.NewTicker(time.Second * 60).C:
		c.ThreatBookSearchMultiengines()
		return
	}
}

func (c *CloudInspection) VirusTotalUploadFile() {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("accept", "application/json")
	writer.WriteField("content-type", "multipart/form-data")

	fw, _ := writer.CreateFormFile("file", c.FilePath)
	f, _ := os.Open(c.FilePath)
	_, err := io.Copy(fw, f)
	if err != nil {
		fmt.Println("error when append file", err.Error())
		return
	}

	writer.Close()

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://www.virustotal.com/api/v3/files", body)

	// Headers
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("x-apikey", c.VirusTotal.Api)

	// Fetch Request
	resp, err := client.Do(req)

	all, _ := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	c.VirusTotal.Id = gjson.GetBytes(all, "data.id").String()
}

func (c *CloudInspection) VirusTotalSearchSha256() {
	url := fmt.Sprintf("https://www.virustotal.com/api/v3/analyses/%s", c.VirusTotal.Id)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", c.VirusTotal.Api)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	c.VirusTotal.Sha256 = gjson.GetBytes(body, "meta.file_info.sha256").String()
}

func (c *CloudInspection) VirusTotalSearchMultiengines() {
	url := fmt.Sprintf("https://www.virustotal.com/api/v3/files/%s", c.VirusTotal.Sha256)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", c.VirusTotal.Api)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	MalwareType := gjson.GetBytes(body, "data.attributes.crowdsourced_yara_results.0.rule_name").String()
	MalwareFamily := gjson.GetBytes(body, "data.attributes.crowdsourced_yara_results.0.ruleset_name").String()

	Total1 := gjson.GetBytes(body, "data.attributes.last_analysis_stats.malicious")
	Total2 := gjson.GetBytes(body, "data.attributes.last_analysis_stats.undetected")
	Positives := Total1.Int() + Total2.Int()

	ThreatLevel := "安全"

	if Total1.Int() > 0 {
		if MalwareType == "" || MalwareFamily == "" {
			ThreatLevel = "可疑"
		}
	}

	if MalwareType != "" || MalwareFamily != "" {
		ThreatLevel = "恶意"
	}

	c.Result <- [][]string{{"VirusTotal", ThreatLevel, strconv.FormatInt(Positives, 10), strconv.FormatInt(Positives, 10), Total1.String(), MalwareType, MalwareFamily}}
}

func (c *CloudInspection) ThreatBookUploadFile() {
	c.ThreatBook.SandBoxType = []string{
		"win7_sp1_enx64_office2013",
		"win7_sp1_enx86_office2013",
		"win7_sp1_enx86_office2010",
		"win7_sp1_enx86_office2007",
		"win7_sp1_enx86_office2003",
		"win10_1903_enx64_office2016",
		"ubuntu_1704_x64",
		"centos_7_x64",
		"kylin_desktop_v10",
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("user_sandbox_type", "win10_1903_enx64_office2016")
	writer.WriteField("apikey", c.ThreatBook.Api)
	writer.WriteField("run_time", "60")

	fw, _ := writer.CreateFormFile("file", c.FilePath)
	f, _ := os.Open(c.FilePath)
	_, err := io.Copy(fw, f)
	if err != nil {
		fmt.Println("error when append file", err.Error())
		return
	}

	writer.Close()

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://api.threatbook.cn/v3/file/upload", body)

	// Headers
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	getBytes := gjson.GetBytes(respBody, "data.sha256")
	c.ThreatBook.Sha256 = getBytes.String()
}

func (c *CloudInspection) ThreatBookSearchMultiengines() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.threatbook.cn/v3/file/report/multiengines?apikey=%s&sha256=%s", c.ThreatBook.Api, c.ThreatBook.Sha256), nil)
	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	ThreatLevel := gjson.GetBytes(respBody, "data.multiengines.threat_level").String()

	switch ThreatLevel {
	case "clean":
		ThreatLevel = "安全"
		break
	case "suspicious":
		ThreatLevel = "可疑"
		break
	case "malicious":
		ThreatLevel = "恶意"
		break
	default:
		break
	}

	Total1 := gjson.GetBytes(respBody, "data.multiengines.total").String()
	Total2 := gjson.GetBytes(respBody, "data.multiengines.total2").String()
	Positives := gjson.GetBytes(respBody, "data.multiengines.positives").String()
	MalwareType := gjson.GetBytes(respBody, "data.multiengines.malware_type").String()
	MalwareFamily := gjson.GetBytes(respBody, "data.multiengines.malware_family").String()

	c.Result <- [][]string{{"微步云", ThreatLevel, Total1, Total2, Positives, MalwareType, MalwareFamily}}
}
