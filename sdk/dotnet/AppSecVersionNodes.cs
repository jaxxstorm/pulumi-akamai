// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    /// <summary>
    /// **Scopes**: Security configuration
    /// 
    /// Updates the version notes for a security configuration.
    /// 
    /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/version-notes](https://techdocs.akamai.com/application-security/reference/put-version-notes)
    /// 
    /// ## Example Usage
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
    ///         var versionNotesAppSecVersionNodes = new Akamai.AppSecVersionNodes("versionNotesAppSecVersionNodes", new Akamai.AppSecVersionNodesArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             VersionNotes = "This version enables reputation profiles.",
    ///         });
    ///         this.VersionNotes = versionNotesAppSecVersionNodes.OutputText;
    ///     }
    /// 
    ///     [Output("versionNotes")]
    ///     public Output&lt;string&gt; VersionNotes { get; set; }
    /// }
    /// ```
    /// ## Output Options
    /// 
    /// The following options can be used to determine the information returned, and how that returned information is formatted:
    /// 
    /// - `output_text`. Tabular report showing the updated version notes.
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecVersionNodes:AppSecVersionNodes")]
    public partial class AppSecVersionNodes : Pulumi.CustomResource
    {
        /// <summary>
        /// . Unique identifier of the security configuration whose version notes are being modified.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// Text representation
        /// </summary>
        [Output("outputText")]
        public Output<string> OutputText { get; private set; } = null!;

        /// <summary>
        /// . Brief description of the security configuration version.
        /// </summary>
        [Output("versionNotes")]
        public Output<string> VersionNotes { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecVersionNodes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecVersionNodes(string name, AppSecVersionNodesArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecVersionNodes:AppSecVersionNodes", name, args ?? new AppSecVersionNodesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecVersionNodes(string name, Input<string> id, AppSecVersionNodesState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecVersionNodes:AppSecVersionNodes", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AppSecVersionNodes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecVersionNodes Get(string name, Input<string> id, AppSecVersionNodesState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecVersionNodes(name, id, state, options);
        }
    }

    public sealed class AppSecVersionNodesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration whose version notes are being modified.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Brief description of the security configuration version.
        /// </summary>
        [Input("versionNotes", required: true)]
        public Input<string> VersionNotes { get; set; } = null!;

        public AppSecVersionNodesArgs()
        {
        }
    }

    public sealed class AppSecVersionNodesState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration whose version notes are being modified.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// Text representation
        /// </summary>
        [Input("outputText")]
        public Input<string>? OutputText { get; set; }

        /// <summary>
        /// . Brief description of the security configuration version.
        /// </summary>
        [Input("versionNotes")]
        public Input<string>? VersionNotes { get; set; }

        public AppSecVersionNodesState()
        {
        }
    }
}
