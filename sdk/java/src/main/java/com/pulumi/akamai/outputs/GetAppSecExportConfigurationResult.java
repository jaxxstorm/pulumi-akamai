// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetAppSecExportConfigurationResult {
    private Integer configId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String json;
    private String outputText;
    private @Nullable List<String> searches;
    private Integer version;

    private GetAppSecExportConfigurationResult() {}
    public Integer configId() {
        return this.configId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String json() {
        return this.json;
    }
    public String outputText() {
        return this.outputText;
    }
    public List<String> searches() {
        return this.searches == null ? List.of() : this.searches;
    }
    public Integer version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppSecExportConfigurationResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer configId;
        private String id;
        private String json;
        private String outputText;
        private @Nullable List<String> searches;
        private Integer version;
        public Builder() {}
        public Builder(GetAppSecExportConfigurationResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.configId = defaults.configId;
    	      this.id = defaults.id;
    	      this.json = defaults.json;
    	      this.outputText = defaults.outputText;
    	      this.searches = defaults.searches;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder configId(Integer configId) {
            this.configId = Objects.requireNonNull(configId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder json(String json) {
            this.json = Objects.requireNonNull(json);
            return this;
        }
        @CustomType.Setter
        public Builder outputText(String outputText) {
            this.outputText = Objects.requireNonNull(outputText);
            return this;
        }
        @CustomType.Setter
        public Builder searches(@Nullable List<String> searches) {
            this.searches = searches;
            return this;
        }
        public Builder searches(String... searches) {
            return searches(List.of(searches));
        }
        @CustomType.Setter
        public Builder version(Integer version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetAppSecExportConfigurationResult build() {
            final var o = new GetAppSecExportConfigurationResult();
            o.configId = configId;
            o.id = id;
            o.json = json;
            o.outputText = outputText;
            o.searches = searches;
            o.version = version;
            return o;
        }
    }
}