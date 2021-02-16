// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * @deprecated akamai.properties.EdgeHostName has been deprecated in favor of akamai.EdgeHostName
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
        pulumi.log.warn("EdgeHostName is deprecated: akamai.properties.EdgeHostName has been deprecated in favor of akamai.EdgeHostName")
        return new EdgeHostName(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:properties/edgeHostName:EdgeHostName';

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

    public readonly certificate!: pulumi.Output<number | undefined>;
    /**
     * @deprecated use "contract_id" attribute instead
     */
    public readonly contract!: pulumi.Output<string>;
    public readonly contractId!: pulumi.Output<string>;
    public readonly edgeHostname!: pulumi.Output<string>;
    /**
     * @deprecated use "group_id" attribute instead
     */
    public readonly group!: pulumi.Output<string>;
    public readonly groupId!: pulumi.Output<string>;
    public readonly ipBehavior!: pulumi.Output<string>;
    /**
     * @deprecated use "product_id" attribute instead
     */
    public readonly product!: pulumi.Output<string>;
    public readonly productId!: pulumi.Output<string>;

    /**
     * Create a EdgeHostName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated akamai.properties.EdgeHostName has been deprecated in favor of akamai.EdgeHostName */
    constructor(name: string, args: EdgeHostNameArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated akamai.properties.EdgeHostName has been deprecated in favor of akamai.EdgeHostName */
    constructor(name: string, argsOrState?: EdgeHostNameArgs | EdgeHostNameState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("EdgeHostName is deprecated: akamai.properties.EdgeHostName has been deprecated in favor of akamai.EdgeHostName")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EdgeHostNameState | undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["contract"] = state ? state.contract : undefined;
            inputs["contractId"] = state ? state.contractId : undefined;
            inputs["edgeHostname"] = state ? state.edgeHostname : undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["ipBehavior"] = state ? state.ipBehavior : undefined;
            inputs["product"] = state ? state.product : undefined;
            inputs["productId"] = state ? state.productId : undefined;
        } else {
            const args = argsOrState as EdgeHostNameArgs | undefined;
            if ((!args || args.edgeHostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edgeHostname'");
            }
            if ((!args || args.ipBehavior === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipBehavior'");
            }
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["contract"] = args ? args.contract : undefined;
            inputs["contractId"] = args ? args.contractId : undefined;
            inputs["edgeHostname"] = args ? args.edgeHostname : undefined;
            inputs["group"] = args ? args.group : undefined;
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["ipBehavior"] = args ? args.ipBehavior : undefined;
            inputs["product"] = args ? args.product : undefined;
            inputs["productId"] = args ? args.productId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EdgeHostName.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EdgeHostName resources.
 */
export interface EdgeHostNameState {
    readonly certificate?: pulumi.Input<number>;
    /**
     * @deprecated use "contract_id" attribute instead
     */
    readonly contract?: pulumi.Input<string>;
    readonly contractId?: pulumi.Input<string>;
    readonly edgeHostname?: pulumi.Input<string>;
    /**
     * @deprecated use "group_id" attribute instead
     */
    readonly group?: pulumi.Input<string>;
    readonly groupId?: pulumi.Input<string>;
    readonly ipBehavior?: pulumi.Input<string>;
    /**
     * @deprecated use "product_id" attribute instead
     */
    readonly product?: pulumi.Input<string>;
    readonly productId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EdgeHostName resource.
 */
export interface EdgeHostNameArgs {
    readonly certificate?: pulumi.Input<number>;
    /**
     * @deprecated use "contract_id" attribute instead
     */
    readonly contract?: pulumi.Input<string>;
    readonly contractId?: pulumi.Input<string>;
    readonly edgeHostname: pulumi.Input<string>;
    /**
     * @deprecated use "group_id" attribute instead
     */
    readonly group?: pulumi.Input<string>;
    readonly groupId?: pulumi.Input<string>;
    readonly ipBehavior: pulumi.Input<string>;
    /**
     * @deprecated use "product_id" attribute instead
     */
    readonly product?: pulumi.Input<string>;
    readonly productId?: pulumi.Input<string>;
}
