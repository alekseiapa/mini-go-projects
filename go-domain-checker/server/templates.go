package server

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
	"sync"
	"time"

	dc "github.com/alekseiapa/mini-go-projects/go-domain-checker/domain-checker"
	"github.com/gin-gonic/gin"
)

var tpl *template.Template
var cache sync.Map

type PageData struct {
	Year       string
	DomainInfo *dc.Domain
	Error      error
}

const patternsPath = "server/templates/*.gohtml"

func init() {
	tpl = template.Must(template.ParseGlob(patternsPath))
}

func (s *Server) HomePage(ctx *gin.Context, name string) {
	if html, ok := cache.Load(name); ok {
		// Serve the HTML from the cache
		ctx.Data(http.StatusOK, "text/html", html.([]byte))
		return
	}
	indexPageData := PageData{
		Year: strconv.Itoa(time.Now().Year()),
	}
	domainToCheck := ctx.PostForm("domainToCheck")
	if domainToCheck != "" {
		domainInfo := dc.CheckDomain(domainToCheck)
		indexPageData.DomainInfo = domainInfo
	}
	buf := new(bytes.Buffer)
	err := tpl.ExecuteTemplate(buf, "index.gohtml", indexPageData)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	cache.Store(name, buf.String())
	ctx.Data(http.StatusOK, "text/html", buf.Bytes())
}
