// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.akamai.inputs.GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSettingArgs;
import com.pulumi.akamai.inputs.GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArgs Empty = new GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArgs();

    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * - (Optional) The end time for this match. Specify the value in UTC in seconds since the epoch.
     * 
     */
    @Import(name="end")
    private @Nullable Output<Integer> end;

    /**
     * @return - (Optional) The end time for this match. Specify the value in UTC in seconds since the epoch.
     * 
     */
    public Optional<Output<Integer>> end() {
        return Optional.ofNullable(this.end);
    }

    /**
     * - (Required) Defines data used to construct a new request URL if all conditions are met. If all of the conditions you set are true, the Edge Server returns an HTTP response from the rewritten URL.
     * 
     */
    @Import(name="forwardSettings", required=true)
    private Output<List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSettingArgs>> forwardSettings;

    /**
     * @return - (Required) Defines data used to construct a new request URL if all conditions are met. If all of the conditions you set are true, the Edge Server returns an HTTP response from the rewritten URL.
     * 
     */
    public Output<List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSettingArgs>> forwardSettings() {
        return this.forwardSettings;
    }

    /**
     * - (Optional) An identifier for Akamai internal use only.
     * 
     */
    @Import(name="id")
    private @Nullable Output<Integer> id;

    /**
     * @return - (Optional) An identifier for Akamai internal use only.
     * 
     */
    public Optional<Output<Integer>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * - (Optional) The URL that the Cloudlet uses to match the incoming request.
     * 
     */
    @Import(name="matchUrl")
    private @Nullable Output<String> matchUrl;

    /**
     * @return - (Optional) The URL that the Cloudlet uses to match the incoming request.
     * 
     */
    public Optional<Output<String>> matchUrl() {
        return Optional.ofNullable(this.matchUrl);
    }

    /**
     * - (Optional) A list of conditions to apply to a Cloudlet, including:
     * 
     */
    @Import(name="matches")
    private @Nullable Output<List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchArgs>> matches;

    /**
     * @return - (Optional) A list of conditions to apply to a Cloudlet, including:
     * 
     */
    public Optional<Output<List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchArgs>>> matches() {
        return Optional.ofNullable(this.matches);
    }

    /**
     * - (Optional) Whether the match supports default rules that apply to all requests.
     * 
     */
    @Import(name="matchesAlways")
    private @Nullable Output<Boolean> matchesAlways;

    /**
     * @return - (Optional) Whether the match supports default rules that apply to all requests.
     * 
     */
    public Optional<Output<Boolean>> matchesAlways() {
        return Optional.ofNullable(this.matchesAlways);
    }

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
     * - (Optional) The start time for this match. Specify the value in UTC in seconds since the epoch.
     * 
     */
    @Import(name="start")
    private @Nullable Output<Integer> start;

    /**
     * @return - (Optional) The start time for this match. Specify the value in UTC in seconds since the epoch.
     * 
     */
    public Optional<Output<Integer>> start() {
        return Optional.ofNullable(this.start);
    }

    /**
     * - (Required) The type of the array, either `object`, `range`, or `simple`. Use the `simple` option when adding only an array of string-based values.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return - (Required) The type of the array, either `object`, `range`, or `simple`. Use the `simple` option when adding only an array of string-based values.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArgs() {}

    private GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArgs(GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArgs $) {
        this.disabled = $.disabled;
        this.end = $.end;
        this.forwardSettings = $.forwardSettings;
        this.id = $.id;
        this.matchUrl = $.matchUrl;
        this.matches = $.matches;
        this.matchesAlways = $.matchesAlways;
        this.name = $.name;
        this.start = $.start;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArgs $;

        public Builder() {
            $ = new GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArgs();
        }

        public Builder(GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArgs defaults) {
            $ = new GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArgs(Objects.requireNonNull(defaults));
        }

        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param end - (Optional) The end time for this match. Specify the value in UTC in seconds since the epoch.
         * 
         * @return builder
         * 
         */
        public Builder end(@Nullable Output<Integer> end) {
            $.end = end;
            return this;
        }

        /**
         * @param end - (Optional) The end time for this match. Specify the value in UTC in seconds since the epoch.
         * 
         * @return builder
         * 
         */
        public Builder end(Integer end) {
            return end(Output.of(end));
        }

        /**
         * @param forwardSettings - (Required) Defines data used to construct a new request URL if all conditions are met. If all of the conditions you set are true, the Edge Server returns an HTTP response from the rewritten URL.
         * 
         * @return builder
         * 
         */
        public Builder forwardSettings(Output<List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSettingArgs>> forwardSettings) {
            $.forwardSettings = forwardSettings;
            return this;
        }

        /**
         * @param forwardSettings - (Required) Defines data used to construct a new request URL if all conditions are met. If all of the conditions you set are true, the Edge Server returns an HTTP response from the rewritten URL.
         * 
         * @return builder
         * 
         */
        public Builder forwardSettings(List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSettingArgs> forwardSettings) {
            return forwardSettings(Output.of(forwardSettings));
        }

        /**
         * @param forwardSettings - (Required) Defines data used to construct a new request URL if all conditions are met. If all of the conditions you set are true, the Edge Server returns an HTTP response from the rewritten URL.
         * 
         * @return builder
         * 
         */
        public Builder forwardSettings(GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSettingArgs... forwardSettings) {
            return forwardSettings(List.of(forwardSettings));
        }

        /**
         * @param id - (Optional) An identifier for Akamai internal use only.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<Integer> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id - (Optional) An identifier for Akamai internal use only.
         * 
         * @return builder
         * 
         */
        public Builder id(Integer id) {
            return id(Output.of(id));
        }

        /**
         * @param matchUrl - (Optional) The URL that the Cloudlet uses to match the incoming request.
         * 
         * @return builder
         * 
         */
        public Builder matchUrl(@Nullable Output<String> matchUrl) {
            $.matchUrl = matchUrl;
            return this;
        }

        /**
         * @param matchUrl - (Optional) The URL that the Cloudlet uses to match the incoming request.
         * 
         * @return builder
         * 
         */
        public Builder matchUrl(String matchUrl) {
            return matchUrl(Output.of(matchUrl));
        }

        /**
         * @param matches - (Optional) A list of conditions to apply to a Cloudlet, including:
         * 
         * @return builder
         * 
         */
        public Builder matches(@Nullable Output<List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchArgs>> matches) {
            $.matches = matches;
            return this;
        }

        /**
         * @param matches - (Optional) A list of conditions to apply to a Cloudlet, including:
         * 
         * @return builder
         * 
         */
        public Builder matches(List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchArgs> matches) {
            return matches(Output.of(matches));
        }

        /**
         * @param matches - (Optional) A list of conditions to apply to a Cloudlet, including:
         * 
         * @return builder
         * 
         */
        public Builder matches(GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchArgs... matches) {
            return matches(List.of(matches));
        }

        /**
         * @param matchesAlways - (Optional) Whether the match supports default rules that apply to all requests.
         * 
         * @return builder
         * 
         */
        public Builder matchesAlways(@Nullable Output<Boolean> matchesAlways) {
            $.matchesAlways = matchesAlways;
            return this;
        }

        /**
         * @param matchesAlways - (Optional) Whether the match supports default rules that apply to all requests.
         * 
         * @return builder
         * 
         */
        public Builder matchesAlways(Boolean matchesAlways) {
            return matchesAlways(Output.of(matchesAlways));
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
         * @param start - (Optional) The start time for this match. Specify the value in UTC in seconds since the epoch.
         * 
         * @return builder
         * 
         */
        public Builder start(@Nullable Output<Integer> start) {
            $.start = start;
            return this;
        }

        /**
         * @param start - (Optional) The start time for this match. Specify the value in UTC in seconds since the epoch.
         * 
         * @return builder
         * 
         */
        public Builder start(Integer start) {
            return start(Output.of(start));
        }

        /**
         * @param type - (Required) The type of the array, either `object`, `range`, or `simple`. Use the `simple` option when adding only an array of string-based values.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type - (Required) The type of the array, either `object`, `range`, or `simple`. Use the `simple` option when adding only an array of string-based values.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArgs build() {
            $.forwardSettings = Objects.requireNonNull($.forwardSettings, "expected parameter 'forwardSettings' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}