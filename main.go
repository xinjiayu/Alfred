package main

import (
	"Alfred/handler"
	"fmt"
	"net/http"
)

//应用入口
func main(){

	//文件路由
	http.HandleFunc("/file/upload",handler.UploadHandler)
	http.HandleFunc("/file/upload/suc",handler.UploadSuc)
	http.HandleFunc("/file/download",handler.DownloadHandler)
	http.HandleFunc("/file/meta",handler.GetFilemeta)
	http.HandleFunc("/file/update",handler.UpdateFileMeta)
	http.HandleFunc("/file/delete",handler.FileDeleteHandler)

	err:=http.ListenAndServe(":8080",nil)
	if err!=nil {
		fmt.Printf("Failed to start server,err %s",err.Error())
	}
}