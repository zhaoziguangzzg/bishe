package controller

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

func GetPage(pageStr string) (page int) {

	page, _ = strconv.Atoi(pageStr)

	if page <= 0 {
		page = 1
	}

	return
}

// adminLayoutCache caches compiled admin page templates
var adminLayoutCache sync.Map
var adminLayoutOnce sync.Once
var adminLayoutBase *template.Template

// indexLayoutCache caches compiled index page templates
var indexLayoutCache sync.Map
var indexLayoutOnce sync.Once
var indexLayoutBase *template.Template

func loadAdminLayoutBase() *template.Template {
	adminLayoutOnce.Do(func() {
		layoutContent, err := os.ReadFile("web/views/admin/_layout.html")
		if err != nil {
			return
		}
		adminLayoutBase = template.New("")
		template.Must(adminLayoutBase.New("admin/_layout.html").Parse(string(layoutContent)))
	})
	return adminLayoutBase
}

// RenderAdminPage renders an admin page with the layout template.
// It creates an independent template set by cloning the layout and adding the page template,
// avoiding conflicts when multiple pages define the same block names.
func RenderAdminPage(c *gin.Context, pagePath string, data any) {
	// Try to get cached template
	if cached, ok := adminLayoutCache.Load(pagePath); ok {
		if tpl, ok := cached.(*template.Template); ok {
			tpl.ExecuteTemplate(c.Writer, pagePath, data)
			return
		}
	}

	// Build template: clone layout + parse page
	base := loadAdminLayoutBase()
	if base == nil {
		// Fallback to default rendering
		c.HTML(http.StatusOK, pagePath, data)
		return
	}

	pageContent, err := os.ReadFile(filepath.Join("web/views", pagePath))
	if err != nil {
		c.HTML(http.StatusOK, pagePath, data)
		return
	}

	pageTpl, err := base.Clone()
	if err != nil {
		c.HTML(http.StatusOK, pagePath, data)
		return
	}

	template.Must(pageTpl.New(pagePath).Parse(string(pageContent)))

	// Cache for future use
	adminLayoutCache.Store(pagePath, pageTpl)

	// Execute the page template
	pageTpl.ExecuteTemplate(c.Writer, pagePath, data)
}

func loadIndexLayoutBase() *template.Template {
	indexLayoutOnce.Do(func() {
		layoutContent, err := os.ReadFile("web/views/layout/layout.html")
		if err != nil {
			return
		}
		indexLayoutBase = template.New("")
		template.Must(indexLayoutBase.New("layout/layout.html").Parse(string(layoutContent)))
	})
	return indexLayoutBase
}

// RenderIndexPage renders an index page with the layout template.
// It creates an independent template set by cloning the layout and adding the page template,
// avoiding conflicts when multiple pages define the same block names.
func RenderIndexPage(c *gin.Context, pagePath string, data any) {
	// Try to get cached template
	if cached, ok := indexLayoutCache.Load(pagePath); ok {
		if tpl, ok := cached.(*template.Template); ok {
			tpl.ExecuteTemplate(c.Writer, pagePath, data)
			return
		}
	}

	// Build template: clone layout + parse page
	base := loadIndexLayoutBase()
	if base == nil {
		// Fallback to default rendering
		c.HTML(http.StatusOK, pagePath, data)
		return
	}

	pageContent, err := os.ReadFile(filepath.Join("web/views", pagePath))
	if err != nil {
		c.HTML(http.StatusOK, pagePath, data)
		return
	}

	pageTpl, err := base.Clone()
	if err != nil {
		c.HTML(http.StatusOK, pagePath, data)
		return
	}

	template.Must(pageTpl.New(pagePath).Parse(string(pageContent)))

	// Cache for future use
	indexLayoutCache.Store(pagePath, pageTpl)

	// Execute the page template
	pageTpl.ExecuteTemplate(c.Writer, pagePath, data)
}
