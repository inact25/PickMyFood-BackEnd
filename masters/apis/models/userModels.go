package models

type User struct {
	UserID        string `json:"userID"`
	UserFirstName string `json:"userFirstName"`
	UserLastName  string `json:"userLastName"`
	UserAddress   string `json:"userAddress"`
	UserPhone     string `json:"userPhone"`
	UserPoin      int    `json:"userPoin"`
	UserImage     string `json:"userImage"`
	UserStatus    string `json:"userStatus"`
	Auth          Auth   `json:"auth"`
}
type Auth struct {
	AuthID      int    `json:"authID"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	UserID      string `json:"userID"`
	UserLevelID int    `json:"userLevelID"`
	UserStatus  string `json:"userStatus"`
	Token       Token  `json:"authentication"`
}

type UserLevel struct {
	UserLevelID int    `json:"userLevelID"`
	Description string `json:"description"`
}
