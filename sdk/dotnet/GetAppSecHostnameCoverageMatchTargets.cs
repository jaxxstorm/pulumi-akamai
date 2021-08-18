// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecHostnameCoverageMatchTargets
    {
        /// <summary>
        /// Use the `akamai.getAppSecHostnameCoverageMatchTargets` data source to retrieve information about the API and website match targets that protect a hostname. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#gethostnamecoveragematchtargets).
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
        ///         var matchTargets = Output.Create(Akamai.GetAppSecHostnameCoverageMatchTargets.InvokeAsync(new Akamai.GetAppSecHostnameCoverageMatchTargetsArgs
        ///         {
        ///             ConfigId = 43253,
        ///             Hostname = "example.com",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecHostnameCoverageMatchTargetsResult> InvokeAsync(GetAppSecHostnameCoverageMatchTargetsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecHostnameCoverageMatchTargetsResult>("akamai:index/getAppSecHostnameCoverageMatchTargets:getAppSecHostnameCoverageMatchTargets", args ?? new GetAppSecHostnameCoverageMatchTargetsArgs(), options.WithVersion());
    }


    public sealed class GetAppSecHostnameCoverageMatchTargetsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The configuration ID.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The hostname for which to retrieve information.
        /// </summary>
        [Input("hostname", required: true)]
        public string Hostname { get; set; } = null!;

        public GetAppSecHostnameCoverageMatchTargetsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecHostnameCoverageMatchTargetsResult
    {
        public readonly int ConfigId;
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A JSON-formatted list of the coverage information.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// A tabular display of the coverage information.
        /// </summary>
        public readonly string OutputText;

        [OutputConstructor]
        private GetAppSecHostnameCoverageMatchTargetsResult(
            int configId,

            string hostname,

            string id,

            string json,

            string outputText)
        {
            ConfigId = configId;
            Hostname = hostname;
            Id = id;
            Json = json;
            OutputText = outputText;
        }
    }
}
