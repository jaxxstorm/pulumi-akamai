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
    public sealed class GetDatastreamActivationHistoryActivationResult
    {
        public readonly string CreatedBy;
        public readonly string CreatedDate;
        public readonly bool IsActive;
        /// <summary>
        /// - (Required) A stream's unique identifier.
        /// </summary>
        public readonly int StreamId;
        public readonly int StreamVersionId;

        [OutputConstructor]
        private GetDatastreamActivationHistoryActivationResult(
            string createdBy,

            string createdDate,

            bool isActive,

            int streamId,

            int streamVersionId)
        {
            CreatedBy = createdBy;
            CreatedDate = createdDate;
            IsActive = isActive;
            StreamId = streamId;
            StreamVersionId = streamVersionId;
        }
    }
}