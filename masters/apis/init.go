package apis

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers"
	productCategoryControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/productCategory"
	productControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/productController"
	storeCategoryControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/storeCategory"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/middlewares"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
	productRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/product"
	productCategoryRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/productCategory"
	storerepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/store"
	storeCategoryRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/storeCategory"
	walletrepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/wallet"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
	productUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/product"
	productCategoryUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/productCategory"
	storeusecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/store"
	storeCategoryUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/storeCatergory"
	walletusecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/wallet"
)

func Init(r *mux.Router, db *sql.DB) {
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
	// productCategory
	productCategoryRepo := productCategoryRepositories.InitProductCategoryRepoImpl(db)
	productCategoryUseCase := productCategoryUsecases.InitProductCategoryUseCase(productCategoryRepo)
	productCategoryController := productCategoryControllers.ProductCategoryController(productCategoryUseCase)
	productCategoryController.ProductCategoryAPI(r)
	// product
	productRepo := productRepositories.InitProductRepoImpl(db)
	productUseCase := productUsecases.InitProductUseCaseImpl(productRepo)
	productController := productControllers.InitProductController(productUseCase)
	productController.ProductAPI(r)

	r.Use(middlewares.ActivityLogMiddleware)
}
