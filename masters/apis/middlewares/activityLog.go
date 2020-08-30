package middlewares

import (
	"github.com/inact25/PickMyFood-BackEnd/utils"
	"github.com/inact25/PickMyFood-BackEnd/utils/device"
	ua "github.com/mileusna/useragent"
	"log"
	"net/http"
)

//type LogHandler struct {
//	logRepo repositories.LogRepositories
//}

func ActivityLogMiddleware(next http.Handler) http.Handler {
	//var l *LogHandler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.UserAgent()
		detect := ua.Parse(userAgent)
		userDevice := device.Detect(detect)
		methodDetect := utils.MethodDetect(r)
		//device := utils.GetDeviceType(detect)
		//currentTime := time.Now()
		//timeNow := currentTime.Format("2006.01.02 15:04:05")
		log.Printf("Accessing path %v using %v with : %v\n", r.RequestURI, methodDetect, userDevice)

		//logsData := &models.LogModels{
		//	LogDate:       timeNow,
		//	ApisUri:       r.RequestURI,
		//	Methods:       methodDetect,
		//	LogDevice:     detect.Name,
		//	LogDevVersion: detect.Version,
		//	DeviceOs:      detect.OS,
		//	DeviceType:    device,
		//}
		//
		//l.logRepo.AddNewLogs(logsData)
		// transactionRepo := repositories.InitLogsRepoImpl()
		// usecases.InitLogsUseCase(transactionRepo)
		// LogHandler := LogHandler{LogHand}
		// LogHandler.logUsecases.AddNewLogs(logsData)
		next.ServeHTTP(w, r)
	})
}

//func InitActivityLogMiddleware(logRepo repositories.LogRepositories) *LogHandler {
//	return &LogHandler{logRepo}
//}

//type LogServices struct {
//	db *sql.DB
//}
//
//func NewService(db *sql.DB) *LogServices {
//	return &LogServices{db}
//}
