package ug

import (
	"bufio"
	"fmt"
	"github.com/dave/jennifer/jen"
	"legend_score/infra/genarator"
	"os"
	"path"
)

func CreateUsecase(cg *genarator.CreateGenerator) error {
	cg.BasePath = "/go/src/app/usecases/"

	err := createCi(cg)
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

func createCi(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("ui")

	f.Type().Id(cg.In).Interface()

	f.Save(path.Join(cg.BasePath, "ui", cg.In+".go"))

	return nil
}

func createImp(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("usecases")

	f.ImportName("app/usecases/ui", "ui")
	f.ImportName("app/repositoriesri", "ri")

	f.Type().Id(cg.Fn).Struct()

	f.Func().Id("New"+cg.In).Params(
		jen.Id("repo").Qual("app/repositories/ri", "InRepository"),
	).Qual("app/usecases/ui", cg.In).Block(
		jen.Return(jen.Op("&").Id(cg.Fn).Values()),
	)

	f.Save(path.Join(cg.BasePath, cg.Fn+".go"))

	return nil
}

func createImpTest(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("usecases_test")

	f.Save(path.Join(cg.BasePath, cg.Fn+"_test.go"))

	return nil
}

func createMock(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("umock")

	f.Type().Id(cg.Mn).Struct(
		jen.Qual("github.com/stretchr/testify/mock", "Mock"),
	)

	f.Save(path.Join(cg.BasePath, "umock", cg.Mn+".go"))

	return nil
}

func addInterface(cg *genarator.CreateGenerator) error {
	path := path.Join(cg.BasePath, "ui", "inUsecase.go")

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