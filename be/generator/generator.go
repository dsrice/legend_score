package main

import (
	"legend_score/controllers/cg"
	"legend_score/infra/genarator"
	"legend_score/repositories/rg"
	"legend_score/usecases/ug"
	"log"
	"os"
	"strings"
)

const (
	controller = "1"
	usecase    = "2"
	repo       = "3"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("引数がたりません")
		os.Exit(9)
	}

	cgs := genarator.CreateGenerator{
		In: os.Args[2],
	}

	cgs.Fn = createImpName(cgs.In)
	cgs.Mn = createMockName(cgs.In)

	ts := os.Args[1]
	var err error

	switch ts {
	case controller:
		err = cg.CreateController(&cgs)
	case usecase:
		err = ug.CreateUsecase(&cgs)
	case repo:
		err = rg.CreateRepository(&cgs)
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}

func createImpName(name string) string {
	sName := strings.Split(name, "")

	sName[0] = strings.ToLower(sName[0])

	return strings.Join(sName, "") + "Imp"
}

func createMockName(name string) string {
	return name + "Mock"
}