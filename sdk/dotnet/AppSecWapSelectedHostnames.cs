// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    [AkamaiResourceType("akamai:index/appSecWapSelectedHostnames:AppSecWapSelectedHostnames")]
    public partial class AppSecWapSelectedHostnames : Pulumi.CustomResource
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the hostnames being protected or evaluated.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        [Output("evaluatedHosts")]
        public Output<ImmutableArray<string>> EvaluatedHosts { get; private set; } = null!;

        [Output("protectedHosts")]
        public Output<ImmutableArray<string>> ProtectedHosts { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy responsible for protecting or evaluating the specified hosts.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecWapSelectedHostnames resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecWapSelectedHostnames(string name, AppSecWapSelectedHostnamesArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecWapSelectedHostnames:AppSecWapSelectedHostnames", name, args ?? new AppSecWapSelectedHostnamesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecWapSelectedHostnames(string name, Input<string> id, AppSecWapSelectedHostnamesState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecWapSelectedHostnames:AppSecWapSelectedHostnames", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecWapSelectedHostnames resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecWapSelectedHostnames Get(string name, Input<string> id, AppSecWapSelectedHostnamesState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecWapSelectedHostnames(name, id, state, options);
        }
    }

    public sealed class AppSecWapSelectedHostnamesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the hostnames being protected or evaluated.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        [Input("evaluatedHosts")]
        private InputList<string>? _evaluatedHosts;
        public InputList<string> EvaluatedHosts
        {
            get => _evaluatedHosts ?? (_evaluatedHosts = new InputList<string>());
            set => _evaluatedHosts = value;
        }

        [Input("protectedHosts")]
        private InputList<string>? _protectedHosts;
        public InputList<string> ProtectedHosts
        {
            get => _protectedHosts ?? (_protectedHosts = new InputList<string>());
            set => _protectedHosts = value;
        }

        /// <summary>
        /// . Unique identifier of the security policy responsible for protecting or evaluating the specified hosts.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        public AppSecWapSelectedHostnamesArgs()
        {
        }
    }

    public sealed class AppSecWapSelectedHostnamesState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the hostnames being protected or evaluated.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        [Input("evaluatedHosts")]
        private InputList<string>? _evaluatedHosts;
        public InputList<string> EvaluatedHosts
        {
            get => _evaluatedHosts ?? (_evaluatedHosts = new InputList<string>());
            set => _evaluatedHosts = value;
        }

        [Input("protectedHosts")]
        private InputList<string>? _protectedHosts;
        public InputList<string> ProtectedHosts
        {
            get => _protectedHosts ?? (_protectedHosts = new InputList<string>());
            set => _protectedHosts = value;
        }

        /// <summary>
        /// . Unique identifier of the security policy responsible for protecting or evaluating the specified hosts.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        public AppSecWapSelectedHostnamesState()
        {
        }
    }
}
