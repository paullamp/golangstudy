package models

import (
	"fmt"

	"github.com/astaxie/beego"

	"github.com/weilaihui/fdfs_client"
)

func BeegoUploadByFilename(filename string) (groupname string, FieldId string, err error) {
	fdfsClient, err := fdfs_client.NewFdfsClient("client.conf")
	if err != nil {
		fmt.Printf("New FdfsClient error %s\n", err.Error())
		return
	}

	uploadResponse, errUpload := fdfsClient.UploadByFilename(filename)
	if errUpload != nil {
		fmt.Printf("UploadByfilename error %s\n", errUpload)
	}

	beego.Info(uploadResponse.GroupName)
	beego.Info(uploadResponse.RemoteFileId)
	// fdfsClient.DeleteFile(uploadResponse.RemoteFileId)
	groupname = uploadResponse.GroupName
	FieldId = uploadResponse.RemoteFileId
	return
}
