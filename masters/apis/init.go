package apis

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/middlewares"
)

func Init(r *mux.Router, db *sql.DB) {
	//menuRepo := repositories.InitMenuRepoImpl(db)
	//menuUseCase := usecases.InitMenuUseCase(menuRepo)
	//controllers.MenuControll(r, menuUseCase)
	//
	//servicesRepo := repositories.InitServiceRepoImpl(db)
	//servicesUseCase := usecases.InitServiceUseCase(servicesRepo)
	//controllers.ServicesControll(r, servicesUseCase)
	//
	//categoriesRepo := repositories.InitCategoriesRepoImpl(db)
	//categoriesUseCase := usecases.InitCategoryUseCase(categoriesRepo)
	//controllers.CategoriesControll(r, categoriesUseCase)
	//
	//transactionRepo := repositories.InitTransactionRepoImpl(db)
	//transactionUseCase := usecases.InitTransactionUseCase(transactionRepo)
	//controllers.TransactionControll(r, transactionUseCase)
	//
	//usersRepo := repositories.InitUserRepoImpl(db)
	//usersUseCases := usecases.InitUsersUseCase(usersRepo)
	//controllers.UsersControll(r, usersUseCases)

	r.Use(middlewares.ActivityLogMiddleware)
}
