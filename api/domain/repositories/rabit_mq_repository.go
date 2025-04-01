// apibasura/api/domain/repositories/rabit_mq_repository.go
package repositories

import "apibasura/api/domain/entities"

type IMessageRabbit interface {
	Publish(message *entities.Message) error
}