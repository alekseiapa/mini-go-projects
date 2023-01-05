package server

import (
	"html/template"
	"log"
	"strconv"
	"time"

	dc "github.com/alekseiapa/mini-go-projects/go-domain-checker/domain-checker"
	"github.com/gin-gonic/gin"
)

var tpl *template.Template

type PageData struct {
	Year       string
	DomainInfo *dc.Domain
}

const patternsPath = "server/templates/*.gohtml"

func init() {
	tpl = template.Must(template.ParseGlob(patternsPath))
}

func (s *Server) HomePage(ctx *gin.Context) {
	indexPageData := PageData{
		Year: strconv.Itoa(time.Now().Year()),
	}
	domainToCheck := ctx.PostForm("domainToCheck")
	if domainToCheck != "" {
		domainInfo := dc.CheckDomain(domainToCheck)
		indexPageData.DomainInfo = domainInfo
	}

	err := tpl.ExecuteTemplate(ctx.Writer, "index.gohtml", indexPageData)
	if err != nil {
		log.Fatal(err)
	}
}
