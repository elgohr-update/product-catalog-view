/*
 * Product Catalog View
 *
 * Product Catalog View
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A use of the ConfigurableSpecificationCharacteristicValueUse is to provide flexibility on configuring Specification Characteristic Value.
type ConfigurableSpecificationCharacteristicValueUse struct {
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	//  hyperlink reference to the schema describing this resource.
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty"`
	// A String. This is used to get configurationStage at productOffering level
	ConfigurationStage string `json:"configurationStage,omitempty"`
	// A narrative that explains in detail what the productSpecificationCharacteristic is
	Description string `json:"description,omitempty"`
	// A String.  This is used to get displayName at productOffering level
	DisplayName string `json:"displayName,omitempty"`
	// This defines on what order the characteristics will be exposed.
	DisplaySequence string `json:"displaySequence,omitempty"`
	// This defines the characteristics will be exposed or not.
	IsVisible string `json:"isVisible,omitempty"`
	// The maximum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where five is the value for the maxCardinality.
	MaxCardinality int32 `json:"maxCardinality,omitempty"`
	// The minimum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where zero is the value for the minCardinality.
	MinCardinality int32 `json:"minCardinality,omitempty"`
	// Name of the associated productSpecificationCharacteristic
	Name string `json:"name"`
	ReferenceId string `json:"referenceId,omitempty"`
	// A number or text that can be assigned to a SpecificationCharacteristic.
	SpecCharacteristicValue []SpecificationCharacteristicValue `json:"specCharacteristicValue"`
	// A Specification is a detailed description of a tangible or intangible object made available externally in the form of a ProductOffering to customers or other parties playing a party role.
	Specification *SpecificationRef `json:"specification,omitempty"`
	// When sub-classing, this defines the sub-class specification type
	SpecificationType string `json:"specificationType,omitempty"`
	// The period for which the productSpecificationCharacteristic is valid
	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// A kind of value that the characteristic can take on, such as numeric, text and so forth
	ValueType string `json:"valueType,omitempty"`
}
