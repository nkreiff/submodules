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
	{"portuguese", "Olá mundo!"},
	{"italian", "Ciao mondo!"},
	{"japanese", "Kon'nichiwa sekai!"},
}
