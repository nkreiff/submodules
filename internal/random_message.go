package internal

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nkreiff/submodules/public"
)

func RandomMessage() string {
	rand.Seed(time.Now().UnixMicro())
	message := public.AllMessages[rand.Intn(len(public.AllMessages))]
	return fmt.Sprintf("[%s:%d] %s", message.Language, len(public.AllMessages), message.Message)
}
