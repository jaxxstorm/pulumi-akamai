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
    /// Modifies SIEM (Security Information and Event Management) integration settings for a security configuration.
    /// 
    /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/siem](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsiemsettings)
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
    ///         var siemDefinition = Output.Create(Akamai.GetAppSecSiemDefinitions.InvokeAsync(new Akamai.GetAppSecSiemDefinitionsArgs
    ///         {
    ///             SiemDefinitionName = "SIEM Version 01",
    ///         }));
    ///         var securityPolicies = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecSecurityPolicy.InvokeAsync(new Akamai.GetAppSecSecurityPolicyArgs
    ///         {
    ///             ConfigId = configuration.ConfigId,
    ///         })));
    ///         var siem = new Akamai.AppSecSiemSettings("siem", new Akamai.AppSecSiemSettingsArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             EnableSiem = true,
    ///             EnableForAllPolicies = false,
    ///             EnableBotmanSiem = true,
    ///             SiemId = siemDefinition.Apply(siemDefinition =&gt; siemDefinition.Id),
    ///             SecurityPolicyIds = securityPolicies.Apply(securityPolicies =&gt; securityPolicies.SecurityPolicyIdLists),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Output Options
    /// 
    /// The following options can be used to determine the information returned, and how that returned information is formatted:
    /// 
    /// - `output_text`. Tabular report showing the updated SIEM integration settings.
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecSiemSettings:AppSecSiemSettings")]
    public partial class AppSecSiemSettings : Pulumi.CustomResource
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the SIEM settings being modified.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
        /// </summary>
        [Output("enableBotmanSiem")]
        public Output<bool> EnableBotmanSiem { get; private set; } = null!;

        /// <summary>
        /// . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `security_policy_ids` argument.
        /// </summary>
        [Output("enableForAllPolicies")]
        public Output<bool> EnableForAllPolicies { get; private set; } = null!;

        /// <summary>
        /// . Set to **true** to enable SIEM; set to **false** to disable SIEM.
        /// </summary>
        [Output("enableSiem")]
        public Output<bool> EnableSiem { get; private set; } = null!;

        /// <summary>
        /// JSON array of IDs for the security policies where SIEM integration is to be enabled.
        /// </summary>
        [Output("securityPolicyIds")]
        public Output<ImmutableArray<string>> SecurityPolicyIds { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the SIEM settings being modified.
        /// </summary>
        [Output("siemId")]
        public Output<int> SiemId { get; private set; } = null!;


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
        /// . Unique identifier of the security configuration associated with the SIEM settings being modified.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
        /// </summary>
        [Input("enableBotmanSiem", required: true)]
        public Input<bool> EnableBotmanSiem { get; set; } = null!;

        /// <summary>
        /// . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `security_policy_ids` argument.
        /// </summary>
        [Input("enableForAllPolicies", required: true)]
        public Input<bool> EnableForAllPolicies { get; set; } = null!;

        /// <summary>
        /// . Set to **true** to enable SIEM; set to **false** to disable SIEM.
        /// </summary>
        [Input("enableSiem", required: true)]
        public Input<bool> EnableSiem { get; set; } = null!;

        [Input("securityPolicyIds")]
        private InputList<string>? _securityPolicyIds;

        /// <summary>
        /// JSON array of IDs for the security policies where SIEM integration is to be enabled.
        /// </summary>
        public InputList<string> SecurityPolicyIds
        {
            get => _securityPolicyIds ?? (_securityPolicyIds = new InputList<string>());
            set => _securityPolicyIds = value;
        }

        /// <summary>
        /// . Unique identifier of the SIEM settings being modified.
        /// </summary>
        [Input("siemId", required: true)]
        public Input<int> SiemId { get; set; } = null!;

        public AppSecSiemSettingsArgs()
        {
        }
    }

    public sealed class AppSecSiemSettingsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the SIEM settings being modified.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
        /// </summary>
        [Input("enableBotmanSiem")]
        public Input<bool>? EnableBotmanSiem { get; set; }

        /// <summary>
        /// . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `security_policy_ids` argument.
        /// </summary>
        [Input("enableForAllPolicies")]
        public Input<bool>? EnableForAllPolicies { get; set; }

        /// <summary>
        /// . Set to **true** to enable SIEM; set to **false** to disable SIEM.
        /// </summary>
        [Input("enableSiem")]
        public Input<bool>? EnableSiem { get; set; }

        [Input("securityPolicyIds")]
        private InputList<string>? _securityPolicyIds;

        /// <summary>
        /// JSON array of IDs for the security policies where SIEM integration is to be enabled.
        /// </summary>
        public InputList<string> SecurityPolicyIds
        {
            get => _securityPolicyIds ?? (_securityPolicyIds = new InputList<string>());
            set => _securityPolicyIds = value;
        }

        /// <summary>
        /// . Unique identifier of the SIEM settings being modified.
        /// </summary>
        [Input("siemId")]
        public Input<int>? SiemId { get; set; }

        public AppSecSiemSettingsState()
        {
        }
    }
}
