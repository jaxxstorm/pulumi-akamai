// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.akamai.outputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValue {
    /**
     * @return - (Optional) If you&#39;re using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
     * 
     */
    private @Nullable String name;
    /**
     * @return - (Optional) Whether the `name` argument should be evaluated based on case sensitivity.
     * 
     */
    private @Nullable Boolean nameCaseSensitive;
    /**
     * @return - (Optional) Whether the `name` argument includes wildcards.
     * 
     */
    private @Nullable Boolean nameHasWildcard;
    /**
     * @return - (Optional) If you set the `type` argument to `object`, use this array to list the values to match on.
     * 
     */
    private @Nullable GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions options;
    /**
     * @return - (Required) The type of the array, either `object` or `simple`. Use the `simple` option when adding only an array of string-based values.
     * 
     */
    private String type;
    /**
     * @return - (Optional) If you set the `type` argument to `simple`, specify the values in the incoming request to match on.
     * 
     */
    private @Nullable List<String> values;

    private GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValue() {}
    /**
     * @return - (Optional) If you&#39;re using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return - (Optional) Whether the `name` argument should be evaluated based on case sensitivity.
     * 
     */
    public Optional<Boolean> nameCaseSensitive() {
        return Optional.ofNullable(this.nameCaseSensitive);
    }
    /**
     * @return - (Optional) Whether the `name` argument includes wildcards.
     * 
     */
    public Optional<Boolean> nameHasWildcard() {
        return Optional.ofNullable(this.nameHasWildcard);
    }
    /**
     * @return - (Optional) If you set the `type` argument to `object`, use this array to list the values to match on.
     * 
     */
    public Optional<GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions> options() {
        return Optional.ofNullable(this.options);
    }
    /**
     * @return - (Required) The type of the array, either `object` or `simple`. Use the `simple` option when adding only an array of string-based values.
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return - (Optional) If you set the `type` argument to `simple`, specify the values in the incoming request to match on.
     * 
     */
    public List<String> values() {
        return this.values == null ? List.of() : this.values;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValue defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private @Nullable Boolean nameCaseSensitive;
        private @Nullable Boolean nameHasWildcard;
        private @Nullable GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions options;
        private String type;
        private @Nullable List<String> values;
        public Builder() {}
        public Builder(GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValue defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.nameCaseSensitive = defaults.nameCaseSensitive;
    	      this.nameHasWildcard = defaults.nameHasWildcard;
    	      this.options = defaults.options;
    	      this.type = defaults.type;
    	      this.values = defaults.values;
        }

        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder nameCaseSensitive(@Nullable Boolean nameCaseSensitive) {
            this.nameCaseSensitive = nameCaseSensitive;
            return this;
        }
        @CustomType.Setter
        public Builder nameHasWildcard(@Nullable Boolean nameHasWildcard) {
            this.nameHasWildcard = nameHasWildcard;
            return this;
        }
        @CustomType.Setter
        public Builder options(@Nullable GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions options) {
            this.options = options;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
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
        public GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValue build() {
            final var o = new GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValue();
            o.name = name;
            o.nameCaseSensitive = nameCaseSensitive;
            o.nameHasWildcard = nameHasWildcard;
            o.options = options;
            o.type = type;
            o.values = values;
            return o;
        }
    }
}
