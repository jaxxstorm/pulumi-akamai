// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecBypassNetworkLists
    {
        /// <summary>
        /// Use the `akamai.AppSecByPassNetworkList` data source to retrieve information about which network
        /// lists are used in the bypass network lists settings.  The information available is described
        /// [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getbypassnetworklistsforawapconfigversion).
        /// Note: this data source is only applicable to WAP (Web Application Protector) configurations.
        /// 
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
        ///         var bypassNetworkLists = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecBypassNetworkLists.InvokeAsync(new Akamai.GetAppSecBypassNetworkListsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.BypassNetworkListsOutput = bypassNetworkLists.Apply(bypassNetworkLists =&gt; bypassNetworkLists.OutputText);
        ///         this.BypassNetworkListsJson = bypassNetworkLists.Apply(bypassNetworkLists =&gt; bypassNetworkLists.Json);
        ///         this.BypassNetworkListsIdList = bypassNetworkLists.Apply(bypassNetworkLists =&gt; bypassNetworkLists.BypassNetworkLists);
        ///     }
        /// 
        ///     [Output("bypassNetworkListsOutput")]
        ///     public Output&lt;string&gt; BypassNetworkListsOutput { get; set; }
        ///     [Output("bypassNetworkListsJson")]
        ///     public Output&lt;string&gt; BypassNetworkListsJson { get; set; }
        ///     [Output("bypassNetworkListsIdList")]
        ///     public Output&lt;string&gt; BypassNetworkListsIdList { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecBypassNetworkListsResult> InvokeAsync(GetAppSecBypassNetworkListsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecBypassNetworkListsResult>("akamai:index/getAppSecBypassNetworkLists:getAppSecBypassNetworkLists", args ?? new GetAppSecBypassNetworkListsArgs(), options.WithVersion());
    }


    public sealed class GetAppSecBypassNetworkListsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The configuration ID to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        public GetAppSecBypassNetworkListsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecBypassNetworkListsResult
    {
        /// <summary>
        /// A list of strings containing the network list IDs.
        /// </summary>
        public readonly ImmutableArray<string> BypassNetworkLists;
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A JSON-formatted list of information about the bypass network lists.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// A tabular display showing the bypass network list information.
        /// </summary>
        public readonly string OutputText;

        [OutputConstructor]
        private GetAppSecBypassNetworkListsResult(
            ImmutableArray<string> bypassNetworkLists,

            int configId,

            string id,

            string json,

            string outputText)
        {
            BypassNetworkLists = bypassNetworkLists;
            ConfigId = configId;
            Id = id;
            Json = json;
            OutputText = outputText;
        }
    }
}
