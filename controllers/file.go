// file
package controllers

import (
	"fmt"
	"github.com/ginuerzh/mbbcloud/errors"
	"github.com/ginuerzh/mbbcloud/models"
	"github.com/ginuerzh/weedo"
	"io"
	"labix.org/v2/mgo/bson"
	"log"
	"strconv"
	//"time"
)

type FileController struct {
	BaseController
}

func (this *FileController) Upload() {
	filedata, header, err := this.GetFile("file")
	if err != nil {
		this.Data["json"] = this.response(nil, &errors.FileNotFoundError)
		this.ServeJson()
		return
	}

	fid, size, err := weedo.AssignUpload(header.Filename, filedata)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = this.response(nil, &errors.FileUploadError)
		this.ServeJson()
		return
	}

	var file models.File
	file.Fid = fid
	file.Name = header.Filename
	file.ContentType = header.Header.Get("Content-Type")
	file.Size = size
	file.Md5 = this.fileMd5(filedata)
	file.UploadDate = bson.Now()
	if err := file.Save(); err != nil {
		fmt.Println(err)
		this.Data["json"] = this.response(nil, err)
		this.ServeJson()
		return
	}

	url := this.fileUrl(fid)

	fileInfo := map[string]interface{}{
		"fid": fid, "name": header.Filename, "size": size,
		"url": url, "thumbnailUrl": url, "deleteUrl": url,
		"deleteType": "DELETE"}
	r := map[string]interface{}{"files": []interface{}{fileInfo}}

	this.Data["json"] = r
	this.ServeJson()
}

func (this *FileController) Download() {
	//fid := this.Ctx.Input.Param[":all"]
	id := this.Ctx.Input.Param[":id"]
	key := this.Ctx.Input.Param[":key"]

	fid := id + "," + key
	var f models.File

	if err := f.FindOneBy("fid", fid); err != nil {
		this.Data["json"] = this.response(nil, err)
		this.ServeJson()
		return
	}

	file, err := weedo.Download(fid)
	if err != nil {
		this.Data["json"] = this.response(nil, &errors.FileNotFoundError)
		this.ServeJson()
		return
	}
	defer file.Close()
	log.Println("content-length:" + strconv.FormatInt(f.Size, 10))
	this.Ctx.ResponseWriter.Header().Set("Content-Length", strconv.FormatInt(f.Size, 10))
	//this.Ctx.ResponseWriter.Header().Set("Content-Type", f.ContentType)
	//this.Ctx.ResponseWriter.Header().Set("Content-Disposition", "filename="+f.Name)
	io.Copy(this.Ctx.ResponseWriter, file)
}

func (this *FileController) Delete() {
	fid := this.Ctx.Input.Param[":all"]

	file := models.File{Fid: fid}
	file.Delete()

	this.Data["json"] = this.response(nil, nil)
	this.ServeJson()
}
