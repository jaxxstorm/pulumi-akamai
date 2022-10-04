// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions Empty = new GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions();

    /**
     * - (Optional) Whether the `value` argument should be evaluated based on case sensitivity.
     * 
     */
    @Import(name="valueCaseSensitive")
    private @Nullable Boolean valueCaseSensitive;

    /**
     * @return - (Optional) Whether the `value` argument should be evaluated based on case sensitivity.
     * 
     */
    public Optional<Boolean> valueCaseSensitive() {
        return Optional.ofNullable(this.valueCaseSensitive);
    }

    /**
     * - (Optional) Whether the `value` argument should be compared in an escaped form.
     * 
     */
    @Import(name="valueEscaped")
    private @Nullable Boolean valueEscaped;

    /**
     * @return - (Optional) Whether the `value` argument should be compared in an escaped form.
     * 
     */
    public Optional<Boolean> valueEscaped() {
        return Optional.ofNullable(this.valueEscaped);
    }

    /**
     * - (Optional) Whether the `value` argument includes wildcards.
     * 
     */
    @Import(name="valueHasWildcard")
    private @Nullable Boolean valueHasWildcard;

    /**
     * @return - (Optional) Whether the `value` argument includes wildcards.
     * 
     */
    public Optional<Boolean> valueHasWildcard() {
        return Optional.ofNullable(this.valueHasWildcard);
    }

    /**
     * - (Optional) If you set the `type` argument to `simple`, specify the values in the incoming request to match on.
     * 
     */
    @Import(name="values")
    private @Nullable List<String> values;

    /**
     * @return - (Optional) If you set the `type` argument to `simple`, specify the values in the incoming request to match on.
     * 
     */
    public Optional<List<String>> values() {
        return Optional.ofNullable(this.values);
    }

    private GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions() {}

    private GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions(GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions $) {
        this.valueCaseSensitive = $.valueCaseSensitive;
        this.valueEscaped = $.valueEscaped;
        this.valueHasWildcard = $.valueHasWildcard;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions $;

        public Builder() {
            $ = new GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions();
        }

        public Builder(GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions defaults) {
            $ = new GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions(Objects.requireNonNull(defaults));
        }

        /**
         * @param valueCaseSensitive - (Optional) Whether the `value` argument should be evaluated based on case sensitivity.
         * 
         * @return builder
         * 
         */
        public Builder valueCaseSensitive(@Nullable Boolean valueCaseSensitive) {
            $.valueCaseSensitive = valueCaseSensitive;
            return this;
        }

        /**
         * @param valueEscaped - (Optional) Whether the `value` argument should be compared in an escaped form.
         * 
         * @return builder
         * 
         */
        public Builder valueEscaped(@Nullable Boolean valueEscaped) {
            $.valueEscaped = valueEscaped;
            return this;
        }

        /**
         * @param valueHasWildcard - (Optional) Whether the `value` argument includes wildcards.
         * 
         * @return builder
         * 
         */
        public Builder valueHasWildcard(@Nullable Boolean valueHasWildcard) {
            $.valueHasWildcard = valueHasWildcard;
            return this;
        }

        /**
         * @param values - (Optional) If you set the `type` argument to `simple`, specify the values in the incoming request to match on.
         * 
         * @return builder
         * 
         */
        public Builder values(@Nullable List<String> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values - (Optional) If you set the `type` argument to `simple`, specify the values in the incoming request to match on.
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        public GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueOptions build() {
            return $;
        }
    }

}
