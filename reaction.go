package mastodon

type PleromaReactions struct {
	EmojiReaction
	Accounts map[string]string `json:"accounts"`
}
