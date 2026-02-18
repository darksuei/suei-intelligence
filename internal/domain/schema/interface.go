package schema

type FieldMappingTypeEnum string

const (
	FieldMappingTypeDirect    FieldMappingTypeEnum = "direct"
	FieldMappingTypeTransform FieldMappingTypeEnum = "transform"
)

type FieldMappingTransformTypeEnum string

const (
	// Type conversion
	FieldMappingTransformTypeToString   FieldMappingTransformTypeEnum = "to_string"
	FieldMappingTransformTypeToInt      FieldMappingTransformTypeEnum = "to_int"
	FieldMappingTransformTypeToFloat    FieldMappingTransformTypeEnum = "to_float"
	FieldMappingTransformTypeToBool     FieldMappingTransformTypeEnum = "to_bool"
	FieldMappingTransformTypeToDateTime FieldMappingTransformTypeEnum = "to_datetime"

	// String normalization
	FieldMappingTransformTypeLowercase  FieldMappingTransformTypeEnum = "lowercase"
	FieldMappingTransformTypeUppercase  FieldMappingTransformTypeEnum = "uppercase"
	FieldMappingTransformTypeTrim       FieldMappingTransformTypeEnum = "trim"
	FieldMappingTransformTypeHashSHA256 FieldMappingTransformTypeEnum = "hash_sha256"

	// Structural transforms
	FieldMappingTransformTypeConcat        FieldMappingTransformTypeEnum = "concat"
	FieldMappingTransformTypeValueMap      FieldMappingTransformTypeEnum = "value_map"
	FieldMappingTransformTypeDefaultIfNull FieldMappingTransformTypeEnum = "default_if_null"
	FieldMappingTransformTypeNullIfEmpty   FieldMappingTransformTypeEnum = "null_if_empty"
)

type SchemaMapping struct {
	IntegrationID         string                 `json:"integrationId"`
	InternalSchemaVersion string                 `json:"internalSchemaVersion"`
	MappingVersion        string                 `json:"mappingVersion"`
	Entities              []SchemaMappingEntity  `json:"entities"`
}

type SchemaMappingEntity struct {
	Name           string                `json:"name"`
	SourceEntity   SchemaMappingSource   `json:"sourceEntity"`
	InternalEntity SchemaMappingInternal `json:"internalEntity"`
	IDStrategy     IDStrategyDefinition  `json:"idStrategy"`
	FieldMappings  []FieldMapping        `json:"fieldMappings"`
}

type SchemaMappingSource struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type SchemaMappingSourceField struct {
	Name   string `json:"name"`
	Entity string `json:"entity,omitempty"` // Optional: allow pointing to another source entity for multi-source transforms
}

type SchemaMappingInternal struct {
	Name string `json:"name"`
}

type SchemaMappingInternalField struct {
	Name string `json:"name"`
}

type IDStrategyTypeEnum string

const (
	IDStrategyTypeSourcePrimaryKey IDStrategyTypeEnum = "source_primary_key"
	IDStrategyTypeComposite        IDStrategyTypeEnum = "composite"
	IDStrategyTypeGeneratedUUID    IDStrategyTypeEnum = "generated_uuid"
)

type IDStrategyDefinition struct {
	Type         IDStrategyTypeEnum          `json:"type"`
	SourceFields []SchemaMappingSourceField  `json:"sourceFields,omitempty"` // used for source_primary_key or composite
}

type FieldMapping struct {
	Required      bool                       `json:"required"`
	Type          FieldMappingTypeEnum       `json:"type"`
	InternalField SchemaMappingInternalField `json:"internalField"`
	SourceField   SchemaMappingSourceField   `json:"sourceField,omitempty"`   // optional for transform-only mappings
	Transforms    []FieldMappingTransform    `json:"transforms,omitempty"`    // applied sequentially
}

type FieldMappingTransform struct {
	Type   FieldMappingTransformTypeEnum `json:"type"`
	Config map[string]interface{}        `json:"config,omitempty"`
	// Example configs:
	// to_datetime -> { "format": "2006-01-02 15:04:05" }
	// concat -> { "fields": [{"entity": "customers", "field": "first_name"}, {"entity": "customers", "field": "last_name"}], "separator": " " }
	// value_map -> { "mapping": { "A": "active", "1": "active" } }
	// default_if_null -> { "value": "unknown" }
}