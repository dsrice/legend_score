package rg

import (
	"github.com/dave/jennifer/jen"
	"legend_score/infra/genarator"
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

	return nil
}

func createRi(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("ri")

	f.Type().Id(cg.In + "Repository").Interface()

	f.Save(path.Join(cg.BasePath, "ri", cg.In+".go"))

	return nil
}

func createImp(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("repositories")

	f.ImportName("database/sql", "sql")
	f.ImportName("app/repositories/ri", "ri")

	f.Type().Id(cg.Fn + "Repository").Struct(
		jen.Id("con").Op("*").Qual("database/sql", "DB"),
	)

	f.Func().Id("New"+cg.In+"Repository").Params(
		jen.Id("con").Op("*").Qual("app/infra/database/connection", "Connection"),
	).Qual("app/repositories/ri", cg.In+"Repository").Block(
		jen.Return(jen.Op("&").Id(cg.Fn + "Repository").Values(
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