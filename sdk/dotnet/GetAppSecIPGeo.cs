// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecIPGeo
    {
        /// <summary>
        /// Use the `akamai.AppSecIPGeo` data source to retrieve information about which network lists are used in the IP/Geo Firewall settings.
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
        ///         var ipGeo = Output.Tuple(configuration, configuration).Apply(values =&gt;
        ///         {
        ///             var configuration = values.Item1;
        ///             var configuration1 = values.Item2;
        ///             return Output.Create(Akamai.GetAppSecIPGeo.InvokeAsync(new Akamai.GetAppSecIPGeoArgs
        ///             {
        ///                 ConfigId = configuration.ConfigId,
        ///                 Version = configuration1.LatestVersion,
        ///                 SecurityPolicyId = @var.Security_policy_id,
        ///             }));
        ///         });
        ///         this.IpGeoMode = ipGeo.Apply(ipGeo =&gt; ipGeo.Mode);
        ///         this.GeoNetworkLists = ipGeo.Apply(ipGeo =&gt; ipGeo.GeoNetworkLists);
        ///         this.IpNetworkLists = ipGeo.Apply(ipGeo =&gt; ipGeo.IpNetworkLists);
        ///         this.ExceptionIpNetworkLists = ipGeo.Apply(ipGeo =&gt; ipGeo.ExceptionIpNetworkLists);
        ///     }
        /// 
        ///     [Output("ipGeoMode")]
        ///     public Output&lt;string&gt; IpGeoMode { get; set; }
        ///     [Output("geoNetworkLists")]
        ///     public Output&lt;string&gt; GeoNetworkLists { get; set; }
        ///     [Output("ipNetworkLists")]
        ///     public Output&lt;string&gt; IpNetworkLists { get; set; }
        ///     [Output("exceptionIpNetworkLists")]
        ///     public Output&lt;string&gt; ExceptionIpNetworkLists { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecIPGeoResult> InvokeAsync(GetAppSecIPGeoArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecIPGeoResult>("akamai:index/getAppSecIPGeo:getAppSecIPGeo", args ?? new GetAppSecIPGeoArgs(), options.WithVersion());
    }


    public sealed class GetAppSecIPGeoArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The ID of the security policy to use.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public string SecurityPolicyId { get; set; } = null!;

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Input("version", required: true)]
        public int Version { get; set; }

        public GetAppSecIPGeoArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecIPGeoResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// The network lists to be allowed regardless of `mode`, `geo_network_lists`, and `ip_network_lists` parameters.
        /// </summary>
        public readonly ImmutableArray<string> ExceptionIpNetworkLists;
        /// <summary>
        /// The network lists to be blocked or allowed geographically.
        /// </summary>
        public readonly ImmutableArray<string> GeoNetworkLists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The network lists to be blocked or allowd by IP address.
        /// </summary>
        public readonly ImmutableArray<string> IpNetworkLists;
        /// <summary>
        /// The mode used for IP/Geo firewall blocking: `block` to block specific IPs, geographies or network lists, or `allow` to allow specific IPs or geographies to be let through while blocking the rest.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// A tabular display of the IP/Geo firewall settings.
        /// </summary>
        public readonly string OutputText;
        public readonly string SecurityPolicyId;
        public readonly int Version;

        [OutputConstructor]
        private GetAppSecIPGeoResult(
            int configId,

            ImmutableArray<string> exceptionIpNetworkLists,

            ImmutableArray<string> geoNetworkLists,

            string id,

            ImmutableArray<string> ipNetworkLists,

            string mode,

            string outputText,

            string securityPolicyId,

            int version)
        {
            ConfigId = configId;
            ExceptionIpNetworkLists = exceptionIpNetworkLists;
            GeoNetworkLists = geoNetworkLists;
            Id = id;
            IpNetworkLists = ipNetworkLists;
            Mode = mode;
            OutputText = outputText;
            SecurityPolicyId = securityPolicyId;
            Version = version;
        }
    }
}
