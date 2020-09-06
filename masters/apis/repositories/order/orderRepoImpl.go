package orderRepositories

import (
	"database/sql"

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

func (o *OrderRepoImpl) AddOrder(storeID string, order *models.Order) error {
	orderID := guuid.New()
	tx, err := o.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.INSERT_ORDER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(orderID, order.OrderCreated, storeID); err != nil {
		tx.Rollback()
		return err
	}

	// productPriceID := guuid.New()
	stmt, err = tx.Prepare(utils.INSERT_ORDER_DETAIl)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, val := range order.OrderDetail.Product {
		_, err = stmt.Exec(order.OrderDetail.Qty, orderID, val.ProductID, order.OrderDetail.User.UserID, order.OrderDetail.Price)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (o *OrderRepoImpl) GetOrderByID(orderID string) (*models.Order, error) {
	stmt, err := o.db.Prepare(utils.SELECT_ORDER_BY_ID)
	order := models.Order{}
	if err != nil {
		return &order, err
	}
	errQuery := stmt.QueryRow(orderID).Scan(&order.OrderID, sad)

	if errQuery != nil {
		return &order, err
	}

	defer stmt.Close()
	return &order, nil
}

func (o *OrderRepoImpl) GetAllOrderByStore(storeID string) ([]*models.Order, error) {
	stmt, err := o.db.Prepare(utils.SELECT_ALL_ORDER_BY_STORE)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	listOrder := []*models.Order{}
	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.OrderCreated, &product.ProductName)
		if err != nil {
			return nil, err
		}
		listOrder = append(listOrder, &order)
	}
	return listOrder, nil
}

func (o *OrderRepoImpl) UpdateOrderPaid(orderID string, order *models.Order) error {
	tx, err := o.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_ORDER_DETAIL_STATUS_PAID)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(order.asdasd, orderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	transactionID := guuid.New()
	stmt, err = tx.Prepare(utils.INSERT_TRANSACTION)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(transactionID, storeID, product.ProductName, productCategory.ProductID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (o *OrderRepoImpl) UpdateOrderCancel(orderID string, payment *models.Payment) error {
	tx, err := o.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_ORDER_DETAIL_STATUS_CANCEL)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(order.asdasd, orderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// transactionID := guuid.New()
	stmt, err = tx.Prepare(utils.UPDATE_TRANSACTION_CANCEL)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(payment.TransactionID, storeID, product.ProductName, productCategory.ProductID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
