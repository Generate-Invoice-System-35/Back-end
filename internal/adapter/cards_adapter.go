package adapter

import "Back-end/internal/model"

type AdapterCardRepository interface {
	GetAllCards() []model.Card
	GetCardByID(id int) (card model.Card, err error)
	UpdateCardByID(id int, card model.Card) error
	DeleteCardByID(id int) error
}

type AdapterCardService interface {
	GetAllCardsService() []model.Card
	GetCardByIDService(id int) (model.Card, error)
	UpdateCardByIDService(id int, card model.Card) error
	DeleteCardByIDService(id int) error
}
