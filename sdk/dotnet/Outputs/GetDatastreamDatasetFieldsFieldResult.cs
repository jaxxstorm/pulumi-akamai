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
    public sealed class GetDatastreamDatasetFieldsFieldResult
    {
        public readonly ImmutableArray<Outputs.GetDatastreamDatasetFieldsFieldDatasetFieldResult> DatasetFields;
        public readonly string DatasetGroupDescription;
        public readonly string DatasetGroupName;

        [OutputConstructor]
        private GetDatastreamDatasetFieldsFieldResult(
            ImmutableArray<Outputs.GetDatastreamDatasetFieldsFieldDatasetFieldResult> datasetFields,

            string datasetGroupDescription,

            string datasetGroupName)
        {
            DatasetFields = datasetFields;
            DatasetGroupDescription = datasetGroupDescription;
            DatasetGroupName = datasetGroupName;
        }
    }
}