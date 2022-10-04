// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetAppSecSelectedHostnamesResult {
    private Integer configId;
    private List<String> hostnames;
    private String hostnamesJson;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String outputText;

    private GetAppSecSelectedHostnamesResult() {}
    public Integer configId() {
        return this.configId;
    }
    public List<String> hostnames() {
        return this.hostnames;
    }
    public String hostnamesJson() {
        return this.hostnamesJson;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String outputText() {
        return this.outputText;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppSecSelectedHostnamesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer configId;
        private List<String> hostnames;
        private String hostnamesJson;
        private String id;
        private String outputText;
        public Builder() {}
        public Builder(GetAppSecSelectedHostnamesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.configId = defaults.configId;
    	      this.hostnames = defaults.hostnames;
    	      this.hostnamesJson = defaults.hostnamesJson;
    	      this.id = defaults.id;
    	      this.outputText = defaults.outputText;
        }

        @CustomType.Setter
        public Builder configId(Integer configId) {
            this.configId = Objects.requireNonNull(configId);
            return this;
        }
        @CustomType.Setter
        public Builder hostnames(List<String> hostnames) {
            this.hostnames = Objects.requireNonNull(hostnames);
            return this;
        }
        public Builder hostnames(String... hostnames) {
            return hostnames(List.of(hostnames));
        }
        @CustomType.Setter
        public Builder hostnamesJson(String hostnamesJson) {
            this.hostnamesJson = Objects.requireNonNull(hostnamesJson);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder outputText(String outputText) {
            this.outputText = Objects.requireNonNull(outputText);
            return this;
        }
        public GetAppSecSelectedHostnamesResult build() {
            final var o = new GetAppSecSelectedHostnamesResult();
            o.configId = configId;
            o.hostnames = hostnames;
            o.hostnamesJson = hostnamesJson;
            o.id = id;
            o.outputText = outputText;
            return o;
        }
    }
}
