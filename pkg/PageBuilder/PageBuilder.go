package pageBuilder

import (
	"fmt"
	"os"
)

type PageBuilder struct {
    componentString string
    fontawesomeKey string
    title string
}

func New(title string) (*PageBuilder) {
    return &PageBuilder{
        componentString: "",
        fontawesomeKey: os.Getenv("FONTAWESOME_KEY"),
        title: title,
    }
}

func (b *PageBuilder) GetBasePageAsBytes() []byte {
    return []byte(fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <head>
            <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no"></meta>
            <link rel="stylesheet" href="/public/output.css"></link>
            <script src=%s crossorigin="anonymous"></script>
            <title>%s</title>
            </head>
            <body>
            %s
            <script src="/public/bundle.js"></script>
        </body>
        </html>
    `, []byte(b.fontawesomeKey), []byte(b.title), []byte(b.componentString)))
}


func (b *PageBuilder) Add(component string) {
    b.componentString = b.componentString + component
}