package main

import "github.com/vugu/vugu"

type SiteHome struct {
	Router *Router
}

type SiteHomeData struct {
	Router *Router
}

func (comp *SiteHome) NewData(props vugu.Props) (interface{}, error) {
	return &SiteHomeData{Router: props["router"].(*Router)}, nil
}

/* notes:


VUGU: The Go+WebAssembly UI Library

Vugu is experimental technology, but a surprising amount of functionality works in your browser right now:


Inspired by modern web UI libraries like Vue and React.
Single-file components.
Runs in most modern browsers.
Simple and sane dev and build environment.  (Say goodbye to that node_modules folder!)

Write UIs with the ease of HTML+CSS for presentation, and the facility of Go for interface logic.  Trust me, it's pretty cool.  See the Getting Started page.


Create an app in

*/
