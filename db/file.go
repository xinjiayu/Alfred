package db
//负责数据库中的文件操作，提供增删改查的功能实现

import (
	mydb "Alfred/db/mysql"
	"database/sql"
	"fmt"
)

//UploadFileToDB:完成文件向数据库上传的功能
func UploadFileToDB(filehash string,filename string,filesize int64,fileaddr string) bool {
	stmt,err:=mydb.DBConn().Prepare("insert into tbl_file (`file_sha1`,`file_name`,`file_size`," +
		"`file_addr`,`status`) values (?,?,?,?,1)")
	if err !=nil {
		fmt.Printf("Failed to prepare sql,err:"+err.Error())
		return false
	}
	defer stmt.Close()

	res,err:=stmt.Exec(filehash,filename,filesize,fileaddr)
	if err !=nil {
		fmt.Printf("Failed to excute sql,err:"+err.Error())
		return false
	}

	if rf,err:=res.RowsAffected();err==nil {
		if rf<=0 {
			fmt.Println("File with filehash"+filehash+"has already been uploaded")
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

//GetMetaFromDB：获取文件元信息
func GetMetaFromDB(filehash string) (*TableFile,error) {
	stmt,err:=mydb.DBConn().Prepare("select file_sha1,file_addr,file_name,file_size from " +
		"tbl_file where file_sha1=? and status =1 limit 1")
	if err !=nil {
		fmt.Printf("Failed to prepare sql,err:"+err.Error())
		return nil,err
	}
	defer stmt.Close()

	resfile:=TableFile{}
	err=stmt.QueryRow(filehash).Scan(&resfile.FileHash,&resfile.FileAddr,&resfile.FileName,&resfile.FileSize)
	if err !=nil {
		fmt.Printf("Failed to excute sql,err:"+err.Error())
		return nil,err
	}
	return &resfile,nil
}