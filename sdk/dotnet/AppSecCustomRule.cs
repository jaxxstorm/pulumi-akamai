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
    /// Creates a custom rule associated with a security configuration. Custom rules are rules that you define yourself and are not part of the Kona Rule Set.
    /// 
    /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/custom-rules]https://techdocs.akamai.com/application-security/reference/get-configs-custom-rules)
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecCustomRule:AppSecCustomRule")]
    public partial class AppSecCustomRule : Pulumi.CustomResource
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the custom rule being modified.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// . Path to a JSON file containing the custom rule definition. To view a sample JSON file, see the [Create a custom rule](https://techdocs.akamai.com/application-security/reference/post-config-custom-rules) section of the Application Security API documentation.
        /// </summary>
        [Output("customRule")]
        public Output<string> CustomRule { get; private set; } = null!;

        [Output("customRuleId")]
        public Output<int> CustomRuleId { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecCustomRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecCustomRule(string name, AppSecCustomRuleArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecCustomRule:AppSecCustomRule", name, args ?? new AppSecCustomRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecCustomRule(string name, Input<string> id, AppSecCustomRuleState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecCustomRule:AppSecCustomRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecCustomRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecCustomRule Get(string name, Input<string> id, AppSecCustomRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecCustomRule(name, id, state, options);
        }
    }

    public sealed class AppSecCustomRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the custom rule being modified.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Path to a JSON file containing the custom rule definition. To view a sample JSON file, see the [Create a custom rule](https://techdocs.akamai.com/application-security/reference/post-config-custom-rules) section of the Application Security API documentation.
        /// </summary>
        [Input("customRule", required: true)]
        public Input<string> CustomRule { get; set; } = null!;

        public AppSecCustomRuleArgs()
        {
        }
    }

    public sealed class AppSecCustomRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the custom rule being modified.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// . Path to a JSON file containing the custom rule definition. To view a sample JSON file, see the [Create a custom rule](https://techdocs.akamai.com/application-security/reference/post-config-custom-rules) section of the Application Security API documentation.
        /// </summary>
        [Input("customRule")]
        public Input<string>? CustomRule { get; set; }

        [Input("customRuleId")]
        public Input<int>? CustomRuleId { get; set; }

        public AppSecCustomRuleState()
        {
        }
    }
}
