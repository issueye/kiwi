package router

import (
	v1 "kiwi/internal/app/project/controller/v1"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	InitProjectRouter(r)
}

func InitProjectRouter(r *gin.RouterGroup) {

	webhook := r.Group("webhook")
	{
		webhook.POST("gitlab", func(ctx *gin.Context) {
		})
	}

	project := r.Group("project")
	// project.Use(middleware.AuthMiddleware())
	{
		project.GET(":id", v1.GetProject)
		project.POST("list", v1.ProjectList)
		project.POST("", v1.CreateProject)
		project.PUT("", v1.UpdateProject)
		project.DELETE(":id", v1.DeleteProject)

		// Branch routes
		branch := project.Group("branch")
		{
			branch.POST("", v1.CreateBranch)
			branch.PUT("", v1.UpdateBranch)
			branch.PUT("status", v1.UpdateBranchStatus)
			branch.DELETE(":id", v1.DeleteBranch)
			branch.POST("list", v1.BranchList)
		}

		// Tag routes
		tag := project.Group("tag")
		{
			tag.POST("", v1.CreateTag)
			tag.PUT("", v1.UpdateTag)
			tag.PUT("release", v1.UpdateTagReleaseStatus)
			tag.DELETE(":id", v1.DeleteTag)
			tag.POST("list", v1.TagList)
		}

		// Version routes
		version := project.Group("version")
		{
			version.POST("", v1.CreateVersion)
			version.PUT("", v1.UpdateVersion)
			version.PUT("build", v1.UpdateVersionBuildStatus)
			version.DELETE(":id", v1.DeleteVersion)
			version.POST("list", v1.VersionList)
		}

		robot := project.Group("robot")
		{
			robot.POST("", v1.CreateRobot)
			robot.PUT("", v1.UpdateRobot)
			robot.PUT("state", v1.UpdateRobotState)
			robot.DELETE(":id", v1.DeleteRobot)
			robot.POST("list", v1.RobotList)
		}
	}
}
