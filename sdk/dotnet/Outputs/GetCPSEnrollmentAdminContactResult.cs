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
    public sealed class GetCPSEnrollmentAdminContactResult
    {
        public readonly string AddressLineOne;
        public readonly string? AddressLineTwo;
        public readonly string City;
        public readonly string CountryCode;
        public readonly string Email;
        public readonly string FirstName;
        public readonly string LastName;
        public readonly string Organization;
        public readonly string Phone;
        public readonly string PostalCode;
        public readonly string Region;
        public readonly string? Title;

        [OutputConstructor]
        private GetCPSEnrollmentAdminContactResult(
            string addressLineOne,

            string? addressLineTwo,

            string city,

            string countryCode,

            string email,

            string firstName,

            string lastName,

            string organization,

            string phone,

            string postalCode,

            string region,

            string? title)
        {
            AddressLineOne = addressLineOne;
            AddressLineTwo = addressLineTwo;
            City = city;
            CountryCode = countryCode;
            Email = email;
            FirstName = firstName;
            LastName = lastName;
            Organization = organization;
            Phone = phone;
            PostalCode = postalCode;
            Region = region;
            Title = title;
        }
    }
}
