package cg

import (
	"github.com/dave/jennifer/jen"
	"legend_score/infra/genarator"
	"path"
)

func CreateController(cg *genarator.CreateGenerator) error {
	cg.BasePath = "/go/src/app/controllers/"

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

	return nil
}

func createCi(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("ci")

	f.Type().Id(cg.In + "Controller").Interface()

	f.Save(path.Join(cg.BasePath, "ci", cg.In+".go"))

	return nil
}

func createImp(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("controllers")

	f.ImportName("legend_score/controllers/ci", "ci")
	f.ImportName("legend_score/usecases/ui", "ui")

	f.Type().Id(cg.Fn + "Controller").Struct()

	f.Func().Id("New"+cg.In).Params().
		Qual("legend_score/controllers/ci", cg.In+"Controller").Block(
		jen.Return(jen.Op("&").Id(cg.Fn + "Controller").Values()),
	)

	f.Save(path.Join(cg.BasePath, cg.Fn+".go"))

	return nil
}

func createImpTest(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("controllers_test")

	f.Save(path.Join(cg.BasePath, cg.Fn+"_test.go"))

	return nil
}