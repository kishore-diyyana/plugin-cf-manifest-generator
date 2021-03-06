package mg
import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
	"fmt"
	"bufio"
	"os"
	"path"
)

type StepAppname struct {
	Step
}

func NewStepAppname(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *StepAppname {
	stepAppname := new(StepAppname)
	stepAppname.appManifest = appManifest
	stepAppname.cliConnection = cliConnection
	return stepAppname
}

func (s *StepAppname) Run() error {
	reader := bufio.NewReader(os.Stdin)
	defaultName, err := s.defaultName()
	if err != nil {
		return err
	}
	fmt.Print(fmt.Sprintf("What is the name of your app <%s> ? ", defaultName))
	appNameBytes, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	appName := string(appNameBytes)
	if (appName == "") {
		appName = defaultName
	}
	session.AppName = appName
	s.appManifest.Memory(appName, 1024)
	s.appManifest.Instances(appName, 1)
	return nil
}
func (s *StepAppname) defaultName() (string, error) {
	defaultName := path.Base(path.Dir(session.ManifestPath))
	if defaultName != "." {
		return defaultName, nil
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path.Base(wd), nil
}