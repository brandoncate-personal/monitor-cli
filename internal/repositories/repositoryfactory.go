package repositories

import (
	"errors"

	"github.com/brandoncate-personal/monitor-cli/internal/core/ports"
)

type ExporterType string

const (
	Http ExporterType = "http"
)

// Here, NewPerson returns an interface, and not the person struct itself
func NewExporter(exporterType ExporterType, params interface{}) (ports.ExporterService, error) {

	switch exporterType {
	case Http:
		httpParams, ok := params.(HttpExporter)
		if !ok {
			return nil, errors.New("required params for http exporter not passed")
		}
		return HttpExporter{
			Url: httpParams.Url,
		}, nil
	default:
		return nil, errors.New("unsupported exporter type")
	}

}
