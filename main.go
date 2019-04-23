package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	docopt "github.com/docopt/docopt-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const usage = `sonarjunit

Usage:
  sonarjunit --in <path_in> --out <path_out>

Options:
  --in <path_in>     incoming PhpUnit junit file
  --out <path_out>     outgoing junit file
  -h --help     Show this screen.
`

var (
	log        *logrus.Logger
	testsuites Testsuites
)

func main() {
	initializeLogger()

	cliArgs := parseCLIArgs()
	inputFile := cliArgs["--in"].(string)
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		log.Fatalln(err)
	}
	xmlFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer xmlFile.Close()
	log.Infof("Successfully Opened %s", inputFile)

	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &testsuites)

	for i := 0; i < len(testsuites.Testsuite.Testsuite); i++ {
		for j := 0; j < len(testsuites.Testsuite.Testsuite[i].Testsuite); j++ {
			testsuite := testsuites.Testsuite.Testsuite[i].Testsuite[j]
			if testsuite.File == "" {
				testsuite.File = testsuite.Testcase[0].File
			}
			testsuites.Testsuite.Testsuite[i].Testsuite[j] = testsuite
		}
	}

	xmlResult, err := xml.MarshalIndent(testsuites, "", " ")
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Create(cliArgs["--out"].(string))
	if err != nil {
		log.Fatalln(err)
	}
	l, err := file.Write(xmlResult)
	if err != nil {
		file.Close()
		log.Fatalln(err)
	}
	log.Infoln(l, "bytes written successfully")
}

func initializeLogger() {
	log = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.TraceLevel,
	}
}

func parseCLIArgs() cliArgs {
	args, err := docopt.Parse(usage, nil, true, "dev", true)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "parse cli args"))
	}
	return args
}
