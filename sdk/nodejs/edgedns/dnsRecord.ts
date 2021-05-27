// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * @deprecated akamai.edgedns.DnsRecord has been deprecated in favor of akamai.DnsRecord
 */
export class DnsRecord extends pulumi.CustomResource {
    /**
     * Get an existing DnsRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsRecordState, opts?: pulumi.CustomResourceOptions): DnsRecord {
        pulumi.log.warn("DnsRecord is deprecated: akamai.edgedns.DnsRecord has been deprecated in favor of akamai.DnsRecord")
        return new DnsRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:edgedns/dnsRecord:DnsRecord';

    /**
     * Returns true if the given object is an instance of DnsRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsRecord.__pulumiType;
    }

    public readonly active!: pulumi.Output<boolean | undefined>;
    public readonly algorithm!: pulumi.Output<number | undefined>;
    public /*out*/ readonly answerType!: pulumi.Output<string>;
    public readonly certificate!: pulumi.Output<string | undefined>;
    public readonly digest!: pulumi.Output<string | undefined>;
    public readonly digestType!: pulumi.Output<number | undefined>;
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    public readonly emailAddress!: pulumi.Output<string | undefined>;
    public readonly expiration!: pulumi.Output<string | undefined>;
    public readonly expiry!: pulumi.Output<number | undefined>;
    public readonly fingerprint!: pulumi.Output<string | undefined>;
    public readonly fingerprintType!: pulumi.Output<number | undefined>;
    public readonly flags!: pulumi.Output<number | undefined>;
    public readonly flagsnaptr!: pulumi.Output<string | undefined>;
    public readonly hardware!: pulumi.Output<string | undefined>;
    public readonly inception!: pulumi.Output<string | undefined>;
    public readonly iterations!: pulumi.Output<number | undefined>;
    public readonly key!: pulumi.Output<string | undefined>;
    public readonly keytag!: pulumi.Output<number | undefined>;
    public readonly labels!: pulumi.Output<number | undefined>;
    public readonly mailbox!: pulumi.Output<string | undefined>;
    public readonly matchType!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly nameServer!: pulumi.Output<string | undefined>;
    public readonly nextHashedOwnerName!: pulumi.Output<string | undefined>;
    public readonly nxdomainTtl!: pulumi.Output<number | undefined>;
    public readonly order!: pulumi.Output<number | undefined>;
    public readonly originalTtl!: pulumi.Output<number | undefined>;
    public readonly port!: pulumi.Output<number | undefined>;
    public readonly preference!: pulumi.Output<number | undefined>;
    public readonly priority!: pulumi.Output<number | undefined>;
    public readonly priorityIncrement!: pulumi.Output<number | undefined>;
    public readonly protocol!: pulumi.Output<number | undefined>;
    public /*out*/ readonly recordSha!: pulumi.Output<string>;
    public readonly recordtype!: pulumi.Output<string>;
    public readonly refresh!: pulumi.Output<number | undefined>;
    public readonly regexp!: pulumi.Output<string | undefined>;
    public readonly replacement!: pulumi.Output<string | undefined>;
    public readonly retry!: pulumi.Output<number | undefined>;
    public readonly salt!: pulumi.Output<string | undefined>;
    public readonly selector!: pulumi.Output<number | undefined>;
    public /*out*/ readonly serial!: pulumi.Output<number>;
    public readonly service!: pulumi.Output<string | undefined>;
    public readonly signature!: pulumi.Output<string | undefined>;
    public readonly signer!: pulumi.Output<string | undefined>;
    public readonly software!: pulumi.Output<string | undefined>;
    public readonly subtype!: pulumi.Output<number | undefined>;
    public readonly svcParams!: pulumi.Output<string | undefined>;
    public readonly svcPriority!: pulumi.Output<number | undefined>;
    public readonly targetName!: pulumi.Output<string | undefined>;
    public readonly targets!: pulumi.Output<string[] | undefined>;
    public readonly ttl!: pulumi.Output<number>;
    public readonly txt!: pulumi.Output<string | undefined>;
    public readonly typeBitmaps!: pulumi.Output<string | undefined>;
    public readonly typeCovered!: pulumi.Output<string | undefined>;
    public readonly typeMnemonic!: pulumi.Output<string | undefined>;
    public readonly typeValue!: pulumi.Output<number | undefined>;
    public readonly usage!: pulumi.Output<number | undefined>;
    public readonly weight!: pulumi.Output<number | undefined>;
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a DnsRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated akamai.edgedns.DnsRecord has been deprecated in favor of akamai.DnsRecord */
    constructor(name: string, args: DnsRecordArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated akamai.edgedns.DnsRecord has been deprecated in favor of akamai.DnsRecord */
    constructor(name: string, argsOrState?: DnsRecordArgs | DnsRecordState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DnsRecord is deprecated: akamai.edgedns.DnsRecord has been deprecated in favor of akamai.DnsRecord")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsRecordState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["algorithm"] = state ? state.algorithm : undefined;
            inputs["answerType"] = state ? state.answerType : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["digest"] = state ? state.digest : undefined;
            inputs["digestType"] = state ? state.digestType : undefined;
            inputs["dnsName"] = state ? state.dnsName : undefined;
            inputs["emailAddress"] = state ? state.emailAddress : undefined;
            inputs["expiration"] = state ? state.expiration : undefined;
            inputs["expiry"] = state ? state.expiry : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["fingerprintType"] = state ? state.fingerprintType : undefined;
            inputs["flags"] = state ? state.flags : undefined;
            inputs["flagsnaptr"] = state ? state.flagsnaptr : undefined;
            inputs["hardware"] = state ? state.hardware : undefined;
            inputs["inception"] = state ? state.inception : undefined;
            inputs["iterations"] = state ? state.iterations : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["keytag"] = state ? state.keytag : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["mailbox"] = state ? state.mailbox : undefined;
            inputs["matchType"] = state ? state.matchType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nameServer"] = state ? state.nameServer : undefined;
            inputs["nextHashedOwnerName"] = state ? state.nextHashedOwnerName : undefined;
            inputs["nxdomainTtl"] = state ? state.nxdomainTtl : undefined;
            inputs["order"] = state ? state.order : undefined;
            inputs["originalTtl"] = state ? state.originalTtl : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["preference"] = state ? state.preference : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["priorityIncrement"] = state ? state.priorityIncrement : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["recordSha"] = state ? state.recordSha : undefined;
            inputs["recordtype"] = state ? state.recordtype : undefined;
            inputs["refresh"] = state ? state.refresh : undefined;
            inputs["regexp"] = state ? state.regexp : undefined;
            inputs["replacement"] = state ? state.replacement : undefined;
            inputs["retry"] = state ? state.retry : undefined;
            inputs["salt"] = state ? state.salt : undefined;
            inputs["selector"] = state ? state.selector : undefined;
            inputs["serial"] = state ? state.serial : undefined;
            inputs["service"] = state ? state.service : undefined;
            inputs["signature"] = state ? state.signature : undefined;
            inputs["signer"] = state ? state.signer : undefined;
            inputs["software"] = state ? state.software : undefined;
            inputs["subtype"] = state ? state.subtype : undefined;
            inputs["svcParams"] = state ? state.svcParams : undefined;
            inputs["svcPriority"] = state ? state.svcPriority : undefined;
            inputs["targetName"] = state ? state.targetName : undefined;
            inputs["targets"] = state ? state.targets : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["txt"] = state ? state.txt : undefined;
            inputs["typeBitmaps"] = state ? state.typeBitmaps : undefined;
            inputs["typeCovered"] = state ? state.typeCovered : undefined;
            inputs["typeMnemonic"] = state ? state.typeMnemonic : undefined;
            inputs["typeValue"] = state ? state.typeValue : undefined;
            inputs["usage"] = state ? state.usage : undefined;
            inputs["weight"] = state ? state.weight : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as DnsRecordArgs | undefined;
            if ((!args || args.recordtype === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recordtype'");
            }
            if ((!args || args.ttl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ttl'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["algorithm"] = args ? args.algorithm : undefined;
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["digest"] = args ? args.digest : undefined;
            inputs["digestType"] = args ? args.digestType : undefined;
            inputs["emailAddress"] = args ? args.emailAddress : undefined;
            inputs["expiration"] = args ? args.expiration : undefined;
            inputs["expiry"] = args ? args.expiry : undefined;
            inputs["fingerprint"] = args ? args.fingerprint : undefined;
            inputs["fingerprintType"] = args ? args.fingerprintType : undefined;
            inputs["flags"] = args ? args.flags : undefined;
            inputs["flagsnaptr"] = args ? args.flagsnaptr : undefined;
            inputs["hardware"] = args ? args.hardware : undefined;
            inputs["inception"] = args ? args.inception : undefined;
            inputs["iterations"] = args ? args.iterations : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["keytag"] = args ? args.keytag : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["mailbox"] = args ? args.mailbox : undefined;
            inputs["matchType"] = args ? args.matchType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nameServer"] = args ? args.nameServer : undefined;
            inputs["nextHashedOwnerName"] = args ? args.nextHashedOwnerName : undefined;
            inputs["nxdomainTtl"] = args ? args.nxdomainTtl : undefined;
            inputs["order"] = args ? args.order : undefined;
            inputs["originalTtl"] = args ? args.originalTtl : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["preference"] = args ? args.preference : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["priorityIncrement"] = args ? args.priorityIncrement : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["recordtype"] = args ? args.recordtype : undefined;
            inputs["refresh"] = args ? args.refresh : undefined;
            inputs["regexp"] = args ? args.regexp : undefined;
            inputs["replacement"] = args ? args.replacement : undefined;
            inputs["retry"] = args ? args.retry : undefined;
            inputs["salt"] = args ? args.salt : undefined;
            inputs["selector"] = args ? args.selector : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["signature"] = args ? args.signature : undefined;
            inputs["signer"] = args ? args.signer : undefined;
            inputs["software"] = args ? args.software : undefined;
            inputs["subtype"] = args ? args.subtype : undefined;
            inputs["svcParams"] = args ? args.svcParams : undefined;
            inputs["svcPriority"] = args ? args.svcPriority : undefined;
            inputs["targetName"] = args ? args.targetName : undefined;
            inputs["targets"] = args ? args.targets : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["txt"] = args ? args.txt : undefined;
            inputs["typeBitmaps"] = args ? args.typeBitmaps : undefined;
            inputs["typeCovered"] = args ? args.typeCovered : undefined;
            inputs["typeMnemonic"] = args ? args.typeMnemonic : undefined;
            inputs["typeValue"] = args ? args.typeValue : undefined;
            inputs["usage"] = args ? args.usage : undefined;
            inputs["weight"] = args ? args.weight : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["answerType"] = undefined /*out*/;
            inputs["dnsName"] = undefined /*out*/;
            inputs["recordSha"] = undefined /*out*/;
            inputs["serial"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DnsRecord.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsRecord resources.
 */
export interface DnsRecordState {
    active?: pulumi.Input<boolean>;
    algorithm?: pulumi.Input<number>;
    answerType?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    digest?: pulumi.Input<string>;
    digestType?: pulumi.Input<number>;
    dnsName?: pulumi.Input<string>;
    emailAddress?: pulumi.Input<string>;
    expiration?: pulumi.Input<string>;
    expiry?: pulumi.Input<number>;
    fingerprint?: pulumi.Input<string>;
    fingerprintType?: pulumi.Input<number>;
    flags?: pulumi.Input<number>;
    flagsnaptr?: pulumi.Input<string>;
    hardware?: pulumi.Input<string>;
    inception?: pulumi.Input<string>;
    iterations?: pulumi.Input<number>;
    key?: pulumi.Input<string>;
    keytag?: pulumi.Input<number>;
    labels?: pulumi.Input<number>;
    mailbox?: pulumi.Input<string>;
    matchType?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    nameServer?: pulumi.Input<string>;
    nextHashedOwnerName?: pulumi.Input<string>;
    nxdomainTtl?: pulumi.Input<number>;
    order?: pulumi.Input<number>;
    originalTtl?: pulumi.Input<number>;
    port?: pulumi.Input<number>;
    preference?: pulumi.Input<number>;
    priority?: pulumi.Input<number>;
    priorityIncrement?: pulumi.Input<number>;
    protocol?: pulumi.Input<number>;
    recordSha?: pulumi.Input<string>;
    recordtype?: pulumi.Input<string>;
    refresh?: pulumi.Input<number>;
    regexp?: pulumi.Input<string>;
    replacement?: pulumi.Input<string>;
    retry?: pulumi.Input<number>;
    salt?: pulumi.Input<string>;
    selector?: pulumi.Input<number>;
    serial?: pulumi.Input<number>;
    service?: pulumi.Input<string>;
    signature?: pulumi.Input<string>;
    signer?: pulumi.Input<string>;
    software?: pulumi.Input<string>;
    subtype?: pulumi.Input<number>;
    svcParams?: pulumi.Input<string>;
    svcPriority?: pulumi.Input<number>;
    targetName?: pulumi.Input<string>;
    targets?: pulumi.Input<pulumi.Input<string>[]>;
    ttl?: pulumi.Input<number>;
    txt?: pulumi.Input<string>;
    typeBitmaps?: pulumi.Input<string>;
    typeCovered?: pulumi.Input<string>;
    typeMnemonic?: pulumi.Input<string>;
    typeValue?: pulumi.Input<number>;
    usage?: pulumi.Input<number>;
    weight?: pulumi.Input<number>;
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DnsRecord resource.
 */
export interface DnsRecordArgs {
    active?: pulumi.Input<boolean>;
    algorithm?: pulumi.Input<number>;
    certificate?: pulumi.Input<string>;
    digest?: pulumi.Input<string>;
    digestType?: pulumi.Input<number>;
    emailAddress?: pulumi.Input<string>;
    expiration?: pulumi.Input<string>;
    expiry?: pulumi.Input<number>;
    fingerprint?: pulumi.Input<string>;
    fingerprintType?: pulumi.Input<number>;
    flags?: pulumi.Input<number>;
    flagsnaptr?: pulumi.Input<string>;
    hardware?: pulumi.Input<string>;
    inception?: pulumi.Input<string>;
    iterations?: pulumi.Input<number>;
    key?: pulumi.Input<string>;
    keytag?: pulumi.Input<number>;
    labels?: pulumi.Input<number>;
    mailbox?: pulumi.Input<string>;
    matchType?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    nameServer?: pulumi.Input<string>;
    nextHashedOwnerName?: pulumi.Input<string>;
    nxdomainTtl?: pulumi.Input<number>;
    order?: pulumi.Input<number>;
    originalTtl?: pulumi.Input<number>;
    port?: pulumi.Input<number>;
    preference?: pulumi.Input<number>;
    priority?: pulumi.Input<number>;
    priorityIncrement?: pulumi.Input<number>;
    protocol?: pulumi.Input<number>;
    recordtype: pulumi.Input<string>;
    refresh?: pulumi.Input<number>;
    regexp?: pulumi.Input<string>;
    replacement?: pulumi.Input<string>;
    retry?: pulumi.Input<number>;
    salt?: pulumi.Input<string>;
    selector?: pulumi.Input<number>;
    service?: pulumi.Input<string>;
    signature?: pulumi.Input<string>;
    signer?: pulumi.Input<string>;
    software?: pulumi.Input<string>;
    subtype?: pulumi.Input<number>;
    svcParams?: pulumi.Input<string>;
    svcPriority?: pulumi.Input<number>;
    targetName?: pulumi.Input<string>;
    targets?: pulumi.Input<pulumi.Input<string>[]>;
    ttl: pulumi.Input<number>;
    txt?: pulumi.Input<string>;
    typeBitmaps?: pulumi.Input<string>;
    typeCovered?: pulumi.Input<string>;
    typeMnemonic?: pulumi.Input<string>;
    typeValue?: pulumi.Input<number>;
    usage?: pulumi.Input<number>;
    weight?: pulumi.Input<number>;
    zone: pulumi.Input<string>;
}
