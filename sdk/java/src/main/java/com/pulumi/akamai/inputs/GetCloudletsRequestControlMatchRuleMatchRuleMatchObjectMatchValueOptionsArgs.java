// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudletsRequestControlMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetCloudletsRequestControlMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs Empty = new GetCloudletsRequestControlMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs();

    /**
     * - (Optional) Whether the `value` argument should be evaluated based on case sensitivity.
     * 
     */
    @Import(name="valueCaseSensitive")
    private @Nullable Output<Boolean> valueCaseSensitive;

    /**
     * @return - (Optional) Whether the `value` argument should be evaluated based on case sensitivity.
     * 
     */
    public Optional<Output<Boolean>> valueCaseSensitive() {
        return Optional.ofNullable(this.valueCaseSensitive);
    }

    /**
     * - (Optional) Whether the `value` argument should be compared in an escaped form.
     * 
     */
    @Import(name="valueEscaped")
    private @Nullable Output<Boolean> valueEscaped;

    /**
     * @return - (Optional) Whether the `value` argument should be compared in an escaped form.
     * 
     */
    public Optional<Output<Boolean>> valueEscaped() {
        return Optional.ofNullable(this.valueEscaped);
    }

    /**
     * - (Optional) Whether the `value` argument includes wildcards.
     * 
     */
    @Import(name="valueHasWildcard")
    private @Nullable Output<Boolean> valueHasWildcard;

    /**
     * @return - (Optional) Whether the `value` argument includes wildcards.
     * 
     */
    public Optional<Output<Boolean>> valueHasWildcard() {
        return Optional.ofNullable(this.valueHasWildcard);
    }

    /**
     * - (Optional) If you set the `type` argument to `simple`, specify the values in the incoming request to match on.
     * 
     */
    @Import(name="values")
    private @Nullable Output<List<String>> values;

    /**
     * @return - (Optional) If you set the `type` argument to `simple`, specify the values in the incoming request to match on.
     * 
     */
    public Optional<Output<List<String>>> values() {
        return Optional.ofNullable(this.values);
    }

    private GetCloudletsRequestControlMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs() {}

    private GetCloudletsRequestControlMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs(GetCloudletsRequestControlMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs $) {
        this.valueCaseSensitive = $.valueCaseSensitive;
        this.valueEscaped = $.valueEscaped;
        this.valueHasWildcard = $.valueHasWildcard;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudletsRequestControlMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudletsRequestControlMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs $;

        public Builder() {
            $ = new GetCloudletsRequestControlMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs();
        }

        public Builder(GetCloudletsRequestControlMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs defaults) {
            $ = new GetCloudletsRequestControlMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param valueCaseSensitive - (Optional) Whether the `value` argument should be evaluated based on case sensitivity.
         * 
         * @return builder
         * 
         */
        public Builder valueCaseSensitive(@Nullable Output<Boolean> valueCaseSensitive) {
            $.valueCaseSensitive = valueCaseSensitive;
            return this;
        }

        /**
         * @param valueCaseSensitive - (Optional) Whether the `value` argument should be evaluated based on case sensitivity.
         * 
         * @return builder
         * 
         */
        public Builder valueCaseSensitive(Boolean valueCaseSensitive) {
            return valueCaseSensitive(Output.of(valueCaseSensitive));
        }

        /**
         * @param valueEscaped - (Optional) Whether the `value` argument should be compared in an escaped form.
         * 
         * @return builder
         * 
         */
        public Builder valueEscaped(@Nullable Output<Boolean> valueEscaped) {
            $.valueEscaped = valueEscaped;
            return this;
        }

        /**
         * @param valueEscaped - (Optional) Whether the `value` argument should be compared in an escaped form.
         * 
         * @return builder
         * 
         */
        public Builder valueEscaped(Boolean valueEscaped) {
            return valueEscaped(Output.of(valueEscaped));
        }

        /**
         * @param valueHasWildcard - (Optional) Whether the `value` argument includes wildcards.
         * 
         * @return builder
         * 
         */
        public Builder valueHasWildcard(@Nullable Output<Boolean> valueHasWildcard) {
            $.valueHasWildcard = valueHasWildcard;
            return this;
        }

        /**
         * @param valueHasWildcard - (Optional) Whether the `value` argument includes wildcards.
         * 
         * @return builder
         * 
         */
        public Builder valueHasWildcard(Boolean valueHasWildcard) {
            return valueHasWildcard(Output.of(valueHasWildcard));
        }

        /**
         * @param values - (Optional) If you set the `type` argument to `simple`, specify the values in the incoming request to match on.
         * 
         * @return builder
         * 
         */
        public Builder values(@Nullable Output<List<String>> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values - (Optional) If you set the `type` argument to `simple`, specify the values in the incoming request to match on.
         * 
         * @return builder
         * 
         */
        public Builder values(List<String> values) {
            return values(Output.of(values));
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

        public GetCloudletsRequestControlMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs build() {
            return $;
        }
    }

}
