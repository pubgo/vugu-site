package main

import (
	"github.com/vugu/vugu"
	"log"
)

type SiteDoc struct {
	Router *Router
}

func init() {
	log.Printf("TODO: add some instructions on where to ask questions/get help")
}

type SiteDocData struct {
	Router *Router
}

func (comp *SiteDoc) NewData(props vugu.Props) (interface{}, error) {
	return &SiteDocData{Router: props["router"].(*Router)}, nil
}
