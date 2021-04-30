// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetPropertyRulesTemplate
    {
        /// <summary>
        /// The `akamai.getPropertyRulesTemplate` data source lets you configure a rule tree through the use of JSON template files. A rule tree is a nested block of property 
        /// rules in JSON format that include match criteria and behaviors. 
        /// 
        /// With this data source you define the location of the JSON template files and provide information about any user-defined variables included within the templates.
        /// 
        /// The template format used in this data source matches those used in the [Property Manager CLI](https://learn.akamai.com/en-us/learn_akamai/getting_started_with_akamai_developers/developer_tools/getstartedpmcli.html#addanewsnippet).
        /// 
        /// You can pass user-defined variables by supplying either: 
        /// 
        /// * paths to `variableDefinitions.json` and `variables.json` with syntax used in Property Manager CLI, or 
        /// * a set of provider variables.
        /// 
        /// ## Referencing sub-files from a template
        /// 
        /// You can split each template out into a series of smaller template files. To add 
        /// them to this data source, you need to include them in the currently loaded file, 
        /// which corresponds to the value in the `template_file` argument.  For example, to 
        /// include `example-file.json` from the `property-snippets` directory, use this syntax 
        /// including the quotes: `"#include:example-file.json"`.  Make sure the `property-snippets` folder contains only `.json` files.
        /// All files are resolved in relation to the directory that contains the starting template file. 
        /// 
        /// ## Inserting variables in a template
        /// 
        /// You can also add variables to a template by using a string like `“${env.&lt;variableName&gt;}"`. You'll need the quotes here too.  
        /// These variables follow the format used in the [Property Manager CLI](https://github.com/akamai/cli-property-manager#update-the-variabledefinitions-file).  They differ from the provider variables which should resolve normally.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% /examples %}}
        /// ## Argument reference
        /// 
        /// * `template_file` - (Required) The absolute path to your top-level JSON template file. The top-level template combines smaller, nested JSON templates to form your property rule tree.
        /// * `variables` - (Optional) A definition of a variable. Variables aren't required and you can use multiple ones if needed. This argument conflicts with the `variable_definition_file` and `variable_values_file` arguments. A `variables` block includes:
        ///     * `name` - The name of the variable used in template.
        ///     * `type` - The type of variable: `string`, `number`, `bool`, or `jsonBlock`.
        ///     * `value` - The value of the variable passed as a string.
        /// * `variable_definition_file` - (Optional) The absolute path to the file containing variable definitions and defaults. This file follows the syntax used in the [Property Manager CLI](https://github.com/akamai/cli-property-manager). This argument is required if you set `variable_values_file` and conflicts with `variables`.
        /// * `variable_values_file` - (Optional) The absolute path to the file containing variable values. This file follows the syntax used in the Property Manager CLI. This argument is required if you set `variable_definition_file` and conflicts with `variables`.
        /// 
        /// ## Attributes reference
        /// 
        /// This data source returns this attribute:
        /// 
        /// * `json` - The fully expanded template with variables and all nested templates resolved.
        /// </summary>
        public static Task<GetPropertyRulesTemplateResult> InvokeAsync(GetPropertyRulesTemplateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPropertyRulesTemplateResult>("akamai:index/getPropertyRulesTemplate:getPropertyRulesTemplate", args ?? new GetPropertyRulesTemplateArgs(), options.WithVersion());
    }


    public sealed class GetPropertyRulesTemplateArgs : Pulumi.InvokeArgs
    {
        [Input("templateFile", required: true)]
        public string TemplateFile { get; set; } = null!;

        [Input("varDefinitionFile")]
        public string? VarDefinitionFile { get; set; }

        [Input("varValuesFile")]
        public string? VarValuesFile { get; set; }

        [Input("variables")]
        private List<Inputs.GetPropertyRulesTemplateVariableArgs>? _variables;
        public List<Inputs.GetPropertyRulesTemplateVariableArgs> Variables
        {
            get => _variables ?? (_variables = new List<Inputs.GetPropertyRulesTemplateVariableArgs>());
            set => _variables = value;
        }

        public GetPropertyRulesTemplateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPropertyRulesTemplateResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly string TemplateFile;
        public readonly string? VarDefinitionFile;
        public readonly string? VarValuesFile;
        public readonly ImmutableArray<Outputs.GetPropertyRulesTemplateVariableResult> Variables;

        [OutputConstructor]
        private GetPropertyRulesTemplateResult(
            string id,

            string json,

            string templateFile,

            string? varDefinitionFile,

            string? varValuesFile,

            ImmutableArray<Outputs.GetPropertyRulesTemplateVariableResult> variables)
        {
            Id = id;
            Json = json;
            TemplateFile = templateFile;
            VarDefinitionFile = varDefinitionFile;
            VarValuesFile = varValuesFile;
            Variables = variables;
        }
    }
}
