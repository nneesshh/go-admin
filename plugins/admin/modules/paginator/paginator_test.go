package paginator

import (
	"testing"

	"github.com/nneesshh/go-admin/modules/config"
	"github.com/nneesshh/go-admin/plugins/admin/modules/parameter"
	_ "github.com/nneesshh/themes/sword"
)

func TestGet(t *testing.T) {
	config.Initialize(&config.Config{Theme: "sword"})
	Get(Config{
		Size:         105,
		Param:        parameter.BaseParam().SetPage("7"),
		PageSizeList: []string{"10", "20", "50", "100"},
	})
}
