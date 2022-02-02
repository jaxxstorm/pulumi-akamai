// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecExportConfiguration
    {
        /// <summary>
        /// **Scopes**: Security configuration and version
        /// 
        /// Returns comprehensive details about a security configuration, including rate policies, security policies, rules, hostnames, and match targets.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/export/configs/{configId}/versions/{versionNumber}](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getconfigurationversionexport)
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
        ///             Name = "Documentation",
        ///         }));
        ///         var export = Output.Tuple(configuration, configuration).Apply(values =&gt;
        ///         {
        ///             var configuration = values.Item1;
        ///             var configuration1 = values.Item2;
        ///             return Output.Create(Akamai.GetAppSecExportConfiguration.InvokeAsync(new Akamai.GetAppSecExportConfigurationArgs
        ///             {
        ///                 ConfigId = configuration.ConfigId,
        ///                 Version = configuration1.LatestVersion,
        ///                 Searches = 
        ///                 {
        ///                     "securityPolicies",
        ///                     "selectedHosts",
        ///                 },
        ///             }));
        ///         });
        ///         this.Json = export.Apply(export =&gt; export.Json);
        ///         this.Text = export.Apply(export =&gt; export.OutputText);
        ///     }
        /// 
        ///     [Output("json")]
        ///     public Output&lt;string&gt; Json { get; set; }
        ///     [Output("text")]
        ///     public Output&lt;string&gt; Text { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `json`. Complete set of information about the specified security configuration version in JSON format. Includes the types available for the `search` parameter as well as additional fields such as `createDate` and `createdBy`.
        /// - `output_text`. Tabular report showing the types of data specified in the `search` parameter. Valid only if the `search` parameter references at least one type.
        /// </summary>
        public static Task<GetAppSecExportConfigurationResult> InvokeAsync(GetAppSecExportConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecExportConfigurationResult>("akamai:index/getAppSecExportConfiguration:getAppSecExportConfiguration", args ?? new GetAppSecExportConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// **Scopes**: Security configuration and version
        /// 
        /// Returns comprehensive details about a security configuration, including rate policies, security policies, rules, hostnames, and match targets.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/export/configs/{configId}/versions/{versionNumber}](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getconfigurationversionexport)
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
        ///             Name = "Documentation",
        ///         }));
        ///         var export = Output.Tuple(configuration, configuration).Apply(values =&gt;
        ///         {
        ///             var configuration = values.Item1;
        ///             var configuration1 = values.Item2;
        ///             return Output.Create(Akamai.GetAppSecExportConfiguration.InvokeAsync(new Akamai.GetAppSecExportConfigurationArgs
        ///             {
        ///                 ConfigId = configuration.ConfigId,
        ///                 Version = configuration1.LatestVersion,
        ///                 Searches = 
        ///                 {
        ///                     "securityPolicies",
        ///                     "selectedHosts",
        ///                 },
        ///             }));
        ///         });
        ///         this.Json = export.Apply(export =&gt; export.Json);
        ///         this.Text = export.Apply(export =&gt; export.OutputText);
        ///     }
        /// 
        ///     [Output("json")]
        ///     public Output&lt;string&gt; Json { get; set; }
        ///     [Output("text")]
        ///     public Output&lt;string&gt; Text { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `json`. Complete set of information about the specified security configuration version in JSON format. Includes the types available for the `search` parameter as well as additional fields such as `createDate` and `createdBy`.
        /// - `output_text`. Tabular report showing the types of data specified in the `search` parameter. Valid only if the `search` parameter references at least one type.
        /// </summary>
        public static Output<GetAppSecExportConfigurationResult> Invoke(GetAppSecExportConfigurationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecExportConfigurationResult>("akamai:index/getAppSecExportConfiguration:getAppSecExportConfiguration", args ?? new GetAppSecExportConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecExportConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration you want to return information for.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        [Input("searches")]
        private List<string>? _searches;

        /// <summary>
        /// . JSON array of strings specifying the types of information to be retrieved. Allowed values include:
        /// &gt; - **AdvancedSettingsLogging**
        /// &gt; - **AdvancedSettingsPrefetch**
        /// &gt; - **ApiRequestConstraints**
        /// &gt; - **AttackGroup**
        /// &gt; - **AttackGroupConditionException**
        /// &gt; - **Eval**
        /// &gt; - **EvalRuleConditionException**
        /// &gt; - **CustomDeny**
        /// &gt; - **CustomRule**
        /// &gt; - **CustomRuleAction**
        /// &gt; - **IPGeoFirewall**
        /// &gt; - **MatchTarget**
        /// &gt; - **PenaltyBox**
        /// &gt; - **RatePolicy**
        /// &gt; - **RatePolicyAction**
        /// &gt; - **ReputationProfile**
        /// &gt; - **ReputationProfileAction**
        /// &gt; - **Rule**
        /// &gt; - **RuleConditionException**
        /// &gt; - **SecurityPolicy**
        /// &gt; - **SiemSettings**
        /// &gt; - **SlowPost**
        /// </summary>
        public List<string> Searches
        {
            get => _searches ?? (_searches = new List<string>());
            set => _searches = value;
        }

        /// <summary>
        /// . Version number of the security configuration.
        /// </summary>
        [Input("version", required: true)]
        public int Version { get; set; }

        public GetAppSecExportConfigurationArgs()
        {
        }
    }

    public sealed class GetAppSecExportConfigurationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration you want to return information for.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        [Input("searches")]
        private InputList<string>? _searches;

        /// <summary>
        /// . JSON array of strings specifying the types of information to be retrieved. Allowed values include:
        /// &gt; - **AdvancedSettingsLogging**
        /// &gt; - **AdvancedSettingsPrefetch**
        /// &gt; - **ApiRequestConstraints**
        /// &gt; - **AttackGroup**
        /// &gt; - **AttackGroupConditionException**
        /// &gt; - **Eval**
        /// &gt; - **EvalRuleConditionException**
        /// &gt; - **CustomDeny**
        /// &gt; - **CustomRule**
        /// &gt; - **CustomRuleAction**
        /// &gt; - **IPGeoFirewall**
        /// &gt; - **MatchTarget**
        /// &gt; - **PenaltyBox**
        /// &gt; - **RatePolicy**
        /// &gt; - **RatePolicyAction**
        /// &gt; - **ReputationProfile**
        /// &gt; - **ReputationProfileAction**
        /// &gt; - **Rule**
        /// &gt; - **RuleConditionException**
        /// &gt; - **SecurityPolicy**
        /// &gt; - **SiemSettings**
        /// &gt; - **SlowPost**
        /// </summary>
        public InputList<string> Searches
        {
            get => _searches ?? (_searches = new InputList<string>());
            set => _searches = value;
        }

        /// <summary>
        /// . Version number of the security configuration.
        /// </summary>
        [Input("version", required: true)]
        public Input<int> Version { get; set; } = null!;

        public GetAppSecExportConfigurationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecExportConfigurationResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly string OutputText;
        public readonly ImmutableArray<string> Searches;
        public readonly int Version;

        [OutputConstructor]
        private GetAppSecExportConfigurationResult(
            int configId,

            string id,

            string json,

            string outputText,

            ImmutableArray<string> searches,

            int version)
        {
            ConfigId = configId;
            Id = id;
            Json = json;
            OutputText = outputText;
            Searches = searches;
            Version = version;
        }
    }
}
