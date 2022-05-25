package public

type (
	Message struct {
		Language string
		Message  string
	}
)

var AllMessages []Message {
	Message("spanish", "Hola Mundo!"),
	Message("english", "Hello World!"),
	Message("french", "Salut Monde!"),
}