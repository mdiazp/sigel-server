package public

import (
	"time"

	"github.com/mdiazp/sirel-server/api/controllers"
)

// InfoController ...
type InfoController struct {
	controllers.BaseController
}

// ServerTime ...
// @Title Retrieve serve time info
// @Description Get server time info
// @Success 200 {object} time.Time
// @Failure 500 Internal Server Error
// @Accept json
// @router /servertime [get]
func (c *InfoController) ServerTime() {
	c.Data["json"] = time.Now()
	c.ServeJSON()
}
