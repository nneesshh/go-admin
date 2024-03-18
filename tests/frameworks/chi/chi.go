package chi

import (
	// add chi adapter
	_ "github.com/nneesshh/go-admin/adapter/chi"
	"github.com/nneesshh/go-admin/modules/config"
	"github.com/nneesshh/go-admin/modules/language"
	"github.com/nneesshh/go-admin/plugins/admin/modules/table"

	// add mysql driver
	_ "github.com/nneesshh/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/nneesshh/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/nneesshh/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/nneesshh/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	"github.com/nneesshh/go-admin/themes/adminlte"

	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/nneesshh/go-admin/engine"
	"github.com/nneesshh/go-admin/plugins/admin"
	"github.com/nneesshh/go-admin/plugins/example"
	"github.com/nneesshh/go-admin/template"
	"github.com/nneesshh/go-admin/template/chartjs"
	"github.com/nneesshh/go-admin/tests/tables"
)

func internalHandler() http.Handler {
	r := chi.NewRouter()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)
	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(r); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return r
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	r := chi.NewRouter()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(gens)
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfig(&config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddPlugins(adminPlugin).Use(r); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return r
}
