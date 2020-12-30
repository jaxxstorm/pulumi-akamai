// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    /// <summary>
    /// The `akamai.DnsRecord` provides the resource for configuring a dns record to integrate easily with your existing DNS infrastructure to provide a secure, high performance, highly available and scalable solution for DNS hosting.
    /// 
    /// ## Example Usage
    /// ### A Record Example
    /// resource "akamai_dns_record" "origin" {
    ///     zone = "origin.org"
    ///     name = "origin.example.org"
    ///     recordtype =  "A"
    ///     active = true
    ///     ttl =  30
    ///     target = ["192.0.2.42"]
    /// }
    /// ## Required Fields Per Record Type
    /// 
    /// In addition to the fields listed in the prior section, type specific fields define the data makeup of each Record's data. This section identfies required fields per type.
    /// 
    /// ### A Record
    /// 
    /// The following field is required:
    /// 
    /// * target - One or more IPv4 addresses, for example, 1.2.3.4.
    /// 
    /// ### AAAA Record
    /// 
    /// The following field is required:
    /// 
    /// * target - One or more IPv6 addresses, for example, 2001:0db8::ff00:0042:8329.
    /// 
    /// ### AFSDB Record
    /// 
    /// The following fields are required:
    /// 
    /// * target - The domain name of the host having a server for the cell named by the owner name of the resource record.
    /// * subtype- An integer between 0 and 65535, indicating the type of service provided by the host.
    /// 
    /// ### AKAMAICDN Record
    /// 
    /// The following field is required:
    /// 
    /// * target - DNS name representing selected Edge Hostname name+domain.
    /// 
    /// ### AKAMAITLC Record
    /// 
    /// No additional fields are required. The following fields are Computed.
    /// 
    /// * dns_name - valid DNS name.
    /// * answer_type - answer type.
    /// 
    /// ### CAA Record
    /// 
    /// The following field are required:
    /// 
    /// * target - One or more CA Authorizations. Each authorization contains three attributes: flags, property tag and property value.
    /// 
    /// Example:
    /// ```csharp
    /// using Pulumi;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### CERT Record
    /// 
    /// The following fields are required:
    /// 
    /// * type_value - numeric certificate type value
    /// * type_mnemonic - mnemonic certificate type value.
    /// * keytag - value computed for the key embedded in the certificate
    /// * algorithm - identifies the cryptographic algorithm used to create the signature.
    /// * certificate - certificate data
    /// 
    /// Note: Type can be configured either a numeric OR menmonic value. If both set, type_mnemonic takes precedent.
    /// 
    /// ### CNAME Record
    /// 
    /// The following field is required:
    /// 
    /// * target - A domain name that specifies the canonical or primary name for the owner. The owner name is an alias.
    /// 
    /// ### DNSKEY Record
    /// 
    /// The following fields are required:
    /// 
    /// * flags
    /// * protocol - Must have the value 3. The DNSKEY resource record must be treated as invalid during signature verification if it contains a value other than 3.
    /// * algorithm - The public key’s cryptographic algorithm and determine the format of the public key field.
    /// * key - Base 64 encoded value representing the public key, the format of which depends on the algorithm being used.
    /// 
    /// ### DS Record
    /// 
    /// The following fields are required:
    /// 
    /// * keytag - The key tag of the DNSKEY resource record referred to by the DS record, in network byte order.
    /// * algorithm - The algorithm number of the DNSKEY resource record referred to by the DS record.
    /// * digest_type - Identifies the algorithm used to construct the digest.
    /// * digest - The base 16 encoded DS record refers to a DNSKEY RR by including a digest of that DNSKEY RR. The digest is calculated by concatenating the canonical form of the fully qualified owner name of the DNSKEY RR with the DNSKEY RDATA, and then applying the digest algorithm.
    /// 
    /// ### HINFO Record
    /// 
    /// The following fields are required:
    /// 
    /// * hardware - Type of hardware the host uses. A machine name or CPU type may be up to 40 characters taken from the set of uppercase letters, digits, and the two punctuation characters hyphen and slash. It must start with a letter, and end with a letter.
    /// * software - Type of software the host uses. A system name may be up to 40 characters taken from the set of uppercase letters, digits, and the two punctuation characters hyphen and slash. It must start with a letter, and end with a letter or digit.
    /// 
    /// ### LOC Record
    /// 
    /// The following field is required:
    /// 
    /// * target - A geographical location associated with a domain name.
    /// 
    /// ### MX Record
    /// 
    /// The following field is required:
    /// 
    /// * target - One or more domain names that specifies a host willing to act as a mail exchange for the owner name.
    /// 
    /// The following fields are optional depending on configuration type. See DNS Getting Started Guide for more information.
    /// 
    /// * priority - The preference value given to the MX record among MX records. When a mailer needs to send mail to a certain DNS domain, it first contacts a DNS server for that domain and retrieves all the MX records. It then contacts the mailer with the lowest preference value. Ignored if embedded priority specified in target
    /// * priority_increment - auto priority increment when multiple targets are provided with no embedded priority.
    /// 
    /// ### NAPTR Record
    /// 
    /// The following fields are required:
    /// 
    /// * order - A 16-bit unsigned integer specifying the order in which the NAPTR records MUST be processed to ensure the correct ordering ofrules. Low numbers are processed before high numbers, and once a NAPTR is found whose rule “matches” the target, the client MUST NOT consider any NAPTRs with a higher value for order (except as noted below for the Flags field).
    /// * preference - A 16-bit unsigned integer that specifies the order in which NAPTR records with equal order values should be processed, low numbers being processed before high numbers.
    /// * flagsnaptr - A &lt;character-string&gt; containing flags to control aspects of the rewriting and interpretation of the fields in the record. Flags are single characters from the set [A-Z0-9]. The case of the alphabetic characters is not significant.
    /// * service - Specifies the services available down this rewrite path.
    /// * regexp - A String containing a substitution expression that is applied to the original string held by the client in order to construct the next domain name to lookup.
    /// * replacement - The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field. This MUST be a fully qualified domain-name.
    /// 
    /// ### NS Record
    /// 
    /// The following field is required:
    /// 
    /// * target - One or more domain names that specify authoritative hosts for the specified class and domain.
    /// 
    /// ### NSEC3 Record
    /// 
    /// The following fields are required:
    /// 
    /// * algorithm - The cryptographic hash algorithm used to construct the hash-value.
    /// * flags - The 8 one-bit flags that can be used to indicate different processing. All undefined flags must be zero.
    /// * iterations - The number of additional times the hash function has been performed.
    /// * salt - The base 16 encoded salt value, which is appended to the original owner name before hashing in order to defend against pre-calculated dictionary attacks.
    /// * next_hashed_owner_name - Base 32 encoded. The next hashed owner name in hash order. This value is in binary format. Given the ordered set of all hashed owner names, the Next Hashed Owner Name field contains the hash of an owner name that immediately follows the owner name of the given NSEC3 RR.
    /// * type_bitmaps - The resource record set types that exist at the original owner name of the NSEC3 RR.
    /// 
    /// ### NSEC3PARAM Record
    /// 
    /// The following fields are required:
    /// 
    /// * algorithm - The cryptographic hash algorithm used to construct the hash-value.
    /// * flags - The 8 one-bit flags that can be used to indicate different processing. All undefined flags must be zero.
    /// * iterations - The number of additional times the hash function has been performed.
    /// * salt - The base 16 encoded salt value, which is appended to the original owner name before hashing in order to defend against pre-calculated dictionary attacks.
    /// 
    /// ### PTR Record
    /// 
    /// The following field is required:
    /// 
    /// * target - A domain name that points to some location in the domain name space.
    /// 
    /// ### RP Record
    /// 
    /// The following fields are required:
    /// 
    /// * mailbox - A domain name that specifies the mailbox for the responsible person.
    /// * txt - A domain name for which TXT resource records exist.
    /// 
    /// ### RRSIG Record
    /// 
    /// The following fields are required:
    /// 
    /// * type_covered - The resource record set type covered by this signature.
    /// * algorithm - The Algorithm Number field identifies the cryptographic algorithm used to create the signature.
    /// * original_ttl - The TTL of the covered record set as it appears in the authoritative zone.
    /// * expiration - The end point of this signature’s validity. The signature cannot be used for authentication past this point.
    /// * inception - The start point of this signature’s validity. The signature cannot be used for authentication prior to this point.
    /// * keytag - The Key Tag field contains the key tag value of the DNSKEY RR that validates this signature, in network byte order.
    /// * signer - The owner of the DSNKEY resource record who validates this signature.
    /// * signature - The base 64 encoded cryptographic signature that covers the RRSIG RDATA and covered record set. Format depends on the TSIG algorithm in use.
    /// * labels - The Labels field specifies the number of labels in the original RRSIG RR owner name. The significance of this field is that a validator uses it to determine whether the answer was synthesized from a wildcard. If so, it can be used to determine what owner name was used in generating the signature.
    /// 
    /// ### SPF Record
    /// 
    /// The following field is required:
    /// 
    /// * target - Indicates which hosts are, and are not, authorized to use a domain name for the “HELO” and “MAIL FROM” identities.
    /// 
    /// ### SRV Record
    /// 
    /// The following fields are required:
    /// 
    /// * target - The domain name of the target host.
    /// * priority - A 16-bit integer that specifies the preference given to this resource record among others at the same owner. Lower values are preferred.
    /// * weight - A server selection mechanism, specifying a relative weight for entries with the same priority. Larger weights should be given a proportionately higher probability of being selected. The range of this number is 0–65535, a 16-bit unsigned integer in network byte order. Domain administrators should use Weight 0 when there isn’t any server selection to do, to make the RR easier to read for humans. In the presence of records containing weights greater than 0, records with weight 0 should have a very small chance of being selected.
    /// * port - The port on this target of this service. The range of this number is 0–65535, a 16-bit unsigned integer in network byte order.
    /// 
    /// ### SSHFP Record
    /// 
    /// The following fields are required:
    /// 
    /// * algorithm - Describes the algorithm of the public key. The following values are assigned: 0 = reserved; 1 = RSA; 2 = DSS, 3 = ECDSA
    /// * fingerprint_type - Describes the message-digest algorithm used to calculate the fingerprint of the public key. The following values are assigned: 0 = reserved, 1 = SHA-1, 2 = SHA-256
    /// * fingerprint - The base 16 encoded fingerprint as calculated over the public key blob. The message-digest algorithm is presumed to produce an opaque octet string output, which is placed as-is in the RDATA fingerprint field.
    /// 
    /// ### SOA Record
    /// 
    /// The following fields are required:
    /// 
    /// * name_server - The domain name of the name server that was the original or primary source of data for this zone.
    /// * email_address - A domain name that specifies the mailbox of this person responsible for this zone.
    /// * serial - The unsigned version number between 0 and 214748364 of the original copy of the zone.
    /// * refresh - A time interval between 0 and 214748364 before the zone should be refreshed.
    /// * retry - A time interval between 0 and 214748364 that should elapse before a failed refresh should be retried.
    /// * expiry - A time value between 0 and 214748364 that specifies the upper limit on the time interval that can elapse before the zone is no longer authoritative.
    /// * nxdomain_ttl - The unsigned minimum TTL between 0 and 214748364 that should be exported with any resource record from this zone.
    /// 
    /// ### TLSA Record
    /// 
    /// The following fields are required:
    /// 
    /// * usage - specifies the provided association that will be used to match the certificate presented in the TLS handshake.
    /// * selector - specifies which part of the TLS certificate presented by the server will be matched against the association data.
    /// * match_type - specifies how the certificate association is presented.
    /// * certificate - specifies the "certificate association data" to be matched.
    /// 
    /// ### TXT Record
    /// 
    /// The following field is required:
    /// 
    /// * target - One or more character strings. TXT RRs are used to hold descriptive text. The semantics of the text depends on the domain where it is found.
    /// </summary>
    public partial class DnsRecord : Pulumi.CustomResource
    {
        /// <summary>
        /// Maintained for backward compatibility
        /// </summary>
        [Output("active")]
        public Output<bool?> Active { get; private set; } = null!;

        [Output("algorithm")]
        public Output<int?> Algorithm { get; private set; } = null!;

        [Output("answerType")]
        public Output<string> AnswerType { get; private set; } = null!;

        [Output("certificate")]
        public Output<string?> Certificate { get; private set; } = null!;

        [Output("digest")]
        public Output<string?> Digest { get; private set; } = null!;

        [Output("digestType")]
        public Output<int?> DigestType { get; private set; } = null!;

        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        [Output("emailAddress")]
        public Output<string?> EmailAddress { get; private set; } = null!;

        [Output("expiration")]
        public Output<string?> Expiration { get; private set; } = null!;

        [Output("expiry")]
        public Output<int?> Expiry { get; private set; } = null!;

        [Output("fingerprint")]
        public Output<string?> Fingerprint { get; private set; } = null!;

        [Output("fingerprintType")]
        public Output<int?> FingerprintType { get; private set; } = null!;

        [Output("flags")]
        public Output<int?> Flags { get; private set; } = null!;

        [Output("flagsnaptr")]
        public Output<string?> Flagsnaptr { get; private set; } = null!;

        [Output("hardware")]
        public Output<string?> Hardware { get; private set; } = null!;

        [Output("inception")]
        public Output<string?> Inception { get; private set; } = null!;

        [Output("iterations")]
        public Output<int?> Iterations { get; private set; } = null!;

        [Output("key")]
        public Output<string?> Key { get; private set; } = null!;

        [Output("keytag")]
        public Output<int?> Keytag { get; private set; } = null!;

        [Output("labels")]
        public Output<int?> Labels { get; private set; } = null!;

        [Output("mailbox")]
        public Output<string?> Mailbox { get; private set; } = null!;

        [Output("matchType")]
        public Output<int?> MatchType { get; private set; } = null!;

        /// <summary>
        /// The name of the record. The name is an owner name, that is, the name of the node to which this resource record pertains.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("nameServer")]
        public Output<string?> NameServer { get; private set; } = null!;

        [Output("nextHashedOwnerName")]
        public Output<string?> NextHashedOwnerName { get; private set; } = null!;

        [Output("nxdomainTtl")]
        public Output<int?> NxdomainTtl { get; private set; } = null!;

        [Output("order")]
        public Output<int?> Order { get; private set; } = null!;

        [Output("originalTtl")]
        public Output<int?> OriginalTtl { get; private set; } = null!;

        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        [Output("preference")]
        public Output<int?> Preference { get; private set; } = null!;

        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        [Output("priorityIncrement")]
        public Output<int?> PriorityIncrement { get; private set; } = null!;

        [Output("protocol")]
        public Output<int?> Protocol { get; private set; } = null!;

        [Output("recordSha")]
        public Output<string> RecordSha { get; private set; } = null!;

        [Output("recordtype")]
        public Output<string> Recordtype { get; private set; } = null!;

        [Output("refresh")]
        public Output<int?> Refresh { get; private set; } = null!;

        [Output("regexp")]
        public Output<string?> Regexp { get; private set; } = null!;

        [Output("replacement")]
        public Output<string?> Replacement { get; private set; } = null!;

        [Output("retry")]
        public Output<int?> Retry { get; private set; } = null!;

        [Output("salt")]
        public Output<string?> Salt { get; private set; } = null!;

        [Output("selector")]
        public Output<int?> Selector { get; private set; } = null!;

        [Output("serial")]
        public Output<int> Serial { get; private set; } = null!;

        [Output("service")]
        public Output<string?> Service { get; private set; } = null!;

        [Output("signature")]
        public Output<string?> Signature { get; private set; } = null!;

        [Output("signer")]
        public Output<string?> Signer { get; private set; } = null!;

        [Output("software")]
        public Output<string?> Software { get; private set; } = null!;

        [Output("subtype")]
        public Output<int?> Subtype { get; private set; } = null!;

        [Output("targets")]
        public Output<ImmutableArray<string>> Targets { get; private set; } = null!;

        /// <summary>
        /// The TTL is a 32-bit signed integer that specifies the time interval that the resource record may be cached before the source of the information should be consulted again. A value of zero means that the RR can only be used for the transaction in progress, and should not be cached. Zero values can also be used for extremely volatile data.
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;

        [Output("txt")]
        public Output<string?> Txt { get; private set; } = null!;

        [Output("typeBitmaps")]
        public Output<string?> TypeBitmaps { get; private set; } = null!;

        [Output("typeCovered")]
        public Output<string?> TypeCovered { get; private set; } = null!;

        [Output("typeMnemonic")]
        public Output<string?> TypeMnemonic { get; private set; } = null!;

        [Output("typeValue")]
        public Output<int?> TypeValue { get; private set; } = null!;

        [Output("usage")]
        public Output<int?> Usage { get; private set; } = null!;

        [Output("weight")]
        public Output<int?> Weight { get; private set; } = null!;

        /// <summary>
        /// Domain zone, encapsulating any nested subdomains.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a DnsRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DnsRecord(string name, DnsRecordArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/dnsRecord:DnsRecord", name, args ?? new DnsRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DnsRecord(string name, Input<string> id, DnsRecordState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/dnsRecord:DnsRecord", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "akamai:edgedns/dnsRecord:DnsRecord"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DnsRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DnsRecord Get(string name, Input<string> id, DnsRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new DnsRecord(name, id, state, options);
        }
    }

    public sealed class DnsRecordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maintained for backward compatibility
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        [Input("algorithm")]
        public Input<int>? Algorithm { get; set; }

        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("digest")]
        public Input<string>? Digest { get; set; }

        [Input("digestType")]
        public Input<int>? DigestType { get; set; }

        [Input("emailAddress")]
        public Input<string>? EmailAddress { get; set; }

        [Input("expiration")]
        public Input<string>? Expiration { get; set; }

        [Input("expiry")]
        public Input<int>? Expiry { get; set; }

        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        [Input("fingerprintType")]
        public Input<int>? FingerprintType { get; set; }

        [Input("flags")]
        public Input<int>? Flags { get; set; }

        [Input("flagsnaptr")]
        public Input<string>? Flagsnaptr { get; set; }

        [Input("hardware")]
        public Input<string>? Hardware { get; set; }

        [Input("inception")]
        public Input<string>? Inception { get; set; }

        [Input("iterations")]
        public Input<int>? Iterations { get; set; }

        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("keytag")]
        public Input<int>? Keytag { get; set; }

        [Input("labels")]
        public Input<int>? Labels { get; set; }

        [Input("mailbox")]
        public Input<string>? Mailbox { get; set; }

        [Input("matchType")]
        public Input<int>? MatchType { get; set; }

        /// <summary>
        /// The name of the record. The name is an owner name, that is, the name of the node to which this resource record pertains.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nameServer")]
        public Input<string>? NameServer { get; set; }

        [Input("nextHashedOwnerName")]
        public Input<string>? NextHashedOwnerName { get; set; }

        [Input("nxdomainTtl")]
        public Input<int>? NxdomainTtl { get; set; }

        [Input("order")]
        public Input<int>? Order { get; set; }

        [Input("originalTtl")]
        public Input<int>? OriginalTtl { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("preference")]
        public Input<int>? Preference { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("priorityIncrement")]
        public Input<int>? PriorityIncrement { get; set; }

        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        [Input("recordtype", required: true)]
        public Input<string> Recordtype { get; set; } = null!;

        [Input("refresh")]
        public Input<int>? Refresh { get; set; }

        [Input("regexp")]
        public Input<string>? Regexp { get; set; }

        [Input("replacement")]
        public Input<string>? Replacement { get; set; }

        [Input("retry")]
        public Input<int>? Retry { get; set; }

        [Input("salt")]
        public Input<string>? Salt { get; set; }

        [Input("selector")]
        public Input<int>? Selector { get; set; }

        [Input("service")]
        public Input<string>? Service { get; set; }

        [Input("signature")]
        public Input<string>? Signature { get; set; }

        [Input("signer")]
        public Input<string>? Signer { get; set; }

        [Input("software")]
        public Input<string>? Software { get; set; }

        [Input("subtype")]
        public Input<int>? Subtype { get; set; }

        [Input("targets")]
        private InputList<string>? _targets;
        public InputList<string> Targets
        {
            get => _targets ?? (_targets = new InputList<string>());
            set => _targets = value;
        }

        /// <summary>
        /// The TTL is a 32-bit signed integer that specifies the time interval that the resource record may be cached before the source of the information should be consulted again. A value of zero means that the RR can only be used for the transaction in progress, and should not be cached. Zero values can also be used for extremely volatile data.
        /// </summary>
        [Input("ttl", required: true)]
        public Input<int> Ttl { get; set; } = null!;

        [Input("txt")]
        public Input<string>? Txt { get; set; }

        [Input("typeBitmaps")]
        public Input<string>? TypeBitmaps { get; set; }

        [Input("typeCovered")]
        public Input<string>? TypeCovered { get; set; }

        [Input("typeMnemonic")]
        public Input<string>? TypeMnemonic { get; set; }

        [Input("typeValue")]
        public Input<int>? TypeValue { get; set; }

        [Input("usage")]
        public Input<int>? Usage { get; set; }

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        /// <summary>
        /// Domain zone, encapsulating any nested subdomains.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public DnsRecordArgs()
        {
        }
    }

    public sealed class DnsRecordState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maintained for backward compatibility
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        [Input("algorithm")]
        public Input<int>? Algorithm { get; set; }

        [Input("answerType")]
        public Input<string>? AnswerType { get; set; }

        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("digest")]
        public Input<string>? Digest { get; set; }

        [Input("digestType")]
        public Input<int>? DigestType { get; set; }

        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        [Input("emailAddress")]
        public Input<string>? EmailAddress { get; set; }

        [Input("expiration")]
        public Input<string>? Expiration { get; set; }

        [Input("expiry")]
        public Input<int>? Expiry { get; set; }

        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        [Input("fingerprintType")]
        public Input<int>? FingerprintType { get; set; }

        [Input("flags")]
        public Input<int>? Flags { get; set; }

        [Input("flagsnaptr")]
        public Input<string>? Flagsnaptr { get; set; }

        [Input("hardware")]
        public Input<string>? Hardware { get; set; }

        [Input("inception")]
        public Input<string>? Inception { get; set; }

        [Input("iterations")]
        public Input<int>? Iterations { get; set; }

        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("keytag")]
        public Input<int>? Keytag { get; set; }

        [Input("labels")]
        public Input<int>? Labels { get; set; }

        [Input("mailbox")]
        public Input<string>? Mailbox { get; set; }

        [Input("matchType")]
        public Input<int>? MatchType { get; set; }

        /// <summary>
        /// The name of the record. The name is an owner name, that is, the name of the node to which this resource record pertains.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nameServer")]
        public Input<string>? NameServer { get; set; }

        [Input("nextHashedOwnerName")]
        public Input<string>? NextHashedOwnerName { get; set; }

        [Input("nxdomainTtl")]
        public Input<int>? NxdomainTtl { get; set; }

        [Input("order")]
        public Input<int>? Order { get; set; }

        [Input("originalTtl")]
        public Input<int>? OriginalTtl { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("preference")]
        public Input<int>? Preference { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("priorityIncrement")]
        public Input<int>? PriorityIncrement { get; set; }

        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        [Input("recordSha")]
        public Input<string>? RecordSha { get; set; }

        [Input("recordtype")]
        public Input<string>? Recordtype { get; set; }

        [Input("refresh")]
        public Input<int>? Refresh { get; set; }

        [Input("regexp")]
        public Input<string>? Regexp { get; set; }

        [Input("replacement")]
        public Input<string>? Replacement { get; set; }

        [Input("retry")]
        public Input<int>? Retry { get; set; }

        [Input("salt")]
        public Input<string>? Salt { get; set; }

        [Input("selector")]
        public Input<int>? Selector { get; set; }

        [Input("serial")]
        public Input<int>? Serial { get; set; }

        [Input("service")]
        public Input<string>? Service { get; set; }

        [Input("signature")]
        public Input<string>? Signature { get; set; }

        [Input("signer")]
        public Input<string>? Signer { get; set; }

        [Input("software")]
        public Input<string>? Software { get; set; }

        [Input("subtype")]
        public Input<int>? Subtype { get; set; }

        [Input("targets")]
        private InputList<string>? _targets;
        public InputList<string> Targets
        {
            get => _targets ?? (_targets = new InputList<string>());
            set => _targets = value;
        }

        /// <summary>
        /// The TTL is a 32-bit signed integer that specifies the time interval that the resource record may be cached before the source of the information should be consulted again. A value of zero means that the RR can only be used for the transaction in progress, and should not be cached. Zero values can also be used for extremely volatile data.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        [Input("txt")]
        public Input<string>? Txt { get; set; }

        [Input("typeBitmaps")]
        public Input<string>? TypeBitmaps { get; set; }

        [Input("typeCovered")]
        public Input<string>? TypeCovered { get; set; }

        [Input("typeMnemonic")]
        public Input<string>? TypeMnemonic { get; set; }

        [Input("typeValue")]
        public Input<int>? TypeValue { get; set; }

        [Input("usage")]
        public Input<int>? Usage { get; set; }

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        /// <summary>
        /// Domain zone, encapsulating any nested subdomains.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DnsRecordState()
        {
        }
    }
}
