package queries

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"
)

type EnterpriseQuery struct {
	OwnerId        string `json:"owner_id"`
	EnterpriseName string `json:"enterprise_name"`
	UserID         string `json:"user_id"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
	Term           string `json:"term"`
	EventID        string `json:"event_id"`
	ChatID         string `json:"chat_id"`
}

// func (enterpriseQuery EnterpriseQuery) Enterprise(
// 	ctx context.Context,
// ) (dto.Enterprise, error) {
// 	mapper := repositories.MapperRegistryInstance().Get("Enterprise")
// 	enterpriseRepository, ok := mapper.(repositories.EnterpriseRepository)
// 	if !ok {
// 		return dto.Enterprise{}, repositories.ErrWrongTypeAssertion
// 	}
// 	enterprise, err := enterpriseRepository.Get(ctx, enterpriseQuery.EnterpriseName)
// 	if err != nil {
// 		return dto.Enterprise{}, err
// 	}
// 	return dto.NewEnterprise(enterprise), nil
// }

// func (q EnterpriseQuery) EnterprisesByOwnerId(
// 	ctx context.Context,
// ) ([]dto.Enterprise, error) {
// 	mapper, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
// 	if !ok {
// 		return nil, ErrWrongTypeAssertion
// 	}
// 	enterprises, err := mapper.FindAll(ctx,
// 		find_all_enterprises.ByOwnerId(q.OwnerId),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	s := enterprises.ToSlice()
// 	result := make([]dto.Enterprise, 0, len(s))
// 	for _, v := range s {
// 		result = append(result, dto.NewEnterpriseDTO(v))
// 	}
// 	slices.SortStableFunc(result, func(a, b dto.Enterprise) int {
// 		return strings.Compare(a.Name, b.Name)
// 	})
// 	return result, nil
// }

func (query EnterpriseQuery) EnterpriseChat(
	ctx context.Context,
) (dto.Chat, error) {
	if query.EnterpriseName == "" || query.ChatID == "" {
		return dto.Chat{}, ErrBadRequest
	}
	chatID, err := enterprise.NewChatID(query.ChatID)
	if err != nil {
		return dto.Chat{}, err
	}
	chat, err := repositories.MapperRegistryInstance().Get("Chat").(repositories.ChatRepository).Get(ctx, *chatID)
	if err != nil {
		return dto.Chat{}, err
	}
	return *dto.NewChat(*chat), nil
}
