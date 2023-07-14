package web

type NodeText struct {
	Text string
}

func Text(text string) *NodeText {
	this := &NodeText{
		Text: text,
	}
	return this
}

func (e *NodeText) Marshal() string {
	return e.Text
}

func (e *NodeText) MarshalIndent(prefix string) string {
	return prefix + e.Text
}
