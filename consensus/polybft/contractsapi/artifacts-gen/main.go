package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/dave/jennifer/jen"

	"github.com/0xPolygon/polygon-edge/consensus/polybft/contractsapi/artifact"
)

const (
	extension = ".sol"
)

func main() {
	_, filename, _, _ := runtime.Caller(0) //nolint: dogsled
	currentPath := path.Dir(filename)
	scpath := path.Join(currentPath, "../../../../core-contracts/artifacts/contracts/")

	f := jen.NewFile("contractsapi")
	f.Comment("This is auto-generated file. DO NOT EDIT.")

	readContracts := []struct {
		Path string
		Name string
	}{
		{
			"child/L2StateSender.sol",
			"L2StateSender",
		},
		{
			"child/StateReceiver.sol",
			"StateReceiver",
		},
		{
			"child/System.sol",
			"System",
		},
		{
			"common/BLS.sol",
			"BLS",
		},
		{
			"common/BN256G2.sol",
			"BN256G2",
		},
		{
			"common/Merkle.sol",
			"Merkle",
		},
		{
			"root/CheckpointManager.sol",
			"CheckpointManager",
		},
		{
			"root/ExitHelper.sol",
			"ExitHelper",
		},
		{
			"root/StateSender.sol",
			"StateSender",
		},
		{
			"child/ChildValidatorSet.sol",
			"ChildValidatorSet",
		},
		{
			"child/LiquidityToken.sol",
			"LiquidityToken",
		},
	}

	for _, v := range readContracts {
		artifactBytes, err := artifact.ReadArtifactData(scpath, v.Path, getContractName(v.Path))
		if err != nil {
			log.Fatal(err)
		}

		f.Var().Id(v.Name + "Artifact").String().Op("=").Lit(string(artifactBytes))
	}

	fl, err := os.Create(currentPath + "/../gen_sc_data.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(fl, "%#v", f)
	if err != nil {
		log.Fatal(err)
	}
}

// getContractName extracts smart contract name from provided path
func getContractName(path string) string {
	pathSegments := strings.Split(path, string([]rune{os.PathSeparator}))
	nameSegment := pathSegments[len(pathSegments)-1]

	return strings.Split(nameSegment, extension)[0]
}
