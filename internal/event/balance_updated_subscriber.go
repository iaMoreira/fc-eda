package event

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/database"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type BalanceUpdatedSubscriber struct {
	ConfigMap *ckafka.ConfigMap
	BalanceDB *database.BalanceDB
}

func NewBalanceUpdatedSubscriber(configMap *ckafka.ConfigMap, balanceDB *database.BalanceDB) *BalanceUpdatedSubscriber {
	return &BalanceUpdatedSubscriber{
		ConfigMap: configMap,
		BalanceDB: balanceDB,
	}
}

func (s *BalanceUpdatedSubscriber) Handle() {
	c := kafka.NewConsumer(s.ConfigMap, []string{"balances"})
	msgChan := make(chan *ckafka.Message)

	go c.Consume(msgChan) // faz a magica

	for {
		msg := <-msgChan
		var balanceUpdated map[string]any
		json.Unmarshal([]byte(string(msg.Value)), &balanceUpdated)
		Payload := balanceUpdated["Payload"].(map[string]any)

		balance := entity.Balance{
			AccountIDFrom:      Payload["account_id_from"].(string),
			AccountIDTo:        Payload["account_id_to"].(string),
			BalanceAccountFrom: Payload["balance_account_id_from"].(float64),
			BalanceAccountTo:   Payload["balance_account_id_to"].(float64),
			CreatedAt:          time.Now(),
		}

		err := s.BalanceDB.Create(&balance)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("OK : " + string(msg.Value))
		}
	}
}
