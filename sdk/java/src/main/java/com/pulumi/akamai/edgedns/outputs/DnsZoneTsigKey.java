// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.edgedns.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class DnsZoneTsigKey {
    /**
     * @return The hashing algorithm.
     * 
     */
    private String algorithm;
    /**
     * @return The key name.
     * 
     */
    private String name;
    /**
     * @return String known between transfer endpoints.
     * 
     */
    private String secret;

    private DnsZoneTsigKey() {}
    /**
     * @return The hashing algorithm.
     * 
     */
    public String algorithm() {
        return this.algorithm;
    }
    /**
     * @return The key name.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return String known between transfer endpoints.
     * 
     */
    public String secret() {
        return this.secret;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DnsZoneTsigKey defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String algorithm;
        private String name;
        private String secret;
        public Builder() {}
        public Builder(DnsZoneTsigKey defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.algorithm = defaults.algorithm;
    	      this.name = defaults.name;
    	      this.secret = defaults.secret;
        }

        @CustomType.Setter
        public Builder algorithm(String algorithm) {
            this.algorithm = Objects.requireNonNull(algorithm);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder secret(String secret) {
            this.secret = Objects.requireNonNull(secret);
            return this;
        }
        public DnsZoneTsigKey build() {
            final var o = new DnsZoneTsigKey();
            o.algorithm = algorithm;
            o.name = name;
            o.secret = secret;
            return o;
        }
    }
}
