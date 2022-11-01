// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CpsDvEnrollmentDnsChallenge {
    private @Nullable String domain;
    private @Nullable String fullPath;
    private @Nullable String responseBody;

    private CpsDvEnrollmentDnsChallenge() {}
    public Optional<String> domain() {
        return Optional.ofNullable(this.domain);
    }
    public Optional<String> fullPath() {
        return Optional.ofNullable(this.fullPath);
    }
    public Optional<String> responseBody() {
        return Optional.ofNullable(this.responseBody);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CpsDvEnrollmentDnsChallenge defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String domain;
        private @Nullable String fullPath;
        private @Nullable String responseBody;
        public Builder() {}
        public Builder(CpsDvEnrollmentDnsChallenge defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.domain = defaults.domain;
    	      this.fullPath = defaults.fullPath;
    	      this.responseBody = defaults.responseBody;
        }

        @CustomType.Setter
        public Builder domain(@Nullable String domain) {
            this.domain = domain;
            return this;
        }
        @CustomType.Setter
        public Builder fullPath(@Nullable String fullPath) {
            this.fullPath = fullPath;
            return this;
        }
        @CustomType.Setter
        public Builder responseBody(@Nullable String responseBody) {
            this.responseBody = responseBody;
            return this;
        }
        public CpsDvEnrollmentDnsChallenge build() {
            final var o = new CpsDvEnrollmentDnsChallenge();
            o.domain = domain;
            o.fullPath = fullPath;
            o.responseBody = responseBody;
            return o;
        }
    }
}