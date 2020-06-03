package reporters

import (
	"os"
	"text/template"

	"github.com/reaandrew/reqit/core"
)

type Section struct {
	Title string
	Items map[string]string
}

type CliReportViewModel struct {
	Sections []Section
}

func mapViewModel(result core.Result) CliReportViewModel {
	viewModel := CliReportViewModel{
		Sections: []Section{},
	}

	timings := Section{
		Title: "Timings",
		Items: map[string]string{},
	}

	timings.Items["Connect"] = result.Timings.ConnectDone.String()
	timings.Items["DNS Lookup"] = result.Timings.DnsDone.String()
	timings.Items["TLS Handshake"] = result.Timings.TlsHandshakeDone.String()
	timings.Items["Time to first byte"] = result.Timings.FirstByteDone.String()
	timings.Items["Time to complete"] = result.Timings.Complete.String()
	viewModel.Sections = append(viewModel.Sections, timings)

	return viewModel
}

type CliReporter struct {
}

func (self CliReporter) Execute(result core.Result) {
	reportTemplate := `
Reponse Headers
----------------------------------
{{range $key, $value := .Headers}}{{$key}} = {{$value}}
{{end}}

Timings
----------------------------------
{{range $key, $value := .Timings}}{{$key}} = {{$value}}
{{end}}`
	tmpl, err := template.New("cli").Parse(reportTemplate)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, mapViewModel(result))
}
