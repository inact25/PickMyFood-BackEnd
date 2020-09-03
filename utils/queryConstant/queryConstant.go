package utils

const (
	INSERT_USER             = "insert into tb_user (user_id,user_firstname,user_lastname,user_address,user_phone,user_image,user_status) values (?,?,?,?,?,?,?)"
	INSERT_AUTH             = "insert into tb_auth(username,password,user_id) values (?,?,?)"
	SELECT_USER_BY_ID       = "select auth.username, auth.password,user.user_image,user.user_poin,user.user_status, user.user_firstname, user.user_lastname, user.user_phone, user.user_address from tb_user user inner join tb_auth auth on auth.user_id = user.user_id where user.user_id = ?"
	SELECT_ALL_USER         = "SELECT tu.user_id,tu.user_firstname,tu.user_lastname,tu.user_address,tu.user_phone,tu.user_poin,tu.user_image,tu.user_status,ta.username,ta.password,ta.user_level_id,ta.user_status FROM tb_user tu JOIN tb_auth ta ON tu.user_id=ta.user_id"
	UPDATE_USER             = "UPDATE tb_user SET user_firstname=?,user_lastname=?,user_address=?,user_phone=?,user_image=?,user_status=? WHERE user_id=?"
	UPDATE_AUTH             = "UPDATE tb_auth SET username=?,password=? WHERE user_id=?"
	DELETE_AUTH             = "UPDATE tb_auth SET user_status = NA WHERE user_id = ?"
	LOGIN                   = "select user_id, user_level_id, user_status from tb_auth where username = ? and password= ?;"
	SELECT_AUTH_BY_USERNAME = "SELECT * FROM tb_auth WHERE username = ?"
	INSERT_WALLET           = "INSERT INTO tb_wallet (wallet_id,user_id) values (?,?)"
	SELECT_WALLET_USER_ID   = "SELECT tw.wallet_id,tw.amount,tu.user_id,tu.user_firstname,tu.user_lastname,tu.user_address,tu.user_phone,tu.user_poin FROM tb_wallet tw JOIN tb_user tu ON tw.user_id=tu.user_id WHERE tu.user_id = ?"
	UPDATE_AMOUNT_WALLET    = "UPDATE tb_wallet SET amount = ? WHERE user_id = ?"
	UPDATE_POIN_USER        = "UPDATE tb_user SET user_poin = ? WHERE user_id = ?"
	INSERT_TOP_UP           = "INSERT INTO tb_top_up (top_up_id,top_up_amount,user_id,top_up_date) VALUES (?,?,?,?)"
)
