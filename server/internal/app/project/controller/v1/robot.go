package v1

import (
	"kiwi/internal/app/project/logic"
	"kiwi/internal/app/project/requests"
	"kiwi/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRobot doc
//
//	@tags			机器人管理
//	@Summary		创建信息
//	@Description	创建信息
//	@Produce		json
//	@Param			body	body		requests.CreateRobot	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/robot [post]
//	@Security		ApiKeyAuth
func CreateRobot(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateRobot()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateRobot(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateRobot doc
//
//	@tags			机器人管理
//	@Summary		更新信息
//	@Description	更新信息
//	@Produce		json
//	@Param			body	body		requests.UpdateRobot	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/robot [put]
//	@Security		ApiKeyAuth
func UpdateRobot(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateRobot()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateRobot(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateRobotState doc
//
//	@tags			机器人管理
//	@Summary		更新状态
//	@Description	更新状态
//	@Produce		json
//	@Param			id		path	int	true	"分支ID"
//	@Param			status	body	string	true	"状态"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/robot/state [put]
//	@Security		ApiKeyAuth
func UpdateRobotState(c *gin.Context) {
	ctl := controller.New(c)

	id := c.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	var State struct {
		State int `json:"state"`
	}

	err = ctl.Bind(&State)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateRobotStatus(uint(i), State.State)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteRobot doc
//
//	@tags			机器人管理
//	@Summary		删除信息
//	@Description	删除信息
//	@Produce		json
//	@Param			id		path	int	true	"分支ID"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/robot/{id} [delete]
//	@Security		ApiKeyAuth
func DeleteRobot(c *gin.Context) {
	ctl := controller.New(c)

	id := c.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.DeleteRobot(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// RobotList doc
//
//	@tags			机器人管理
//	@Summary		信息列表
//	@Description	信息列表
//	@Produce		json
//	@Param			body	body		requests.QueryRobot	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/project/robot/list [post]
//	@Security		ApiKeyAuth
func RobotList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryRobot()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.RobotList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}
