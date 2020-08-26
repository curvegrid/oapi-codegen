package codegen

import (
	"bufio"
	"bytes"
	"sort"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
)

// A parsed SecurityScheme.
type SecurityScheme struct {
	Name string
	Spec *openapi3.SecurityScheme
}

// IsStandardBearer returns whether the security scheme is a standard Bearer Authorization header.
func (s *SecurityScheme) IsStandardBearer() bool {
	return s.Spec.Type == "http" && s.Spec.Scheme == "bearer"
}

// IsAPIKey returns whether the security scheme is an API key.
func (s *SecurityScheme) IsAPIKey() bool {
	return s.Spec.Type == "apiKey"
}

// GenerateSecuritySchemes generates the needed security scheme descriptions.
func GenerateSecuritySchemes(t *template.Template, schemes map[string]*openapi3.SecuritySchemeRef) (string, error) {
	ss := make([]SecurityScheme, 0, len(schemes))
	for name, scheme := range schemes {
		if scheme.Ref != "" {
			return "", errors.New("we do not yet support referential security schemes")
		}
		ss = append(ss, SecurityScheme{
			Name: "SecurityScheme" + ToCamelCase(name),
			Spec: scheme.Value,
		})
	}
	sort.Slice(ss, func(i, j int) bool { return ss[i].Name < ss[j].Name })
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	if err := t.ExecuteTemplate(w, "security.tmpl", ss); err != nil {
		return "", errors.Wrapf(err, "error generating")
	}
	if err := w.Flush(); err != nil {
		return "", errors.Wrapf(err, "error flushing after generating")
	}

	return buf.String(), nil
}
