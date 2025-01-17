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
public final class GtmGeomapDefaultDatacenter {
    /**
     * @return A unique identifier for an existing data center in the domain.
     * 
     */
    private Integer datacenterId;
    /**
     * @return A descriptive label for the group.
     * 
     */
    private @Nullable String nickname;

    private GtmGeomapDefaultDatacenter() {}
    /**
     * @return A unique identifier for an existing data center in the domain.
     * 
     */
    public Integer datacenterId() {
        return this.datacenterId;
    }
    /**
     * @return A descriptive label for the group.
     * 
     */
    public Optional<String> nickname() {
        return Optional.ofNullable(this.nickname);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GtmGeomapDefaultDatacenter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer datacenterId;
        private @Nullable String nickname;
        public Builder() {}
        public Builder(GtmGeomapDefaultDatacenter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.datacenterId = defaults.datacenterId;
    	      this.nickname = defaults.nickname;
        }

        @CustomType.Setter
        public Builder datacenterId(Integer datacenterId) {
            this.datacenterId = Objects.requireNonNull(datacenterId);
            return this;
        }
        @CustomType.Setter
        public Builder nickname(@Nullable String nickname) {
            this.nickname = nickname;
            return this;
        }
        public GtmGeomapDefaultDatacenter build() {
            final var o = new GtmGeomapDefaultDatacenter();
            o.datacenterId = datacenterId;
            o.nickname = nickname;
            return o;
        }
    }
}
