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
    public sealed class DatastreamSumologicConnector
    {
        /// <summary>
        /// **Secret**. The unique HTTP collector code of your Sumo Logic `endpoint`.
        /// </summary>
        public readonly string CollectorCode;
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
        /// The Sumo Logic collection endpoint where you want to send your logs. You should follow the `https://&lt;SumoEndpoint&gt;/receiver/v1/http` format and pass the collector code in the `collectorCode` argument.
        /// </summary>
        public readonly string Endpoint;

        [OutputConstructor]
        private DatastreamSumologicConnector(
            string collectorCode,

            bool? compressLogs,

            int? connectorId,

            string connectorName,

            string endpoint)
        {
            CollectorCode = collectorCode;
            CompressLogs = compressLogs;
            ConnectorId = connectorId;
            ConnectorName = connectorName;
            Endpoint = endpoint;
        }
    }
}
