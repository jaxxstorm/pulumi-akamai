// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Properties
{
    [Obsolete(@"akamai.properties.getPropertyRules has been deprecated in favor of akamai.getPropertyRules")]
    public static class GetPropertyRules
    {
        /// <summary>
        /// Use the `akamai.getPropertyRules` data source to query and retrieve the rule tree of
        /// an existing property version. This data source lets you search across the contracts
        /// and groups you have access to.
        /// 
        /// ## Basic usage
        /// 
        /// This example returns the rule tree for version 3 of a property based on the selected contract and group:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var my_example = Output.Create(Akamai.GetPropertyRules.InvokeAsync(new Akamai.GetPropertyRulesArgs
        ///         {
        ///             PropertyId = "prp_123",
        ///             GroupId = "grp_12345",
        ///             ContractId = "ctr_1-AB123",
        ///             Version = 3,
        ///         }));
        ///         this.PropertyMatch = my_example;
        ///     }
        /// 
        ///     [Output("propertyMatch")]
        ///     public Output&lt;string&gt; PropertyMatch { get; set; }
        /// }
        /// ```
        /// 
        /// ## Attributes reference
        /// 
        /// This data source returns these attributes:
        /// 
        /// * `rule_format` - The rule tree version used. Property rule objects are versioned infrequently, and are known as rule formats. See [Rule format schemas](https://techdocs.akamai.com/property-mgr/reference/rule-format-schemas) to learn more.
        /// * `rules` - A JSON-encoded rule tree for the property.
        /// * `errors` - A list of validation errors for the rule tree object returned. For more information see [Errors](https://techdocs.akamai.com/property-mgr/reference/api-errors) in the Property Manager API documentation.
        /// </summary>
        public static Task<GetPropertyRulesResult> InvokeAsync(GetPropertyRulesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPropertyRulesResult>("akamai:properties/getPropertyRules:getPropertyRules", args ?? new GetPropertyRulesArgs(), options.WithDefaults());

        /// <summary>
        /// Use the `akamai.getPropertyRules` data source to query and retrieve the rule tree of
        /// an existing property version. This data source lets you search across the contracts
        /// and groups you have access to.
        /// 
        /// ## Basic usage
        /// 
        /// This example returns the rule tree for version 3 of a property based on the selected contract and group:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var my_example = Output.Create(Akamai.GetPropertyRules.InvokeAsync(new Akamai.GetPropertyRulesArgs
        ///         {
        ///             PropertyId = "prp_123",
        ///             GroupId = "grp_12345",
        ///             ContractId = "ctr_1-AB123",
        ///             Version = 3,
        ///         }));
        ///         this.PropertyMatch = my_example;
        ///     }
        /// 
        ///     [Output("propertyMatch")]
        ///     public Output&lt;string&gt; PropertyMatch { get; set; }
        /// }
        /// ```
        /// 
        /// ## Attributes reference
        /// 
        /// This data source returns these attributes:
        /// 
        /// * `rule_format` - The rule tree version used. Property rule objects are versioned infrequently, and are known as rule formats. See [Rule format schemas](https://techdocs.akamai.com/property-mgr/reference/rule-format-schemas) to learn more.
        /// * `rules` - A JSON-encoded rule tree for the property.
        /// * `errors` - A list of validation errors for the rule tree object returned. For more information see [Errors](https://techdocs.akamai.com/property-mgr/reference/api-errors) in the Property Manager API documentation.
        /// </summary>
        public static Output<GetPropertyRulesResult> Invoke(GetPropertyRulesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPropertyRulesResult>("akamai:properties/getPropertyRules:getPropertyRules", args ?? new GetPropertyRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPropertyRulesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// - (Required) A contract's unique ID, including the `ctr_` prefix.
        /// </summary>
        [Input("contractId")]
        public string? ContractId { get; set; }

        /// <summary>
        /// - (Required) A group's unique ID, including the `grp_` prefix.
        /// </summary>
        [Input("groupId")]
        public string? GroupId { get; set; }

        /// <summary>
        /// - (Required) A property's unique ID, including the `prp_` prefix.
        /// </summary>
        [Input("propertyId", required: true)]
        public string PropertyId { get; set; } = null!;

        [Input("ruleFormat")]
        public string? RuleFormat { get; set; }

        /// <summary>
        /// - (Optional) The version to return. Returns the latest version by default.
        /// </summary>
        [Input("version")]
        public int? Version { get; set; }

        public GetPropertyRulesArgs()
        {
        }
    }

    public sealed class GetPropertyRulesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// - (Required) A contract's unique ID, including the `ctr_` prefix.
        /// </summary>
        [Input("contractId")]
        public Input<string>? ContractId { get; set; }

        /// <summary>
        /// - (Required) A group's unique ID, including the `grp_` prefix.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// - (Required) A property's unique ID, including the `prp_` prefix.
        /// </summary>
        [Input("propertyId", required: true)]
        public Input<string> PropertyId { get; set; } = null!;

        [Input("ruleFormat")]
        public Input<string>? RuleFormat { get; set; }

        /// <summary>
        /// - (Optional) The version to return. Returns the latest version by default.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public GetPropertyRulesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPropertyRulesResult
    {
        public readonly string ContractId;
        public readonly string Errors;
        public readonly string GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string PropertyId;
        public readonly string? RuleFormat;
        public readonly string Rules;
        public readonly int Version;

        [OutputConstructor]
        private GetPropertyRulesResult(
            string contractId,

            string errors,

            string groupId,

            string id,

            string propertyId,

            string? ruleFormat,

            string rules,

            int version)
        {
            ContractId = contractId;
            Errors = errors;
            GroupId = groupId;
            Id = id;
            PropertyId = propertyId;
            RuleFormat = ruleFormat;
            Rules = rules;
            Version = version;
        }
    }
}
