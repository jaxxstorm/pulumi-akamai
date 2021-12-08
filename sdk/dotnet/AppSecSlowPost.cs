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
    /// **Scopes**: Security policy
    /// 
    /// Modifies slow POST protection settings for a security configuration and security policy. Slow POST protections help defend a site against attacks that try to tie up the site by using extremely slow requests and responses.
    /// 
    /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/slow-post](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putslowpostprotectionsettings)
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
    ///         var slowPost = new Akamai.AppSecSlowPost("slowPost", new Akamai.AppSecSlowPostArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             SecurityPolicyId = "gms1_134637",
    ///             SlowRateAction = "alert",
    ///             SlowRateThresholdRate = 10,
    ///             SlowRateThresholdPeriod = 30,
    ///             DurationThresholdTimeout = 20,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecSlowPost:AppSecSlowPost")]
    public partial class AppSecSlowPost : Pulumi.CustomResource
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the slow POST settings being modified.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// . Maximum amount of time (in seconds) that the first eight kilobytes of the POST body must be received in to avoid triggering the specified action.
        /// </summary>
        [Output("durationThresholdTimeout")]
        public Output<int?> DurationThresholdTimeout { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the slow POST settings being modified.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;

        /// <summary>
        /// . Action to be taken if slow POST protection is triggered. Allowed values are:
        /// - **alert**. Record the event.
        /// - **abort**. Block the request.
        /// </summary>
        [Output("slowRateAction")]
        public Output<string> SlowRateAction { get; private set; } = null!;

        /// <summary>
        /// . Amount of time (in seconds) that the server should allow a request before marking the request as being too slow.
        /// </summary>
        [Output("slowRateThresholdPeriod")]
        public Output<int?> SlowRateThresholdPeriod { get; private set; } = null!;

        /// <summary>
        /// . Average rate (in bytes per second over the specified time period) allowed before the specified action is triggered.
        /// </summary>
        [Output("slowRateThresholdRate")]
        public Output<int?> SlowRateThresholdRate { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecSlowPost resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecSlowPost(string name, AppSecSlowPostArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecSlowPost:AppSecSlowPost", name, args ?? new AppSecSlowPostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecSlowPost(string name, Input<string> id, AppSecSlowPostState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecSlowPost:AppSecSlowPost", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecSlowPost resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecSlowPost Get(string name, Input<string> id, AppSecSlowPostState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecSlowPost(name, id, state, options);
        }
    }

    public sealed class AppSecSlowPostArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the slow POST settings being modified.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Maximum amount of time (in seconds) that the first eight kilobytes of the POST body must be received in to avoid triggering the specified action.
        /// </summary>
        [Input("durationThresholdTimeout")]
        public Input<int>? DurationThresholdTimeout { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy associated with the slow POST settings being modified.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        /// <summary>
        /// . Action to be taken if slow POST protection is triggered. Allowed values are:
        /// - **alert**. Record the event.
        /// - **abort**. Block the request.
        /// </summary>
        [Input("slowRateAction", required: true)]
        public Input<string> SlowRateAction { get; set; } = null!;

        /// <summary>
        /// . Amount of time (in seconds) that the server should allow a request before marking the request as being too slow.
        /// </summary>
        [Input("slowRateThresholdPeriod")]
        public Input<int>? SlowRateThresholdPeriod { get; set; }

        /// <summary>
        /// . Average rate (in bytes per second over the specified time period) allowed before the specified action is triggered.
        /// </summary>
        [Input("slowRateThresholdRate")]
        public Input<int>? SlowRateThresholdRate { get; set; }

        public AppSecSlowPostArgs()
        {
        }
    }

    public sealed class AppSecSlowPostState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the slow POST settings being modified.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// . Maximum amount of time (in seconds) that the first eight kilobytes of the POST body must be received in to avoid triggering the specified action.
        /// </summary>
        [Input("durationThresholdTimeout")]
        public Input<int>? DurationThresholdTimeout { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy associated with the slow POST settings being modified.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        /// <summary>
        /// . Action to be taken if slow POST protection is triggered. Allowed values are:
        /// - **alert**. Record the event.
        /// - **abort**. Block the request.
        /// </summary>
        [Input("slowRateAction")]
        public Input<string>? SlowRateAction { get; set; }

        /// <summary>
        /// . Amount of time (in seconds) that the server should allow a request before marking the request as being too slow.
        /// </summary>
        [Input("slowRateThresholdPeriod")]
        public Input<int>? SlowRateThresholdPeriod { get; set; }

        /// <summary>
        /// . Average rate (in bytes per second over the specified time period) allowed before the specified action is triggered.
        /// </summary>
        [Input("slowRateThresholdRate")]
        public Input<int>? SlowRateThresholdRate { get; set; }

        public AppSecSlowPostState()
        {
        }
    }
}
