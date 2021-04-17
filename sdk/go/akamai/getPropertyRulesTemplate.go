// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getPropertyRulesTemplate` data source lets you configure a rule tree through the use of JSON template files. A rule tree is a nested block of property
// rules in JSON format that include match criteria and behaviors.
//
// With this data source you define the location of the JSON template files and provide information about any user-defined variables included within the templates.
//
// The template format used in this data source matches those used in the [Property Manager CLI](https://learn.akamai.com/en-us/learn_akamai/getting_started_with_akamai_developers/developer_tools/getstartedpmcli.html#addanewsnippet).
//
// You can pass user-defined variables by supplying either:
//
// * paths to `variableDefinitions.json` and `variables.json` with syntax used in Property Manager CLI, or
// * a set of provider variables.
//
// ## Referencing sub-files from a template
//
// You can split each template out into a series of smaller template files. To add
// them to this data source, you need to include them in the currently loaded file,
// which corresponds to the value in the `templateFile` argument.  For example, to
// include `example-file.json` from the `template` directory, use this syntax
// including the quotes: `"#include:example-file.json"`.  All files are resolved in
// relation to the directory that contains the starting template file.
//
// ## Inserting variables in a template
//
// You can also add variables to a template by using a string like `“${env.<variableName>}"`. You'll need the quotes here too.\
// These variables follow the format used in the [Property Manager CLI](https://github.com/akamai/cli-property-manager#update-the-variabledefinitions-file).  They differ from the provider variables which should resolve normally.
//
// ## Example Usage
// ## Argument reference
//
// * `templateFile` - (Required) The absolute path to your top-level JSON template file. The top-level template combines smaller, nested JSON templates to form your property rule tree.
// * `variables` - (Optional) A definition of a variable. Variables aren't required and you can use multiple ones if needed. This argument conflicts with the `variableDefinitionFile` and `variableValuesFile` arguments. A `variables` block includes:
//     * `name` - The name of the variable used in template.
//     * `type` - The type of variable: `string`, `number`, `bool`, or `jsonBlock`.
//     * `value` - The value of the variable passed as a string.
// * `variableDefinitionFile` - (Optional) The absolute path to the file containing variable definitions and defaults. This file follows the syntax used in the [Property Manager CLI](https://github.com/akamai/cli-property-manager). This argument is required if you set `variableValuesFile` and conflicts with `variables`.
// * `variableValuesFile` - (Optional) The absolute path to the file containing variable values. This file follows the syntax used in the Property Manager CLI. This argument is required if you set `variableDefinitionFile` and conflicts with `variables`.
//
// ## Attributes reference
//
// This data source returns this attribute:
//
// * `json` - The fully expanded template with variables and all nested templates resolved.
func GetPropertyRulesTemplate(ctx *pulumi.Context, args *GetPropertyRulesTemplateArgs, opts ...pulumi.InvokeOption) (*GetPropertyRulesTemplateResult, error) {
	var rv GetPropertyRulesTemplateResult
	err := ctx.Invoke("akamai:index/getPropertyRulesTemplate:getPropertyRulesTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPropertyRulesTemplate.
type GetPropertyRulesTemplateArgs struct {
	TemplateFile      string                             `pulumi:"templateFile"`
	VarDefinitionFile *string                            `pulumi:"varDefinitionFile"`
	VarValuesFile     *string                            `pulumi:"varValuesFile"`
	Variables         []GetPropertyRulesTemplateVariable `pulumi:"variables"`
}

// A collection of values returned by getPropertyRulesTemplate.
type GetPropertyRulesTemplateResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                string                             `pulumi:"id"`
	Json              string                             `pulumi:"json"`
	TemplateFile      string                             `pulumi:"templateFile"`
	VarDefinitionFile *string                            `pulumi:"varDefinitionFile"`
	VarValuesFile     *string                            `pulumi:"varValuesFile"`
	Variables         []GetPropertyRulesTemplateVariable `pulumi:"variables"`
}
