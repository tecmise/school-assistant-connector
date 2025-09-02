package assistant

type Response struct {
	MessageContent
	Error    *string `json:"error"`
	Override *bool   `json:"override"`
}

// type Request struct {
// 	Messages []MessageContent `json:"messages"`
// }

type MessageContent struct {
	Role  string         `json:"role"`
	Text  string         `json:"text"`
	Files *[]MessageFile `json:"files"`
	Html  *string        `json:"html"`
}

type MessageFile struct {
	Name string `json:"name"`
	Src  string `json:"src"`
	Type string `json:"type"`
}
