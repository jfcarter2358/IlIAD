package page

import (
	"iad/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectIndexPage(c *gin.Context) {
	c.Redirect(301, "/ui/dashboard")
}

func showPage(c *gin.Context, page string, header gin.H) {
	header["version"] = constants.VERSION

	render(c, header, page)
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
