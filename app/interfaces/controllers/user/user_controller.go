package user

// import (
// 	"github.com/{}/interfaces/database"
// 	"github.com/{}/interfaces/database/config"
// 	"github.com/{}/models"
// 	"github.com/{}/pkg/module/dto/response"

// 	// "github.com/{}/models"
// 	// "github.com/{}/pkg/module/dto/response"
// 	"github.com/{}/usecase/service"
// )

// type UserController struct {
// 	Interactor service.UserInteractor
// }

// func NewUserController(db config.DB) *UserController {
// 	return &UserController{
// 		Interactor: service.UserInteractor{
// 			DB:              &config.DBRepository{DB: db},
// 			UserRepo:        &database.ActiveUserRepository{},
// 			HistoryUserRepo: &database.HistoryUserRepository{},
// 		},
// 	}
// }

// func (uc *UserController) convertActiveUserToDTO(obj *models.ActiveUserModel) response.ActiveUserEntity {
// 	return response.ActiveUserEntity{

// 	}

// }

// func (uc *UserController) convertActiveUserToDTOs(obj []*models.ActiveUserModel) []response.ActiveUserEntity {
// 	UEs := make([]response.ActiveUserEntity, len(obj))
// 	for i, v := range obj {
// 		// ue := response.ActiveUserEntity{

// 		// }
// 		UEs[i] = uc.convertActiveUserToDTO(v)
// 	}
// 	return UEs
// }
