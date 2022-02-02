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
    public sealed class CpsDvEnrollmentTechContact
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
        /// The email address of the technical contact at Akamai, accessible at the `akamai.com` domain.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The first name of the technical contact at Akamai.
        /// </summary>
        public readonly string FirstName;
        /// <summary>
        /// The last name of the technical contact at Akamai.
        /// </summary>
        public readonly string LastName;
        /// <summary>
        /// Your organization information.
        /// </summary>
        public readonly string Organization;
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
        /// <summary>
        /// The title of the technical contact at Akamai.
        /// </summary>
        public readonly string? Title;

        [OutputConstructor]
        private CpsDvEnrollmentTechContact(
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
