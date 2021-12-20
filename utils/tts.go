package utils

import (
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/hegedustibor/htgo-tts/voices"
)

/**
todo 使用的是google服务，国内访问不通
*/
func TTS(path string, text string) error {
	speech := htgotts.Speech{Folder: path, Language: voices.Chinese, Handler: &handlers.MPlayer{}}
	return speech.Speak(text)
}
