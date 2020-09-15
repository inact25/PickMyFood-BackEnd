package utils

const (
	INSERT_USER             = "insert into tb_user (user_id,user_firstname,user_lastname,user_address,user_phone,user_email,user_status) values (?,?,?,?,?,?,?)"
	INSERT_AUTH             = "insert into tb_auth(username,password,user_id) values (?,?,?)"
	SELECT_USER_BY_ID       = "select user.user_id,auth.username, auth.password,user.user_email,user.user_image,user.user_poin,user.user_status, user.user_firstname, user.user_lastname, user.user_phone, user.user_address from tb_user user inner join tb_auth auth on auth.user_id = user.user_id where user.user_id = ?"
	SELECT_ALL_USER         = "SELECT tu.user_id,tu.user_firstname,tu.user_lastname,tu.user_address,tu.user_phone,tu.user_poin,tu.user_email,tu.user_image,tu.user_status,ta.username,ta.password,ta.user_level_id,ta.user_status FROM tb_user tu JOIN tb_auth ta ON tu.user_id=ta.user_id WHERE tu.user_firstname OR tu.user_lastname LIKE ? LIMIT %s ,%s"
	UPDATE_USER             = "UPDATE tb_user SET user_firstname=?,user_lastname=?,user_address=?,user_phone=?,user_image=?,user_status=? WHERE user_id=?"
	UPDATE_AUTH             = "UPDATE tb_auth SET username=?,password=? WHERE user_id=?"
	DELETE_AUTH             = "UPDATE tb_auth SET user_status = NA WHERE user_id = ?"
	LOGIN                   = "select user_id, user_level_id, user_status from tb_auth where username = ? and password= ?;"
	SELECT_AUTH_BY_USERNAME = "SELECT tu.user_id,tu.user_firstname,tu.user_lastname,tu.user_address,tu.user_phone,tu.user_poin,tw.amount,tu.user_email,tu.user_image,tu.user_status,ta.username,ta.password,ta.user_level_id,ta.user_status FROM tb_user tu JOIN tb_auth ta ON tu.user_id=ta.user_id JOIN tb_wallet tw ON tu.user_id=tw.user_id WHERE ta.username = ?"
	INSERT_WALLET           = "INSERT INTO tb_wallet (wallet_id,user_id) values (?,?)"
	SELECT_WALLET_USER_ID   = "SELECT wallet_id,amount,user_id FROM tb_wallet WHERE user_id = ?"
	UPDATE_AMOUNT_WALLET    = "UPDATE tb_wallet SET amount = ? WHERE user_id = ?"
	UPDATE_POIN_USER        = "UPDATE tb_user SET user_poin = ? WHERE user_id = ?"
	INSERT_TOP_UP           = "INSERT INTO tb_top_up (top_up_id,top_up_amount,user_id,top_up_date) VALUES (?,?,?,?)"
	UPDATE_STATUS_TOP_UP    = "UPDATE tb_top_up SET top_up_status = ? WHERE user_id = ?"
	//STORE
	INSERT_STORE                = "INSERT INTO tb_store (store_id,store_name,store_category_id,store_address,store_owner,store_username,store_password,qr_path) VALUES (?,?,?,?,?,?,?,000)"
	SELECT_STORE_BY_ID          = "SELECT ts.store_id,ts.store_name,ts.store_address,ts.store_owner,ts.store_status,ts.store_username,ts.store_password,ts.store_images,ts.qr_path,tsc.store_category_id,tsc.store_category_name FROM tb_store ts JOIN tb_store_category tsc ON ts.store_category_id=tsc.store_category_id WHERE store_id = ?"
	SELECT_ALL_STORE            = "SELECT ts.store_id,ts.store_name,ts.store_address,ts.store_owner,ts.store_status,ts.store_username,ts.store_password,ts.store_images,ts.qr_path,tsc.store_category_id,tsc.store_category_name FROM tb_store ts JOIN tb_store_category tsc ON ts.store_category_id=tsc.store_category_id WHERE ts.store_status = 'A'"
	UPDATE_STORE                = "UPDATE tb_store SET store_name=?,store_category_id=?,store_address=?,store_owner=?,store_username=?,store_password=?,store_images=?,qr_path=? WHERE store_id=?"
	DELETE_STORE                = "UPDATE tb_store SET store_status = NA WHERE store_id = ?"
	STORE_AUTH                  = "SELECT store_id,store_name,store_address,store_owner,store_status,store_username,store_password,store_images,qr_path,store_category_id FROM tb_store WHERE store_username = ?"
	INSERT_STORE_CATEGORY       = "INSERT INTO tb_store_category VALUES(?,?)"
	SELECT_STORE_CATEGORY_BY_ID = "SELECT * FROM tb_store_category WHERE store_category_id=?"
	SELECT_ALL_STORE_CATEGORY   = "SELECT * FROM tb_store_category"
	UPDATE_STORE_CATEGORY       = "UPDATE tb_store_category SET store_category_name=? WHERE store_category_id=?"
	DELETE_STORE_CATEGORY       = "UPDATE tb_store_category SET store_category_status = 'NA' WHERE store_category_id = ?"
	//QUERY PRODUCT CATEGORY
	INSERT_PRODUCT_CATEGORY       = "INSERT INTO tb_product_category VALUES (?,?)"
	SELECT_PRODUCT_CATEGORY_BY_ID = "SELECT * FROM tb_product_category WHERE product_category_id = ?"
	SELECT_ALL_PRODUCT_CATEGORY   = "SELECT * FROM tb_product_category"
	UPDATE_PRODUCT_CATEGORY       = "UPDATE tb_product_category SET product_category_name=? WHERE product_category_id=?"
	DELETE_PRODUCT_CATEGORY       = "UPDATE tb_product_category SET product_category_status = 'NA' WHERE product_category_id=? "
	//QUERY PRODUCT
	SELECT_ALL_PRODUCT_BY_STORE = "select pp.product_id, p.product_name, p.product_stock,p.product_images,p.product_status,pc.product_category_name, pp.price, pp.date_modified from tb_product_price pp inner join tb_product p on p.product_id = pp.product_id inner join tb_product_category pc on p.product_category_id = pc.product_category_id inner join (select product_id, max(date_modified) as maxDate from tb_product_price group by product_id) pj on pp.product_id = pj.product_id and pp.date_modified = pj.maxDate where p.store_id= ? AND p.product_status = 'A'"
	INSERT_PRODUCT              = "INSERT INTO tb_product(product_id,store_id,product_name,product_category_id,product_stock,product_images) VALUES (?,?,?,?,?,?)"
	INSERT_PRODUCT_PRICE        = "INSERT INTO tb_product_price VALUES(?,?,?,?)"
	SELECT_PRODUCT_BY_ID        = "select pp.product_id, p.product_name, p.product_stock,p.product_images,p.product_status,pc.product_category_id,pc.product_category_name, pp.price, pp.date_modified from tb_product_price pp inner join tb_product p on p.product_id = pp.product_id inner join tb_product_category pc on p.product_category_id = pc.product_category_id inner join (select product_id, max(date_modified) as maxDate from tb_product_price group by product_id) pj on pp.product_id = pj.product_id and pp.date_modified = pj.maxDate where p.product_id= ? AND p.product_status = 'A'"
	UPDATE_PRODUCT_WITH_PRICE   = "UPDATE tb_product SET product_name = ?,product_stock = ?,product_images = ?,product_category_id = ? WHERE product_id = ?"
	DELETE_PRODUCT              = "UPDATE tb_product SET product_status = 'NA' WHERE product_id = ?"

	//Feedback
	GET_ALL_FEEDBACK   = "SELECT * FROM tb_feedback"
	GET_FEEDBACK_BY_ID = "SELECT * FROM tb_feedback WHERE feedback_id = ?"
	POST_FEEDBACK      = "INSERT INTO tb_feedback(feedback_id,store_id,user_id,feedback_value,feedback_created) VALUES (?, ?, ?, ?, ?)"
	UPDATE_FEEDBACK    = "UPDATE tb_feedback SET store_id=?, feedback_value=?, feedback_created=? WHERE feedback_id=?"
	DELETE_FEEDBACK    = "DELETE FROM tb_feedback WHERE feedback_id = ?"
	GET_ALL_POINT      = "SELECT * FROM tb_poin"
	GET_POINT_BY_ID    = "SELECT * FROM tb_poin WHERE poin_id = ?"
	POST_POINT         = "INSERT INTO tb_poin(poin_id, store_id) VALUES(?, ?)"
	UPDATE_POINT       = "UPDATE tb_poin SET store_id=? WHERE product_id=?"
	DELETE_POINT       = "DELETE FROM tb_poin WHERE poin_id = ?"
	UPDATE_USER_POINT  = "UPDATE tb_user SET user_poin = user_poin + 1 WHERE user_id=?"
	GET_ALL_RATING     = "SELECT * FROM tb_rating"
	GET_RATING_BY_ID   = "SELECT * FROM tb_rating WHERE rating_id = ?"
	POST_RATING        = "INSERT INTO tb_rating(rating_id, store_id, user_id, rating_value, rating_description, rating_created) VALUES (?, ?, ?, ?, ?, ?)"
	UPDATE_RATING      = "UPDATE tb_rating SET store_id=?, user_id=?, rating_value=?, rating_description=?, rating_created=? WHERE rating_id=?"
	DELETE_RATING      = "DELETE FROM tb_rating WHERE rating_id = ?"

	//ORDER
	INSERT_ORDER         = "insert into tb_order value (?,?,?)"
	INSERT_ORDER_DETAIl  = "insert into tb_order_detail (qty, order_id, product_id, user_id ,price,description) value (?,?,?,?,?,?)"
	GET_NEW_PRICE        = "select pp.price from tb_product_price pp inner join tb_product p on p.product_id = pp.product_id inner join ( select product_id, max(date_modified) as maxDate from tb_product_price group by product_id ) pj on pp.product_id = pj.product_id and pp.date_modified = pj.maxDate where p.product_id = ?"
	UPDATE_PRODUCT_STOCK = "UPDATE tb_product SET product_stock=product_stock - ? WHERE product_id = ? "
	//GET ORDER BY ID
	SELECT_ORDER_BY_ID               = "SELECT * FROM tb_order WHERE order_id = ?"
	SELECT_SOLD_ITEM_ORDER_BY_ID     = "select od.qty,od.product_id,p.product_name,od.user_id ,tu.user_firstname,od.price,od.description ,od.order_detail_status from tb_order_detail od inner join tb_product p on p.product_id = od.product_id JOIN tb_user tu ON tu.user_id=od.user_id where order_id = ?"
	SELECT_ALL_ORDER_BY_STORE        = "SELECT * FROM tb_order WHERE store_id = ? ORDER BY order_created ASC"
	SELECT_ALL_SOLD_ITEM_BY_ORDER_ID = "SELECT u.user_firstname,p.product_name,od.price,od.qty,od.order_detail_status FROM tb_order_detail od JOIN tb_product p ON od.product_id=p.product_id JOIN tb_user u ON od.user_id=u.user_id JOIN tb_order o ON o.order_id=od.order_id WHERE o.order_id = ?"
	//GET ALL ORDER BY USER
	SELECT_ALL_ORDER_BY_USER       = "SELECT distinct o.order_id,o.order_created,o.store_id FROM tb_order o JOIN tb_order_detail od ON o.order_id=od.order_id WHERE od.user_id = ? ORDER BY o.order_created ASC"
	INSERT_TRANSACTION             = "INSERT INTO tb_transaction (transaction_id,order_id,user_id,amount,transaction_created) VALUES (?,?,?,?,?)"
	UPDATE_WALLET_AMOUNT_USER      = "UPDATE tb_wallet SET amount = amount - ? WHERE user_id = ? "
	UPDATE_ORDER_DETAIL_STATUS     = "UPDATE tb_order_detail SET order_detail_status = 'Paid' WHERE order_id = ?"
	UPDATE_TRANSACTION_PICK        = "UPDATE tb_transaction SET transaction_status = 'Picked' WHERE order_id = ?"
	UPDATE_POIN_USER_AFTER_PAYMENT = "UPDATE tb_user SET user_poin = user_poin + '1' WHERE user_id = ?"
	// TRANSACTIOn
	SELECT_ALL_TRANSACTION_BY_STORE = "SELECT tr.transaction_id,tr.order_id,tr.user_id,tu.user_firstname,tr.amount,tr.transaction_created,tr.transaction_status FROM tb_transaction tr JOIN tb_order o ON tr.order_id=o.order_id JOIN tb_user tu ON tr.user_id=tu.user_id WHERE o.store_id = ? ORDER BY tr.transaction_created"
	SELECT_ALL_TRANSACTION_BY_USER  = "SELECT tr.transaction_id,tr.order_id,tr.user_id,tu.user_firstname,tr.amount,tr.transaction_created,tr.transaction_status FROM tb_transaction tr JOIN tb_order o ON tr.order_id=o.order_id JOIN tb_user tu ON tr.user_id=tu.user_id WHERE tr.user_id = ? ORDER BY tr.transaction_created"
	SELECT_TRANSACTION_BY_ID        = "SELECT tr.transaction_id,tr.order_id,tr.user_id,tu.user_firstname,tr.amount,tr.transaction_created,tr.transaction_status FROM tb_transaction tr JOIN tb_order o ON tr.order_id=o.order_id JOIN tb_user tu ON tr.user_id=tu.user_id WHERE tr.transaction_id = ?"
	// TOP UP
	SELECT_ALL_TOP_UP       = "SELECT tp.top_up_id,tp.top_up_amount,tp.user_id,u.user_firstname,tp.top_up_date,tp.top_up_status FROM tb_top_up tp JOIN tb_user u ON tp.user_id=u.user_id ORDER BY tp.top_up_date"
	SELECT_VALIDATION_ORDER = "select * from tb_order where order_id = ? and store_id = ?"
)
