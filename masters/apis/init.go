package apis

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers"

	// productCategoryControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/productCategory"
	// productControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/productController"
	// storeCategoryControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/storeCategory"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/middlewares"

	feedbackRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
	poinRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
	ratingRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"

	// productCategoryRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/productCategory"
	// storeCategoryRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/storeCategory"
	feedbackUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
	poinUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
	ratingUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
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

	// // user
	// userRepo := repositories.InitUserRepoImpl(db)
	// userUseCases := usecases.InitUsersUseCase(userRepo)
	// usersController := controllers.UsersController(userUseCases)
	// usersController.Authenticate(r)
	// // wallet
	// walletRepo := walletrepositories.InitWalletRepoImpl(db)
	// walletUseCases := walletusecases.InitWalletUseCase(walletRepo)
	// walletController := controllers.WalletController(walletUseCases)
	// walletController.WalletApi(r)
	// // store
	// storeRepo := storerepositories.InitStoreRepoImpl(db)
	// storeUseCase := storeusecases.InitStoreUseCase(storeRepo)
	// storeController := controllers.StoreController(storeUseCase)
	// storeController.StoreAPI(r)
	// // storeCategory
	// storeCategoryRepo := storeCategoryRepositories.InitStoreCategoryRepoImpl(db)
	// storeCategoryUseCase := storeCategoryUsecases.InitStoreCategoryUseCase(storeCategoryRepo)
	// storeCategoryController := storeCategoryControllers.StoreCategoryController(storeCategoryUseCase)
	// storeCategoryController.StoreCategoryAPI(r)

	//feedback
	feedbackRepo := feedbackRepositories.InitFeedbackImpl(db)
	feedbackUseCase := feedbackUsecases.InitFeedbackUsecase(feedbackRepo)
	feedbackController := controllers.FeedbacksController(feedbackUseCase)
	feedbackController.FeedbackAPI(r)

	//poin
	poinRepo := poinRepositories.InitPoinRepoImpl(db)
	poinUseCase := poinUsecases.InitPoinUsecase(poinRepo)
	poinController := controllers.PointsController(poinUseCase)
	poinController.PointAPI(r)

	//rating
	ratingRepo := ratingRepositories.InitRatingRepoImpl(db)
	ratingUseCase := ratingUsecases.InitRatingUsecase(ratingRepo)
	ratingController := controllers.RatingController(ratingUseCase)
	ratingController.RatingAPI(r)

	r.Use(middlewares.ActivityLogMiddleware)
}
