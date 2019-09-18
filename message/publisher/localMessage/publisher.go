package localMessage

import (
	"fmt"
	"github.com/linmadan/egglib-go/core/application"
	"github.com/linmadan/egglib-go/message/publisher/localMessage/beego"
	beegoTransaction "github.com/linmadan/egglib-go/transaction/beego"
)

type Publisher struct {
	messageStore MessageStore
}

func (publisher *Publisher) PublishMessages(messages []*application.Message, option map[string]interface{}) error {
	for _, message := range messages {
		if err := publisher.messageStore.AppendMessage(message); err != nil {
			return err
		}
	}
	if err := dispatcher.MessagePublishedNotice(); err != nil {
		return err
	}
	return nil
}

func NewLocalMessagePublisher(storeType string, storeOption map[string]interface{}, ) (*Publisher, error) {
	var messageStore MessageStore
	switch storeType {
	case "beego":
		var tc *beegoTransaction.TransactionContext
		if transactionContext, ok := storeOption["transactionContext"]; ok {
			tc = transactionContext.(*beegoTransaction.TransactionContext)
		} else {
			tc = nil
		}
		messageStore = &beego.MessagesStore{
			TransactionContext: tc,
		}
	default:
		return nil, fmt.Errorf("无效的storeType: %s", storeType)
	}
	return &Publisher{
		messageStore: messageStore,
	}, nil
}