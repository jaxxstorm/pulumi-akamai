// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `akamai.EdgeHostName` resource lets you configure a secure edge hostname. Your
 * edge hostname determines how requests for your site, app, or content are mapped to
 * Akamai edge servers.
 *
 * An edge hostname is the CNAME target you use when directing your end user traffic to
 * Akamai. Each hostname assigned to a property has a corresponding edge hostname.
 *
 * Akamai supports three types of edge hostnames, depending on the level of security
 * you need for your traffic: Standard TLS, Enhanced TLS, and Shared Certificate. When
 * entering the `edgeHostname` attribute, you need to include a specific domain suffix
 * for your edge hostname type:
 *
 * | Edge hostname type | Domain suffix |
 * |------|-------|
 * | Enhanced TLS | edgekey.net |
 * | Standard TLS | edgesuite.net |
 * | Shared Cert | akamaized.net |
 *
 * For example, if you use Standard TLS and have `www.example.com` as a hostname, your edge hostname would be `www.example.com.edgesuite.net`. If you wanted to use Enhanced TLS with the same hostname, your edge hostname would be `www.example.com.edgekey.net`. See the [Property Manager API (PAPI)](https://developer.akamai.com/api/core_features/property_manager/v1.html#createedgehostnames) for more information.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const provider_demo = new akamai.EdgeHostName("provider-demo", {
 *     contractId: "ctr_1-AB123",
 *     edgeHostname: "www.example.org.edgesuite.net",
 *     groupId: "grp_123",
 *     productId: "prd_Object_Delivery",
 * });
 * ```
 * ## Attributes reference
 *
 * This resource returns this attribute:
 *
 * * `ipBehavior` - Returns the IP protocol the hostname will use, either `IPV4` for version 4, IPV6_PERFORMANCE`for version 6, or`IPV6_COMPLIANCE` for both.
 *
 * ## Import
 *
 * Basic Usagehcl resource "akamai_edge_hostname" "example" {
 *
 * # (resource arguments)
 *
 *  } You can import Akamai edge hostnames using a comma-delimited string of edge hostname, contract ID, and group ID. You have to enter the values in this order:
 *
 * `edge_hostname, contract_id, group_id` For example
 *
 * ```sh
 *  $ pulumi import akamai:index/edgeHostName:EdgeHostName example ehn_123,ctr_1-AB123,grp_123
 * ```
 */
export class EdgeHostName extends pulumi.CustomResource {
    /**
     * Get an existing EdgeHostName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EdgeHostNameState, opts?: pulumi.CustomResourceOptions): EdgeHostName {
        return new EdgeHostName(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/edgeHostName:EdgeHostName';

    /**
     * Returns true if the given object is an instance of EdgeHostName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EdgeHostName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EdgeHostName.__pulumiType;
    }

    /**
     * Required only when creating an Enhanced TLS edge hostname. This argument sets the certificate enrollment ID. Edge hostnames for Enhanced TLS end in `edgekey.net`. You can retrieve this ID from the [Certificate Provisioning Service CLI](https://github.com/akamai/cli-cps) .
     */
    public readonly certificate!: pulumi.Output<number | undefined>;
    /**
     * Replaced by `contractId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "contract" has been deprecated.
     */
    public readonly contract!: pulumi.Output<string>;
    /**
     * - (Required) A contract's unique ID, including the `ctr_` prefix.
     */
    public readonly contractId!: pulumi.Output<string>;
    /**
     * One or more edge hostnames. The number of edge hostnames must be less than or equal to the number of public hostnames.
     */
    public readonly edgeHostname!: pulumi.Output<string>;
    /**
     * Replaced by `groupId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "group" has been deprecated.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * - (Required) A group's unique ID, including the `grp_` prefix.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * Which version of the IP protocol to use: `IPV4` for version 4 only, `IPV6_PERFORMANCE` for version 6 only, or `IPV6_COMPLIANCE` for both 4 and 6.
     */
    public readonly ipBehavior!: pulumi.Output<string>;
    /**
     * Replaced by `productId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "product" has been deprecated.
     */
    public readonly product!: pulumi.Output<string>;
    /**
     * - (Required) A product's unique ID, including the `prd_` prefix.
     */
    public readonly productId!: pulumi.Output<string>;
    /**
     * A JSON encoded list of use cases
     */
    public readonly useCases!: pulumi.Output<string | undefined>;

