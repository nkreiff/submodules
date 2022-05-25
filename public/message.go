package public

type (
	Message struct {
		Language string
		Message  string
	}
)

var AllMessages []Message = []Message{
	{"spanish", "Hola Mundo!"},
	{"english", "Hello World!"},
	{"french", "Salut Monde!"},
}
