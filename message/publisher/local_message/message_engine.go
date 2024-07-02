package local_message

import (
	"github.com/stevechan1993/egglib-go/core/application"
)

type MessageEngine interface {
	PublishToMessageSystem(messages []*application.Message, option map[string]interface{}) ([]int64, error)
}
