// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecConfigurationVersion
    {
        /// <summary>
        /// Use the `akamai.getAppSecConfigurationVersion` data source to retrieve information about the versions of a security configuration, or about a specific version.
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
        ///         var specificConfiguration = Output.Create(Akamai.GetAppSecConfiguration.InvokeAsync(new Akamai.GetAppSecConfigurationArgs
        ///         {
        ///             Name = "Akamai Tools",
        ///         }));
        ///         var versions = specificConfiguration.Apply(specificConfiguration =&gt; Output.Create(Akamai.GetAppSecConfigurationVersion.InvokeAsync(new Akamai.GetAppSecConfigurationVersionArgs
        ///         {
        ///             ConfigId = specificConfiguration.ConfigId,
        ///         })));
        ///         this.VersionsOutputText = versions.Apply(versions =&gt; versions.OutputText);
        ///         this.VersionsLatest = versions.Apply(versions =&gt; versions.LatestVersion);
        ///         var specificVersion = specificConfiguration.Apply(specificConfiguration =&gt; Output.Create(Akamai.GetAppSecConfigurationVersion.InvokeAsync(new Akamai.GetAppSecConfigurationVersionArgs
        ///         {
        ///             ConfigId = specificConfiguration.ConfigId,
        ///             Version = 42,
        ///         })));
        ///         this.SpecificVersionVersion = specificVersion.Apply(specificVersion =&gt; specificVersion.Version);
        ///         this.SpecificVersionStaging = specificVersion.Apply(specificVersion =&gt; specificVersion.StagingStatus);
        ///         this.SpecificVersionProduction = specificVersion.Apply(specificVersion =&gt; specificVersion.ProductionStatus);
        ///     }
        /// 
        ///     [Output("versionsOutputText")]
        ///     public Output&lt;string&gt; VersionsOutputText { get; set; }
        ///     [Output("versionsLatest")]
        ///     public Output&lt;string&gt; VersionsLatest { get; set; }
        ///     [Output("specificVersionVersion")]
        ///     public Output&lt;string&gt; SpecificVersionVersion { get; set; }
        ///     [Output("specificVersionStaging")]
        ///     public Output&lt;string&gt; SpecificVersionStaging { get; set; }
        ///     [Output("specificVersionProduction")]
        ///     public Output&lt;string&gt; SpecificVersionProduction { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecConfigurationVersionResult> InvokeAsync(GetAppSecConfigurationVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecConfigurationVersionResult>("akamai:index/getAppSecConfigurationVersion:getAppSecConfigurationVersion", args ?? new GetAppSecConfigurationVersionArgs(), options.WithVersion());
    }


    public sealed class GetAppSecConfigurationVersionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The version number of the security configuration to use. If not supplied, information about all versions of the specified security configuration is returned.
        /// </summary>
        [Input("version")]
        public int? Version { get; set; }

        public GetAppSecConfigurationVersionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecConfigurationVersionResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The last version of the security configuration created.
        /// </summary>
        public readonly int LatestVersion;
        /// <summary>
        /// A tabular display showing the following information about all versions of the security configuration: version number, staging status, and production status.
        /// </summary>
        public readonly string OutputText;
        /// <summary>
        /// The status of the specified version in production: "Active", "Inactive", or "Deactivated". Returned only if `version` was specified.
        /// </summary>
        public readonly string ProductionStatus;
        /// <summary>
        /// The status of the specified version in staging: "Active", "Inactive", or "Deactivated". Returned only if `version` was specified.
        /// </summary>
        public readonly string StagingStatus;
        public readonly int? Version;

        [OutputConstructor]
        private GetAppSecConfigurationVersionResult(
            int configId,

            string id,

            int latestVersion,

            string outputText,

            string productionStatus,

            string stagingStatus,

            int? version)
        {
            ConfigId = configId;
            Id = id;
            LatestVersion = latestVersion;
            OutputText = outputText;
            ProductionStatus = productionStatus;
            StagingStatus = stagingStatus;
            Version = version;
        }
    }
}
