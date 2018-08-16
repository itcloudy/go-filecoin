package commands

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
)

// The file used to build these addresses can be found in:
// $GOPATH/src/github.com/filecoin-project/go-filecoin/testhelpers/testfiles/walletGenFile.toml
//
// If said file is modified these addresses will need to change as well
//
// The method to generate a new file can be found in:
// $GOPATH/src/github.com/filecoin-project/go-filecoin/testhelpers/util.go
//    GenerateWalletFile(numAddrs)
//    WriteWalletFile(file string, wf WalletFile)
const testAddress1 = "fcqh6v455ywfrjz70zny6n893qjyasq4ypgwk5vhu" // nolint: deadcode
const testAddress2 = "fcqmjv8n8zekqz62gq5ra9fn0zapadq0fjeqd32sj" // nolint: deadcode
const testAddress3 = "fcqrn3nwxlpqng6ms8kp4tk44zrjyh4nurrmg6wth" // nolint: deadcode
const testAddress4 = "fcqvadmgjj3a3v8dx53m946ex7cva3vdxnk34s0ky" // nolint: deadcode, varcheck, megacheck
const testAddress5 = "fcqnmxl42szla5qkczwg053fnuuttsrd29nr2gftw" // nolint: deadcode, varcheck, megacheck

func requireSchemaConformance(t *testing.T, jsonBytes []byte, schemaName string) { // nolint: deadcode
	wdir, _ := os.Getwd()
	rLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s/schema/%s.schema.json", wdir, schemaName))
	jLoader := gojsonschema.NewBytesLoader(jsonBytes)

	result, err := gojsonschema.Validate(rLoader, jLoader)
	require.NoError(t, err)

	for _, desc := range result.Errors() {
		t.Errorf("- %s\n", desc)
	}

	require.Truef(t, result.Valid(), "Error schema validating: %s", string(jsonBytes))
}