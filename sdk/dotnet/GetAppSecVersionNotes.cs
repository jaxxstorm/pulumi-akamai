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
    public static class GetAppSecVersionNotes
    {
        /// <summary>
        /// Use the `akamai.AppSecVersionNodes` data source to retrieve the most recent version notes for a configuration.
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
        ///         var versionNotes = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecVersionNotes.InvokeAsync(new Akamai.GetAppSecVersionNotesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.VersionNotesText = versionNotes.Apply(versionNotes =&gt; versionNotes.OutputText);
        ///         this.VersionNotesJson = versionNotes.Apply(versionNotes =&gt; versionNotes.Json);
        ///     }
        /// 
        ///     [Output("versionNotesText")]
        ///     public Output&lt;string&gt; VersionNotesText { get; set; }
        ///     [Output("versionNotesJson")]
        ///     public Output&lt;string&gt; VersionNotesJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecVersionNotesResult> InvokeAsync(GetAppSecVersionNotesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecVersionNotesResult>("akamai:index/getAppSecVersionNotes:getAppSecVersionNotes", args ?? new GetAppSecVersionNotesArgs(), options.WithVersion());

        /// <summary>
        /// Use the `akamai.AppSecVersionNodes` data source to retrieve the most recent version notes for a configuration.
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
        ///         var versionNotes = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecVersionNotes.InvokeAsync(new Akamai.GetAppSecVersionNotesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.VersionNotesText = versionNotes.Apply(versionNotes =&gt; versionNotes.OutputText);
        ///         this.VersionNotesJson = versionNotes.Apply(versionNotes =&gt; versionNotes.Json);
        ///     }
        /// 
        ///     [Output("versionNotesText")]
        ///     public Output&lt;string&gt; VersionNotesText { get; set; }
        ///     [Output("versionNotesJson")]
        ///     public Output&lt;string&gt; VersionNotesJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAppSecVersionNotesResult> Invoke(GetAppSecVersionNotesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecVersionNotesResult>("akamai:index/getAppSecVersionNotes:getAppSecVersionNotes", args ?? new GetAppSecVersionNotesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAppSecVersionNotesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The configuration ID to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        public GetAppSecVersionNotesArgs()
        {
        }
    }

    public sealed class GetAppSecVersionNotesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The configuration ID to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        public GetAppSecVersionNotesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecVersionNotesResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A JSON-formatted list showing the version notes.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// A tabular display showing the version notes.
        /// </summary>
        public readonly string OutputText;

        [OutputConstructor]
        private GetAppSecVersionNotesResult(
            int configId,

            string id,

            string json,

            string outputText)
        {
            ConfigId = configId;
            Id = id;
            Json = json;
            OutputText = outputText;
        }
    }
}
