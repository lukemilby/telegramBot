package bot

// InScope is used to look up the scope of the message. This can be public, a channel or private
func (b *Bot) InScope(scope string) bool {
	if _, ok := b.Scope[scope]; ok {
		return true
	}
	return false
}