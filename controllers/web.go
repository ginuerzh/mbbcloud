// web
package controllers

type WebController struct {
	BaseController
}

func (this *WebController) Get() {
	this.Layout = "base.html"
	this.TplNames = "pub.html"
}
