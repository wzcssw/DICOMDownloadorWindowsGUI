package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"justtest/lib"
	"encoding/json"
	"fmt"
)
// 网络不好会直接挂掉
// go build -ldflags="-H windowsgui"
const DICOMServerURL string = "http://dicomup.tongxinyiliao.com/api/getByFilmNo"
var con int = 30
var inTE *walk.LineEdit
var outTE *walk.TextEdit
var btn *walk.PushButton


func main() {
	MainWindow{
		Title:   "下载DICOM影像",
		MinSize: Size{320, 290},
		Layout:  VBox{},
		Children: []Widget{
			GroupBox{
				Layout: HBox{},
				Children: []Widget{
					LineEdit{AssignTo: &inTE},
					PushButton{
						AssignTo: &btn,
						Text: "开始下载",
						OnClicked: func() {
							// outTE.SetText(strings.ToUpper(inTE.Text()))
							addText("开始下载...")
							btn.SetEnabled(false)
							addText("------------------")
							addText("影像号: "+inTE.Text())

							m := make(map[string]string)
							m["filmno"] = inTE.Text()
							responsedDataBody := lib.SendDicomAPIRequest(DICOMServerURL, m)
							var responsedData lib.DicomAPIRequest
							json.Unmarshal([]byte(responsedDataBody), &responsedData)

							addText(fmt.Sprintf("共发现DICOM文件 %d 个",lib.CountSeriesFile(responsedData.List)))
							addText(fmt.Sprintf("正在下载DICOM文件  并发数:%d",con))
							go lib.DownloadSeriesFile(responsedData.List, inTE.Text(), con, outTE,btn)
							
						},
					},
				},
			},
			TextEdit{
				AssignTo: &outTE,
			},
		},
	}.Run()
}

func addText(str string){
	outTE.SetText(outTE.Text()+str+"\r\n")
}