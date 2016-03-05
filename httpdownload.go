package httpdownload

import (
	"io"
    "log"
	"net/http"
	"os"
    "strings"
    "strconv"
)
//http GetandSave
func GetandSave(url ,filename string) bool {
	filename = strings.TrimSpace(filename)
	resp, err := http.Get(url)
	if err != nil {
		// panic(err)
		log.Println("Download"+url+"失败!")
		return false
	}
    size:= strconv.FormatInt(resp.ContentLength,10)
	log.Println("Download size:"+size)
	defer resp.Body.Close()


	log.Println("Downloading..."+filename)

	// create file
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// downloading
	if _, err := io.Copy(file, resp.Body); err != nil {
		log.Println("Download 被中断了")
		return false
	}

	log.Println("Download Success!"+url)
	return true
}