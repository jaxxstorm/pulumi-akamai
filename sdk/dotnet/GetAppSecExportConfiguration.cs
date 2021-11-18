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
    public static class GetAppSecExportConfiguration
    {
        /// <summary>
        /// Use the `akamai.getAppSecExportConfiguration` data source to retrieve comprehensive details about a security configuration version, including rate and security policies, rules, hostnames, and other settings. You can retrieve the entire set of information in JSON format, or a subset of the information in tabular format.
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
        ///             Name = "Akamai Tools",
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
        /// </summary>
        public static Task<GetAppSecExportConfigurationResult> InvokeAsync(GetAppSecExportConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecExportConfigurationResult>("akamai:index/getAppSecExportConfiguration:getAppSecExportConfiguration", args ?? new GetAppSecExportConfigurationArgs(), options.WithVersion());

        /// <summary>
        /// Use the `akamai.getAppSecExportConfiguration` data source to retrieve comprehensive details about a security configuration version, including rate and security policies, rules, hostnames, and other settings. You can retrieve the entire set of information in JSON format, or a subset of the information in tabular format.
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
        ///             Name = "Akamai Tools",
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
        /// </summary>
        public static Output<GetAppSecExportConfigurationResult> Invoke(GetAppSecExportConfigurationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecExportConfigurationResult>("akamai:index/getAppSecExportConfiguration:getAppSecExportConfiguration", args ?? new GetAppSecExportConfigurationInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAppSecExportConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        [Input("searches")]
        private List<string>? _searches;

        /// <summary>
        /// A bracket-delimited list of quoted strings specifying the types of information to be retrieved and made available for display in the `output_text` format. The following types are available:
        /// * customRules
        /// * matchTargets
        /// * ratePolicies
        /// * reputationProfiles
        /// * rulesets
        /// * securityPolicies
        /// * selectableHosts
        /// * selectedHosts
        /// </summary>
        public List<string> Searches
        {
            get => _searches ?? (_searches = new List<string>());
            set => _searches = value;
        }

        /// <summary>
        /// The version number of the security configuration to use.
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
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        [Input("searches")]
        private InputList<string>? _searches;

        /// <summary>
        /// A bracket-delimited list of quoted strings specifying the types of information to be retrieved and made available for display in the `output_text` format. The following types are available:
        /// * customRules
        /// * matchTargets
        /// * ratePolicies
        /// * reputationProfiles
        /// * rulesets
        /// * securityPolicies
        /// * selectableHosts
        /// * selectedHosts
        /// </summary>
        public InputList<string> Searches
        {
            get => _searches ?? (_searches = new InputList<string>());
            set => _searches = value;
        }

        /// <summary>
        /// The version number of the security configuration to use.
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
        /// <summary>
        /// The complete set of information about the specified security configuration version, in JSON format. This includes the types available for the `search` parameter, plus several additional fields such as createDate and createdBy.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// A tabular display showing the types of data specified in the `search` parameter. Included only if the `search` parameter specifies at least one type.
        /// </summary>
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
