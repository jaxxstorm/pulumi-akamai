// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.akamai.inputs.GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueArgs Empty = new GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueArgs();

    /**
     * - (Optional) If you&#39;re using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return - (Optional) If you&#39;re using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * - (Optional) Whether the `name` argument should be evaluated based on case sensitivity.
     * 
     */
    @Import(name="nameCaseSensitive")
    private @Nullable Output<Boolean> nameCaseSensitive;

    /**
     * @return - (Optional) Whether the `name` argument should be evaluated based on case sensitivity.
     * 
     */
    public Optional<Output<Boolean>> nameCaseSensitive() {
        return Optional.ofNullable(this.nameCaseSensitive);
    }

    /**
     * - (Optional) Whether the `name` argument includes wildcards.
     * 
     */
    @Import(name="nameHasWildcard")
    private @Nullable Output<Boolean> nameHasWildcard;

    /**
     * @return - (Optional) Whether the `name` argument includes wildcards.
     * 
     */
    public Optional<Output<Boolean>> nameHasWildcard() {
        return Optional.ofNullable(this.nameHasWildcard);
    }

    /**
     * - (Optional) If you set the `type` argument to `object`, use this array to list the values to match on.
     * 
     */
    @Import(name="options")
    private @Nullable Output<GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs> options;

    /**
     * @return - (Optional) If you set the `type` argument to `object`, use this array to list the values to match on.
     * 
     */
    public Optional<Output<GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs>> options() {
        return Optional.ofNullable(this.options);
    }

    /**
     * - (Required) The type of the array, either `object` or `simple`. Use the `simple` option when adding only an array of string-based values.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return - (Required) The type of the array, either `object` or `simple`. Use the `simple` option when adding only an array of string-based values.
     * 
     */
    public Output<String> type() {
        return this.type;
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

    private GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueArgs() {}

    private GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueArgs(GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueArgs $) {
        this.name = $.name;
        this.nameCaseSensitive = $.nameCaseSensitive;
        this.nameHasWildcard = $.nameHasWildcard;
        this.options = $.options;
        this.type = $.type;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueArgs $;

        public Builder() {
            $ = new GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueArgs();
        }

        public Builder(GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueArgs defaults) {
            $ = new GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name - (Optional) If you&#39;re using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name - (Optional) If you&#39;re using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nameCaseSensitive - (Optional) Whether the `name` argument should be evaluated based on case sensitivity.
         * 
         * @return builder
         * 
         */
        public Builder nameCaseSensitive(@Nullable Output<Boolean> nameCaseSensitive) {
            $.nameCaseSensitive = nameCaseSensitive;
            return this;
        }

        /**
         * @param nameCaseSensitive - (Optional) Whether the `name` argument should be evaluated based on case sensitivity.
         * 
         * @return builder
         * 
         */
        public Builder nameCaseSensitive(Boolean nameCaseSensitive) {
            return nameCaseSensitive(Output.of(nameCaseSensitive));
        }

        /**
         * @param nameHasWildcard - (Optional) Whether the `name` argument includes wildcards.
         * 
         * @return builder
         * 
         */
        public Builder nameHasWildcard(@Nullable Output<Boolean> nameHasWildcard) {
            $.nameHasWildcard = nameHasWildcard;
            return this;
        }

        /**
         * @param nameHasWildcard - (Optional) Whether the `name` argument includes wildcards.
         * 
         * @return builder
         * 
         */
        public Builder nameHasWildcard(Boolean nameHasWildcard) {
            return nameHasWildcard(Output.of(nameHasWildcard));
        }

        /**
         * @param options - (Optional) If you set the `type` argument to `object`, use this array to list the values to match on.
         * 
         * @return builder
         * 
         */
        public Builder options(@Nullable Output<GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs> options) {
            $.options = options;
            return this;
        }

        /**
         * @param options - (Optional) If you set the `type` argument to `object`, use this array to list the values to match on.
         * 
         * @return builder
         * 
         */
        public Builder options(GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs options) {
            return options(Output.of(options));
        }

        /**
         * @param type - (Required) The type of the array, either `object` or `simple`. Use the `simple` option when adding only an array of string-based values.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type - (Required) The type of the array, either `object` or `simple`. Use the `simple` option when adding only an array of string-based values.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
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

        public GetCloudletsForwardRewriteMatchRuleMatchRuleMatchObjectMatchValueArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
