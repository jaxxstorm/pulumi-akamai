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
    /// **Scopes**: Rate policy
    /// 
    /// Creates, modifies, or deletes the actions associated with a rate policy.
    /// By default, rate policies take no action when triggered.
    /// Note that you must set separate actions for requests originating from an IPv4 IP address and for requests originating from an IPv6 address.
    /// 
    /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rate-policies/{ratePolicyId}](https://techdocs.akamai.com/application-security/reference/put-rate-policy-action)
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
    ///             Name = "Documentation",
    ///         }));
    ///         var appsecRatePolicy = new Akamai.AppSecRatePolicy("appsecRatePolicy", new Akamai.AppSecRatePolicyArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             RatePolicy = File.ReadAllText($"{path.Module}/rate_policy.json"),
    ///         });
    ///         var appsecRatePolicyAction = new Akamai.AppSecRatePolicyAction("appsecRatePolicyAction", new Akamai.AppSecRatePolicyActionArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             SecurityPolicyId = "gms1_134637",
    ///             RatePolicyId = appsecRatePolicy.RatePolicyId,
    ///             Ipv4Action = "deny",
    ///             Ipv6Action = "deny",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecRatePolicyAction:AppSecRatePolicyAction")]
    public partial class AppSecRatePolicyAction : Pulumi.CustomResource
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the rate policy action being modified.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// . Rate policy action for requests coming from an IPv4 IP address. Allowed actions are:
        /// - **alert**. Record the event.
        /// - **deny**. Block the request.
        /// - **deny_custom{custom_deny_id}**. Take the action specified by the custom deny.
        /// - **none**. Take no action.
        /// </summary>
        [Output("ipv4Action")]
        public Output<string> Ipv4Action { get; private set; } = null!;

        /// <summary>
        /// . Rate policy action for requests coming from an IPv6 IP address. Allowed actions are:
        /// - **alert**. Record the event.
        /// - **deny**. Block the request.
        /// - **deny_custom{custom_deny_id}**. Take the action specified by the custom deny.
        /// </summary>
        [Output("ipv6Action")]
        public Output<string> Ipv6Action { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the rate policy whose action is being modified.
        /// </summary>
        [Output("ratePolicyId")]
        public Output<int> RatePolicyId { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the rate policy whose action is being modified.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecRatePolicyAction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecRatePolicyAction(string name, AppSecRatePolicyActionArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecRatePolicyAction:AppSecRatePolicyAction", name, args ?? new AppSecRatePolicyActionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecRatePolicyAction(string name, Input<string> id, AppSecRatePolicyActionState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecRatePolicyAction:AppSecRatePolicyAction", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecRatePolicyAction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecRatePolicyAction Get(string name, Input<string> id, AppSecRatePolicyActionState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecRatePolicyAction(name, id, state, options);
        }
    }

    public sealed class AppSecRatePolicyActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the rate policy action being modified.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Rate policy action for requests coming from an IPv4 IP address. Allowed actions are:
        /// - **alert**. Record the event.
        /// - **deny**. Block the request.
        /// - **deny_custom{custom_deny_id}**. Take the action specified by the custom deny.
        /// - **none**. Take no action.
        /// </summary>
        [Input("ipv4Action", required: true)]
        public Input<string> Ipv4Action { get; set; } = null!;

        /// <summary>
        /// . Rate policy action for requests coming from an IPv6 IP address. Allowed actions are:
        /// - **alert**. Record the event.
        /// - **deny**. Block the request.
        /// - **deny_custom{custom_deny_id}**. Take the action specified by the custom deny.
        /// </summary>
        [Input("ipv6Action", required: true)]
        public Input<string> Ipv6Action { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the rate policy whose action is being modified.
        /// </summary>
        [Input("ratePolicyId", required: true)]
        public Input<int> RatePolicyId { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the rate policy whose action is being modified.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        public AppSecRatePolicyActionArgs()
        {
        }
    }

    public sealed class AppSecRatePolicyActionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the rate policy action being modified.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// . Rate policy action for requests coming from an IPv4 IP address. Allowed actions are:
        /// - **alert**. Record the event.
        /// - **deny**. Block the request.
        /// - **deny_custom{custom_deny_id}**. Take the action specified by the custom deny.
        /// - **none**. Take no action.
        /// </summary>
        [Input("ipv4Action")]
        public Input<string>? Ipv4Action { get; set; }

        /// <summary>
        /// . Rate policy action for requests coming from an IPv6 IP address. Allowed actions are:
        /// - **alert**. Record the event.
        /// - **deny**. Block the request.
        /// - **deny_custom{custom_deny_id}**. Take the action specified by the custom deny.
        /// </summary>
        [Input("ipv6Action")]
        public Input<string>? Ipv6Action { get; set; }

        /// <summary>
        /// . Unique identifier of the rate policy whose action is being modified.
        /// </summary>
        [Input("ratePolicyId")]
        public Input<int>? RatePolicyId { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy associated with the rate policy whose action is being modified.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        public AppSecRatePolicyActionState()
        {
        }
    }
}
