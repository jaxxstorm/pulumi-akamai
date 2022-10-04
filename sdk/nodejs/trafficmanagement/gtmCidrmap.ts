// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use the `akamai.GtmCidrmap` resource to create, configure, and import a GTM Classless Inter-Domain Routing (CIDR) map. CIDR mapping uses the IP addresses of the requesting name server to provide IP-specific CNAME entries. CNAMEs let you direct internal users to a specific environment or direct them to the origin. This lets you provide different responses to an internal corporate DNS infrastructure, such as internal test environments and another answer for all other name servers (`defaultDatacenter`).
 *
 *  CIDR maps split the Internet into multiple CIDR block zones. Properties that use a map can specify a handout CNAME for each zone on the property's editing page. To configure a property for CIDR mapping, your domain needs at least one CIDR map defined.
 *
 * > **Note** Import requires an ID with this format: `existingDomainName`:`existingMapName`.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const demoCidrmap = new akamai.GtmCidrmap("demo_cidrmap", {
 *     defaultDatacenter: {
 *         datacenterId: 5400,
 *         nickname: "All Other CIDR Blocks",
 *     },
 *     domain: "demo_domain.akadns.net",
 * });
 * ```
 *
 * @deprecated akamai.trafficmanagement.GtmCidrmap has been deprecated in favor of akamai.GtmCidrmap
 */
export class GtmCidrmap extends pulumi.CustomResource {
    /**
     * Get an existing GtmCidrmap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GtmCidrmapState, opts?: pulumi.CustomResourceOptions): GtmCidrmap {
        pulumi.log.warn("GtmCidrmap is deprecated: akamai.trafficmanagement.GtmCidrmap has been deprecated in favor of akamai.GtmCidrmap")
        return new GtmCidrmap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:trafficmanagement/gtmCidrmap:GtmCidrmap';

    /**
     * Returns true if the given object is an instance of GtmCidrmap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GtmCidrmap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GtmCidrmap.__pulumiType;
    }

    /**
     * Contains information about the CIDR zone groupings of CIDR blocks. You can have multiple entries with this argument. If used, requires these additional arguments:
     */
    public readonly assignments!: pulumi.Output<outputs.trafficmanagement.GtmCidrmapAssignment[] | undefined>;
    /**
     * A placeholder for all other CIDR zones not found in these CIDR zones. Requires these additional arguments:
     */
    public readonly defaultDatacenter!: pulumi.Output<outputs.trafficmanagement.GtmCidrmapDefaultDatacenter>;
    /**
     * GTM Domain name for the AS Map.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * A descriptive label for the CIDR map, up to 255 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A boolean that, if set to `true`, waits for transaction to complete.
     */
    public readonly waitOnComplete!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GtmCidrmap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated akamai.trafficmanagement.GtmCidrmap has been deprecated in favor of akamai.GtmCidrmap */
    constructor(name: string, args: GtmCidrmapArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated akamai.trafficmanagement.GtmCidrmap has been deprecated in favor of akamai.GtmCidrmap */
    constructor(name: string, argsOrState?: GtmCidrmapArgs | GtmCidrmapState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("GtmCidrmap is deprecated: akamai.trafficmanagement.GtmCidrmap has been deprecated in favor of akamai.GtmCidrmap")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GtmCidrmapState | undefined;
            resourceInputs["assignments"] = state ? state.assignments : undefined;
            resourceInputs["defaultDatacenter"] = state ? state.defaultDatacenter : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["waitOnComplete"] = state ? state.waitOnComplete : undefined;
        } else {
            const args = argsOrState as GtmCidrmapArgs | undefined;
            if ((!args || args.defaultDatacenter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultDatacenter'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            resourceInputs["assignments"] = args ? args.assignments : undefined;
            resourceInputs["defaultDatacenter"] = args ? args.defaultDatacenter : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["waitOnComplete"] = args ? args.waitOnComplete : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GtmCidrmap.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GtmCidrmap resources.
 */
export interface GtmCidrmapState {
    /**
     * Contains information about the CIDR zone groupings of CIDR blocks. You can have multiple entries with this argument. If used, requires these additional arguments:
     */
    assignments?: pulumi.Input<pulumi.Input<inputs.trafficmanagement.GtmCidrmapAssignment>[]>;
    /**
     * A placeholder for all other CIDR zones not found in these CIDR zones. Requires these additional arguments:
     */
    defaultDatacenter?: pulumi.Input<inputs.trafficmanagement.GtmCidrmapDefaultDatacenter>;
    /**
     * GTM Domain name for the AS Map.
     */
    domain?: pulumi.Input<string>;
    /**
     * A descriptive label for the CIDR map, up to 255 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * A boolean that, if set to `true`, waits for transaction to complete.
     */
    waitOnComplete?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GtmCidrmap resource.
 */
export interface GtmCidrmapArgs {
    /**
     * Contains information about the CIDR zone groupings of CIDR blocks. You can have multiple entries with this argument. If used, requires these additional arguments:
     */
    assignments?: pulumi.Input<pulumi.Input<inputs.trafficmanagement.GtmCidrmapAssignment>[]>;
    /**
     * A placeholder for all other CIDR zones not found in these CIDR zones. Requires these additional arguments:
     */
    defaultDatacenter: pulumi.Input<inputs.trafficmanagement.GtmCidrmapDefaultDatacenter>;
    /**
     * GTM Domain name for the AS Map.
     */
    domain: pulumi.Input<string>;
    /**
     * A descriptive label for the CIDR map, up to 255 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * A boolean that, if set to `true`, waits for transaction to complete.
     */
    waitOnComplete?: pulumi.Input<boolean>;
}
