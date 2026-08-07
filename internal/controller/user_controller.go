package controller

import (
	"github.com/klass-lk/ginboot"
	"github.com/klass-lk/test/internal/model"
	"github.com/klass-lk/test/internal/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) Register(group *ginboot.ControllerGroup) {
	group.GET("/:id", c.GetUser)
	group.POST("", c.CreateUser)
}

func (c *UserController) GetUser(ctx *ginboot.Context) (model.User, error) {
	//id := ctx.Param("id")

	// Example of using auth context
	authCtx, err := ctx.GetAuthContext()
	if err != nil {
		return model.User{}, err
	}
	// Use auth context data if needed
	_ = authCtx.UserID

	user, err := c.userService.GetUser(authCtx.UserID)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (c *UserController) CreateUser(ctx *ginboot.Context, request model.User) (model.User, error) {
	user, err := c.userService.CreateUser(request)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
