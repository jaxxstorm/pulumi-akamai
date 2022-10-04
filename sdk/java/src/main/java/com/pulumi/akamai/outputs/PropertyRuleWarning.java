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
public final class PropertyRuleWarning {
    private @Nullable String behaviorName;
    private @Nullable String detail;
    private @Nullable String errorLocation;
    private @Nullable String instance;
    private @Nullable Integer statusCode;
    private @Nullable String title;
    private @Nullable String type;

    private PropertyRuleWarning() {}
    public Optional<String> behaviorName() {
        return Optional.ofNullable(this.behaviorName);
    }
    public Optional<String> detail() {
        return Optional.ofNullable(this.detail);
    }
    public Optional<String> errorLocation() {
        return Optional.ofNullable(this.errorLocation);
    }
    public Optional<String> instance() {
        return Optional.ofNullable(this.instance);
    }
    public Optional<Integer> statusCode() {
        return Optional.ofNullable(this.statusCode);
    }
    public Optional<String> title() {
        return Optional.ofNullable(this.title);
    }
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PropertyRuleWarning defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String behaviorName;
        private @Nullable String detail;
        private @Nullable String errorLocation;
        private @Nullable String instance;
        private @Nullable Integer statusCode;
        private @Nullable String title;
        private @Nullable String type;
        public Builder() {}
        public Builder(PropertyRuleWarning defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.behaviorName = defaults.behaviorName;
    	      this.detail = defaults.detail;
    	      this.errorLocation = defaults.errorLocation;
    	      this.instance = defaults.instance;
    	      this.statusCode = defaults.statusCode;
    	      this.title = defaults.title;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder behaviorName(@Nullable String behaviorName) {
            this.behaviorName = behaviorName;
            return this;
        }
        @CustomType.Setter
        public Builder detail(@Nullable String detail) {
            this.detail = detail;
            return this;
        }
        @CustomType.Setter
        public Builder errorLocation(@Nullable String errorLocation) {
            this.errorLocation = errorLocation;
            return this;
        }
        @CustomType.Setter
        public Builder instance(@Nullable String instance) {
            this.instance = instance;
            return this;
        }
        @CustomType.Setter
        public Builder statusCode(@Nullable Integer statusCode) {
            this.statusCode = statusCode;
            return this;
        }
        @CustomType.Setter
        public Builder title(@Nullable String title) {
            this.title = title;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        public PropertyRuleWarning build() {
            final var o = new PropertyRuleWarning();
            o.behaviorName = behaviorName;
            o.detail = detail;
            o.errorLocation = errorLocation;
            o.instance = instance;
            o.statusCode = statusCode;
            o.title = title;
            o.type = type;
            return o;
        }
    }
}
