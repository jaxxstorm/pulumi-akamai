// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetIamCountries
    {
        /// <summary>
        /// Use `akamai.getIamCountries` to retrieve all the possible countries that Akamai supports. Use the values from this data source to add or update a user's country information.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic usage:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var countries = Output.Create(Akamai.GetIamCountries.InvokeAsync());
        ///         this.SupportedCountries = countries;
        ///     }
        /// 
        ///     [Output("supportedCountries")]
        ///     public Output&lt;string&gt; SupportedCountries { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes reference
        /// 
        /// These attributes are returned:
        /// 
        /// * `countries` — A list of countries.
        /// 
        /// [API Reference](https://techdocs.akamai.com/iam-api/reference/get-common-countries)
        /// </summary>
        public static Task<GetIamCountriesResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIamCountriesResult>("akamai:index/getIamCountries:getIamCountries", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetIamCountriesResult
    {
        public readonly ImmutableArray<string> Countries;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetIamCountriesResult(
            ImmutableArray<string> countries,

            string id)
        {
            Countries = countries;
            Id = id;
        }
    }
}
