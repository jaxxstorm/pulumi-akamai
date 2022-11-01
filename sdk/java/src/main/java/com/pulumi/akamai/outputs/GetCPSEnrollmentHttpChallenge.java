// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCPSEnrollmentHttpChallenge {
    private String domain;
    private String fullPath;
    private String responseBody;

    private GetCPSEnrollmentHttpChallenge() {}
    public String domain() {
        return this.domain;
    }
    public String fullPath() {
        return this.fullPath;
    }
    public String responseBody() {
        return this.responseBody;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCPSEnrollmentHttpChallenge defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String domain;
        private String fullPath;
        private String responseBody;
        public Builder() {}
        public Builder(GetCPSEnrollmentHttpChallenge defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.domain = defaults.domain;
    	      this.fullPath = defaults.fullPath;
    	      this.responseBody = defaults.responseBody;
        }

        @CustomType.Setter
        public Builder domain(String domain) {
            this.domain = Objects.requireNonNull(domain);
            return this;
        }
        @CustomType.Setter
        public Builder fullPath(String fullPath) {
            this.fullPath = Objects.requireNonNull(fullPath);
            return this;
        }
        @CustomType.Setter
        public Builder responseBody(String responseBody) {
            this.responseBody = Objects.requireNonNull(responseBody);
            return this;
        }
        public GetCPSEnrollmentHttpChallenge build() {
            final var o = new GetCPSEnrollmentHttpChallenge();
            o.domain = domain;
            o.fullPath = fullPath;
            o.responseBody = responseBody;
            return o;
        }
    }
}