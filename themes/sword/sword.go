package sword

import (
	"embed"
	"strings"

	adminTemplate "github.com/nneesshh/go-admin/template"
	"github.com/nneesshh/go-admin/template/components"
	"github.com/nneesshh/go-admin/template/types"
	"github.com/nneesshh/go-admin/themes/common"
	"github.com/nneesshh/go-admin/themes/sword/resource"
)

type Theme struct {
	ThemeName string
	components.Base
	*common.BaseTheme
}

var Sword = Theme{
	ThemeName: "sword",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: TemplateList,
		},
	},
	BaseTheme: &common.BaseTheme{
		AssetPaths:   resource.AssetPaths,
		TemplateList: TemplateList,
	},
}

func init() {
	adminTemplate.Add("sword", &Sword)
}

func Get() *Theme {
	return &Sword
}

func (t *Theme) Name() string {
	return t.ThemeName
}

func (t *Theme) GetTmplList() map[string]string {
	return TemplateList
}

//go:embed resource/assets/dist
var fAdminlteResourceAssetsDist embed.FS

func (t *Theme) GetAsset(path string) ([]byte, error) {
	embedPath := strings.ReplaceAll(path, "/assets/dist/", "resource/assets/dist/")
	return fAdminlteResourceAssetsDist.ReadFile(embedPath)
}

func (t *Theme) GetAssetList() []string {
	return resource.AssetsList
}
