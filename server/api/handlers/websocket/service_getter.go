package websocket

import (
	"SnappChatroom/api/service"
	"SnappChatroom/app"
	"context"
)

func chatServiceGetter(appContainer app.App) ServiceGetter[*service.ChatService] {

	return func(ctx context.Context) *service.ChatService {
		return service.NewChatService(appContainer.ChatService(ctx))
	}
}
