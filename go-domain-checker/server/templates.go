package server

import (
	"html/template"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var tpl *template.Template

type PageData struct {
	Year string
}

const patternsPath = "server/templates/*.gohtml"

func init() {
	tpl = template.Must(template.ParseGlob(patternsPath))
}

func (s *Server) HomePage(ctx *gin.Context) {
	indexPageData := PageData{
		Year: strconv.Itoa(time.Now().Year()),
	}
	err := tpl.ExecuteTemplate(ctx.Writer, "index.gohtml", indexPageData)
	if err != nil {
		log.Fatal(err)
	}
}
