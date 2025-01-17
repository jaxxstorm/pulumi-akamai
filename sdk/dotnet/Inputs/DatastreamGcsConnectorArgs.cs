// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class DatastreamGcsConnectorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Oracle Cloud Storage bucket. See [Working with Oracle Cloud Storage buckets](https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/managingbuckets.htm).
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

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
        /// The path to the folder within your Oracle Cloud Storage bucket where you want to store your logs.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// **Secret**. The contents of the JSON private key you generated and downloaded in your Google Cloud Storage account.
        /// </summary>
        [Input("privateKey", required: true)]
        public Input<string> PrivateKey { get; set; } = null!;

        /// <summary>
        /// The unique ID of your Google Cloud project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The name of the service account with the storage.object.create permission or Storage Object Creator role.
        /// </summary>
        [Input("serviceAccountName", required: true)]
        public Input<string> ServiceAccountName { get; set; } = null!;

        public DatastreamGcsConnectorArgs()
        {
        }
    }
}
