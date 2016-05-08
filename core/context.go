package bots

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"net/http"
)

type WebhookContext interface {
	GetLogger() Logger
	BotInputProvider
	Translate(key string) string
	TranslateNoWarning(key string) string

	Init(w http.ResponseWriter, r *http.Request) error
	Context() context.Context

	BotChatID() interface{}

	ChatEntity() BotChat
	ReplyByBot(m MessageFromBot) error

	CommandTitle(title, icon string) string

	Locale() Locale
	SetLocale(code5 string) error

	NewMessage(text string) MessageFromBot
	NewMessageByCode(messageCode string, a ...interface{}) MessageFromBot

	GetHttpClient() *http.Client
	UpdateLastProcessed(chatEntity BotChat) error

	GetOrCreateUserEntity() (BotUser, error)
	AppUserID() int64
	GetUser() (*datastore.Key, AppUser, error)
	GetOrCreateUser() (*datastore.Key, AppUser, error)

	ApiUser() BotApiUser
	BotState
	BotChatStore
	WebhookInput
}

type BotState interface {
	IsNewerThen(chatEntity BotChat) bool
}

type BotInputProvider interface {
	MessageText() string
}

type BotApiUser interface {
	//IdAsString() string
	IdAsInt64() int64
	FirstName() string
	LastName() string
}