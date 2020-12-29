// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Properties
{
    [Obsolete(@"akamai.properties.Property has been deprecated in favor of akamai.Property")]
    public partial class Property : Pulumi.CustomResource
    {
        [Output("contacts")]
        public Output<ImmutableArray<string>> Contacts { get; private set; } = null!;

        [Output("contract")]
        public Output<string> Contract { get; private set; } = null!;

        /// <summary>
        /// Contract ID to be assigned to the Property
        /// </summary>
        [Output("contractId")]
        public Output<string> ContractId { get; private set; } = null!;

        [Output("cpCode")]
        public Output<string?> CpCode { get; private set; } = null!;

        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// Group ID to be assigned to the Property
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// Mapping of edge hostname CNAMEs to other CNAMEs
        /// </summary>
        [Output("hostnames")]
        public Output<ImmutableDictionary<string, string>?> Hostnames { get; private set; } = null!;

        [Output("isSecure")]
        public Output<bool?> IsSecure { get; private set; } = null!;

        /// <summary>
        /// Property's current latest version number
        /// </summary>
        [Output("latestVersion")]
        public Output<int> LatestVersion { get; private set; } = null!;

        /// <summary>
        /// Name to give to the Property (must be unique)
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("origins")]
        public Output<ImmutableArray<Outputs.PropertyOrigin>> Origins { get; private set; } = null!;

        [Output("product")]
        public Output<string> Product { get; private set; } = null!;

        /// <summary>
        /// Product ID to be assigned to the Property
        /// </summary>
        [Output("productId")]
        public Output<string> ProductId { get; private set; } = null!;

        /// <summary>
        /// Property's version currently activated in production (zero when not active in production)
        /// </summary>
        [Output("productionVersion")]
        public Output<int> ProductionVersion { get; private set; } = null!;

        [Output("ruleErrors")]
        public Output<ImmutableArray<Outputs.PropertyRuleError>> RuleErrors { get; private set; } = null!;

        /// <summary>
        /// Specify the rule format version (defaults to latest version available when created)
        /// </summary>
        [Output("ruleFormat")]
        public Output<string> RuleFormat { get; private set; } = null!;

        [Output("ruleWarnings")]
        public Output<ImmutableArray<Outputs.PropertyRuleWarning>> RuleWarnings { get; private set; } = null!;

        /// <summary>
        /// Property Rules as JSON
        /// </summary>
        [Output("rules")]
        public Output<string> Rules { get; private set; } = null!;

        /// <summary>
        /// Property's version currently activated in staging (zero when not active in staging)
        /// </summary>
        [Output("stagingVersion")]
        public Output<int> StagingVersion { get; private set; } = null!;

        [Output("variables")]
        public Output<string?> Variables { get; private set; } = null!;


        /// <summary>
        /// Create a Property resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Property(string name, PropertyArgs? args = null, CustomResourceOptions? options = null)
            : base("akamai:properties/property:Property", name, args ?? new PropertyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Property(string name, Input<string> id, PropertyState? state = null, CustomResourceOptions? options = null)
            : base("akamai:properties/property:Property", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Property resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Property Get(string name, Input<string> id, PropertyState? state = null, CustomResourceOptions? options = null)
        {
            return new Property(name, id, state, options);
        }
    }

    public sealed class PropertyArgs : Pulumi.ResourceArgs
    {
        [Input("contacts")]
        private InputList<string>? _contacts;
        [Obsolete(@"""contact"" is no longer supported by this resource type - See Akamai Terraform Upgrade Guide")]
        public InputList<string> Contacts
        {
            get => _contacts ?? (_contacts = new InputList<string>());
            set => _contacts = value;
        }

        [Input("contract")]
        public Input<string>? Contract { get; set; }

        /// <summary>
        /// Contract ID to be assigned to the Property
        /// </summary>
        [Input("contractId")]
        public Input<string>? ContractId { get; set; }

        [Input("cpCode")]
        public Input<string>? CpCode { get; set; }

        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Group ID to be assigned to the Property
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("hostnames")]
        private InputMap<string>? _hostnames;

        /// <summary>
        /// Mapping of edge hostname CNAMEs to other CNAMEs
        /// </summary>
        public InputMap<string> Hostnames
        {
            get => _hostnames ?? (_hostnames = new InputMap<string>());
            set => _hostnames = value;
        }

        [Input("isSecure")]
        public Input<bool>? IsSecure { get; set; }

        /// <summary>
        /// Name to give to the Property (must be unique)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("origins")]
        private InputList<Inputs.PropertyOriginArgs>? _origins;
        [Obsolete(@"""origin"" is no longer supported by this resource type - See Akamai Terraform Upgrade Guide")]
        public InputList<Inputs.PropertyOriginArgs> Origins
        {
            get => _origins ?? (_origins = new InputList<Inputs.PropertyOriginArgs>());
            set => _origins = value;
        }

        [Input("product")]
        public Input<string>? Product { get; set; }

        /// <summary>
        /// Product ID to be assigned to the Property
        /// </summary>
        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        /// <summary>
        /// Specify the rule format version (defaults to latest version available when created)
        /// </summary>
        [Input("ruleFormat")]
        public Input<string>? RuleFormat { get; set; }

        /// <summary>
        /// Property Rules as JSON
        /// </summary>
        [Input("rules")]
        public Input<string>? Rules { get; set; }

        [Input("variables")]
        public Input<string>? Variables { get; set; }

        public PropertyArgs()
        {
        }
    }

    public sealed class PropertyState : Pulumi.ResourceArgs
    {
        [Input("contacts")]
        private InputList<string>? _contacts;
        [Obsolete(@"""contact"" is no longer supported by this resource type - See Akamai Terraform Upgrade Guide")]
        public InputList<string> Contacts
        {
            get => _contacts ?? (_contacts = new InputList<string>());
            set => _contacts = value;
        }

        [Input("contract")]
        public Input<string>? Contract { get; set; }

        /// <summary>
        /// Contract ID to be assigned to the Property
        /// </summary>
        [Input("contractId")]
        public Input<string>? ContractId { get; set; }

        [Input("cpCode")]
        public Input<string>? CpCode { get; set; }

        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Group ID to be assigned to the Property
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("hostnames")]
        private InputMap<string>? _hostnames;

        /// <summary>
        /// Mapping of edge hostname CNAMEs to other CNAMEs
        /// </summary>
        public InputMap<string> Hostnames
        {
            get => _hostnames ?? (_hostnames = new InputMap<string>());
            set => _hostnames = value;
        }

        [Input("isSecure")]
        public Input<bool>? IsSecure { get; set; }

        /// <summary>
        /// Property's current latest version number
        /// </summary>
        [Input("latestVersion")]
        public Input<int>? LatestVersion { get; set; }

        /// <summary>
        /// Name to give to the Property (must be unique)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("origins")]
        private InputList<Inputs.PropertyOriginGetArgs>? _origins;
        [Obsolete(@"""origin"" is no longer supported by this resource type - See Akamai Terraform Upgrade Guide")]
        public InputList<Inputs.PropertyOriginGetArgs> Origins
        {
            get => _origins ?? (_origins = new InputList<Inputs.PropertyOriginGetArgs>());
            set => _origins = value;
        }

        [Input("product")]
        public Input<string>? Product { get; set; }

        /// <summary>
        /// Product ID to be assigned to the Property
        /// </summary>
        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        /// <summary>
        /// Property's version currently activated in production (zero when not active in production)
        /// </summary>
        [Input("productionVersion")]
        public Input<int>? ProductionVersion { get; set; }

        [Input("ruleErrors")]
        private InputList<Inputs.PropertyRuleErrorGetArgs>? _ruleErrors;
        public InputList<Inputs.PropertyRuleErrorGetArgs> RuleErrors
        {
            get => _ruleErrors ?? (_ruleErrors = new InputList<Inputs.PropertyRuleErrorGetArgs>());
            set => _ruleErrors = value;
        }

        /// <summary>
        /// Specify the rule format version (defaults to latest version available when created)
        /// </summary>
        [Input("ruleFormat")]
        public Input<string>? RuleFormat { get; set; }

        [Input("ruleWarnings")]
        private InputList<Inputs.PropertyRuleWarningGetArgs>? _ruleWarnings;
        public InputList<Inputs.PropertyRuleWarningGetArgs> RuleWarnings
        {
            get => _ruleWarnings ?? (_ruleWarnings = new InputList<Inputs.PropertyRuleWarningGetArgs>());
            set => _ruleWarnings = value;
        }

        /// <summary>
        /// Property Rules as JSON
        /// </summary>
        [Input("rules")]
        public Input<string>? Rules { get; set; }

        /// <summary>
        /// Property's version currently activated in staging (zero when not active in staging)
        /// </summary>
        [Input("stagingVersion")]
        public Input<int>? StagingVersion { get; set; }

        [Input("variables")]
        public Input<string>? Variables { get; set; }

        public PropertyState()
        {
        }
    }
}
