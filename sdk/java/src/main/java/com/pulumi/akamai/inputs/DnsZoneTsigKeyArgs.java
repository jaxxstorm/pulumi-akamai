// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class DnsZoneTsigKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final DnsZoneTsigKeyArgs Empty = new DnsZoneTsigKeyArgs();

    /**
     * The hashing algorithm.
     * 
     */
    @Import(name="algorithm", required=true)
    private Output<String> algorithm;

    /**
     * @return The hashing algorithm.
     * 
     */
    public Output<String> algorithm() {
        return this.algorithm;
    }

    /**
     * The key name.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The key name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * String known between transfer endpoints.
     * 
     */
    @Import(name="secret", required=true)
    private Output<String> secret;

    /**
     * @return String known between transfer endpoints.
     * 
     */
    public Output<String> secret() {
        return this.secret;
    }

    private DnsZoneTsigKeyArgs() {}

    private DnsZoneTsigKeyArgs(DnsZoneTsigKeyArgs $) {
        this.algorithm = $.algorithm;
        this.name = $.name;
        this.secret = $.secret;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DnsZoneTsigKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DnsZoneTsigKeyArgs $;

        public Builder() {
            $ = new DnsZoneTsigKeyArgs();
        }

        public Builder(DnsZoneTsigKeyArgs defaults) {
            $ = new DnsZoneTsigKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param algorithm The hashing algorithm.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm The hashing algorithm.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param name The key name.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The key name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param secret String known between transfer endpoints.
         * 
         * @return builder
         * 
         */
        public Builder secret(Output<String> secret) {
            $.secret = secret;
            return this;
        }

        /**
         * @param secret String known between transfer endpoints.
         * 
         * @return builder
         * 
         */
        public Builder secret(String secret) {
            return secret(Output.of(secret));
        }

        public DnsZoneTsigKeyArgs build() {
            $.algorithm = Objects.requireNonNull($.algorithm, "expected parameter 'algorithm' to be non-null");
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.secret = Objects.requireNonNull($.secret, "expected parameter 'secret' to be non-null");
            return $;
        }
    }

}
