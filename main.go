package main

import (
	"Alfred/handler"
	"fmt"
	"net/http"
)

//应用入口
func main(){
	//处理静态页面
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	//文件路由
	http.HandleFunc("/file/upload",handler.UploadHandler)
	http.HandleFunc("/file/upload/suc",handler.UploadSuc)
	http.HandleFunc("/file/download",handler.DownloadHandler)
	http.HandleFunc("/file/meta",handler.GetFilemeta)
	http.HandleFunc("/file/update",handler.UpdateFileMeta)
	http.HandleFunc("/file/delete",handler.FileDeleteHandler)

	//用户路由
	http.HandleFunc("/user/signup",handler.SignUpHandler)
	http.HandleFunc("/user/signin",handler.UserSignInHandler)
	http.HandleFunc("/user/info",handler.UserInfoHandler)



	err:=http.ListenAndServe(":8080",nil)
	if err!=nil {
		fmt.Printf("Failed to start server,err %s",err.Error())
	}
}
