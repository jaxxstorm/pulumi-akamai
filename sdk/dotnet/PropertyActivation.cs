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
    /// The `akamai.PropertyActivation` resource lets you activate a property version. An activation deploys the version to either the Akamai staging or production network. You can activate a specific version multiple times if you need to.
    /// 
    /// Before activating on production, activate on staging first. This way you can detect any problems in staging before your changes progress to production.
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
    ///         var email = "user@example.org";
    ///         var ruleFormat = "v2020-03-04";
    ///         var example = new Akamai.Property("example", new Akamai.PropertyArgs
    ///         {
    ///             Contacts = 
    ///             {
    ///                 "user@example.org",
    ///             },
    ///             ProductId = "prd_SPM",
    ///             ContractId = @var.Contractid,
    ///             GroupId = @var.Groupid,
    ///             Hostnames = 
    ///             {
    ///                 { "example.org", "example.org.edgesuite.net" },
    ///                 { "www.example.org", "example.org.edgesuite.net" },
    ///                 { "sub.example.org", "sub.example.org.edgesuite.net" },
    ///             },
    ///             RuleFormat = ruleFormat,
    ///             Rules = File.ReadAllText($"{path.Module}/main.json"),
    ///         });
    ///         var exampleStaging = new Akamai.PropertyActivation("exampleStaging", new Akamai.PropertyActivationArgs
    ///         {
    ///             PropertyId = example.Id,
    ///             Contacts = 
    ///             {
    ///                 email,
    ///             },
    ///             Version = example.LatestVersion,
    ///         });
    ///         // not specifying network will target STAGING
    ///         var exampleProd = new Akamai.PropertyActivation("exampleProd", new Akamai.PropertyActivationArgs
    ///         {
    ///             PropertyId = example.Id,
    ///             Network = "PRODUCTION",
    ///             Version = 3,
    ///             Contacts = 
    ///             {
    ///                 email,
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleStaging,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Argument reference
    /// 
    /// The following arguments are supported:
    /// 
    /// * `property_id` - (Required) The property’s unique identifier, including the `prp_` prefix.
    /// * `contact` - (Required) One or more email addresses to send activation status changes to.
    /// * `version` - (Required) The property version to activate. Previously this field was optional. It now depends on the `akamai.Property` resource to identify latest instead of calculating it locally.  This association helps keep the dependency tree properly aligned. To always use the latest version, enter this value `{resource}.{resource identifier}.{field name}`. Using the example code above, the entry would be `akamai_property.example.latest_version` since we want the value of the `latest_version` attribute in the `akamai.Property` resource labeled `example`.
    /// * `network` - (Optional) Akamai network to activate on, either `STAGING` or `PRODUCTION`. `STAGING` is the default.
    /// * `property` - (Deprecated) Replaced by `property_id`. Maintained for legacy purposes.
    /// 
    /// ## Attribute reference
    /// 
    /// The following attributes are returned:
    /// 
    /// * `id` - The unique identifier for this activation.
    /// * `warnings` - The contents of `warnings` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
    /// * `errors` - The contents of `errors` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
    /// * `activation_id` - The ID given to the activation event while it's in progress.
    /// * `status` - The property version’s activation status on the selected network.
    /// </summary>
    [AkamaiResourceType("akamai:index/propertyActivation:PropertyActivation")]
    public partial class PropertyActivation : Pulumi.CustomResource
    {
        [Output("activationId")]
        public Output<string> ActivationId { get; private set; } = null!;

        [Output("contacts")]
        public Output<ImmutableArray<string>> Contacts { get; private set; } = null!;

        [Output("errors")]
        public Output<string> Errors { get; private set; } = null!;

        [Output("network")]
        public Output<string?> Network { get; private set; } = null!;

        [Output("property")]
        public Output<string> Property { get; private set; } = null!;

        [Output("propertyId")]
        public Output<string> PropertyId { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;

        [Output("warnings")]
        public Output<string> Warnings { get; private set; } = null!;


        /// <summary>
        /// Create a PropertyActivation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PropertyActivation(string name, PropertyActivationArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/propertyActivation:PropertyActivation", name, args ?? new PropertyActivationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PropertyActivation(string name, Input<string> id, PropertyActivationState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/propertyActivation:PropertyActivation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "akamai:properties/propertyActivation:PropertyActivation"},
                },
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
        [Input("activationId")]
        public Input<string>? ActivationId { get; set; }

        [Input("contacts", required: true)]
        private InputList<string>? _contacts;
        public InputList<string> Contacts
        {
            get => _contacts ?? (_contacts = new InputList<string>());
            set => _contacts = value;
        }

        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("property")]
        public Input<string>? Property { get; set; }

        [Input("propertyId")]
        public Input<string>? PropertyId { get; set; }

        [Input("version", required: true)]
        public Input<int> Version { get; set; } = null!;

        public PropertyActivationArgs()
        {
        }
    }

    public sealed class PropertyActivationState : Pulumi.ResourceArgs
    {
        [Input("activationId")]
        public Input<string>? ActivationId { get; set; }

        [Input("contacts")]
        private InputList<string>? _contacts;
        public InputList<string> Contacts
        {
            get => _contacts ?? (_contacts = new InputList<string>());
            set => _contacts = value;
        }

        [Input("errors")]
        public Input<string>? Errors { get; set; }

        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("property")]
        public Input<string>? Property { get; set; }

        [Input("propertyId")]
        public Input<string>? PropertyId { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("version")]
        public Input<int>? Version { get; set; }

        [Input("warnings")]
        public Input<string>? Warnings { get; set; }

        public PropertyActivationState()
        {
        }
    }
}
