package api

import (
	"github.com/umtaktpe/iyzipay-go/model"
	"github.com/umtaktpe/iyzipay-go/utils"
)

func RetriveCards(req utils.RequestConvertible, opt *model.Options) (*retriveCards, error) {
	resp := &retriveCards{}
	if err := request("POST", "/cardstorage/cards", req, opt, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func CreateCard(req utils.RequestConvertible, opt *model.Options) (*createCard, error) {
	resp := &createCard{}
	if err := request("POST", "/cardstorage/card", req, opt, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func DeleteCard(req utils.RequestConvertible, opt *model.Options) (*baseResponse, error) {
	resp := &baseResponse{}
	if err := request("DELETE", "/cardstorage/card", req, opt, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
