package driven_adapter

import "fmt"

//IMPLEMENTATION OF DRIVEN PORT TO VIEW
type ConsoleWriter struct {
}

func (c ConsoleWriter) WriteLines(obj interface{}) {
	fmt.Println(obj)
}
