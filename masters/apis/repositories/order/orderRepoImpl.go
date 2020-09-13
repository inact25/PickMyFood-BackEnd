package orderRepositories

import (
	"database/sql"
	"log"

	guuid "github.com/google/uuid"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	utils "github.com/inact25/PickMyFood-BackEnd/utils/queryConstant"
)

type OrderRepoImpl struct {
	db *sql.DB
}

func InitOrderRepoImpl(db *sql.DB) OrderRepo {
	return &OrderRepoImpl{db: db}
}

func (o OrderRepoImpl) AddOrder(order *models.Order) (*models.Order, error) {
	println("MASUK REPO")
	orderID := guuid.New()
	tx, err := o.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(utils.INSERT_ORDER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if _, err := stmt.Exec(orderID, order.OrderCreated, order.StoreID); err != nil {
		tx.Rollback()
		return nil, err

	}
	println("MASUK TB ORDER")

	stmt, err = tx.Prepare(utils.INSERT_ORDER_DETAIl)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return nil, err

	}

	for _, val := range order.SoldItems {
		_, err = stmt.Exec(val.Qty, orderID, val.ProductID, val.UserID, val.Price, val.Note)
		if err != nil {
			tx.Rollback()
			return nil, err

		}
		_, err = tx.Exec(utils.UPDATE_PRODUCT_STOCK, val.Qty, val.ProductID)
	}

	newOrderID := orderID.String()
	tx.Commit()
	println("MASUK TB ORDER DETAIL")
	return o.GetOrderByID(newOrderID)
}

func (o OrderRepoImpl) GetOrderByID(orderID string) (*models.Order, error) {
	stmt, err := o.db.Prepare(utils.SELECT_ORDER_BY_ID)
	order := models.Order{}
	if err != nil {
		return &order, err
	}
	errQuery := stmt.QueryRow(orderID).Scan(&order.OrderID, &order.OrderCreated, &order.StoreID)
	if errQuery != nil {
		return &order, err
	}
	defer stmt.Close()

	stmt, err = o.db.Prepare(utils.SELECT_SOLD_ITEM_ORDER_BY_ID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(orderID)
	if err != nil {
		return nil, err
	}

	var soldItem models.SoldItems
	for rows.Next() {
		err := rows.Scan(&soldItem.Qty, &soldItem.ProductID, &soldItem.ProductName, &soldItem.UserID, &soldItem.UserFirstName, &soldItem.Subtotal, &soldItem.Note, &soldItem.OrderDetailStatus)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		order.SoldItems = append(order.SoldItems, soldItem)
	}

	return &order, nil
}

// get all order by store
func (o *OrderRepoImpl) GetAllOrderByStore(storeID string) ([]*models.Order, error) {
	println("MASUK REPo")
	stmt, err := o.db.Prepare(utils.SELECT_ALL_ORDER_BY_STORE)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(storeID)
	if err != nil {
		return nil, err
	}
	listOrder := []*models.Order{}
	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.OrderID, &order.OrderCreated, &order.StoreID)
		if err != nil {
			return nil, err
		}
		listOrder = append(listOrder, &order)
		println("MASUK ORDER")

		//solditems
		stmt, err = o.db.Prepare(utils.SELECT_ALL_SOLD_ITEM_BY_ORDER_ID)
		if err != nil {
			log.Print(err)
			return nil, err
		}

		rows, err := stmt.Query(order.OrderID)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		println("MASUK SINI 1")

		soldItem := models.SoldItems{}
		for rows.Next() {
			println("MASUK SINI")
			err := rows.Scan(&soldItem.UserFirstName, &soldItem.ProductName, &soldItem.Subtotal, &soldItem.Qty, &soldItem.OrderDetailStatus)
			if err != nil {
				log.Print(err)
				return nil, err
			}
			order.SoldItems = append(order.SoldItems, soldItem)
		}
	}
	return listOrder, nil
}

// get all order by user
func (o *OrderRepoImpl) GetAllOrderByUser(userID string) ([]*models.Order, error) {
	println("MASUK REPO")
	stmt, err := o.db.Prepare(utils.SELECT_ALL_ORDER_BY_USER)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}
	listOrder := []*models.Order{}
	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.OrderID, &order.OrderCreated, &order.StoreID)
		if err != nil {
			return nil, err
		}
		listOrder = append(listOrder, &order)
		println("MASUK ORDER")

		//solditems
		stmt, err = o.db.Prepare(utils.SELECT_ALL_SOLD_ITEM_BY_ORDER_ID)
		if err != nil {
			log.Print(err)
			return nil, err
		}

		rows, err := stmt.Query(order.OrderID)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		println("MASUK SINI 1")

		soldItem := models.SoldItems{}
		for rows.Next() {
			println("MASUK SINI")
			err := rows.Scan(&soldItem.UserFirstName, &soldItem.ProductName, &soldItem.Subtotal, &soldItem.Qty, &soldItem.OrderDetailStatus)
			if err != nil {
				log.Print(err)
				return nil, err
			}
			order.SoldItems = append(order.SoldItems, soldItem)
		}
	}
	return listOrder, nil
}

// func (o *OrderRepoImpl) UpdateOrderCancel(orderID string, payment *models.Payment) error {
// 	tx, err := o.db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	stmt, err := tx.Prepare(utils.UPDATE_ORDER_DETAIL_STATUS_CANCEL)
// 	defer stmt.Close()
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	_, err = stmt.Exec(order.asdasd, orderID)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	// transactionID := guuid.New()
// 	stmt, err = tx.Prepare(utils.UPDATE_TRANSACTION_CANCEL)
// 	defer stmt.Close()
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	if _, err := stmt.Exec(payment.TransactionID, storeID, product.ProductName, productCategory.ProductID); err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	return tx.Commit()
// }
