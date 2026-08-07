package di

import (
	"github.com/klass-lk/test/internal/controller"
	"github.com/klass-lk/test/internal/model"
	"github.com/klass-lk/test/internal/service"

	"github.com/klass-lk/ginboot"
	"github.com/klass-lk/ginboot/db/inmemory"
)

type Container struct {
	Services Services
}

type Services struct {
	UserService service.UserService
}

type Repository struct {
	UserRepository *inmemory.InMemoryRepository[model.User]
}

func NewContainer(engine *ginboot.Server) {
	repos := InitializeRepositories()
	services := InitializeServices(repos)
	InitializeControllers(services, engine)
}

func InitializeRepositories() *Repository {

	userRepository := inmemory.NewInMemoryRepository[model.User]()
	return &Repository{
		UserRepository: userRepository,
	}

}

func InitializeServices(repos *Repository) *Services {
	userService := service.NewUserService(repos.UserRepository)
	return &Services{
		UserService: userService,
	}
}

func InitializeControllers(services *Services, engine *ginboot.Server) {
	userController := controller.NewUserController(services.UserService)
	engine.RegisterController("users", userController)
}
