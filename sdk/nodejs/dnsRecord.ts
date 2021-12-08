// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the `akamai.DnsRecord` resource to configure a DNS record that can integrate with your existing DNS infrastructure.
 *
 * ## Example Usage
 *
 * Here are examples of an A record and a CNAME record.
 * ### An A record example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const origin = new akamai.DnsRecord("origin", {
 *     active: true,
 *     recordtype: "A",
 *     targets: ["192.0.2.42"],
 *     ttl: 30,
 *     zone: "origin.org",
 * });
 * ```
 * ### CNAME Record Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const www = new akamai.DnsRecord("www", {
 *     active: true,
 *     recordtype: "CNAME",
 *     targets: "origin.example.org.edgesuite.net",
 *     ttl: 600,
 *     zone: "example.com",
 * });
 * ```
 * ## Argument reference [argument-reference]
 *
 * This resource supports these arguments for all record types:
 *
 * * `name` - (Required) The DNS record name. This is the node this DNS record is associated with. Also known as an owner name.
 * * `zone` - (Required) The domain zone, including any nested subdomains.
 * * `recordType` - (Required) The DNS record type.
 * * `ttl` - (Required) The time to live (TTL) is a 32-bit signed integer for the time the resource record is cached. <br /> A value of `0` means that the resource record is not cached. It's only used for the transaction in progress and may be useful for extremely volatile data.
 *
 * ## Additional arguments by record type
 *
 * This section lists additional required and optional arguments for specific record types.
 *
 * ### A record
 *
 * An A record requires this argument:
 *
 * * `target` - One or more IPv4 addresses, for example, 192.0.2.0.
 *
 * ### AAAA record
 *
 * An AAAA record requires this argument:
 *
 * * `target` - One or more IPv6 addresses, for example, 2001:0db8::ff00:0042:8329.
 *
 * ### AFSDB record
 *
 * An AFSDB record requires these arguments:
 *
 * * `target` - The domain name of the host having a server for the cell named by the owner name of the resource record.
 * * `subtype` - An integer between `0` and `65535` that indicates the type of service provided by the host.
 *
 * ### AKAMAICDN record
 *
 * An AKAMAICDN record requires this argument:
 *
 * * `target` - A DNS name representing the selected edge hostname and domain.
 *
 * ### AKAMAITLC record
 *
 * No additional arguments are needed for AKAMAITLC records. This resource returns these computed attributes for this record type:
 *
 * * `dnsName` - A valid DNS name.
 * * `answerType` - The answer type.
 *
 * ### CAA record
 *
 * A certificate authority authorization (CAA) record requires this argument:
 *
 * * `target` - One or more certificate authority authorizations. Each authorization contains three attributes: flags, property tag, and property value.
 *
 * Example:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ### CERT record
 *
 * A CERT record requires these arguments:
 *
 * * `typeValue` - A numeric certificate type value.
 * * `typeMnemonic` - A mnemonic certificate type value.
 * * `keytag` - A value computed for the key embedded in the certificate.
 * * `algorithm` - The cryptographic algorithm used to create the signature.
 * * `certificate` - Certificate data.
 *
 * > **Note:** When entering the certificate type, you can enter `typeValue`, `typeMnemonic`, or  both arguments. If you use both, `typeMnemonic` takes precedence.
 *
 * ### CNAME record
 *
 * A CNAME record requires this argument:
 *
 * * ` target  `- A domain name that specifies the canonical or primary name for the owner. The owner name is an alias.
 *
 * ### DNSKEY record
 *
 * A DNSKEY record requires these arguments:
 *
 * * `flags`
 * * `protocol` - Set to `3`. If the value isn't `3`, the DNSKEY resource record is treated as invalid during signature verification.
 * * `algorithm` - The public key's cryptographic algorithm. This algorithm determines the format of the public key field.
 * * `key` - A Base64 encoded value representing the public key. The format used depends on the `algorithm`.
 *
 * ### DS record
 *
 * A DS record requires these arguments:
 *
 * * `keytag` - The key tag of the DNSKEY record that the DS record refers to, in network byte order.
 * * `algorithm` - The algorithm number of the DNSKEY resource record referred to by the DS record.
 * * `digestType` - Identifies the algorithm used to construct the digest.
 * * `digest` - A base 16 encoded DS record includes a digest of the DNSKEY record it refers to. The digest is conifgured the canonical form of the DNSKEY record's fully qualified owner name with the DNSKEY RDATA, and then applying the digest algorithm.
 *
 * ### HINFO record
 *
 * A HINFO record requires these arguments:
 *
 * * `hardware` - The type of hardware the host uses. A machine name or CPU type may be up to 40 characters long and include uppercase letters, digits, hyphens, and slashes. The entry needs to start and to end with an uppercase letter.
 * * `software` - The type of software the host uses. A system name may be up to 40 characters long and include uppercase letters, digits, hyphens, and slashes. The entry needs to start with an uppercase letter and end with an uppercase letter or a digit.
 *
 * ### HTTPS Record
 *
 * The following fields are required:
 *
 * * `svcPriority` - Service priority associated with endpoint. Value mist be between 0 and 65535. A piority of 0 enables alias mode.
 * * `svcParams` - Space separated list of endpoint parameters. Not allowed if service priority is 0.
 * * `targetName` - Domain name of the service endpoint.
 *
 * ### LOC record
 *
 * A LOC record requires this argument:
 *
 * * `target` - A geographical location associated with a domain name.
 *
 * ### MX record
 *
 * An MX record supports these arguments:
 *
 * * `target` - (Required) One or more domain names that specify a host willing to act as a mail exchange for the owner name.
 * * `priority` - (Optional) The preference value given to this MX record in relation to all other MX records. When a mailer needs to send mail to a certain DNS domain, it first contacts a DNS server for that domain and retrieves all the MX records. It then contacts the mailer with the lowest preference value. This value is ignored if an embedded priority exists in the target.
 * * `priorityIncrement` - (Optional) An auto priority increment when multiple targets are provided with no embedded priority.
 *
 * See Working with MX records in the DNS Getting Started Guide for more information.
 *
 * ### NAPTR record
 *
 * An NAPTR record requires these arguments:
 *
 * * `order` - A 16-bit unsigned integer specifying the order in which the NAPTR records need to be processed to ensure the correct ordering of rules. Low numbers are processed before high numbers. Once a NAPTR is found whose rule matches the target, the client shouldn't consider any NAPTRs with a higher value for order (except as noted below for the flagsnapter field).
 * * `preference` - A 16-bit unsigned integer that specifies the order in which NAPTR records with equal order values are processed. Low numbers are processed before high numbers.
 * * `flagsnaptr` - A character string containing flags that control how fields in the record are rewritten and interpreted. Flags are single alphanumeric characters.
 * * `service` - Specifies the services available down this rewrite path.
 * * `regexp` - A regular expression string containing a substitution expression. This substitution expression is applied to the original client string in order to construct the next domain name to lookup.
 * * `replacement` - Depending on the value of the `flags` attribute, the next NAME to query for NAPTR, SRV, or address records. Enter a fully qualified domain name as the value.
 *
 * ### NS record
 *
 * An NS record requires these arguments:
 *
 * * `target` - One or more domain names that specify authoritative hosts for the specified class and domain.
 *
 * ### NSEC3 record
 *
 * An NSEC3 record requires these arguments:
 *
 * * `algorithm` - The cryptographic hash algorithm used to construct the hash-value.
 * * `flags` - Eight one-bit flags you can use to indicate different processing. All undefined flags must be zero.
 * * `iterations` - The number of additional times the hash function has been performed.
 * * `salt` - The base 16 encoded salt value, which is appended to the original owner name before hashing. Used to defend against pre-calculated dictionary attacks.
 * * `nextHashedOwnerName` - Base 32 encoded. The next hashed owner name in hash order. This value is in binary format. Given the ordered set of all hashed owner names, the Next Hashed Owner Name field contains the hash of an owner name that immediately follows the owner name of the given NSEC3 RR.
 * * `typeBitmaps` - The resource record set types that exist at the original owner name of the NSEC3 RR.
 *
 * ### NSEC3PARAM record
 *
 * An NSEC3PARAM record requires these arguments:
 *
 * * `algorithm` - The cryptographic hash algorithm used to construct the hash-value.
 * * `flags` - Eight one-bit flags that can be used to indicate different processing. All undefined flags must be zero.
 * * `iterations` - The number of additional times the hash function has been performed.
 * * `salt` - The base 16 encoded salt value, which is appended to the original owner name before hashing in order to defend against pre-calculated dictionary attacks.
 *
 * ### PTR record
 *
 * A PTR record requires this argument:
 *
 * * `target` - A domain name that points to some location in the domain name space.
 *
 * ### RP record
 *
 * An RP record requires these arguments:
 *
 * * `mailbox` - A domain name that specifies the mailbox for the responsible person.
 * * `txt` - A domain name for which TXT resource records exist.
 *
 * ### RRSIG record
 *
 * An RRSIG record requires these arguments:
 *
 * * `typeCovered` - The resource record set type covered by this signature.
 * * `algorithm` - Identifies the cryptographic algorithm used to create the signature.
 * * `originalTtl` - The TTL of the covered record set as it appears in the authoritative zone.
 * * `expiration` - The end point of this signature's validity. The signature can`t be used for authentication past this point in time.
 * * `inception` - The start point of this signature's validity. The signature can`t be used for authentication prior to this point in time.
 * * `keytag` - The Key Tag field contains the key tag value of the DNSKEY RR that validates this signature, in network byte order.
 * * `signer` - The owner of the DNSKEY resource record who validates this signature.
 * * `signature` - The base 64 encoded cryptographic signature that covers the RRSIG RDATA and covered record set. Format depends on the TSIG algorithm in use.
 * * `labels` - The Labels field specifies the number of labels in the original RRSIG RR owner name. The significance of this field is that a validator uses it to determine whether the answer was synthesized from a wildcard. If so, it can be used to determine what owner name was used in generating the signature.
 *
 * ### SPF record
 *
 * An SPF record requires this argument:
 *
 * * `target` - Indicates which hosts are, and are not, authorized to use a domain name for the “HELO” and “MAIL FROM” identities.
 *
 * ### SRV record
 *
 * An SRV record requires these arguments:
 *
 * * `target` - The domain name of the target host.
 * * `priority` - A 16-bit integer that specifies the preference given to this resource record among others at the same owner. Lower values are preferred.
 * * `weight` - A server selection mechanism that specifies a relative weight for entries with the same priority. Larger weights are given a proportionately higher probability of being selected. The range of this number is 0–65535, a 16-bit unsigned integer in network byte order. Domain administrators should use Weight 0 when there isn't any server selection to do, to make the RR easier to read for humans. In the presence of records containing weights greater than 0, records with weight 0 should have a very small chance of being selected.
 * * `port` - The port on this target of this service. The range of this number is 0–65535, a 16-bit unsigned integer in network byte order.
 *
 * ### SSHFP record
 *
 * An SSHFP record requires these arguments:
 *
 * * `algorithm` - Describes the algorithm of the public key. The following values are assigned: `0` is reserved, `1` is for RSA, `2` is for DSS, and `3` is for ECDSA.
 * * `fingerprintType` - Describes the message-digest algorithm used to calculate the fingerprint of the public key. The following values are assigned: 0 = reserved, 1 = SHA-1, 2 = SHA-256.
 * * `fingerprint` - The base 16 encoded fingerprint as calculated over the public key blob. The message-digest algorithm is presumed to produce an opaque octet string output, which is placed as-is in the RDATA fingerprint field.
 *
 * ### SOA record
 *
 * An SOA record requires these arguments:
 *
 * * `nameServer` - The domain name of the name server that was the original or primary source of data for this zone.
 * * `emailAddress` - A domain name that specifies the mailbox of this person responsible for this zone.
 * * `serial` - The unsigned version number between 0 and 214748364 of the original copy of the zone.
 * * `refresh` - A time interval between 0 and 214748364 before the zone should be refreshed.
 * * `retry` - A time interval between 0 and 214748364 that should elapse before a failed refresh should be retried.
 * * `expiry` - A time value between 0 and 214748364 that specifies the upper limit on the time interval that can elapse before the zone is no longer authoritative.
 * * `nxdomainTtl` - The unsigned minimum TTL between 0 and 214748364 that should be exported with any resource record from this zone.
 *
 * ### SVCB record
 *
 * An SVCB record requires these arguments:
 *
 * * `svcPriority` - Service priority associated with endpoint. Value mist be between 0 and 65535. A piority of 0 enables alias mode.
 * * `svcParams` - Space separated list of endpoint parameters. Not allowed if service priority is 0.
 * * `targetName` - Domain name of the service endpoint.
 *
 * ### TLSA record
 *
 * A TLSA record requires these arguments:
 *
 * * `usage` - Specifies the association used to match the certificate presented in the TLS handshake.
 * * `selector` - Specifies the part of the TLS certificate presented by the server that is matched against the association data.
 * * `matchType` - Specifies how the certificate association is presented.
 * * `certificate` - Specifies the certificate association data to be matched.
 *
 * ### TXT record
 *
 * A TXT record requires this argument:
 *
 * * `target` - One or more character strings. TXT resource records hold descriptive text. The semantics of the text depends on the domain where it is found.
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
        return new DnsRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/dnsRecord:DnsRecord';

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
    constructor(name: string, args: DnsRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnsRecordArgs | DnsRecordState, opts?: pulumi.CustomResourceOptions) {
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
        const aliasOpts = { aliases: [{ type: "akamai:edgedns/dnsRecord:DnsRecord" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
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
