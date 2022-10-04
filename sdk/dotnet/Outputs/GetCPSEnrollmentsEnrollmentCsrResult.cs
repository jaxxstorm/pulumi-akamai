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
    public sealed class GetCPSEnrollmentsEnrollmentCsrResult
    {
        public readonly string City;
        public readonly string CountryCode;
        public readonly string Organization;
        public readonly string OrganizationalUnit;
        public readonly string State;

        [OutputConstructor]
        private GetCPSEnrollmentsEnrollmentCsrResult(
            string city,

            string countryCode,

            string organization,

            string organizationalUnit,

            string state)
        {
            City = city;
            CountryCode = countryCode;
            Organization = organization;
            OrganizationalUnit = organizationalUnit;
            State = state;
        }
    }
}
