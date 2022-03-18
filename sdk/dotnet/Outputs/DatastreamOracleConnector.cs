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
    public sealed class DatastreamOracleConnector
    {
        /// <summary>
        /// **Secret**. The access key identifier that you use to authenticate requests to your Oracle Cloud account. See [Managing user credentials in OCS](https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm).
        /// </summary>
        public readonly string AccessKey;
        /// <summary>
        /// The name of the Oracle Cloud Storage bucket. See [Working with Oracle Cloud Storage buckets](https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/managingbuckets.htm).
        /// </summary>
        public readonly string Bucket;
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
        /// The namespace of your Oracle Cloud Storage account. See [Understanding Object Storage namespaces](https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/understandingnamespaces.htm).
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// The path to the folder within your Oracle Cloud Storage bucket where you want to store your logs.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The Oracle Cloud Storage region where your bucket resides. See [Regions and availability domains in OCS](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm).
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// **Secret**. The secret access key identifier that you use to authenticate requests to your Oracle Cloud account.
        /// </summary>
        public readonly string SecretAccessKey;

        [OutputConstructor]
        private DatastreamOracleConnector(
            string accessKey,

            string bucket,

            bool? compressLogs,

            int? connectorId,

            string connectorName,

            string @namespace,

            string path,

            string region,

            string secretAccessKey)
        {
            AccessKey = accessKey;
            Bucket = bucket;
            CompressLogs = compressLogs;
            ConnectorId = connectorId;
            ConnectorName = connectorName;
            Namespace = @namespace;
            Path = path;
            Region = region;
            SecretAccessKey = secretAccessKey;
        }
    }
}
