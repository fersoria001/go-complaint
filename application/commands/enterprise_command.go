package commands

type EnterpriseCommand struct {
}

// func (command EnterpriseCommand) ReplyChat(
// 	ctx context.Context,
// ) error {
// 	if command.ID == "" ||
// 		command.SenderID == "" ||
// 		command.Content == "" {
// 		return ErrBadRequest
// 	}
// 	sender, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(
// 		ctx,
// 		command.SenderID,
// 	)
// 	if err != nil {
// 		return ErrNotFound
// 	}
// 	chatID, err := enterprise.NewChatID(command.ID)
// 	if err != nil {
// 		return err
// 	}
// 	dbChat, err := repositories.MapperRegistryInstance().Get("Chat").(repositories.ChatRepository).Get(
// 		ctx,
// 		*chatID,
// 	)
// 	var reply enterprise.Reply
// 	if err != nil {
// 		if errors.Is(err, pgx.ErrNoRows) {
// 			dbChat = enterprise.NewChatEntity(
// 				*chatID,
// 			)
// 			reply = dbChat.Reply(uuid.New(), *sender, command.Content)
// 			err = repositories.MapperRegistryInstance().Get("Chat").(repositories.ChatRepository).Save(ctx, dbChat)
// 			if err != nil {
// 				return err
// 			}
// 		} else {
// 			return err
// 		}
// 	} else {
// 		reply = dbChat.Reply(uuid.New(), *sender, command.Content)
// 		err = repositories.MapperRegistryInstance().Get("Chat").(repositories.ChatRepository).Update(ctx, dbChat)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	cache.RequestChannel <- cache.Request{
// 		Type:    cache.WRITE,
// 		Key:     command.ID,
// 		Payload: dto.NewChatReply(reply),
// 	}
// 	return nil
// }
// func (command EnterpriseCommand) MarkAsSeen(ctx context.Context) error {
// 	if command.RepliesID == nil || len(command.RepliesID) == 0 || command.ID == "" {
// 		return ErrBadRequest
// 	}
// 	parseIds := make([]uuid.UUID, len(command.RepliesID))
// 	for _, replyID := range command.RepliesID {
// 		parsedReplyID, err := uuid.Parse(replyID)
// 		if err != nil {
// 			return ErrBadRequest
// 		}
// 		parseIds = append(parseIds, parsedReplyID)
// 	}
// 	chatID, err := enterprise.NewChatID(command.ID)
// 	if err != nil {
// 		return err
// 	}
// 	parsedIds := mapset.NewSet(parseIds...)
// 	r := repositories.MapperRegistryInstance().Get("Chat").(repositories.ChatRepository)
// 	dbChat, err := r.Get(ctx, *chatID)
// 	if err != nil {
// 		return err
// 	}
// 	dbChat.MarkAsSeen(parsedIds)
// 	err = r.Update(ctx, dbChat)
// 	if err != nil {
// 		return err
// 	}
// 	replies := dbChat.Replies()
// 	dtos := make([]*dto.ChatReply, 0, len(replies))
// 	for _, reply := range replies {
// 		dtos = append(dtos, dto.NewChatReply(*reply))
// 	}
// 	cache.RequestChannel <- cache.Request{
// 		Key:     fmt.Sprintf("complaintLastReply:%s", chatID.String()),
// 		Payload: dtos,
// 	}
// 	return nil
// }
