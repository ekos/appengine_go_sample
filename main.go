package main

import (
    "net/http"
    "github.com/go-martini/martini"
    "appengine"
)

func init() {
  m := martini.Classic()
  m.Use(AppEngine)
  m.Get("/", func() string {
    return "Hello world!"
  })
  http.Handle("/", m)
}
func AppEngine(c martini.Context, r *http.Request) {
    c.MapTo(appengine.NewContext(r), (*appengine.Context)(nil))
}