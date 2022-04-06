// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Properties.Inputs
{

    public sealed class PropertyHostnameGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate's provisioning type, either the default `CPS_MANAGED` type for the custom certificates you provision with the [Certificate Provisioning System (CPS)](https://learn.akamai.com/en-us/products/core_features/certificate_provisioning_system.html), or `DEFAULT` for certificates provisioned automatically.
        /// </summary>
        [Input("certProvisioningType", required: true)]
        public Input<string> CertProvisioningType { get; set; } = null!;

        [Input("certStatuses")]
        private InputList<Inputs.PropertyHostnameCertStatusGetArgs>? _certStatuses;
        public InputList<Inputs.PropertyHostnameCertStatusGetArgs> CertStatuses
        {
            get => _certStatuses ?? (_certStatuses = new InputList<Inputs.PropertyHostnameCertStatusGetArgs>());
            set => _certStatuses = value;
        }

        /// <summary>
        /// A string containing the original origin's hostname. For example, `"example.org"`.
        /// </summary>
        [Input("cnameFrom", required: true)]
        public Input<string> CnameFrom { get; set; } = null!;

        /// <summary>
        /// A string containing the hostname for edge content. For example,  `"example.org.edgesuite.net"`.
        /// </summary>
        [Input("cnameTo", required: true)]
        public Input<string> CnameTo { get; set; } = null!;

        [Input("cnameType")]
        public Input<string>? CnameType { get; set; }

        [Input("edgeHostnameId")]
        public Input<string>? EdgeHostnameId { get; set; }

        public PropertyHostnameGetArgs()
        {
        }
    }
}