    /**
     * Create a EdgeHostName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EdgeHostNameArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EdgeHostNameArgs | EdgeHostNameState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EdgeHostNameState | undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["contract"] = state ? state.contract : undefined;
            resourceInputs["contractId"] = state ? state.contractId : undefined;
            resourceInputs["edgeHostname"] = state ? state.edgeHostname : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["ipBehavior"] = state ? state.ipBehavior : undefined;
            resourceInputs["product"] = state ? state.product : undefined;
            resourceInputs["productId"] = state ? state.productId : undefined;
            resourceInputs["useCases"] = state ? state.useCases : undefined;
        } else {
            const args = argsOrState as EdgeHostNameArgs | undefined;
            if ((!args || args.edgeHostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edgeHostname'");
            }
            if ((!args || args.ipBehavior === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipBehavior'");
            }
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["contract"] = args ? args.contract : undefined;
            resourceInputs["contractId"] = args ? args.contractId : undefined;
            resourceInputs["edgeHostname"] = args ? args.edgeHostname : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["ipBehavior"] = args ? args.ipBehavior : undefined;
            resourceInputs["product"] = args ? args.product : undefined;
            resourceInputs["productId"] = args ? args.productId : undefined;
            resourceInputs["useCases"] = args ? args.useCases : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "akamai:properties/edgeHostName:EdgeHostName" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(EdgeHostName.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EdgeHostName resources.
 */
export interface EdgeHostNameState {
    /**
     * Required only when creating an Enhanced TLS edge hostname. This argument sets the certificate enrollment ID. Edge hostnames for Enhanced TLS end in `edgekey.net`. You can retrieve this ID from the [Certificate Provisioning Service CLI](https://github.com/akamai/cli-cps) .
     */
    certificate?: pulumi.Input<number>;
    /**
     * Replaced by `contractId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "contract" has been deprecated.
     */
    contract?: pulumi.Input<string>;
    /**
     * - (Required) A contract's unique ID, including the `ctr_` prefix.
     */
    contractId?: pulumi.Input<string>;
    /**
     * One or more edge hostnames. The number of edge hostnames must be less than or equal to the number of public hostnames.
     */
    edgeHostname?: pulumi.Input<string>;
    /**
     * Replaced by `groupId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "group" has been deprecated.
     */
    group?: pulumi.Input<string>;
    /**
     * - (Required) A group's unique ID, including the `grp_` prefix.
     */
    groupId?: pulumi.Input<string>;
    /**
     * Which version of the IP protocol to use: `IPV4` for version 4 only, `IPV6_PERFORMANCE` for version 6 only, or `IPV6_COMPLIANCE` for both 4 and 6.
     */
    ipBehavior?: pulumi.Input<string>;
    /**
     * Replaced by `productId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "product" has been deprecated.
     */
    product?: pulumi.Input<string>;
    /**
     * - (Required) A product's unique ID, including the `prd_` prefix.
     */
    productId?: pulumi.Input<string>;
    /**
     * A JSON encoded list of use cases
     */
    useCases?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EdgeHostName resource.
 */
export interface EdgeHostNameArgs {
    /**
     * Required only when creating an Enhanced TLS edge hostname. This argument sets the certificate enrollment ID. Edge hostnames for Enhanced TLS end in `edgekey.net`. You can retrieve this ID from the [Certificate Provisioning Service CLI](https://github.com/akamai/cli-cps) .
     */
    certificate?: pulumi.Input<number>;
    /**
     * Replaced by `contractId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "contract" has been deprecated.
     */
    contract?: pulumi.Input<string>;
    /**
     * - (Required) A contract's unique ID, including the `ctr_` prefix.
     */
    contractId?: pulumi.Input<string>;
    /**
     * One or more edge hostnames. The number of edge hostnames must be less than or equal to the number of public hostnames.
     */
    edgeHostname: pulumi.Input<string>;
    /**
     * Replaced by `groupId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "group" has been deprecated.
     */
    group?: pulumi.Input<string>;
    /**
     * - (Required) A group's unique ID, including the `grp_` prefix.
     */
    groupId?: pulumi.Input<string>;
    /**
     * Which version of the IP protocol to use: `IPV4` for version 4 only, `IPV6_PERFORMANCE` for version 6 only, or `IPV6_COMPLIANCE` for both 4 and 6.
     */
    ipBehavior: pulumi.Input<string>;
    /**
     * Replaced by `productId`. Maintained for legacy purposes.
     *
     * @deprecated The setting "product" has been deprecated.
     */
    product?: pulumi.Input<string>;
    /**
     * - (Required) A product's unique ID, including the `prd_` prefix.
     */
    productId?: pulumi.Input<string>;
    /**
     * A JSON encoded list of use cases
     */
    useCases?: pulumi.Input<string>;
}
