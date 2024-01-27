package health_route

import (
	"github.com/evanhongo/happy-golang/pkg/schema"
	"github.com/softbrewery/gojoi/pkg/joi"
)

type PingRequestSchema struct {
	schema *joi.StructSchema
}

func (s *PingRequestSchema) Parse(data any) error {
	if err := s.schema.Validate(data); err != nil {
		return err
	}
	return nil
}

func NewPingRequestSchema() schema.ISchema {
	return &PingRequestSchema{
		schema: joi.Struct().Keys(joi.StructKeys{
			"Hello": joi.String().Min(1),
		}),
	}
}
