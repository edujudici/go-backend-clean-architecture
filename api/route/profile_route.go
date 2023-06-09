package route

import (
	"time"

	"github.com/edujudici/go-backend-clean-architecture/api/controller"
	"github.com/edujudici/go-backend-clean-architecture/bootstrap"
	"github.com/edujudici/go-backend-clean-architecture/domain"
	"github.com/edujudici/go-backend-clean-architecture/mongo"
	"github.com/edujudici/go-backend-clean-architecture/repository"
	"github.com/edujudici/go-backend-clean-architecture/usecase"
	"github.com/gin-gonic/gin"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
