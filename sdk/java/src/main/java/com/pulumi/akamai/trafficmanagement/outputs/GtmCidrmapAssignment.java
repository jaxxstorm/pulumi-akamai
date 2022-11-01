// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.trafficmanagement.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GtmCidrmapAssignment {
    /**
     * @return Specifies an array of CIDR blocks.
     * 
     */
    private @Nullable List<String> blocks;
    /**
     * @return A unique identifier for an existing data center in the domain.
     * 
     */
    private Integer datacenterId;
    /**
     * @return A descriptive label for the CIDR zone group, up to 256 characters.
     * 
     */
    private String nickname;

    private GtmCidrmapAssignment() {}
    /**
     * @return Specifies an array of CIDR blocks.
     * 
     */
    public List<String> blocks() {
        return this.blocks == null ? List.of() : this.blocks;
    }
    /**
     * @return A unique identifier for an existing data center in the domain.
     * 
     */
    public Integer datacenterId() {
        return this.datacenterId;
    }
    /**
     * @return A descriptive label for the CIDR zone group, up to 256 characters.
     * 
     */
    public String nickname() {
        return this.nickname;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GtmCidrmapAssignment defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> blocks;
        private Integer datacenterId;
        private String nickname;
        public Builder() {}
        public Builder(GtmCidrmapAssignment defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.blocks = defaults.blocks;
    	      this.datacenterId = defaults.datacenterId;
    	      this.nickname = defaults.nickname;
        }

        @CustomType.Setter
        public Builder blocks(@Nullable List<String> blocks) {
            this.blocks = blocks;
            return this;
        }
        public Builder blocks(String... blocks) {
            return blocks(List.of(blocks));
        }
        @CustomType.Setter
        public Builder datacenterId(Integer datacenterId) {
            this.datacenterId = Objects.requireNonNull(datacenterId);
            return this;
        }
        @CustomType.Setter
        public Builder nickname(String nickname) {
            this.nickname = Objects.requireNonNull(nickname);
            return this;
        }
        public GtmCidrmapAssignment build() {
            final var o = new GtmCidrmapAssignment();
            o.blocks = blocks;
            o.datacenterId = datacenterId;
            o.nickname = nickname;
            return o;
        }
    }
}