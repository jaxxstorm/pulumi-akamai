// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecHostnameCoverageOverlapping
    {
        /// <summary>
        /// **Scopes**: Security configuration; hostname
        /// 
        /// Returns information about any other configuration versions that contain a hostname found in the current configuration version. The returned information is described in the [List hostname overlaps](https://developer.akamai.com/api/cloud_security/application_security/v1.html#gethostnamecoverageoverlapping) section of the Application Security API.
        /// 
        /// **Related API Endpoint**:[/appsec/v1/configs/{configId}/versions/{versionNumber}/hostname-coverage/overlapping](https://techdocs.akamai.com/application-security/reference/get-hostname-coverage-overlapping)
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
        ///         var test = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecHostnameCoverageOverlapping.InvokeAsync(new Akamai.GetAppSecHostnameCoverageOverlappingArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             Hostname = "documentation.akamai.com",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `json`. JSON-formatted list of the overlap information.
        /// - `output_text`. Tabular report of the overlap information.
        /// </summary>
        public static Task<GetAppSecHostnameCoverageOverlappingResult> InvokeAsync(GetAppSecHostnameCoverageOverlappingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecHostnameCoverageOverlappingResult>("akamai:index/getAppSecHostnameCoverageOverlapping:getAppSecHostnameCoverageOverlapping", args ?? new GetAppSecHostnameCoverageOverlappingArgs(), options.WithDefaults());

        /// <summary>
        /// **Scopes**: Security configuration; hostname
        /// 
        /// Returns information about any other configuration versions that contain a hostname found in the current configuration version. The returned information is described in the [List hostname overlaps](https://developer.akamai.com/api/cloud_security/application_security/v1.html#gethostnamecoverageoverlapping) section of the Application Security API.
        /// 
        /// **Related API Endpoint**:[/appsec/v1/configs/{configId}/versions/{versionNumber}/hostname-coverage/overlapping](https://techdocs.akamai.com/application-security/reference/get-hostname-coverage-overlapping)
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
        ///         var test = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecHostnameCoverageOverlapping.InvokeAsync(new Akamai.GetAppSecHostnameCoverageOverlappingArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             Hostname = "documentation.akamai.com",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `json`. JSON-formatted list of the overlap information.
        /// - `output_text`. Tabular report of the overlap information.
        /// </summary>
        public static Output<GetAppSecHostnameCoverageOverlappingResult> Invoke(GetAppSecHostnameCoverageOverlappingInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecHostnameCoverageOverlappingResult>("akamai:index/getAppSecHostnameCoverageOverlapping:getAppSecHostnameCoverageOverlapping", args ?? new GetAppSecHostnameCoverageOverlappingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecHostnameCoverageOverlappingArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration you want to return information for.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// . Name of the host you want to return information for.
        /// </summary>
        [Input("hostname", required: true)]
        public string Hostname { get; set; } = null!;

        public GetAppSecHostnameCoverageOverlappingArgs()
        {
        }
    }

    public sealed class GetAppSecHostnameCoverageOverlappingInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration you want to return information for.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Name of the host you want to return information for.
        /// </summary>
        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        public GetAppSecHostnameCoverageOverlappingInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecHostnameCoverageOverlappingResult
    {
        public readonly int ConfigId;
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly string OutputText;

        [OutputConstructor]
        private GetAppSecHostnameCoverageOverlappingResult(
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
