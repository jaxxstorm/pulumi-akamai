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
    /// Use the `akamai.AppSecSecurityPolicyProtections` resource to create or modify ...
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
    ///             Name = @var.Security_configuration,
    ///         }));
    ///         var protections = new Akamai.AppSecSecurityPolicyProtections("protections", new Akamai.AppSecSecurityPolicyProtectionsArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             Version = configuration.Apply(configuration =&gt; configuration.LatestVersion),
    ///             SecurityPolicyId = @var.Security_policy_id,
    ///             ApplyApplicationLayerControls = @var.Apply_application_layer_controls,
    ///             ApplyNetworkLayerControls = @var.Apply_network_layer_controls,
    ///             ApplyRateControls = @var.Apply_rate_controls,
    ///             ApplyReputationControls = @var.Apply_reputation_controls,
    ///             ApplyBotmanControls = @var.Apply_botman_controls,
    ///             ApplyApiConstraints = @var.Apply_api_constraints,
    ///             ApplySlowPostControls = @var.Apply_slow_post_controls,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecSecurityPolicyProtections:AppSecSecurityPolicyProtections")]
    public partial class AppSecSecurityPolicyProtections : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to enable api constraints: either `true` or `false`.
        /// </summary>
        [Output("applyApiConstraints")]
        public Output<bool?> ApplyApiConstraints { get; private set; } = null!;

        /// <summary>
        /// Whether to enable application layer controls: either `true` or `false`.
        /// </summary>
        [Output("applyApplicationLayerControls")]
        public Output<bool?> ApplyApplicationLayerControls { get; private set; } = null!;

        /// <summary>
        /// Whether to enable botman controls: either `true` or `false`.
        /// </summary>
        [Output("applyBotmanControls")]
        public Output<bool?> ApplyBotmanControls { get; private set; } = null!;

        /// <summary>
        /// Whether to enable network layer controls: either `true` or `false`.
        /// </summary>
        [Output("applyNetworkLayerControls")]
        public Output<bool?> ApplyNetworkLayerControls { get; private set; } = null!;

        /// <summary>
        /// Whether to enable rate controls: either `true` or `false`.
        /// </summary>
        [Output("applyRateControls")]
        public Output<bool?> ApplyRateControls { get; private set; } = null!;

        /// <summary>
        /// Whether to enable reputation controls: either `true` or `false`.
        /// </summary>
        [Output("applyReputationControls")]
        public Output<bool?> ApplyReputationControls { get; private set; } = null!;

        /// <summary>
        /// Whether to enable slow post controls: either `true` or `false`.
        /// </summary>
        [Output("applySlowPostControls")]
        public Output<bool?> ApplySlowPostControls { get; private set; } = null!;

        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// The ID of the security policy to use.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecSecurityPolicyProtections resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecSecurityPolicyProtections(string name, AppSecSecurityPolicyProtectionsArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecSecurityPolicyProtections:AppSecSecurityPolicyProtections", name, args ?? new AppSecSecurityPolicyProtectionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecSecurityPolicyProtections(string name, Input<string> id, AppSecSecurityPolicyProtectionsState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecSecurityPolicyProtections:AppSecSecurityPolicyProtections", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecSecurityPolicyProtections resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecSecurityPolicyProtections Get(string name, Input<string> id, AppSecSecurityPolicyProtectionsState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecSecurityPolicyProtections(name, id, state, options);
        }
    }

    public sealed class AppSecSecurityPolicyProtectionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable api constraints: either `true` or `false`.
        /// </summary>
        [Input("applyApiConstraints")]
        public Input<bool>? ApplyApiConstraints { get; set; }

        /// <summary>
        /// Whether to enable application layer controls: either `true` or `false`.
        /// </summary>
        [Input("applyApplicationLayerControls")]
        public Input<bool>? ApplyApplicationLayerControls { get; set; }

        /// <summary>
        /// Whether to enable botman controls: either `true` or `false`.
        /// </summary>
        [Input("applyBotmanControls")]
        public Input<bool>? ApplyBotmanControls { get; set; }

        /// <summary>
        /// Whether to enable network layer controls: either `true` or `false`.
        /// </summary>
        [Input("applyNetworkLayerControls")]
        public Input<bool>? ApplyNetworkLayerControls { get; set; }

        /// <summary>
        /// Whether to enable rate controls: either `true` or `false`.
        /// </summary>
        [Input("applyRateControls")]
        public Input<bool>? ApplyRateControls { get; set; }

        /// <summary>
        /// Whether to enable reputation controls: either `true` or `false`.
        /// </summary>
        [Input("applyReputationControls")]
        public Input<bool>? ApplyReputationControls { get; set; }

        /// <summary>
        /// Whether to enable slow post controls: either `true` or `false`.
        /// </summary>
        [Input("applySlowPostControls")]
        public Input<bool>? ApplySlowPostControls { get; set; }

        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// The ID of the security policy to use.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Input("version", required: true)]
        public Input<int> Version { get; set; } = null!;

        public AppSecSecurityPolicyProtectionsArgs()
        {
        }
    }

    public sealed class AppSecSecurityPolicyProtectionsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable api constraints: either `true` or `false`.
        /// </summary>
        [Input("applyApiConstraints")]
        public Input<bool>? ApplyApiConstraints { get; set; }

        /// <summary>
        /// Whether to enable application layer controls: either `true` or `false`.
        /// </summary>
        [Input("applyApplicationLayerControls")]
        public Input<bool>? ApplyApplicationLayerControls { get; set; }

        /// <summary>
        /// Whether to enable botman controls: either `true` or `false`.
        /// </summary>
        [Input("applyBotmanControls")]
        public Input<bool>? ApplyBotmanControls { get; set; }

        /// <summary>
        /// Whether to enable network layer controls: either `true` or `false`.
        /// </summary>
        [Input("applyNetworkLayerControls")]
        public Input<bool>? ApplyNetworkLayerControls { get; set; }

        /// <summary>
        /// Whether to enable rate controls: either `true` or `false`.
        /// </summary>
        [Input("applyRateControls")]
        public Input<bool>? ApplyRateControls { get; set; }

        /// <summary>
        /// Whether to enable reputation controls: either `true` or `false`.
        /// </summary>
        [Input("applyReputationControls")]
        public Input<bool>? ApplyReputationControls { get; set; }

        /// <summary>
        /// Whether to enable slow post controls: either `true` or `false`.
        /// </summary>
        [Input("applySlowPostControls")]
        public Input<bool>? ApplySlowPostControls { get; set; }

        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// The ID of the security policy to use.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public AppSecSecurityPolicyProtectionsState()
        {
        }
    }
}
