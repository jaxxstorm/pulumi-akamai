// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PropertyOrigin {
    private @Nullable String cacheKeyHostname;
    private @Nullable Boolean compress;
    private @Nullable Boolean enableTrueClientIp;
    private @Nullable String forwardHostname;
    private @Nullable String hostname;
    private @Nullable Integer port;

    private PropertyOrigin() {}
    public Optional<String> cacheKeyHostname() {
        return Optional.ofNullable(this.cacheKeyHostname);
    }
    public Optional<Boolean> compress() {
        return Optional.ofNullable(this.compress);
    }
    public Optional<Boolean> enableTrueClientIp() {
        return Optional.ofNullable(this.enableTrueClientIp);
    }
    public Optional<String> forwardHostname() {
        return Optional.ofNullable(this.forwardHostname);
    }
    public Optional<String> hostname() {
        return Optional.ofNullable(this.hostname);
    }
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PropertyOrigin defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String cacheKeyHostname;
        private @Nullable Boolean compress;
        private @Nullable Boolean enableTrueClientIp;
        private @Nullable String forwardHostname;
        private @Nullable String hostname;
        private @Nullable Integer port;
        public Builder() {}
        public Builder(PropertyOrigin defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cacheKeyHostname = defaults.cacheKeyHostname;
    	      this.compress = defaults.compress;
    	      this.enableTrueClientIp = defaults.enableTrueClientIp;
    	      this.forwardHostname = defaults.forwardHostname;
    	      this.hostname = defaults.hostname;
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder cacheKeyHostname(@Nullable String cacheKeyHostname) {
            this.cacheKeyHostname = cacheKeyHostname;
            return this;
        }
        @CustomType.Setter
        public Builder compress(@Nullable Boolean compress) {
            this.compress = compress;
            return this;
        }
        @CustomType.Setter
        public Builder enableTrueClientIp(@Nullable Boolean enableTrueClientIp) {
            this.enableTrueClientIp = enableTrueClientIp;
            return this;
        }
        @CustomType.Setter
        public Builder forwardHostname(@Nullable String forwardHostname) {
            this.forwardHostname = forwardHostname;
            return this;
        }
        @CustomType.Setter
        public Builder hostname(@Nullable String hostname) {
            this.hostname = hostname;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {
            this.port = port;
            return this;
        }
        public PropertyOrigin build() {
            final var o = new PropertyOrigin();
            o.cacheKeyHostname = cacheKeyHostname;
            o.compress = compress;
            o.enableTrueClientIp = enableTrueClientIp;
            o.forwardHostname = forwardHostname;
            o.hostname = hostname;
            o.port = port;
            return o;
        }
    }
}
