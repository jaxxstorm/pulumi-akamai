// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecMatchTargets
    {
        /// <summary>
        /// Use the `akamai.getAppSecMatchTargets` data source to retrieve information about the match targets associated with a given configuration version.
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
        ///         var matchTargetsAppSecMatchTargets = Output.Tuple(configuration, configuration).Apply(values =&gt;
        ///         {
        ///             var configuration = values.Item1;
        ///             var configuration1 = values.Item2;
        ///             return Output.Create(Akamai.GetAppSecMatchTargets.InvokeAsync(new Akamai.GetAppSecMatchTargetsArgs
        ///             {
        ///                 ConfigId = configuration.ConfigId,
        ///                 Version = configuration1.LatestVersion,
        ///             }));
        ///         });
        ///         this.MatchTargets = matchTargetsAppSecMatchTargets.Apply(matchTargetsAppSecMatchTargets =&gt; matchTargetsAppSecMatchTargets.OutputText);
        ///     }
        /// 
        ///     [Output("matchTargets")]
        ///     public Output&lt;string&gt; MatchTargets { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecMatchTargetsResult> InvokeAsync(GetAppSecMatchTargetsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecMatchTargetsResult>("akamai:index/getAppSecMatchTargets:getAppSecMatchTargets", args ?? new GetAppSecMatchTargetsArgs(), options.WithVersion());
    }


    public sealed class GetAppSecMatchTargetsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Input("version", required: true)]
        public int Version { get; set; }

        public GetAppSecMatchTargetsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecMatchTargetsResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A tabular display showing the ID and Policy ID of all match targets associated with the specified security configuration and version.
        /// </summary>
        public readonly string OutputText;
        public readonly int Version;

        [OutputConstructor]
        private GetAppSecMatchTargetsResult(
            int configId,

            string id,

            string outputText,

            int version)
        {
            ConfigId = configId;
            Id = id;
            OutputText = outputText;
            Version = version;
        }
    }
}
