//go:build wireinject
// +build wireinject

package controller

import (
	"context"

	"github.com/aquaproj/aqua/pkg/config"
	finder "github.com/aquaproj/aqua/pkg/config-finder"
	reader "github.com/aquaproj/aqua/pkg/config-reader"
	"github.com/aquaproj/aqua/pkg/controller/exec"
	"github.com/aquaproj/aqua/pkg/controller/generate"
	"github.com/aquaproj/aqua/pkg/controller/initcmd"
	"github.com/aquaproj/aqua/pkg/controller/install"
	"github.com/aquaproj/aqua/pkg/controller/list"
	"github.com/aquaproj/aqua/pkg/controller/which"
	"github.com/aquaproj/aqua/pkg/download"
	"github.com/aquaproj/aqua/pkg/github"
	registry "github.com/aquaproj/aqua/pkg/install-registry"
	"github.com/aquaproj/aqua/pkg/installpackage"
	"github.com/aquaproj/aqua/pkg/log"
	"github.com/aquaproj/aqua/pkg/runtime"
	"github.com/google/wire"
)

func InitializeListCommandController(ctx context.Context, aquaVersion string, param *config.Param) *list.Controller {
	wire.Build(list.NewController, finder.NewConfigFinder, log.NewLogger, github.New, config.NewRootDir, registry.New, download.NewRegistryDownloader, reader.New, reader.NewFileReader)
	return &list.Controller{}
}

func InitializeInitCommandController(ctx context.Context, aquaVersion string, param *config.Param) *initcmd.Controller {
	wire.Build(initcmd.New, log.NewLogger, github.New)
	return &initcmd.Controller{}
}

func InitializeGenerateCommandController(ctx context.Context, aquaVersion string, param *config.Param) *generate.Controller {
	wire.Build(generate.New, finder.NewConfigFinder, log.NewLogger, github.New, config.NewRootDir, registry.New, download.NewRegistryDownloader, reader.New, reader.NewFileReader)
	return &generate.Controller{}
}

func InitializeInstallCommandController(ctx context.Context, aquaVersion string, param *config.Param) *install.Controller {
	wire.Build(install.New, finder.NewConfigFinder, log.NewLogger, github.New, config.NewRootDir, registry.New, download.NewRegistryDownloader, reader.New, reader.NewFileReader, installpackage.New, download.NewPackageDownloader, runtime.New)
	return &install.Controller{}
}

func InitializeWhichCommandController(ctx context.Context, aquaVersion string, param *config.Param) which.Controller {
	wire.Build(which.New, finder.NewConfigFinder, log.NewLogger, github.New, config.NewRootDir, registry.New, download.NewRegistryDownloader, reader.New, reader.NewFileReader, runtime.New)
	return nil
}

func InitializeExecCommandController(ctx context.Context, aquaVersion string, param *config.Param) *exec.Controller {
	wire.Build(exec.New, finder.NewConfigFinder, log.NewLogger, download.NewPackageDownloader, installpackage.New, github.New, config.NewRootDir, registry.New, download.NewRegistryDownloader, reader.New, reader.NewFileReader, which.New, runtime.New)
	return &exec.Controller{}
}
