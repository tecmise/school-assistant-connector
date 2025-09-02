package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := f.File
			if pwd, err := os.Getwd(); err == nil {
				if rel, err := filepath.Rel(pwd, filename); err == nil {
					filename = rel
				}
			}

			funcName := path.Base(f.Function)
			return fmt.Sprintf("%s()", funcName), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	logrus.Debug("Log de depuração ativado")
}

func main() {
	// client := assistant.Rest("http://localhost:3003")

	// message := assistant.MessageContent{Role: "user", Text: "Olá"}
	// messages := [1]assistant.MessageContent{message}
	// request := assistant.Request{Messages: messages[:]}

	// response, err := client.ChatPrompt(context.TODO(), request)
	// if err != nil {
	// 	fmt.Printf("err: ", err)
	// }
	// fmt.Printf("Response", response)
}
