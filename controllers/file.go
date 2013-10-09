// file
package controllers

import (
	"fmt"
	"github.com/ginuerzh/mbbcloud/errors"
	"github.com/ginuerzh/mbbcloud/models"
	"github.com/ginuerzh/weedo"
	"io"
	"log"
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

	log.Println(header.Filename)
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
	file.Size = size
	if err := file.Save(); err != nil {
		fmt.Println(err)
		this.Data["json"] = this.response(nil, err)
		this.ServeJson()
		return
	}

	url := "http://" + this.Ctx.Request.Host + "/file/" + fid

	fileInfo := map[string]interface{}{
		"fid": fid, "name": header.Filename, "size": size,
		"url": url, "thumbnailUrl": url, "deleteUrl": url,
		"deleteType": "DELETE"}
	r := map[string]interface{}{"files": []interface{}{fileInfo}}

	this.Data["json"] = r
	//log.Println(this.Data["json"])
	this.ServeJson()
}

func (this *FileController) Download() {
	fid := this.Ctx.Input.Param[":all"]
	file, err := weedo.Download(fid)
	if err != nil {
		this.Data["json"] = this.response(nil, &errors.FileNotFoundError)
		this.ServeJson()
		return
	}
	//url, _ := weedo.GetUrl(fid)
	//this.Redirect(url, 302)
	defer file.Close()

	io.Copy(this.Ctx.ResponseWriter, file)
}

func (this *FileController) Delete() {
	fid := this.Ctx.Input.Param[":all"]

	file := models.File{Fid: fid}
	if err := file.Delete(); err != nil {
		log.Println(err)
	}

	if err := weedo.Delete(fid); err != nil {
		log.Println(err)
		this.Data["json"] = this.response(nil, &errors.DbError)
	}

	log.Println("delete file: " + fid)
	this.Data["json"] = this.response(nil, nil)
	this.ServeJson()
}
