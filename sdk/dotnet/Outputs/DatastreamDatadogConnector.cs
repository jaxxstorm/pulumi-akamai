// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Outputs
{

    [OutputType]
    public sealed class DatastreamDatadogConnector
    {
        /// <summary>
        /// - (Required) **Secret**. The API key associated with your Datadog account. See [View API keys in Datadog](https://docs.datadoghq.com/account_management/api-app-keys/#api-keys).
        /// * `compress logs` - (Optional) Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `false`.
        /// </summary>
        public readonly string AuthToken;
        /// <summary>
        /// Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `true`.
        /// </summary>
        public readonly bool? CompressLogs;
        public readonly int? ConnectorId;
        /// <summary>
        /// The name of the connector.
        /// </summary>
        public readonly string ConnectorName;
        /// <summary>
        /// The service of the Datadog connector. A service groups together endpoints, queries, or jobs for the purposes of scaling instances. See [View Datadog reserved attribute list](https://docs.datadoghq.com/logs/log_configuration/attributes_naming_convention/#reserved-attributes).
        /// </summary>
        public readonly string? Service;
        /// <summary>
        /// The source of the Datadog connector. See [View Datadog reserved attribute list](https://docs.datadoghq.com/logs/log_collection/?tab=http#reserved-attributes).
        /// </summary>
        public readonly string? Source;
        /// <summary>
        /// The tags of the Datadog connector. See [View Datadog tags](https://docs.datadoghq.com/getting_started/tagging/).
        /// </summary>
        public readonly string? Tags;
        /// <summary>
        /// Enter the secure URL where you want to send and store your logs.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private DatastreamDatadogConnector(
            string authToken,

            bool? compressLogs,

            int? connectorId,

            string connectorName,

            string? service,

            string? source,

            string? tags,

            string url)
        {
            AuthToken = authToken;
            CompressLogs = compressLogs;
            ConnectorId = connectorId;
            ConnectorName = connectorName;
            Service = service;
            Source = source;
            Tags = tags;
            Url = url;
        }
    }
}