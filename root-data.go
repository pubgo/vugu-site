package main

import "github.com/vugu/vugu"

type Root struct {
	Router *Router
}

type RootData struct {
	Router *Router
}

func (comp *Root) NewData(props vugu.Props) (interface{}, error) {
	//return &RootData{Router: props["router"].(*Router)}, nil
	return &RootData{Router: comp.Router}, nil
}
