// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class DnsZone extends pulumi.CustomResource {
    /**
     * Get an existing DnsZone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsZoneState, opts?: pulumi.CustomResourceOptions): DnsZone {
        return new DnsZone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/dnsZone:DnsZone';

    /**
     * Returns true if the given object is an instance of DnsZone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsZone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsZone.__pulumiType;
    }

    public /*out*/ readonly activationState!: pulumi.Output<string>;
    public /*out*/ readonly aliasCount!: pulumi.Output<number>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly contract!: pulumi.Output<string>;
    public readonly endCustomerId!: pulumi.Output<string | undefined>;
    public readonly group!: pulumi.Output<string | undefined>;
    public readonly masters!: pulumi.Output<string[] | undefined>;
    public readonly signAndServe!: pulumi.Output<boolean | undefined>;
    public readonly signAndServeAlgorithm!: pulumi.Output<string | undefined>;
    public readonly target!: pulumi.Output<string | undefined>;
    public readonly tsigKey!: pulumi.Output<outputs.DnsZoneTsigKey | undefined>;
    public readonly type!: pulumi.Output<string>;
    public /*out*/ readonly versionId!: pulumi.Output<string>;
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a DnsZone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DnsZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnsZoneArgs | DnsZoneState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsZoneState | undefined;
            inputs["activationState"] = state ? state.activationState : undefined;
            inputs["aliasCount"] = state ? state.aliasCount : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["contract"] = state ? state.contract : undefined;
            inputs["endCustomerId"] = state ? state.endCustomerId : undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["masters"] = state ? state.masters : undefined;
            inputs["signAndServe"] = state ? state.signAndServe : undefined;
            inputs["signAndServeAlgorithm"] = state ? state.signAndServeAlgorithm : undefined;
            inputs["target"] = state ? state.target : undefined;
            inputs["tsigKey"] = state ? state.tsigKey : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["versionId"] = state ? state.versionId : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as DnsZoneArgs | undefined;
            if ((!args || args.contract === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contract'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["comment"] = args ? args.comment : undefined;
            inputs["contract"] = args ? args.contract : undefined;
            inputs["endCustomerId"] = args ? args.endCustomerId : undefined;
            inputs["group"] = args ? args.group : undefined;
            inputs["masters"] = args ? args.masters : undefined;
            inputs["signAndServe"] = args ? args.signAndServe : undefined;
            inputs["signAndServeAlgorithm"] = args ? args.signAndServeAlgorithm : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["tsigKey"] = args ? args.tsigKey : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["activationState"] = undefined /*out*/;
            inputs["aliasCount"] = undefined /*out*/;
            inputs["versionId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "akamai:edgedns/dnsZone:DnsZone" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(DnsZone.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsZone resources.
 */
export interface DnsZoneState {
    activationState?: pulumi.Input<string>;
    aliasCount?: pulumi.Input<number>;
    comment?: pulumi.Input<string>;
    contract?: pulumi.Input<string>;
    endCustomerId?: pulumi.Input<string>;
    group?: pulumi.Input<string>;
    masters?: pulumi.Input<pulumi.Input<string>[]>;
    signAndServe?: pulumi.Input<boolean>;
    signAndServeAlgorithm?: pulumi.Input<string>;
    target?: pulumi.Input<string>;
    tsigKey?: pulumi.Input<inputs.DnsZoneTsigKey>;
    type?: pulumi.Input<string>;
    versionId?: pulumi.Input<string>;
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DnsZone resource.
 */
export interface DnsZoneArgs {
    comment?: pulumi.Input<string>;
    contract: pulumi.Input<string>;
    endCustomerId?: pulumi.Input<string>;
    group?: pulumi.Input<string>;
    masters?: pulumi.Input<pulumi.Input<string>[]>;
    signAndServe?: pulumi.Input<boolean>;
    signAndServeAlgorithm?: pulumi.Input<string>;
    target?: pulumi.Input<string>;
    tsigKey?: pulumi.Input<inputs.DnsZoneTsigKey>;
    type: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
