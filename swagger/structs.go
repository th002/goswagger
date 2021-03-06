package swagger

//SwaggerVersion swagger version
const SwaggerVersion = "2.0"

const (
	//ContentTypeJSON content type json
	ContentTypeJSON = "application/json"
	//ContentTypeXML content type XML
	ContentTypeXML = "application/xml"
	//ContentTypePlain content type plain
	ContentTypePlain = "text/plain"
	//ContentTypeHTML content type html
	ContentTypeHTML = "text/html"
	//ContentTypeMultiPartFormData content type form-data
	ContentTypeMultiPartFormData = "multipart/form-data"
	//ContentTypeXFORMURLENCODE content type x-www-form-urlencoded
	ContentTypeXFORMURLENCODE = "application/x-www-form-urlencoded"
)

//SwaggerObject swagger object
type SwaggerObject struct {
	Swagger             string                       `json:"swagger"`
	Info                *InfoObject                  `json:"info"`
	Host                string                       `json:"host,omitempty"`
	BasePath            string                       `json:"basePath,omitempty"`
	Schemes             []string                     `json:"schemes,omitempty"`
	Consumes            []string                     `json:"consumes,omitempty"`
	Produces            []string                     `json:"produces,omitempty"`
	Paths               map[string]*PathItemObject   `json:"paths"`
	Definitions         map[string]*SchemaObject     `json:"definitions,omitempty"`
	Parameters          map[string]interface{}       `json:"patameters,omitempty"`
	Responses           map[string]interface{}       `json:"responses,omitempty"`
	SecurityDefinitions map[string]interface{}       `json:"securityDefinitions,omitempty"`
	Security            map[string][]string          `json:"security,omitempty"`
	Tags                []*TagObject                 `json:"tags,omitempty"`
	ExternalDocs        *ExternalDocumentationObject `json:"externalDocs,omitempty"`
}

//InfoObject swagger info
type InfoObject struct {
	Title          string   `json:"title"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Version        string   `json:"version,omitempty"`
}

//Contact swagger contact
type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

//License swagger License
type License struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

//PathItemObject swagger path item
type PathItemObject struct {
	Ref        string           `json:"$ref,omitempty"`
	Get        *OperationObject `json:"get,omitempty"`
	Put        *OperationObject `json:"put,omitempty"`
	Post       *OperationObject `json:"post,omitempty"`
	Delete     *OperationObject `json:"delete,omitempty"`
	Options    *OperationObject `json:"options,omitempty"`
	Head       *OperationObject `json:"head,omitempty"`
	Patch      *OperationObject `json:"patch,omitempty"`
	Parameters []interface{}    `json:"parameters,omitempty"`
}

//ParameterObject swagger parameter
type ParameterObject struct {
	Ref              string        `json:"$ref,omitempty"`
	Name             string        `json:"name"`
	In               string        `json:"in"`
	Description      string        `json:"description,omitempty"`
	Required         bool          `json:"required,omitempty"`
	Schema           *SchemaObject `json:"schema,omitempty"`
	Type             string        `json:"type,omitempty"`
	Format           string        `json:"format,omitempty"`
	AllowEmptyValue  bool          `json:"allowEmptyValue,omitempty"`
	Items            []interface{} `json:"items,omitempty"`
	CollectionFormat string        `json:"collectFormat,omitempty"`
	// TODO ...
}

//OperationObject swagger operation
type OperationObject struct {
	Tags         []string                     `json:"tags,omitempty"`
	Summary      string                       `json:"summary,omitempty"`
	Description  string                       `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentationObject `json:"externalDocs,omitempty"`
	OperationID  string                       `json:"operationId,omitempty"`
	Consumes     []string                     `json:"consumes,omitempty"`
	Produces     []string                     `json:"produces,omitempty"`
	Parameters   []interface{}                `json:"parameters,omitempty"`
	Responses    map[string]*ResponseObject   `json:"responses,omitempty"`
	Schemes      []string                     `json:"schemes,omitempty"`
	Deprecated   bool                         `json:"deprecated,omitempty"`
	Security     map[string][]string          `json:"security,omitempty"`
	parser       *Parser
	packageName  string
}

//ReferenceObject swagger reference
type ReferenceObject struct {
	Ref string `json:"$ref,omitempty"`
}

//ResponseObject or ReferenceObject
type ResponseObject struct {
	Ref         string                 `json:"$ref,omitempty"`
	Description string                 `json:"description,omitempty"`
	Schema      *SchemaObject          `json:"schema,omitempty"`
	Headers     HeadersObject          `json:"headers,omitempty"`
	Examples    map[string]interface{} `json:"examples,omitempty"`
}

//SchemaObject swagger schema
type SchemaObject struct {
	Ref        string                 `json:"$ref,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Required   []string               `json:"required,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Items      *ReferenceObject       `json:"items,omitempty"`
}

//HeadersObject swagger headers
type HeadersObject map[string]*HeaderObject

//HeaderObject swagger header
type HeaderObject struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}

//TagObject swagger tag
type TagObject struct {
	Name         string                      `json:"name"`
	Description  string                      `json:"description,omitempty"`
	ExternalDocs ExternalDocumentationObject `json:"externalDocs,omitempty"`
}

//ExternalDocumentationObject swagger external documentation
type ExternalDocumentationObject struct {
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
}

// type Property struct {
// 	Type        string `json:"type"`
// 	Format      string `json:"format,omitempty"`
// 	Description string `json:"description,omitempty"`
// }
