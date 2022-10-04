// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAppSecCustomDenyResult {
    private Integer configId;
    private @Nullable String customDenyId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String json;
    private String outputText;

    private GetAppSecCustomDenyResult() {}
    public Integer configId() {
        return this.configId;
    }
    public Optional<String> customDenyId() {
        return Optional.ofNullable(this.customDenyId);
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppSecCustomDenyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer configId;
        private @Nullable String customDenyId;
        private String id;
        private String json;
        private String outputText;
        public Builder() {}
        public Builder(GetAppSecCustomDenyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.configId = defaults.configId;
    	      this.customDenyId = defaults.customDenyId;
    	      this.id = defaults.id;
    	      this.json = defaults.json;
    	      this.outputText = defaults.outputText;
        }

        @CustomType.Setter
        public Builder configId(Integer configId) {
            this.configId = Objects.requireNonNull(configId);
            return this;
        }
        @CustomType.Setter
        public Builder customDenyId(@Nullable String customDenyId) {
            this.customDenyId = customDenyId;
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
        public GetAppSecCustomDenyResult build() {
            final var o = new GetAppSecCustomDenyResult();
            o.configId = configId;
            o.customDenyId = customDenyId;
            o.id = id;
            o.json = json;
            o.outputText = outputText;
            return o;
        }
    }
}
