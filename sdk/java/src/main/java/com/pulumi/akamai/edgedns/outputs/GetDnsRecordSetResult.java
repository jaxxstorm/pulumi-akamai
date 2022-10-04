// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.edgedns.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDnsRecordSetResult {
    private String host;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> rdatas;
    private String recordType;
    private String zone;

    private GetDnsRecordSetResult() {}
    public String host() {
        return this.host;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<String> rdatas() {
        return this.rdatas;
    }
    public String recordType() {
        return this.recordType;
    }
    public String zone() {
        return this.zone;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDnsRecordSetResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String host;
        private String id;
        private List<String> rdatas;
        private String recordType;
        private String zone;
        public Builder() {}
        public Builder(GetDnsRecordSetResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.host = defaults.host;
    	      this.id = defaults.id;
    	      this.rdatas = defaults.rdatas;
    	      this.recordType = defaults.recordType;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder host(String host) {
            this.host = Objects.requireNonNull(host);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder rdatas(List<String> rdatas) {
            this.rdatas = Objects.requireNonNull(rdatas);
            return this;
        }
        public Builder rdatas(String... rdatas) {
            return rdatas(List.of(rdatas));
        }
        @CustomType.Setter
        public Builder recordType(String recordType) {
            this.recordType = Objects.requireNonNull(recordType);
            return this;
        }
        @CustomType.Setter
        public Builder zone(String zone) {
            this.zone = Objects.requireNonNull(zone);
            return this;
        }
        public GetDnsRecordSetResult build() {
            final var o = new GetDnsRecordSetResult();
            o.host = host;
            o.id = id;
            o.rdatas = rdatas;
            o.recordType = recordType;
            o.zone = zone;
            return o;
        }
    }
}
