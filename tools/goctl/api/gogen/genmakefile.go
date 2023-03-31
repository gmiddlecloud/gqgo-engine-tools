package gogen

import (
	_ "embed"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

//go:embed makefile.tpl
var makefileTemplate string

func genMakefile(dir string, api *spec.ApiSpec, g *GenContext) error {
	service := api.Service

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          "",
		filename:        "Makefile",
		templateName:    "makefileTemplate",
		category:        category,
		templateFile:    makefileTemplateFile,
		builtinTemplate: makefileTemplate,
		data: map[string]any{
			"serviceName": service.Name,
			"useEnt":      g.UseEnt,
		},
	})
}
