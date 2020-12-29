// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * @deprecated akamai.trafficmanagement.GtmGeomap has been deprecated in favor of akamai.GtmGeomap
 */
export class GtmGeomap extends pulumi.CustomResource {
    /**
     * Get an existing GtmGeomap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GtmGeomapState, opts?: pulumi.CustomResourceOptions): GtmGeomap {
        pulumi.log.warn("GtmGeomap is deprecated: akamai.trafficmanagement.GtmGeomap has been deprecated in favor of akamai.GtmGeomap")
        return new GtmGeomap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:trafficmanagement/gtmGeomap:GtmGeomap';

    /**
     * Returns true if the given object is an instance of GtmGeomap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GtmGeomap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GtmGeomap.__pulumiType;
    }

    public readonly assignments!: pulumi.Output<outputs.trafficmanagement.GtmGeomapAssignment[] | undefined>;
    public readonly defaultDatacenter!: pulumi.Output<outputs.trafficmanagement.GtmGeomapDefaultDatacenter>;
    public readonly domain!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly waitOnComplete!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GtmGeomap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated akamai.trafficmanagement.GtmGeomap has been deprecated in favor of akamai.GtmGeomap */
    constructor(name: string, args: GtmGeomapArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated akamai.trafficmanagement.GtmGeomap has been deprecated in favor of akamai.GtmGeomap */
    constructor(name: string, argsOrState?: GtmGeomapArgs | GtmGeomapState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("GtmGeomap is deprecated: akamai.trafficmanagement.GtmGeomap has been deprecated in favor of akamai.GtmGeomap")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GtmGeomapState | undefined;
            inputs["assignments"] = state ? state.assignments : undefined;
            inputs["defaultDatacenter"] = state ? state.defaultDatacenter : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["waitOnComplete"] = state ? state.waitOnComplete : undefined;
        } else {
            const args = argsOrState as GtmGeomapArgs | undefined;
            if ((!args || args.defaultDatacenter === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'defaultDatacenter'");
            }
            if ((!args || args.domain === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'domain'");
            }
            inputs["assignments"] = args ? args.assignments : undefined;
            inputs["defaultDatacenter"] = args ? args.defaultDatacenter : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["waitOnComplete"] = args ? args.waitOnComplete : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GtmGeomap.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GtmGeomap resources.
 */
export interface GtmGeomapState {
    readonly assignments?: pulumi.Input<pulumi.Input<inputs.trafficmanagement.GtmGeomapAssignment>[]>;
    readonly defaultDatacenter?: pulumi.Input<inputs.trafficmanagement.GtmGeomapDefaultDatacenter>;
    readonly domain?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly waitOnComplete?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GtmGeomap resource.
 */
export interface GtmGeomapArgs {
    readonly assignments?: pulumi.Input<pulumi.Input<inputs.trafficmanagement.GtmGeomapAssignment>[]>;
    readonly defaultDatacenter: pulumi.Input<inputs.trafficmanagement.GtmGeomapDefaultDatacenter>;
    readonly domain: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly waitOnComplete?: pulumi.Input<boolean>;
}
