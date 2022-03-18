// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class DatastreamDatadogConnectorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// - (Required) **Secret**. The API key associated with your Datadog account. See [View API keys in Datadog](https://docs.datadoghq.com/account_management/api-app-keys/#api-keys).
        /// * `compress logs` - (Optional) Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `false`.
        /// </summary>
        [Input("authToken", required: true)]
        public Input<string> AuthToken { get; set; } = null!;

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
        /// The service of the Datadog connector. A service groups together endpoints, queries, or jobs for the purposes of scaling instances. See [View Datadog reserved attribute list](https://docs.datadoghq.com/logs/log_configuration/attributes_naming_convention/#reserved-attributes).
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// The source of the Datadog connector. See [View Datadog reserved attribute list](https://docs.datadoghq.com/logs/log_collection/?tab=http#reserved-attributes).
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// The tags of the Datadog connector. See [View Datadog tags](https://docs.datadoghq.com/getting_started/tagging/).
        /// </summary>
        [Input("tags")]
        public Input<string>? Tags { get; set; }

        /// <summary>
        /// Enter the secure URL where you want to send and store your logs.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public DatastreamDatadogConnectorArgs()
        {
        }
    }
}
