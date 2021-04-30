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
    /// Use the `akamai.AppSecSiemSettings` resource to mpdate the SIEM integration settings for a specific configuration.
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
    ///         var siemDefinition = Output.Create(Akamai.GetAppSecSiemDefinitions.InvokeAsync(new Akamai.GetAppSecSiemDefinitionsArgs
    ///         {
    ///             SiemDefinitionName = @var.Siem_definition_name,
    ///         }));
    ///         var securityPolicies = Output.Tuple(configuration, configuration).Apply(values =&gt;
    ///         {
    ///             var configuration = values.Item1;
    ///             var configuration1 = values.Item2;
    ///             return Output.Create(Akamai.GetAppSecSecurityPolicy.InvokeAsync(new Akamai.GetAppSecSecurityPolicyArgs
    ///             {
    ///                 ConfigId = configuration.ConfigId,
    ///                 Version = configuration1.LatestVersion,
    ///             }));
    ///         });
    ///         var siem = new Akamai.AppSecSiemSettings("siem", new Akamai.AppSecSiemSettingsArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             Version = configuration.Apply(configuration =&gt; configuration.LatestVersion),
    ///             EnableSiem = true,
    ///             EnableForAllPolicies = false,
    ///             EnableBotmanSiem = true,
    ///             SiemId = siemDefinition.Apply(siemDefinition =&gt; siemDefinition.Id),
    ///             SecurityPolicyIds = securityPolicies.Apply(securityPolicies =&gt; securityPolicies.PolicyLists),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecSiemSettings:AppSecSiemSettings")]
    public partial class AppSecSiemSettings : Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration ID to use.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// Whether you enabled SIEM for the Bot Manager events.
        /// </summary>
        [Output("enableBotmanSiem")]
        public Output<bool?> EnableBotmanSiem { get; private set; } = null!;

        /// <summary>
        /// Whether you enabled SIEM for all the security policies in the configuration version.
        /// </summary>
        [Output("enableForAllPolicies")]
        public Output<bool> EnableForAllPolicies { get; private set; } = null!;

        /// <summary>
        /// Whether you enabled SIEM in a security configuration version.
        /// </summary>
        [Output("enableSiem")]
        public Output<bool> EnableSiem { get; private set; } = null!;

        /// <summary>
        /// A tabular display showing the updated SIEM integration settings.
        /// </summary>
        [Output("outputText")]
        public Output<string> OutputText { get; private set; } = null!;

        /// <summary>
        /// The list of security policy identifiers for which to enable the SIEM integration.
        /// </summary>
        [Output("securityPolicyIds")]
        public Output<ImmutableArray<string>> SecurityPolicyIds { get; private set; } = null!;

        /// <summary>
        /// An integer that uniquely identifies the SIEM settings.
        /// </summary>
        [Output("siemId")]
        public Output<int> SiemId { get; private set; } = null!;

        /// <summary>
        /// The version number of the configuration to use.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecSiemSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecSiemSettings(string name, AppSecSiemSettingsArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecSiemSettings:AppSecSiemSettings", name, args ?? new AppSecSiemSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecSiemSettings(string name, Input<string> id, AppSecSiemSettingsState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecSiemSettings:AppSecSiemSettings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecSiemSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecSiemSettings Get(string name, Input<string> id, AppSecSiemSettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecSiemSettings(name, id, state, options);
        }
    }

    public sealed class AppSecSiemSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration ID to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// Whether you enabled SIEM for the Bot Manager events.
        /// </summary>
        [Input("enableBotmanSiem")]
        public Input<bool>? EnableBotmanSiem { get; set; }

        /// <summary>
        /// Whether you enabled SIEM for all the security policies in the configuration version.
        /// </summary>
        [Input("enableForAllPolicies", required: true)]
        public Input<bool> EnableForAllPolicies { get; set; } = null!;

        /// <summary>
        /// Whether you enabled SIEM in a security configuration version.
        /// </summary>
        [Input("enableSiem", required: true)]
        public Input<bool> EnableSiem { get; set; } = null!;

        [Input("securityPolicyIds", required: true)]
        private InputList<string>? _securityPolicyIds;

        /// <summary>
        /// The list of security policy identifiers for which to enable the SIEM integration.
        /// </summary>
        public InputList<string> SecurityPolicyIds
        {
            get => _securityPolicyIds ?? (_securityPolicyIds = new InputList<string>());
            set => _securityPolicyIds = value;
        }

        /// <summary>
        /// An integer that uniquely identifies the SIEM settings.
        /// </summary>
        [Input("siemId", required: true)]
        public Input<int> SiemId { get; set; } = null!;

        /// <summary>
        /// The version number of the configuration to use.
        /// </summary>
        [Input("version", required: true)]
        public Input<int> Version { get; set; } = null!;

        public AppSecSiemSettingsArgs()
        {
        }
    }

    public sealed class AppSecSiemSettingsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration ID to use.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// Whether you enabled SIEM for the Bot Manager events.
        /// </summary>
        [Input("enableBotmanSiem")]
        public Input<bool>? EnableBotmanSiem { get; set; }

        /// <summary>
        /// Whether you enabled SIEM for all the security policies in the configuration version.
        /// </summary>
        [Input("enableForAllPolicies")]
        public Input<bool>? EnableForAllPolicies { get; set; }

        /// <summary>
        /// Whether you enabled SIEM in a security configuration version.
        /// </summary>
        [Input("enableSiem")]
        public Input<bool>? EnableSiem { get; set; }

        /// <summary>
        /// A tabular display showing the updated SIEM integration settings.
        /// </summary>
        [Input("outputText")]
        public Input<string>? OutputText { get; set; }

        [Input("securityPolicyIds")]
        private InputList<string>? _securityPolicyIds;

        /// <summary>
        /// The list of security policy identifiers for which to enable the SIEM integration.
        /// </summary>
        public InputList<string> SecurityPolicyIds
        {
            get => _securityPolicyIds ?? (_securityPolicyIds = new InputList<string>());
            set => _securityPolicyIds = value;
        }

        /// <summary>
        /// An integer that uniquely identifies the SIEM settings.
        /// </summary>
        [Input("siemId")]
        public Input<int>? SiemId { get; set; }

        /// <summary>
        /// The version number of the configuration to use.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public AppSecSiemSettingsState()
        {
        }
    }
}
