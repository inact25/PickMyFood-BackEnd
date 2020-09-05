package apis

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers"
	storeCategoryControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/storeCategory"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/middlewares"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
	storerepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/store"
	storeCategoryRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/storeCategory"
	walletrepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/wallet"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
	storeusecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/store"
	storeCategoryUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/storeCatergory"
	walletusecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/wallet"
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
	//transactionRepo := repositories.InitTransactionRepoImpl(db)
	//transactionUseCase := usecases.InitTransactionUseCase(transactionRepo)
	//controllers.TransactionControll(r, transactionUseCase)

	// user
	userRepo := repositories.InitUserRepoImpl(db)
	userUseCases := usecases.InitUsersUseCase(userRepo)
	usersController := controllers.UsersController(userUseCases)
	usersController.Authenticate(r)
	// wallet
	walletRepo := walletrepositories.InitWalletRepoImpl(db)
	walletUseCases := walletusecases.InitWalletUseCase(walletRepo)
	walletController := controllers.WalletController(walletUseCases)
	walletController.WalletApi(r)
	// store
	storeRepo := storerepositories.InitStoreRepoImpl(db)
	storeUseCase := storeusecases.InitStoreUseCase(storeRepo)
	storeController := controllers.StoreController(storeUseCase)
	storeController.StoreAPI(r)
	// storeCategory
	storeCategoryRepo := storeCategoryRepositories.InitStoreCategoryRepoImpl(db)
	storeCategoryUseCase := storeCategoryUsecases.InitStoreCategoryUseCase(storeCategoryRepo)
	storeCategoryController := storeCategoryControllers.StoreCategoryController(storeCategoryUseCase)
	storeCategoryController.StoreCategoryAPI(r)

	r.Use(middlewares.ActivityLogMiddleware)
}
