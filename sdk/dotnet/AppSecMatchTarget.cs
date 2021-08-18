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
    /// The `akamai.AppSecMatchTarget` resource allows you to create or modify a match target associated with a given security configuration.
    /// 
    /// ## Example Usage
    /// 
    /// Basic usage:
    /// 
    /// ```csharp
    /// using System.IO;
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
    ///         var matchTarget = new Akamai.AppSecMatchTarget("matchTarget", new Akamai.AppSecMatchTargetArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             MatchTarget = File.ReadAllText($"{path.Module}/match_targets.json"),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecMatchTarget:AppSecMatchTarget")]
    public partial class AppSecMatchTarget : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// The name of a JSON file containing one or more match target definitions ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postmatchtargets)).
        /// </summary>
        [Output("matchTarget")]
        public Output<string> MatchTarget { get; private set; } = null!;

        /// <summary>
        /// The ID of the match target.
        /// </summary>
        [Output("matchTargetId")]
        public Output<int> MatchTargetId { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecMatchTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecMatchTarget(string name, AppSecMatchTargetArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecMatchTarget:AppSecMatchTarget", name, args ?? new AppSecMatchTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecMatchTarget(string name, Input<string> id, AppSecMatchTargetState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecMatchTarget:AppSecMatchTarget", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecMatchTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecMatchTarget Get(string name, Input<string> id, AppSecMatchTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecMatchTarget(name, id, state, options);
        }
    }

    public sealed class AppSecMatchTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// The name of a JSON file containing one or more match target definitions ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postmatchtargets)).
        /// </summary>
        [Input("matchTarget", required: true)]
        public Input<string> MatchTarget { get; set; } = null!;

        public AppSecMatchTargetArgs()
        {
        }
    }

    public sealed class AppSecMatchTargetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// The name of a JSON file containing one or more match target definitions ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postmatchtargets)).
        /// </summary>
        [Input("matchTarget")]
        public Input<string>? MatchTarget { get; set; }

        /// <summary>
        /// The ID of the match target.
        /// </summary>
        [Input("matchTargetId")]
        public Input<int>? MatchTargetId { get; set; }

        public AppSecMatchTargetState()
        {
        }
    }
}
