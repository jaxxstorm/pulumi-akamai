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
    public sealed class GetContractsContractResult
    {
        public readonly string ContractId;
        public readonly string ContractTypeName;

        [OutputConstructor]
        private GetContractsContractResult(
            string contractId,

            string contractTypeName)
        {
            ContractId = contractId;
            ContractTypeName = contractTypeName;
        }
    }
}