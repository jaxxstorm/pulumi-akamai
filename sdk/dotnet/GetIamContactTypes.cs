// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetIamContactTypes
    {
        /// <summary>
        /// Use `akamai.getIamContactTypes` to retrieve all the possible `contact_types` that Akamai supports. Use the values from this data source to add or update a user's contact type.
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
        ///         var contactTypes = Output.Create(Akamai.GetIamContactTypes.InvokeAsync());
        ///         this.SupportedContactTypes = contactTypes;
        ///     }
        /// 
        ///     [Output("supportedContactTypes")]
        ///     public Output&lt;string&gt; SupportedContactTypes { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes reference
        /// 
        /// These attributes are returned:
        /// 
        /// * `contact_types` — A list of contact types.
        /// 
        /// [API Reference](https://techdocs.akamai.com/iam-api/reference/get-user-contact-types)
        /// </summary>
        public static Task<GetIamContactTypesResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIamContactTypesResult>("akamai:index/getIamContactTypes:getIamContactTypes", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetIamContactTypesResult
    {
        public readonly ImmutableArray<string> ContactTypes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetIamContactTypesResult(
            ImmutableArray<string> contactTypes,

            string id)
        {
            ContactTypes = contactTypes;
            Id = id;
        }
    }
}
