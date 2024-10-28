package settings

import (
	"github.com/gin-gonic/gin"
	"gvb-server/models/res"
)

func (s *SettingsApi) SettingsInfoView(c *gin.Context) {
	//res.Ok(map[string]any{}, "aurora", c)
	//res.OkWithData(map[string]string{}, c)
	res.FailWithCode(1, c)
}
