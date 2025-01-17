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
public final class GetAppSecMatchTargetsResult {
    private Integer configId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String json;
    private @Nullable Integer matchTargetId;
    private String outputText;

    private GetAppSecMatchTargetsResult() {}
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
    public Optional<Integer> matchTargetId() {
        return Optional.ofNullable(this.matchTargetId);
    }
    public String outputText() {
        return this.outputText;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppSecMatchTargetsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer configId;
        private String id;
        private String json;
        private @Nullable Integer matchTargetId;
        private String outputText;
        public Builder() {}
        public Builder(GetAppSecMatchTargetsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.configId = defaults.configId;
    	      this.id = defaults.id;
    	      this.json = defaults.json;
    	      this.matchTargetId = defaults.matchTargetId;
    	      this.outputText = defaults.outputText;
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
        public Builder matchTargetId(@Nullable Integer matchTargetId) {
            this.matchTargetId = matchTargetId;
            return this;
        }
        @CustomType.Setter
        public Builder outputText(String outputText) {
            this.outputText = Objects.requireNonNull(outputText);
            return this;
        }
        public GetAppSecMatchTargetsResult build() {
            final var o = new GetAppSecMatchTargetsResult();
            o.configId = configId;
            o.id = id;
            o.json = json;
            o.matchTargetId = matchTargetId;
            o.outputText = outputText;
            return o;
        }
    }
}
