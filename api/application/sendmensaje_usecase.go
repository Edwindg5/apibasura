// apibasura/api/application/sendmensaje_usecase.goa
package application

import (
	"apibasura/api/domain/entities"
	"apibasura/api/domain/repositories"
)

type PublishMessageUseCase struct {
	repository repositories.IMessageRabbit
}
func NewSaveMessage(repository repositories.IMessageRabbit) *PublishMessageUseCase {
	return &PublishMessageUseCase{
		repository: repository,
	}
}

func (sm *PublishMessageUseCase) Execute(message string, action string) (*entities.Message,error) {
	mess := entities.NewMessage(message, action)
	err := sm.repository.Publish(mess)
	if err != nil {
		return nil,err
	}
	return mess,nil

}