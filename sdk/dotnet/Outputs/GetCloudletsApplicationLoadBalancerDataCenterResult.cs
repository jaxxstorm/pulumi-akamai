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
    public sealed class GetCloudletsApplicationLoadBalancerDataCenterResult
    {
        public readonly string City;
        public readonly bool CloudServerHostHeaderOverride;
        public readonly bool CloudService;
        public readonly string Continent;
        public readonly string Country;
        public readonly string Hostname;
        public readonly double Latitude;
        public readonly ImmutableArray<string> LivenessHosts;
        public readonly double Longitude;
        /// <summary>
        /// - (Required) A unique identifier for the Conditional Origin that supports the load balancing configuration. The Conditional Origin type must be set to `APPLICATION_LOAD_BALANCER` in the `origin` behavior. See property rules for more information.
        /// </summary>
        public readonly string OriginId;
        public readonly double Percent;
        public readonly string StateOrProvince;

        [OutputConstructor]
        private GetCloudletsApplicationLoadBalancerDataCenterResult(
            string city,

            bool cloudServerHostHeaderOverride,

            bool cloudService,

            string continent,

            string country,

            string hostname,

            double latitude,

            ImmutableArray<string> livenessHosts,

            double longitude,

            string originId,

            double percent,

            string stateOrProvince)
        {
            City = city;
            CloudServerHostHeaderOverride = cloudServerHostHeaderOverride;
            CloudService = cloudService;
            Continent = continent;
            Country = country;
            Hostname = hostname;
            Latitude = latitude;
            LivenessHosts = livenessHosts;
            Longitude = longitude;
            OriginId = originId;
            Percent = percent;
            StateOrProvince = stateOrProvince;
        }
    }
}
