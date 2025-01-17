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
    /// Use the `akamai.CloudletsApplicationLoadBalancerActivation` resource to activate the Application Load Balancer Cloudlet configuration. An activation deploys the configuration version to either the Akamai staging or production network. You can activate a specific version multiple times if you need to.
    /// 
    /// Before activating on production, activate on staging first. This way you can detect any problems in staging before your changes progress to production.
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
    ///         var example = new Akamai.CloudletsApplicationLoadBalancerActivation("example", new Akamai.CloudletsApplicationLoadBalancerActivationArgs
    ///         {
    ///             OriginId = "alb_test_1",
    ///             Network = "staging",
    ///             Version = 1,
    ///         });
    ///         this.Status = example.Status;
    ///     }
    /// 
    ///     [Output("status")]
    ///     public Output&lt;string&gt; Status { get; set; }
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/cloudletsApplicationLoadBalancerActivation:CloudletsApplicationLoadBalancerActivation")]
    public partial class CloudletsApplicationLoadBalancerActivation : Pulumi.CustomResource
    {
        /// <summary>
        /// The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
        /// </summary>
        [Output("originId")]
        public Output<string> OriginId { get; private set; } = null!;

        /// <summary>
        /// The activation status for this load balancing configuration.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The Application Load Balancer Cloudlet configuration version you want to activate.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a CloudletsApplicationLoadBalancerActivation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudletsApplicationLoadBalancerActivation(string name, CloudletsApplicationLoadBalancerActivationArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/cloudletsApplicationLoadBalancerActivation:CloudletsApplicationLoadBalancerActivation", name, args ?? new CloudletsApplicationLoadBalancerActivationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudletsApplicationLoadBalancerActivation(string name, Input<string> id, CloudletsApplicationLoadBalancerActivationState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/cloudletsApplicationLoadBalancerActivation:CloudletsApplicationLoadBalancerActivation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CloudletsApplicationLoadBalancerActivation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudletsApplicationLoadBalancerActivation Get(string name, Input<string> id, CloudletsApplicationLoadBalancerActivationState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudletsApplicationLoadBalancerActivation(name, id, state, options);
        }
    }

    public sealed class CloudletsApplicationLoadBalancerActivationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
        /// </summary>
        [Input("originId", required: true)]
        public Input<string> OriginId { get; set; } = null!;

        /// <summary>
        /// The Application Load Balancer Cloudlet configuration version you want to activate.
        /// </summary>
        [Input("version", required: true)]
        public Input<int> Version { get; set; } = null!;

        public CloudletsApplicationLoadBalancerActivationArgs()
        {
        }
    }

    public sealed class CloudletsApplicationLoadBalancerActivationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
        /// </summary>
        [Input("originId")]
        public Input<string>? OriginId { get; set; }

        /// <summary>
        /// The activation status for this load balancing configuration.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The Application Load Balancer Cloudlet configuration version you want to activate.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public CloudletsApplicationLoadBalancerActivationState()
        {
        }
    }
}
