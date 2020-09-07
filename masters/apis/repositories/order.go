package repositories

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
)

type OrderRepoImpl struct {
	db *sql.DB
}

func (s *OrderRepoImpl) GetOrders() ([]*models.OrderModels, error) {
	var orders []*models.OrderModels
	query := "SELECT * FROM tb_order"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		order := models.OrderModels{}
		err := rows.Scan(&order.OrderID, &order.OrderCreated, &order.StoreID)

		if err != nil {
			return nil, err
		}

		orders = append(orders, &order)

	}

	return orders, nil
}

func (s *OrderRepoImpl) GetOrderByID(ID string) (*models.OrderModels, error) {
	results := s.db.QueryRow("SELECT * FROM tb_order WHERE order_id = ?", ID)

	var d models.OrderModels
	err := results.Scan(&d.OrderID, &d.OrderCreated, &d.StoreID)
	if err != nil {
		return nil, errors.New("ID Not Found")
	}

	return &d, nil
}

func (s *OrderRepoImpl) PostOrder(d models.OrderModels) (*models.OrderModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmnt, _ := tx.Prepare(`INSERT INTO tb_order(order_id, order_created, store_id) VALUES (?, ?, ?)`)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.OrderID, d.OrderCreated, d.StoreID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	return s.GetOrderByID(strconv.Itoa(int(lastInsertID)))
}

func (s *OrderRepoImpl) UpdateOrder(ID string, data models.OrderModels) (*models.OrderModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(`UPDATE tb_order SET order_created=?, store_id=? WHERE order_id=?`,
		data.OrderID, data.OrderCreated, data.StoreID, ID)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return s.GetOrderByID(ID)
}

func (s *OrderRepoImpl) DeleteOrder(ID string) (*models.OrderModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("DELETE FROM tb_order WHERE order_id = ?", ID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return s.GetOrderByID(ID)

}

func InitOrderImpl(db *sql.DB) OrderRepo {
	return &OrderRepoImpl{db}

}
