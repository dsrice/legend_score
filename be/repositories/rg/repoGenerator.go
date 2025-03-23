package rg

import (
	"bufio"
	"fmt"
	"github.com/dave/jennifer/jen"
	"legend_score/infra/genarator"
	"os"
	"path"
)

func CreateRepository(cg *genarator.CreateGenerator) error {
	cg.BasePath = "/go/src/app/repositories/"

	err := createRi(cg)
	if err != nil {
		return err
	}

	err = createImp(cg)
	if err != nil {
		return err
	}

	err = createImpTest(cg)
	if err != nil {
		return err
	}

	err = createMock(cg)
	if err != nil {
		return err
	}

	err = addInterface(cg)
	if err != nil {
		return err
	}

	return nil
}

func createRi(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("ri")

	f.Type().Id(cg.In).Interface()

	f.Save(path.Join(cg.BasePath, "ri", cg.In+".go"))

	return nil
}

func createImp(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("repositories")

	f.ImportName("database/sql", "sql")
	f.ImportName("app/repositories/ri", "ri")

	f.Type().Id(cg.Fn).Struct(
		jen.Id("con").Op("*").Qual("database/sql", "DB"),
	)

	f.Func().Id("New"+cg.In).Params(
		jen.Id("con").Op("*").Qual("app/infra/database/connection", "Connection"),
	).Qual("app/repositories/ri", cg.In).Block(
		jen.Return(jen.Op("&").Id(cg.Fn).Values(
			jen.Dict{jen.Id("con"): jen.Id("con.Con")}),
		),
	)

	// fmt.Printf("%#v", f)

	f.Save(path.Join(cg.BasePath, cg.Fn+".go"))

	return nil
}

func createImpTest(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("repositories_test")

	f.Save(path.Join(cg.BasePath, cg.Fn+"_test.go"))

	return nil
}

func createMock(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("rmock")

	f.Type().Id(cg.Mn).Struct(
		jen.Qual("github.com/stretchr/testify/mock", "Mock"),
	)

	f.Save(path.Join(cg.BasePath, "rmock", cg.Mn+".go"))

	return nil
}

func addInterface(cg *genarator.CreateGenerator) error {
	path := path.Join(cg.BasePath, "ri", "inRepository.go")

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0775)

	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}

	for scanner.Scan() {
		// ここで一行ずつ処理
		t := scanner.Text()

		if t == "}" {
			add := fmt.Sprintf("	%s %s", cg.In, cg.In)
			lines = append(lines, add)
		}

		lines = append(lines, t)
	}

	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0775)

	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range lines {
		f.WriteString(line + "\n")
	}

	return nil
}