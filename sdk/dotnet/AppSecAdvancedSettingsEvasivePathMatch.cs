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
    /// The `resource_akamai_appsec_advanced_settings_evasive_path_match` resource allows you to enable, disable, or update the evasive path match setting for a configuration. This setting determines whether fuzzy matching is used to make URL matching more inclusive.
    /// This operation applies at the configuration level, and therefore applies to all policies within a configuration. You may override this setting for a particular policy by specifying the policy using the security_policy_id parameter.
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
    ///         var configEvasivePathMatch = new Akamai.AppSecAdvancedSettingsEvasivePathMatch("configEvasivePathMatch", new Akamai.AppSecAdvancedSettingsEvasivePathMatchArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             EnablePathMatch = true,
    ///         });
    ///         // USE CASE: user wants to override the evasive path match setting for a security policy
    ///         var policyOverride = new Akamai.AppSecAdvancedSettingsEvasivePathMatch("policyOverride", new Akamai.AppSecAdvancedSettingsEvasivePathMatchArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             SecurityPolicyId = @var.Security_policy_id,
    ///             EnablePathMatch = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecAdvancedSettingsEvasivePathMatch:AppSecAdvancedSettingsEvasivePathMatch")]
    public partial class AppSecAdvancedSettingsEvasivePathMatch : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// Whether to enable path match.
        /// </summary>
        [Output("enablePathMatch")]
        public Output<bool> EnablePathMatch { get; private set; } = null!;

        /// <summary>
        /// The ID of a specific security policy to which the evasive path match setting should be applied. If not supplied, the indicated setting will be applied to all policies within the configuration.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string?> SecurityPolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecAdvancedSettingsEvasivePathMatch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecAdvancedSettingsEvasivePathMatch(string name, AppSecAdvancedSettingsEvasivePathMatchArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecAdvancedSettingsEvasivePathMatch:AppSecAdvancedSettingsEvasivePathMatch", name, args ?? new AppSecAdvancedSettingsEvasivePathMatchArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecAdvancedSettingsEvasivePathMatch(string name, Input<string> id, AppSecAdvancedSettingsEvasivePathMatchState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecAdvancedSettingsEvasivePathMatch:AppSecAdvancedSettingsEvasivePathMatch", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecAdvancedSettingsEvasivePathMatch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecAdvancedSettingsEvasivePathMatch Get(string name, Input<string> id, AppSecAdvancedSettingsEvasivePathMatchState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecAdvancedSettingsEvasivePathMatch(name, id, state, options);
        }
    }

    public sealed class AppSecAdvancedSettingsEvasivePathMatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// Whether to enable path match.
        /// </summary>
        [Input("enablePathMatch", required: true)]
        public Input<bool> EnablePathMatch { get; set; } = null!;

        /// <summary>
        /// The ID of a specific security policy to which the evasive path match setting should be applied. If not supplied, the indicated setting will be applied to all policies within the configuration.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        public AppSecAdvancedSettingsEvasivePathMatchArgs()
        {
        }
    }

    public sealed class AppSecAdvancedSettingsEvasivePathMatchState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// Whether to enable path match.
        /// </summary>
        [Input("enablePathMatch")]
        public Input<bool>? EnablePathMatch { get; set; }

        /// <summary>
        /// The ID of a specific security policy to which the evasive path match setting should be applied. If not supplied, the indicated setting will be applied to all policies within the configuration.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        public AppSecAdvancedSettingsEvasivePathMatchState()
        {
        }
    }
}
