package model

type BaseModel struct {
	Locale         string
	ConversationId string
}

func (b *BaseModel) GetLocale() string {
	return b.Locale
}

func (b *BaseModel) SetLocale(locale string) {
	b.Locale = locale
}

func (b *BaseModel) GetConversationId() string {
	return b.ConversationId
}

func (b *BaseModel) SetConversationId(conversationId string) {
	b.ConversationId = conversationId
}
