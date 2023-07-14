package macos

type nodeText struct {
	Text string
}

func text(text string) *nodeText {
	this := &nodeText{
		Text: text,
	}
	return this
}

func (e *nodeText) Marshal() string {
	return e.Text
}

func (e *nodeText) MarshalIndent(prefix string) string {
	return prefix + e.Text
}
