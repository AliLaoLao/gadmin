package api

import (
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/code"
)

const (
	UNDEFIND_POLICY_NAME = "未命名"
)

type PolicyController struct {
	BaseController
}

func (c *PolicyController) Get() {
	page := c.Request.GetInt("page", 1)
	if page < 1 {
		page = 1
	}
	limit := c.Request.GetInt("limit", 10)

	var list struct {
		List  []model.Policy `json:"items"`
		Total int            `json:"total"`
	}

	list.List, list.Total = model.GetPolicyList(page, limit, UNDEFIND_POLICY_NAME)

	Success(c.Controller, list)
}

func (c *PolicyController) Post() {
	Success(c.Controller, "Post")
}

func (c *PolicyController) Put() {
	data := c.Request.GetJson()
	name := data.GetString("name")
	path := data.GetString("policy")
	glog.Debug(name, path)
	if name == UNDEFIND_POLICY_NAME {
		Fail(c.Controller, code.RESPONSE_ERROR)
	} else {
		err := model.UpdatePolicyByFullPath(path, name)
		if err != nil {
			Fail(c.Controller, code.RESPONSE_ERROR, err.Error())
		}
		Success(c.Controller, "修改成功")
	}

}

func (c *PolicyController) Delete() {
	Success(c.Controller, "Delete")
}