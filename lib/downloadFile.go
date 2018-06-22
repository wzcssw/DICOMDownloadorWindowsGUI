package lib

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/lxn/walk"
)

const DownloadsDir = "./"

func DownloadFile(url, dir string, ch chan bool) error {
	defer func() {
		<-ch
	}()
	stringArray := strings.Split(url, "/")
	fileName := stringArray[len(stringArray)-1]

	os.MkdirAll(DownloadsDir+dir, os.ModePerm)

	out, err := os.Create(DownloadsDir + dir + "/" + fileName)
	if err != nil {
		return err
	}
	defer out.Close()
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// concurrent 下载并发数
func DownloadSeriesFile(series []Series, dir string, concurrent int, outTE *walk.TextEdit,btn *walk.PushButton) {
	// 下载文件
	fmt.Println("downloading...")
	ch := make(chan bool, concurrent)
	for _, serie := range series {
		for _, instance := range serie.InstanceList {
			ch <- true
			go DownloadFile(instance.ImageId, dir, ch)
		}
	}
	outTE.SetText(outTE.Text()+"下载完成！"+"\r\n")
	btn.SetEnabled(true)
}

// CountSeriesFile
func CountSeriesFile(series []Series) int {
	count := 0
	for _, serie := range series {
		count += len(serie.InstanceList)
	}
	return count
}


/*
	go build -ldflags="-H windowsgui"
	
	如给exe文件加上自己喜欢的图标，命令为：rsrc -manifest main.manifest –ico icon.ico -o rsrc.syso
*/