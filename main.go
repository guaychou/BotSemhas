package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"strconv"
)

const AppVersion = "Lordchou Semhas Bot 1.0.0"
func main() {
	h, m, _ := time.Now().Clock()
	var result map[string]interface{}
	namaPtr := flag.String("nama", "", "Nama mahasiswa \nContoh: John Doe")
	nimPtr := flag.String("nim", "", "Nomer Induk Mahasiswa \nContoh: 16515020xxxxxxx")
	dokumenPtr := flag.String("dokumen", "", "Nomor Dokumen undangan semhas \nContoh: 279/UN10.F15.11/PP/2020")
	timePtr:=flag.String("time", "", "Waktu pembukaan audiensi semhas\nContoh: 15.00")
	version := flag.Bool("version", false, "prints current Lordchou Semhas Bot Version")
	help := flag.Bool("help", false, "prints usage of Lordchou Semhas Bot")
	flag.Parse()
	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *namaPtr == "" || *nimPtr == "" || *dokumenPtr == ""|| *timePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println("Lordchou Semhas CLI Tools v1.0")
	fmt.Println("Now :",time.Now().Local())
	timeValues:=strings.Split(*timePtr,".")
	jam64,_:=strconv.ParseInt(timeValues[0],10,0)
	jam:=int(jam64)
	menit64,_:=strconv.ParseInt(timeValues[1],10,0)
	menit:=int(menit64)
        fmt.Println("Waiting for the time to execute . . . ")
	for h!=jam || m!=menit{
		h, m, _ = time.Now().Clock()
		//log.Println(time.Now())
	}
	if h == jam && m == menit {
		log.Println("========================INFO=========================")
		log.Println("Nama: " + *namaPtr + " | Nim: " + *nimPtr)
		hmifUrl := "http://hmif.filkom.ub.ac.id/pendataan-audiens-semhas/daftar?undangan=" + *dokumenPtr
		formData := url.Values{
			"nama": {*namaPtr},
			"nim":  {*nimPtr},
		}
		var client = &http.Client{}
		payload := bytes.NewBufferString(formData.Encode())
		request, err := http.NewRequest("POST", hmifUrl, payload)
		if err != nil {
			log.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			json.Unmarshal([]byte(bodyString), &result)
			status := result["status"].(string)
			message := result["msg"].(string)
			fmt.Println("Status: ", status)
			fmt.Println("Message: ", message)
		}
	}
}
