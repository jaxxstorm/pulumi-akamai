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
    public sealed class CpsDvEnrollmentOrganization
    {
        /// <summary>
        /// The address of your organization.
        /// </summary>
        public readonly string AddressLineOne;
        /// <summary>
        /// The address of your organization.
        /// </summary>
        public readonly string? AddressLineTwo;
        /// <summary>
        /// The city where your organization resides.
        /// </summary>
        public readonly string City;
        /// <summary>
        /// The code for the country where your organization resides.
        /// </summary>
        public readonly string CountryCode;
        /// <summary>
        /// The name of your organization.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The phone number of the administrator who you want to use as a contact at your company.
        /// </summary>
        public readonly string Phone;
        /// <summary>
        /// The postal code of your organization.
        /// </summary>
        public readonly string PostalCode;
        /// <summary>
        /// The region of your organization, typically a state or province.
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private CpsDvEnrollmentOrganization(
            string addressLineOne,

            string? addressLineTwo,

            string city,

            string countryCode,

            string name,

            string phone,

            string postalCode,

            string region)
        {
            AddressLineOne = addressLineOne;
            AddressLineTwo = addressLineTwo;
            City = city;
            CountryCode = countryCode;
            Name = name;
            Phone = phone;
            PostalCode = postalCode;
            Region = region;
        }
    }
}
