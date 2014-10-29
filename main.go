package main 

import(
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func main(){

    m := martini.Classic()
    m.Use(render.Renderer(render.Options{
        Directory: "templates",
        Layout: "layout",
        Extensions: []string{".tmpl", ".html"},
        Charset: "UTF-8",
    }))
	
    m.Get("/", func(r render.Render){
		r.HTML(200, "main", nil)
    })
    
	m.Get("/page1", func(r render.Render) {
		r.HTML(200, "page1", nil)
    })
    
	m.Get("/page2", func(r render.Render) {
		r.HTML(200, "page2", nil)
    })
	
	m.NotFound(func(r render.Render){
		r.HTML(200, "notfound", nil)
	})
	
	m.Use(martini.Static("assets"))
	m.Run()
}
