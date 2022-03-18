// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class DatastreamHttpsConnectorGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Either `NONE` for no authentication, or `BASIC`. For basic authentication, provide the `user_name` and `password` you set in your custom HTTPS endpoint.
        /// </summary>
        [Input("authenticationType", required: true)]
        public Input<string> AuthenticationType { get; set; } = null!;

        /// <summary>
        /// Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `true`.
        /// </summary>
        [Input("compressLogs")]
        public Input<bool>? CompressLogs { get; set; }

        [Input("connectorId")]
        public Input<int>? ConnectorId { get; set; }

        /// <summary>
        /// The name of the connector.
        /// </summary>
        [Input("connectorName", required: true)]
        public Input<string> ConnectorName { get; set; } = null!;

        /// <summary>
        /// **Secret**. Enter the password you set in your custom HTTPS endpoint for authentication.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Enter the secure URL where you want to send and store your logs.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// **Secret**. Enter the valid username you set in your custom HTTPS endpoint for authentication.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public DatastreamHttpsConnectorGetArgs()
        {
        }
    }
}
