package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourbasic/bit"
	"pmkit/internal/app/model"
	"pmkit/internal/app/service"
	"pmkit/internal/pkg"
)

var projectService service.ProjectService

type ProjectController struct {
}

func (pc *ProjectController) RestController() string {
	return "/v1/project"
}

func (pc *ProjectController) SaveProject() (string, string, *bit.Set, func(*fiber.Ctx) error) {
	var permissions = bit.New()
	return fiber.MethodPost, "/add", permissions, func(c *fiber.Ctx) error {
		var project = new(model.Project)
		err := c.BodyParser(project)
		if err != nil {
			return err
		}
		uid := pkg.GetCurrentUserId(c)
		project.CreateBy = uid
		project, err = projectService.SaveProject(project)
		if err != nil {
			return err
		}
		return c.JSON(pkg.Success(project))
	}
}
