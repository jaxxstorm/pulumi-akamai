// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    [AkamaiResourceType("akamai:index/appSecActivations:AppSecActivations")]
    public partial class AppSecActivations : Pulumi.CustomResource
    {
        /// <summary>
        /// . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
        /// </summary>
        [Output("activate")]
        public Output<bool?> Activate { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
        /// * **PRODUCTION**
        /// * **STAGING**
        /// </summary>
        [Output("network")]
        public Output<string?> Network { get; private set; } = null!;

        /// <summary>
        /// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
        /// </summary>
        [Output("note")]
        public Output<string?> Note { get; private set; } = null!;

        /// <summary>
        /// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        /// <summary>
        /// . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
        /// </summary>
        [Output("notificationEmails")]
        public Output<ImmutableArray<string>> NotificationEmails { get; private set; } = null!;

        /// <summary>
        /// The results of the activation
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecActivations resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecActivations(string name, AppSecActivationsArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecActivations:AppSecActivations", name, args ?? new AppSecActivationsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecActivations(string name, Input<string> id, AppSecActivationsState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecActivations:AppSecActivations", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecActivations resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecActivations Get(string name, Input<string> id, AppSecActivationsState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecActivations(name, id, state, options);
        }
    }

    public sealed class AppSecActivationsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
        /// </summary>
        [Input("activate")]
        public Input<bool>? Activate { get; set; }

        /// <summary>
        /// . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
        /// * **PRODUCTION**
        /// * **STAGING**
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
        /// </summary>
        [Input("note")]
        public Input<string>? Note { get; set; }

        /// <summary>
        /// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("notificationEmails", required: true)]
        private InputList<string>? _notificationEmails;

        /// <summary>
        /// . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
        /// </summary>
        public InputList<string> NotificationEmails
        {
            get => _notificationEmails ?? (_notificationEmails = new InputList<string>());
            set => _notificationEmails = value;
        }

        /// <summary>
        /// . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
        /// </summary>
        [Input("version", required: true)]
        public Input<int> Version { get; set; } = null!;

        public AppSecActivationsArgs()
        {
        }
    }

    public sealed class AppSecActivationsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
        /// </summary>
        [Input("activate")]
        public Input<bool>? Activate { get; set; }

        /// <summary>
        /// . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
        /// * **PRODUCTION**
        /// * **STAGING**
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
        /// </summary>
        [Input("note")]
        public Input<string>? Note { get; set; }

        /// <summary>
        /// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("notificationEmails")]
        private InputList<string>? _notificationEmails;

        /// <summary>
        /// . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
        /// </summary>
        public InputList<string> NotificationEmails
        {
            get => _notificationEmails ?? (_notificationEmails = new InputList<string>());
            set => _notificationEmails = value;
        }

        /// <summary>
        /// The results of the activation
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public AppSecActivationsState()
        {
        }
    }
}
