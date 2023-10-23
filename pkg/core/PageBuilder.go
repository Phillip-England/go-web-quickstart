package core

import (
	"fmt"
	"os"
)

type PageBuilder struct {
    componentString string
    fontawesomeTag string
    title string
}

func NewPageBuilder(title string) (*PageBuilder) {
    return &PageBuilder{
        componentString: "",
        fontawesomeTag: os.Getenv("FONTAWESOME_TAG"),
        title: title,
    }
}

func (b *PageBuilder) GetHTMLBytes() []byte {
    return []byte(fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <head>
            <script src="https://unpkg.com/htmx.org@1.9.6"></script>
            <script src="https://unpkg.com/hyperscript.org@0.9.11"></script>
            <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no"></meta>
            <link rel="stylesheet" href="/static/output.css"></link>
            %s
            <title>%s</title>
            </head>
            <body hx-boost='true' class='bg-gray-100'>
            %s
        </body>
        </html>
    `, []byte(b.fontawesomeTag), []byte(b.title), []byte(b.componentString)))
}


func (b *PageBuilder) Add(component string) {
    b.componentString = b.componentString + component
}