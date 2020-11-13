//+build wasm,js

package ui

import (
	"fmt"

	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/html"
	"github.com/gowebapi/webapi/html/htmlevent"
)

func Demo() {
	element := webapi.GetWindow().Document().GetElementById("foo")
	button := html.HTMLButtonElementFromJS(element)
	button.SetInnerText("Press me!")

	count := 1
	buttonClickHandler := button.SetOnClick(func(event *htmlevent.MouseEvent, currentTarget *html.HTMLElement) {
		button.SetInnerText(fmt.Sprint("Count: ", count))
		count++
	})
	defer buttonClickHandler.Release()

	c := make(chan struct{})
	<-c
}
