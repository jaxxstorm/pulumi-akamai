// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.trafficmanagement.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetGtmDefaultDatacenterResult {
    private @Nullable Integer datacenter;
    private Integer datacenterId;
    private String domain;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String nickname;

    private GetGtmDefaultDatacenterResult() {}
    public Optional<Integer> datacenter() {
        return Optional.ofNullable(this.datacenter);
    }
    public Integer datacenterId() {
        return this.datacenterId;
    }
    public String domain() {
        return this.domain;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String nickname() {
        return this.nickname;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGtmDefaultDatacenterResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer datacenter;
        private Integer datacenterId;
        private String domain;
        private String id;
        private String nickname;
        public Builder() {}
        public Builder(GetGtmDefaultDatacenterResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.datacenter = defaults.datacenter;
    	      this.datacenterId = defaults.datacenterId;
    	      this.domain = defaults.domain;
    	      this.id = defaults.id;
    	      this.nickname = defaults.nickname;
        }

        @CustomType.Setter
        public Builder datacenter(@Nullable Integer datacenter) {
            this.datacenter = datacenter;
            return this;
        }
        @CustomType.Setter
        public Builder datacenterId(Integer datacenterId) {
            this.datacenterId = Objects.requireNonNull(datacenterId);
            return this;
        }
        @CustomType.Setter
        public Builder domain(String domain) {
            this.domain = Objects.requireNonNull(domain);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder nickname(String nickname) {
            this.nickname = Objects.requireNonNull(nickname);
            return this;
        }
        public GetGtmDefaultDatacenterResult build() {
            final var o = new GetGtmDefaultDatacenterResult();
            o.datacenter = datacenter;
            o.datacenterId = datacenterId;
            o.domain = domain;
            o.id = id;
            o.nickname = nickname;
            return o;
        }
    }
}
