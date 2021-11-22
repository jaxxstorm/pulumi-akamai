// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Akamai
{
    public static class GetAppSecCustomRules
    {
        /// <summary>
        /// Use the `akamai.getAppSecCustomRules` data source to retrieve a list of the custom rules defined for a security configuration.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic usage:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var configuration = Output.Create(Akamai.GetAppSecConfiguration.InvokeAsync(new Akamai.GetAppSecConfigurationArgs
        ///         {
        ///             Name = @var.Security_configuration,
        ///         }));
        ///         var customRules = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecCustomRules.InvokeAsync(new Akamai.GetAppSecCustomRulesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.CustomRulesOutputText = customRules.Apply(customRules =&gt; customRules.OutputText);
        ///         this.CustomRulesJson = customRules.Apply(customRules =&gt; customRules.Json);
        ///         this.CustomRulesConfigId = customRules.Apply(customRules =&gt; customRules.ConfigId);
        ///         var specificCustomRule = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecCustomRules.InvokeAsync(new Akamai.GetAppSecCustomRulesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             CustomRuleId = @var.Custom_rule_id,
        ///         })));
        ///         this.SpecificCustomRuleJson = specificCustomRule.Apply(specificCustomRule =&gt; specificCustomRule.Json);
        ///     }
        /// 
        ///     [Output("customRulesOutputText")]
        ///     public Output&lt;string&gt; CustomRulesOutputText { get; set; }
        ///     [Output("customRulesJson")]
        ///     public Output&lt;string&gt; CustomRulesJson { get; set; }
        ///     [Output("customRulesConfigId")]
        ///     public Output&lt;string&gt; CustomRulesConfigId { get; set; }
        ///     [Output("specificCustomRuleJson")]
        ///     public Output&lt;string&gt; SpecificCustomRuleJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecCustomRulesResult> InvokeAsync(GetAppSecCustomRulesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecCustomRulesResult>("akamai:index/getAppSecCustomRules:getAppSecCustomRules", args ?? new GetAppSecCustomRulesArgs(), options.WithVersion());

        /// <summary>
        /// Use the `akamai.getAppSecCustomRules` data source to retrieve a list of the custom rules defined for a security configuration.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic usage:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var configuration = Output.Create(Akamai.GetAppSecConfiguration.InvokeAsync(new Akamai.GetAppSecConfigurationArgs
        ///         {
        ///             Name = @var.Security_configuration,
        ///         }));
        ///         var customRules = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecCustomRules.InvokeAsync(new Akamai.GetAppSecCustomRulesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.CustomRulesOutputText = customRules.Apply(customRules =&gt; customRules.OutputText);
        ///         this.CustomRulesJson = customRules.Apply(customRules =&gt; customRules.Json);
        ///         this.CustomRulesConfigId = customRules.Apply(customRules =&gt; customRules.ConfigId);
        ///         var specificCustomRule = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecCustomRules.InvokeAsync(new Akamai.GetAppSecCustomRulesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             CustomRuleId = @var.Custom_rule_id,
        ///         })));
        ///         this.SpecificCustomRuleJson = specificCustomRule.Apply(specificCustomRule =&gt; specificCustomRule.Json);
        ///     }
        /// 
        ///     [Output("customRulesOutputText")]
        ///     public Output&lt;string&gt; CustomRulesOutputText { get; set; }
        ///     [Output("customRulesJson")]
        ///     public Output&lt;string&gt; CustomRulesJson { get; set; }
        ///     [Output("customRulesConfigId")]
        ///     public Output&lt;string&gt; CustomRulesConfigId { get; set; }
        ///     [Output("specificCustomRuleJson")]
        ///     public Output&lt;string&gt; SpecificCustomRuleJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAppSecCustomRulesResult> Invoke(GetAppSecCustomRulesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecCustomRulesResult>("akamai:index/getAppSecCustomRules:getAppSecCustomRules", args ?? new GetAppSecCustomRulesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAppSecCustomRulesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The ID of a specific custom rule to use. If not supplied, information about all custom rules associated with the given security configuration will be returned.
        /// </summary>
        [Input("customRuleId")]
        public int? CustomRuleId { get; set; }

        public GetAppSecCustomRulesArgs()
        {
        }
    }

    public sealed class GetAppSecCustomRulesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// The ID of a specific custom rule to use. If not supplied, information about all custom rules associated with the given security configuration will be returned.
        /// </summary>
        [Input("customRuleId")]
        public Input<int>? CustomRuleId { get; set; }

        public GetAppSecCustomRulesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecCustomRulesResult
    {
        public readonly int ConfigId;
        public readonly int? CustomRuleId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A JSON-formatted display of the custom rule information.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// A tabular display showing the ID and name of the custom rule(s).
        /// </summary>
        public readonly string OutputText;

        [OutputConstructor]
        private GetAppSecCustomRulesResult(
            int configId,

            int? customRuleId,

            string id,

            string json,

            string outputText)
        {
            ConfigId = configId;
            CustomRuleId = customRuleId;
            Id = id;
            Json = json;
            OutputText = outputText;
        }
    }
}
