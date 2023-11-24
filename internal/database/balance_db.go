package database

import (
	"database/sql"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (c *BalanceDB) Init() {
	_, err := c.DB.Exec("CREATE TABLE IF NOT EXISTS balances (account_id_from varchar(255), account_id_to varchar(255), balanceAccountFrom int, balanceAccountTo int, created_at datetime)")
	if err != nil {
		panic(err)
	}
}

func (c *BalanceDB) Get(id string) (*entity.Balance, error) {
	balance := &entity.Balance{}
	stmt, err := c.DB.Prepare("SELECT account_id_from, account_id_to, balanceAccountFrom, balanceAccountTo, MAX(created_at) AS CreatedAt  FROM balances WHERE account_id_from = ? OR account_id_to = ? GROUP BY account_id_from, account_id_to, balanceAccountFrom, balanceAccountTo")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id, id)
	if err := row.Scan(&balance.AccountIDFrom, &balance.AccountIDTo, &balance.BalanceAccountFrom, &balance.BalanceAccountTo, &balance.CreatedAt); err != nil {
		return nil, err
	}
	return balance, nil
}

func (t *BalanceDB) Create(balance *entity.Balance) error {
    stmt, err := t.DB.Prepare("INSERT INTO balances (account_id_from, account_id_to, balanceAccountFrom, balanceAccountTo, created_at) VALUES (?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(balance.AccountIDFrom, balance.AccountIDTo, balance.BalanceAccountFrom, balance.BalanceAccountTo, balance.CreatedAt) 
    if err != nil {
        return err
    }
    return nil
}		