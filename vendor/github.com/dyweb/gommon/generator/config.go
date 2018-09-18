package generator

import (
	"bytes"
	"fmt"
	"go/format"

	"github.com/dyweb/gommon/errors"
)

type Config struct {
	Loggers     []LoggerConfig     `yaml:"loggers"`
	GoTemplates []GoTemplateConfig `yaml:"gotmpls"`
	Shells      []ShellConfig      `yaml:"shells"`
	// set when traversing the folders
	pkg  string
	file string
}

func NewConfig(pkg string, file string) *Config {
	return &Config{pkg: pkg, file: file}
}

func (c *Config) RenderGommon() ([]byte, error) {
	body := &bytes.Buffer{}
	header := &bytes.Buffer{}
	fmt.Fprintf(header, Header(generatorName, c.file))
	fmt.Fprint(header, "\n")
	fmt.Fprintf(header, "package %s\n\n", c.pkg)
	if len(c.Loggers) > 0 {
		fmt.Fprintln(header, "import dlog \"github.com/dyweb/gommon/log\"")
		for _, l := range c.Loggers {
			if err := l.RenderTo(body); err != nil {
				return nil, err
			}
		}
	}
	header.Write(body.Bytes())
	//log.Debug(string(header.Bytes()))
	// format go code
	if formatted, err := format.Source(header.Bytes()); err != nil {
		return formatted, errors.Wrap(err, "can't format generated code")
	} else {
		//log.Debugf("formatted len %d", len(formatted))
		return formatted, nil
	}
}

func (c *Config) RenderGoTemplate(root string) error {
	if len(c.GoTemplates) == 0 {
		log.Debugf("no go template specified in file %s", c.file)
		return nil
	}
	for _, t := range c.GoTemplates {
		if err := t.Render(root); err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) RenderShell(root string) error {
	if len(c.Shells) == 0 {
		log.Debugf("no shell specified in file %s", c.file)
		return nil
	}
	for _, s := range c.Shells {
		if err := s.Render(root); err != nil {
			return err
		}
	}
	return nil
}
