package controllers

import (
	"api/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// TargetController operations about target
type TargetController struct {
	beego.Controller
}

// Post function
// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (c *TargetController) Post() {
	//var ob models.Object
	//json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	//objectid := models.AddOne(ob)
	//c.Data["json"] = map[string]string{"ObjectId": objectid}

	var targets []models.Target
	json.Unmarshal(c.Ctx.Input.RequestBody, &targets)

	fmt.Println("---> ", targets)

	c.Data["json"] = `{"id": 301}`
	c.ServeJSON()
}
