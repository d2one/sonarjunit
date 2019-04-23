package main

import "encoding/xml"

type cliArgs map[string]interface{}

//Testsuites struct of junit
type Testsuites struct {
	XMLName   xml.Name `xml:"testsuites"`
	Text      string   `xml:",chardata"`
	Testsuite struct {
		Text       string `xml:",chardata"`
		Name       string `xml:"name,attr"`
		Tests      string `xml:"tests,attr"`
		Assertions string `xml:"assertions,attr"`
		Errors     string `xml:"errors,attr"`
		Failures   string `xml:"failures,attr"`
		Skipped    string `xml:"skipped,attr"`
		Time       string `xml:"time,attr"`
		Testsuite  []struct {
			Text       string `xml:",chardata"`
			Name       string `xml:"name,attr"`
			File       string `xml:"file,attr"`
			Tests      string `xml:"tests,attr"`
			Assertions string `xml:"assertions,attr"`
			Errors     string `xml:"errors,attr"`
			Failures   string `xml:"failures,attr"`
			Skipped    string `xml:"skipped,attr"`
			Time       string `xml:"time,attr"`
			Testsuite  []struct {
				Text       string `xml:",chardata"`
				Name       string `xml:"name,attr"`
				File       string `xml:"file,attr,omitempty"`
				Tests      string `xml:"tests,attr"`
				Assertions string `xml:"assertions,attr"`
				Errors     string `xml:"errors,attr"`
				Failures   string `xml:"failures,attr"`
				Skipped    string `xml:"skipped,attr"`
				Time       string `xml:"time,attr"`
				Testcase   []struct {
					Text       string `xml:",chardata"`
					Name       string `xml:"name,attr"`
					Class      string `xml:"class,attr"`
					Classname  string `xml:"classname,attr"`
					File       string `xml:"file,attr"`
					Line       string `xml:"line,attr"`
					Assertions string `xml:"assertions,attr"`
					Time       string `xml:"time,attr"`
				} `xml:"testcase"`
			} `xml:"testsuite"`
			Testcase []struct {
				Text       string `xml:",chardata"`
				Name       string `xml:"name,attr"`
				Class      string `xml:"class,attr"`
				Classname  string `xml:"classname,attr"`
				File       string `xml:"file,attr"`
				Line       string `xml:"line,attr"`
				Assertions string `xml:"assertions,attr"`
				Time       string `xml:"time,attr"`
			} `xml:"testcase"`
		} `xml:"testsuite"`
	} `xml:"testsuite"`
}
