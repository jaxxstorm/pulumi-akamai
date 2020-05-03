// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Properties
{
    /// <summary>
    /// The `akamai.Properties.PropertyActivation` provides the resource for activating a property in the appropriate environment. Once you are satisfied with any version of a property, an activation deploys it, either to the Akamai staging or production network. You activate a specific version, but the same version can be activated separately more than once.
    /// </summary>
    public partial class PropertyActivation : Pulumi.CustomResource
    {
        /// <summary>
        /// — (Optional, boolean) Whether to activate the property on the network. (Default: `true`).
        /// </summary>
        [Output("activate")]
        public Output<bool?> Activate { get; private set; } = null!;

        /// <summary>
        /// — (Required) One or more email addresses to inform about activation changes.
        /// </summary>
        [Output("contacts")]
        public Output<ImmutableArray<string>> Contacts { get; private set; } = null!;

        /// <summary>
        /// — (Optional) Akamai network to activate on. Allowed values `staging` or `production` (Default: `staging`).
        /// </summary>
        [Output("network")]
        public Output<string?> Network { get; private set; } = null!;

        /// <summary>
        /// — (Required) The property ID.
        /// </summary>
        [Output("property")]
        public Output<string> Property { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// — (Optional) The version to activate. When unset it will activate the latest version of the property.
        /// </summary>
        [Output("version")]
        public Output<int?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a PropertyActivation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PropertyActivation(string name, PropertyActivationArgs args, CustomResourceOptions? options = null)
            : base("akamai:Properties/propertyActivation:PropertyActivation", name, args ?? new PropertyActivationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PropertyActivation(string name, Input<string> id, PropertyActivationState? state = null, CustomResourceOptions? options = null)
            : base("akamai:Properties/propertyActivation:PropertyActivation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PropertyActivation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PropertyActivation Get(string name, Input<string> id, PropertyActivationState? state = null, CustomResourceOptions? options = null)
        {
            return new PropertyActivation(name, id, state, options);
        }
    }

    public sealed class PropertyActivationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// — (Optional, boolean) Whether to activate the property on the network. (Default: `true`).
        /// </summary>
        [Input("activate")]
        public Input<bool>? Activate { get; set; }

        [Input("contacts", required: true)]
        private InputList<string>? _contacts;

        /// <summary>
        /// — (Required) One or more email addresses to inform about activation changes.
        /// </summary>
        public InputList<string> Contacts
        {
            get => _contacts ?? (_contacts = new InputList<string>());
            set => _contacts = value;
        }

        /// <summary>
        /// — (Optional) Akamai network to activate on. Allowed values `staging` or `production` (Default: `staging`).
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// — (Required) The property ID.
        /// </summary>
        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        /// <summary>
        /// — (Optional) The version to activate. When unset it will activate the latest version of the property.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public PropertyActivationArgs()
        {
        }
    }

    public sealed class PropertyActivationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// — (Optional, boolean) Whether to activate the property on the network. (Default: `true`).
        /// </summary>
        [Input("activate")]
        public Input<bool>? Activate { get; set; }

        [Input("contacts")]
        private InputList<string>? _contacts;

        /// <summary>
        /// — (Required) One or more email addresses to inform about activation changes.
        /// </summary>
        public InputList<string> Contacts
        {
            get => _contacts ?? (_contacts = new InputList<string>());
            set => _contacts = value;
        }

        /// <summary>
        /// — (Optional) Akamai network to activate on. Allowed values `staging` or `production` (Default: `staging`).
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// — (Required) The property ID.
        /// </summary>
        [Input("property")]
        public Input<string>? Property { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// — (Optional) The version to activate. When unset it will activate the latest version of the property.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public PropertyActivationState()
        {
        }
    }
}
