package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"html/template"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
)

// 1. go:embed
//
//go:embed template
var templates embed.FS

type Generator struct {
	pkg string
	typ string
}

type Target struct {
	Pkg string
	Typ string
	Arr string
}

func NewGenerator(dir string, typ string) (*Generator, error) {
	pkg, err := packageInfo(dir)
	if err != nil {
		return nil, err
	}

	return &Generator{
		pkg: pkg.Name,
		typ: typ,
	}, nil
}

func packageInfo(dir string) (*packages.Package, error) {
	// 1. パッケージ名を取得するために、指定したディレクトリを静的解析
	pkgs, err := packages.Load(&packages.Config{
		Mode:  packages.NeedName,
		Tests: false,
	}, dir)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to load package dir=%v", dir)
	}
	if len(pkgs) != 1 {
		return nil, errors.Wrapf(err, "%d packages found", len(pkgs))
	}

	return pkgs[0], nil
}

func (g *Generator) Run() ([]byte, error) {
	w := &bytes.Buffer{}

	// 2. テンプレート読み込み
	tmpl, err := template.ParseFS(templates, "template/*")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse templetes")
	}

	t := &Target{
		Pkg: g.pkg,
		Typ: g.typ,
		Arr: fmt.Sprintf("%ss", g.typ),
	}

	// 3. テンプレート処理
	err = tmpl.ExecuteTemplate(w, "function", t)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}

	fmted, err := format.Source(w.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "failed to format code")
	}

	return fmted, nil
}
