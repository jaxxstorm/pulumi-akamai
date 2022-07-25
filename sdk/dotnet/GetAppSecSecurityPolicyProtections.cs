// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecSecurityPolicyProtections
    {
        /// <summary>
        /// **Scopes**: Security policy
        /// 
        /// Returns information about the protections in effect for the specified security policy.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/protections](https://techdocs.akamai.com/application-security/reference/get-policy-protections)
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
        ///             Name = "Documentation",
        ///         }));
        ///         var protections = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecSecurityPolicyProtections.InvokeAsync(new Akamai.GetAppSecSecurityPolicyProtectionsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = "gms1_134637",
        ///         })));
        ///         this.ProtectionsJson = protections.Apply(protections =&gt; protections.Json);
        ///         this.ProtectionsApplyApiConstraints = protections.Apply(protections =&gt; protections.ApplyApiConstraints);
        ///         this.ProtectionsApplyApplicationLayerControls = protections.Apply(protections =&gt; protections.ApplyApplicationLayerControls);
        ///         this.ProtectionsApplyBotmanControls = protections.Apply(protections =&gt; protections.ApplyBotmanControls);
        ///         this.ProtectionsApplyNetworkLayerControls = protections.Apply(protections =&gt; protections.ApplyNetworkLayerControls);
        ///         this.ProtectionsApplyRateControls = protections.Apply(protections =&gt; protections.ApplyRateControls);
        ///         this.ProtectionsApplyReputationControls = protections.Apply(protections =&gt; protections.ApplyReputationControls);
        ///         this.ProtectionsApplySlowPostControls = protections.Apply(protections =&gt; protections.ApplySlowPostControls);
        ///     }
        /// 
        ///     [Output("protectionsJson")]
        ///     public Output&lt;string&gt; ProtectionsJson { get; set; }
        ///     [Output("protectionsApplyApiConstraints")]
        ///     public Output&lt;string&gt; ProtectionsApplyApiConstraints { get; set; }
        ///     [Output("protectionsApplyApplicationLayerControls")]
        ///     public Output&lt;string&gt; ProtectionsApplyApplicationLayerControls { get; set; }
        ///     [Output("protectionsApplyBotmanControls")]
        ///     public Output&lt;string&gt; ProtectionsApplyBotmanControls { get; set; }
        ///     [Output("protectionsApplyNetworkLayerControls")]
        ///     public Output&lt;string&gt; ProtectionsApplyNetworkLayerControls { get; set; }
        ///     [Output("protectionsApplyRateControls")]
        ///     public Output&lt;string&gt; ProtectionsApplyRateControls { get; set; }
        ///     [Output("protectionsApplyReputationControls")]
        ///     public Output&lt;string&gt; ProtectionsApplyReputationControls { get; set; }
        ///     [Output("protectionsApplySlowPostControls")]
        ///     public Output&lt;string&gt; ProtectionsApplySlowPostControls { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned and how that returned information is formatted:
        /// 
        /// - `apply_application_layer_controls`. Returns **true** if application layer controls are enabled; returns **false** if they are not.
        /// - `apply_network_layer_controls`. Returns **true** if network layer controls are enabled; returns **false** if they are not.
        /// - `apply_rate_controls`. Returns **true** if rate controls are enabled; returns **false** if they are not.
        /// - `apply_reputation_controls`. Returns **true** if reputation controls are enabled; returns **false** if they are not.
        /// - `apply_botman_controls`. Returns **true** if Bot Manager controls are enabled; returns **false** if they are not.
        /// - `apply_api_constraints`. Returns **true** if API constraints are enabled; returns **false** if they are not.
        /// - `apply_slow_post_controls`. Returns **true** if slow POST controls are enabled; returns **false** if they are not.
        /// - `json`. JSON-formatted list showing the status of the protection settings.
        /// - `output_text`. Tabular report showing the status of the protection settings.
        /// </summary>
        public static Task<GetAppSecSecurityPolicyProtectionsResult> InvokeAsync(GetAppSecSecurityPolicyProtectionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecSecurityPolicyProtectionsResult>("akamai:index/getAppSecSecurityPolicyProtections:getAppSecSecurityPolicyProtections", args ?? new GetAppSecSecurityPolicyProtectionsArgs(), options.WithDefaults());

        /// <summary>
        /// **Scopes**: Security policy
        /// 
        /// Returns information about the protections in effect for the specified security policy.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/protections](https://techdocs.akamai.com/application-security/reference/get-policy-protections)
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
        ///             Name = "Documentation",
        ///         }));
        ///         var protections = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecSecurityPolicyProtections.InvokeAsync(new Akamai.GetAppSecSecurityPolicyProtectionsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = "gms1_134637",
        ///         })));
        ///         this.ProtectionsJson = protections.Apply(protections =&gt; protections.Json);
        ///         this.ProtectionsApplyApiConstraints = protections.Apply(protections =&gt; protections.ApplyApiConstraints);
        ///         this.ProtectionsApplyApplicationLayerControls = protections.Apply(protections =&gt; protections.ApplyApplicationLayerControls);
        ///         this.ProtectionsApplyBotmanControls = protections.Apply(protections =&gt; protections.ApplyBotmanControls);
        ///         this.ProtectionsApplyNetworkLayerControls = protections.Apply(protections =&gt; protections.ApplyNetworkLayerControls);
        ///         this.ProtectionsApplyRateControls = protections.Apply(protections =&gt; protections.ApplyRateControls);
        ///         this.ProtectionsApplyReputationControls = protections.Apply(protections =&gt; protections.ApplyReputationControls);
        ///         this.ProtectionsApplySlowPostControls = protections.Apply(protections =&gt; protections.ApplySlowPostControls);
        ///     }
        /// 
        ///     [Output("protectionsJson")]
        ///     public Output&lt;string&gt; ProtectionsJson { get; set; }
        ///     [Output("protectionsApplyApiConstraints")]
        ///     public Output&lt;string&gt; ProtectionsApplyApiConstraints { get; set; }
        ///     [Output("protectionsApplyApplicationLayerControls")]
        ///     public Output&lt;string&gt; ProtectionsApplyApplicationLayerControls { get; set; }
        ///     [Output("protectionsApplyBotmanControls")]
        ///     public Output&lt;string&gt; ProtectionsApplyBotmanControls { get; set; }
        ///     [Output("protectionsApplyNetworkLayerControls")]
        ///     public Output&lt;string&gt; ProtectionsApplyNetworkLayerControls { get; set; }
        ///     [Output("protectionsApplyRateControls")]
        ///     public Output&lt;string&gt; ProtectionsApplyRateControls { get; set; }
        ///     [Output("protectionsApplyReputationControls")]
        ///     public Output&lt;string&gt; ProtectionsApplyReputationControls { get; set; }
        ///     [Output("protectionsApplySlowPostControls")]
        ///     public Output&lt;string&gt; ProtectionsApplySlowPostControls { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned and how that returned information is formatted:
        /// 
        /// - `apply_application_layer_controls`. Returns **true** if application layer controls are enabled; returns **false** if they are not.
        /// - `apply_network_layer_controls`. Returns **true** if network layer controls are enabled; returns **false** if they are not.
        /// - `apply_rate_controls`. Returns **true** if rate controls are enabled; returns **false** if they are not.
        /// - `apply_reputation_controls`. Returns **true** if reputation controls are enabled; returns **false** if they are not.
        /// - `apply_botman_controls`. Returns **true** if Bot Manager controls are enabled; returns **false** if they are not.
        /// - `apply_api_constraints`. Returns **true** if API constraints are enabled; returns **false** if they are not.
        /// - `apply_slow_post_controls`. Returns **true** if slow POST controls are enabled; returns **false** if they are not.
        /// - `json`. JSON-formatted list showing the status of the protection settings.
        /// - `output_text`. Tabular report showing the status of the protection settings.
        /// </summary>
        public static Output<GetAppSecSecurityPolicyProtectionsResult> Invoke(GetAppSecSecurityPolicyProtectionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecSecurityPolicyProtectionsResult>("akamai:index/getAppSecSecurityPolicyProtections:getAppSecSecurityPolicyProtections", args ?? new GetAppSecSecurityPolicyProtectionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecSecurityPolicyProtectionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the security policy protections.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy you want to return protections information for.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public string SecurityPolicyId { get; set; } = null!;

        public GetAppSecSecurityPolicyProtectionsArgs()
        {
        }
    }

    public sealed class GetAppSecSecurityPolicyProtectionsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the security policy protections.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy you want to return protections information for.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        public GetAppSecSecurityPolicyProtectionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecSecurityPolicyProtectionsResult
    {
        public readonly bool ApplyApiConstraints;
        public readonly bool ApplyApplicationLayerControls;
        public readonly bool ApplyBotmanControls;
        public readonly bool ApplyNetworkLayerControls;
        public readonly bool ApplyRateControls;
        public readonly bool ApplyReputationControls;
        public readonly bool ApplySlowPostControls;
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly string OutputText;
        public readonly string SecurityPolicyId;

        [OutputConstructor]
        private GetAppSecSecurityPolicyProtectionsResult(
            bool applyApiConstraints,

            bool applyApplicationLayerControls,

            bool applyBotmanControls,

            bool applyNetworkLayerControls,

            bool applyRateControls,

            bool applyReputationControls,

            bool applySlowPostControls,

            int configId,

            string id,

            string json,

            string outputText,

            string securityPolicyId)
        {
            ApplyApiConstraints = applyApiConstraints;
            ApplyApplicationLayerControls = applyApplicationLayerControls;
            ApplyBotmanControls = applyBotmanControls;
            ApplyNetworkLayerControls = applyNetworkLayerControls;
            ApplyRateControls = applyRateControls;
            ApplyReputationControls = applyReputationControls;
            ApplySlowPostControls = applySlowPostControls;
            ConfigId = configId;
            Id = id;
            Json = json;
            OutputText = outputText;
            SecurityPolicyId = securityPolicyId;
        }
    }
}
