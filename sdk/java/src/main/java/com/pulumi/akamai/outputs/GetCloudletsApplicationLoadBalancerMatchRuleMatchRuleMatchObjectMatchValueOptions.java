// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchObjectMatchValueOptions {
    /**
     * @return - (Optional) Whether the `value` argument should be evaluated based on case sensitivity.
     * 
     */
    private @Nullable Boolean valueCaseSensitive;
    /**
     * @return - (Optional) Whether the `value` argument should be compared in an escaped form.
     * 
     */
    private @Nullable Boolean valueEscaped;
    /**
     * @return - (Optional) Whether the `value` argument includes wildcards.
     * 
     */
    private @Nullable Boolean valueHasWildcard;
    /**
     * @return - (Optional) If you set the `type` argument to `simple` or `range`, specify the values in the incoming request to match on. With `range`, you can only specify an array of integers, for example `[1, 2]`.
     * 
     */
    private @Nullable List<String> values;

    private GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchObjectMatchValueOptions() {}
    /**
     * @return - (Optional) Whether the `value` argument should be evaluated based on case sensitivity.
     * 
     */
    public Optional<Boolean> valueCaseSensitive() {
        return Optional.ofNullable(this.valueCaseSensitive);
    }
    /**
     * @return - (Optional) Whether the `value` argument should be compared in an escaped form.
     * 
     */
    public Optional<Boolean> valueEscaped() {
        return Optional.ofNullable(this.valueEscaped);
    }
    /**
     * @return - (Optional) Whether the `value` argument includes wildcards.
     * 
     */
    public Optional<Boolean> valueHasWildcard() {
        return Optional.ofNullable(this.valueHasWildcard);
    }
    /**
     * @return - (Optional) If you set the `type` argument to `simple` or `range`, specify the values in the incoming request to match on. With `range`, you can only specify an array of integers, for example `[1, 2]`.
     * 
     */
    public List<String> values() {
        return this.values == null ? List.of() : this.values;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchObjectMatchValueOptions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean valueCaseSensitive;
        private @Nullable Boolean valueEscaped;
        private @Nullable Boolean valueHasWildcard;
        private @Nullable List<String> values;
        public Builder() {}
        public Builder(GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchObjectMatchValueOptions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.valueCaseSensitive = defaults.valueCaseSensitive;
    	      this.valueEscaped = defaults.valueEscaped;
    	      this.valueHasWildcard = defaults.valueHasWildcard;
    	      this.values = defaults.values;
        }

        @CustomType.Setter
        public Builder valueCaseSensitive(@Nullable Boolean valueCaseSensitive) {
            this.valueCaseSensitive = valueCaseSensitive;
            return this;
        }
        @CustomType.Setter
        public Builder valueEscaped(@Nullable Boolean valueEscaped) {
            this.valueEscaped = valueEscaped;
            return this;
        }
        @CustomType.Setter
        public Builder valueHasWildcard(@Nullable Boolean valueHasWildcard) {
            this.valueHasWildcard = valueHasWildcard;
            return this;
        }
        @CustomType.Setter
        public Builder values(@Nullable List<String> values) {
            this.values = values;
            return this;
        }
        public Builder values(String... values) {
            return values(List.of(values));
        }
        public GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchObjectMatchValueOptions build() {
            final var o = new GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchObjectMatchValueOptions();
            o.valueCaseSensitive = valueCaseSensitive;
            o.valueEscaped = valueEscaped;
            o.valueHasWildcard = valueHasWildcard;
            o.values = values;
            return o;
        }
    }
}
