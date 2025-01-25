package app

import (
	"SnappChatroom/config"
	"SnappChatroom/internal/port"
	"SnappChatroom/internal/service"
	"SnappChatroom/pkg/adapters/storage"
	appCtx "SnappChatroom/pkg/context"
	"SnappChatroom/pkg/postgres"
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
	"log"
)

type app struct {
	cfg            config.Config
	db             *gorm.DB
	natsConnection *nats.Conn
	chatService    port.Chat
}

func (a *app) Nats() *nats.Conn {
	return a.natsConnection
}

func (a *app) ChatService(ctx context.Context) port.Chat {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.chatService == nil {
			a.chatService = a.chatServiceWithDB(a.db)
		}
		return a.chatService
	}

	return a.chatServiceWithDB(db)
}

func (a *app) chatServiceWithDB(db *gorm.DB) port.Chat {
	return service.NewChatService(storage.NewChatRepo(db))
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Password,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		DBName: a.cfg.DB.Database,
		Schema: a.cfg.DB.Schema,
	})

	if err != nil {
		return err
	}

	if err := postgres.Migrate(db); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	a.db = db
	return nil
}

func (a *app) SetNats() error {
	natsURl := fmt.Sprintf("nats://localhost:" + a.cfg.Nats.HostPort)
	conn, err := nats.Connect(natsURl)
	if err != nil {
		return err
	}
	log.Println("Connected to nats at", natsURl)
	a.natsConnection = conn
	return nil
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	if err := a.SetNats(); err != nil {
		return nil, err
	}

	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
