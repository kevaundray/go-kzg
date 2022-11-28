package eth_test

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	"github.com/protolambda/go-kzg/eth"
)

type PrecompileInput struct {
	Input         string
	Expected      string
	ExpectedError string
	Name          string
	Gas           int
	NoBenchmark   bool
}

// Package level variable to try and stop golang from doing
// benchmark optimisations on the unused variable
var _noOpt = make([]byte, 30)

func BenchmarkPrecompiles(b *testing.B) {

	content, err := ioutil.ReadFile("./fail-PointEval.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload []PrecompileInput
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	for _, pl := range payload {
		inp, _ := hex.DecodeString(pl.Input)
		var res []byte
		b.Run(pl.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ = eth.PointEvaluationPrecompile(inp)
			}
			_noOpt = res
		})
	}
}
