package main

import (
	"database/sql"
	"fmt"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/database"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/event"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/web"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	balanceDB := database.NewBalanceDB(db)
	balanceDB.Init()

	balanceUpdatedSubscriber := event.NewBalanceUpdatedSubscriber(&configMap, balanceDB)
	go balanceUpdatedSubscriber.Handle()

	balanceWebHandler := web.NewWebBalanceHandler(balanceDB)

	router := gin.Default()
	router.GET("/", index)
	router.GET("/balances/:AccountId", balanceWebHandler.GetBalance)

	fmt.Println("Server is running")
	router.Run(":3003")
}

func index(c *gin.Context) {
	c.IndentedJSON(200, "{}")
}
