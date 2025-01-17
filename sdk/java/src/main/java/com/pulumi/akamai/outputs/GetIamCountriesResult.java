// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetIamCountriesResult {
    private List<String> countries;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;

    private GetIamCountriesResult() {}
    public List<String> countries() {
        return this.countries;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIamCountriesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> countries;
        private String id;
        public Builder() {}
        public Builder(GetIamCountriesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.countries = defaults.countries;
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder countries(List<String> countries) {
            this.countries = Objects.requireNonNull(countries);
            return this;
        }
        public Builder countries(String... countries) {
            return countries(List.of(countries));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public GetIamCountriesResult build() {
            final var o = new GetIamCountriesResult();
            o.countries = countries;
            o.id = id;
            return o;
        }
    }
}
