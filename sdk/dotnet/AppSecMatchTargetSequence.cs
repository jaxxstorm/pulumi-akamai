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
    /// The `akamai.AppSecMatchTargetSequence` resource allows you to specify the order in which match targets are applied within a given security configuration version.
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
    ///         var matchTargetSequence = new Akamai.AppSecMatchTargetSequence("matchTargetSequence", new Akamai.AppSecMatchTargetSequenceArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             Version = configuration.Apply(configuration =&gt; configuration.LatestVersion),
    ///             Type = "website",
    ///             Json = File.ReadAllText($"{path.Module}/match_targets.json"),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class AppSecMatchTargetSequence : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// The name of a JSON file containing the sequence of all match targets defined for the specified security configuration version ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
        /// </summary>
        [Output("json")]
        public Output<string?> Json { get; private set; } = null!;

        [Output("sequenceMap")]
        public Output<ImmutableDictionary<string, string>?> SequenceMap { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecMatchTargetSequence resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecMatchTargetSequence(string name, AppSecMatchTargetSequenceArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecMatchTargetSequence:AppSecMatchTargetSequence", name, args ?? new AppSecMatchTargetSequenceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecMatchTargetSequence(string name, Input<string> id, AppSecMatchTargetSequenceState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecMatchTargetSequence:AppSecMatchTargetSequence", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecMatchTargetSequence resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecMatchTargetSequence Get(string name, Input<string> id, AppSecMatchTargetSequenceState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecMatchTargetSequence(name, id, state, options);
        }
    }

    public sealed class AppSecMatchTargetSequenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// The name of a JSON file containing the sequence of all match targets defined for the specified security configuration version ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
        /// </summary>
        [Input("json")]
        public Input<string>? Json { get; set; }

        [Input("sequenceMap")]
        private InputMap<string>? _sequenceMap;
        public InputMap<string> SequenceMap
        {
            get => _sequenceMap ?? (_sequenceMap = new InputMap<string>());
            set => _sequenceMap = value;
        }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Input("version", required: true)]
        public Input<int> Version { get; set; } = null!;

        public AppSecMatchTargetSequenceArgs()
        {
        }
    }

    public sealed class AppSecMatchTargetSequenceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// The name of a JSON file containing the sequence of all match targets defined for the specified security configuration version ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
        /// </summary>
        [Input("json")]
        public Input<string>? Json { get; set; }

        [Input("sequenceMap")]
        private InputMap<string>? _sequenceMap;
        public InputMap<string> SequenceMap
        {
            get => _sequenceMap ?? (_sequenceMap = new InputMap<string>());
            set => _sequenceMap = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public AppSecMatchTargetSequenceState()
        {
        }
    }
}