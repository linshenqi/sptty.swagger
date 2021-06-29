package swagger

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	sw "github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/linshenqi/sptty"
	"github.com/linshenqi/sptty.swagger/src/base"
)

type Service struct {
	sptty.BaseService

	cfg Config
}

func (s *Service) Init(app sptty.ISptty) error {
	if err := app.GetConfig(base.ServiceSwagger, &s.cfg); err != nil {
		return err
	}

	if !s.cfg.Enable {
		sptty.Log(sptty.InfoLevel, "Swagger Is Disabled", s.ServiceName())
		return nil
	}

	s.generateDoc()

	app.AddRoute("GET", "/swagger/{any:path}", sw.WrapHandler(swaggerFiles.Handler, sw.URL(s.cfg.Url)))
	return nil
}

func (s *Service) ServiceName() string {
	return base.ServiceSwagger
}

func (s *Service) generateDoc() {

	workingDir, err := os.Getwd()
	if err != nil {
		return
	}

	mainDir := getDirOfMain(workingDir)

	workingDir = strings.ReplaceAll(workingDir, "\\", "/")
	mainDir = strings.ReplaceAll(mainDir, "\\", "/")

	rt := exec.Command("swag", "init", "-d", mainDir, "-o", path.Join(workingDir, "doc"))
	if _, err = rt.Output(); err != nil {
		sptty.Log(sptty.ErrorLevel, err.Error(), s.ServiceName())
		return
	}

}

func getDirOfMain(dir string) string {

	rt := ""
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return nil
		}

		if !info.IsDir() && info.Name() == "main.go" {
			rt = filepath.Dir(path)
			return fmt.Errorf("")
		}

		return nil
	})

	return rt
}
