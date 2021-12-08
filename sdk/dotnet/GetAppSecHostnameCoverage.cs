// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecHostnameCoverage
    {
        /// <summary>
        /// **Scopes**: Individual account
        /// 
        /// Returns information about the hostnames associated with your account; the returned data includes the hostname's protections, activation status, and other summary information. This information is described in the [HostnameCoverage members](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getfailoverhostnames) section of the Application Security API.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/hostname-coverage](https://developer.akamai.com/api/cloud_security/application_security/v1.html#gethostnamecoverage)
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
        ///         var hostnameCoverage = Output.Create(Akamai.GetAppSecHostnameCoverage.InvokeAsync());
        ///         this.HostnameCoverageListJson = hostnameCoverage.Apply(hostnameCoverage =&gt; hostnameCoverage.Json);
        ///         this.HostnameCoverageListOutput = hostnameCoverage.Apply(hostnameCoverage =&gt; hostnameCoverage.OutputText);
        ///     }
        /// 
        ///     [Output("hostnameCoverageListJson")]
        ///     public Output&lt;string&gt; HostnameCoverageListJson { get; set; }
        ///     [Output("hostnameCoverageListOutput")]
        ///     public Output&lt;string&gt; HostnameCoverageListOutput { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `json`. JSON-formatted list of the hostname coverage information.
        /// - `output_text`. Tabular report of the hostname coverage information.
        /// </summary>
        public static Task<GetAppSecHostnameCoverageResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecHostnameCoverageResult>("akamai:index/getAppSecHostnameCoverage:getAppSecHostnameCoverage", InvokeArgs.Empty, options.WithVersion());
    }


    [OutputType]
    public sealed class GetAppSecHostnameCoverageResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly string OutputText;

        [OutputConstructor]
        private GetAppSecHostnameCoverageResult(
            string id,

            string json,

            string outputText)
        {
            Id = id;
            Json = json;
            OutputText = outputText;
        }
    }
}
