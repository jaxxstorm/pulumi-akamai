// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    [AkamaiResourceType("akamai:index/appSecSelectedHostnames:AppSecSelectedHostnames")]
    public partial class AppSecSelectedHostnames : Pulumi.CustomResource
    {
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        [Output("hostnames")]
        public Output<ImmutableArray<string>> Hostnames { get; private set; } = null!;

        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecSelectedHostnames resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecSelectedHostnames(string name, AppSecSelectedHostnamesArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecSelectedHostnames:AppSecSelectedHostnames", name, args ?? new AppSecSelectedHostnamesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecSelectedHostnames(string name, Input<string> id, AppSecSelectedHostnamesState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecSelectedHostnames:AppSecSelectedHostnames", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecSelectedHostnames resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecSelectedHostnames Get(string name, Input<string> id, AppSecSelectedHostnamesState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecSelectedHostnames(name, id, state, options);
        }
    }

    public sealed class AppSecSelectedHostnamesArgs : Pulumi.ResourceArgs
    {
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        [Input("hostnames", required: true)]
        private InputList<string>? _hostnames;
        public InputList<string> Hostnames
        {
            get => _hostnames ?? (_hostnames = new InputList<string>());
            set => _hostnames = value;
        }

        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        [Input("version", required: true)]
        public Input<int> Version { get; set; } = null!;

        public AppSecSelectedHostnamesArgs()
        {
        }
    }

    public sealed class AppSecSelectedHostnamesState : Pulumi.ResourceArgs
    {
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        [Input("hostnames")]
        private InputList<string>? _hostnames;
        public InputList<string> Hostnames
        {
            get => _hostnames ?? (_hostnames = new InputList<string>());
            set => _hostnames = value;
        }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public AppSecSelectedHostnamesState()
        {
        }
    }
}
