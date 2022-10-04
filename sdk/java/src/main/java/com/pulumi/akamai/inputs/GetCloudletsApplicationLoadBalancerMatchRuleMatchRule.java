// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.akamai.inputs.GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSetting;
import com.pulumi.akamai.inputs.GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatch;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudletsApplicationLoadBalancerMatchRuleMatchRule extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudletsApplicationLoadBalancerMatchRuleMatchRule Empty = new GetCloudletsApplicationLoadBalancerMatchRuleMatchRule();

    @Import(name="disabled")
    private @Nullable Boolean disabled;

    public Optional<Boolean> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * - (Optional) The end time for this match. Specify the value in UTC in seconds since the epoch.
     * 
     */
    @Import(name="end")
    private @Nullable Integer end;

    /**
     * @return - (Optional) The end time for this match. Specify the value in UTC in seconds since the epoch.
     * 
     */
    public Optional<Integer> end() {
        return Optional.ofNullable(this.end);
    }

    /**
     * - (Required) Defines data used to construct a new request URL if all conditions are met. If all of the conditions you set are true, the Edge Server returns an HTTP response from the rewritten URL.
     * 
     */
    @Import(name="forwardSettings", required=true)
    private List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSetting> forwardSettings;

    /**
     * @return - (Required) Defines data used to construct a new request URL if all conditions are met. If all of the conditions you set are true, the Edge Server returns an HTTP response from the rewritten URL.
     * 
     */
    public List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSetting> forwardSettings() {
        return this.forwardSettings;
    }

    /**
     * - (Optional) An identifier for Akamai internal use only.
     * 
     */
    @Import(name="id")
    private @Nullable Integer id;

    /**
     * @return - (Optional) An identifier for Akamai internal use only.
     * 
     */
    public Optional<Integer> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * - (Optional) The URL that the Cloudlet uses to match the incoming request.
     * 
     */
    @Import(name="matchUrl")
    private @Nullable String matchUrl;

    /**
     * @return - (Optional) The URL that the Cloudlet uses to match the incoming request.
     * 
     */
    public Optional<String> matchUrl() {
        return Optional.ofNullable(this.matchUrl);
    }

    /**
     * - (Optional) A list of conditions to apply to a Cloudlet, including:
     * 
     */
    @Import(name="matches")
    private @Nullable List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatch> matches;

    /**
     * @return - (Optional) A list of conditions to apply to a Cloudlet, including:
     * 
     */
    public Optional<List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatch>> matches() {
        return Optional.ofNullable(this.matches);
    }

    /**
     * - (Optional) Whether the match supports default rules that apply to all requests.
     * 
     */
    @Import(name="matchesAlways")
    private @Nullable Boolean matchesAlways;

    /**
     * @return - (Optional) Whether the match supports default rules that apply to all requests.
     * 
     */
    public Optional<Boolean> matchesAlways() {
        return Optional.ofNullable(this.matchesAlways);
    }

    /**
     * - (Optional) If you&#39;re using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return - (Optional) If you&#39;re using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * - (Optional) The start time for this match. Specify the value in UTC in seconds since the epoch.
     * 
     */
    @Import(name="start")
    private @Nullable Integer start;

    /**
     * @return - (Optional) The start time for this match. Specify the value in UTC in seconds since the epoch.
     * 
     */
    public Optional<Integer> start() {
        return Optional.ofNullable(this.start);
    }

    /**
     * - (Required) The type of the array, either `object`, `range`, or `simple`. Use the `simple` option when adding only an array of string-based values.
     * 
     */
    @Import(name="type", required=true)
    private String type;

    /**
     * @return - (Required) The type of the array, either `object`, `range`, or `simple`. Use the `simple` option when adding only an array of string-based values.
     * 
     */
    public String type() {
        return this.type;
    }

    private GetCloudletsApplicationLoadBalancerMatchRuleMatchRule() {}

    private GetCloudletsApplicationLoadBalancerMatchRuleMatchRule(GetCloudletsApplicationLoadBalancerMatchRuleMatchRule $) {
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
    public static Builder builder(GetCloudletsApplicationLoadBalancerMatchRuleMatchRule defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudletsApplicationLoadBalancerMatchRuleMatchRule $;

        public Builder() {
            $ = new GetCloudletsApplicationLoadBalancerMatchRuleMatchRule();
        }

        public Builder(GetCloudletsApplicationLoadBalancerMatchRuleMatchRule defaults) {
            $ = new GetCloudletsApplicationLoadBalancerMatchRuleMatchRule(Objects.requireNonNull(defaults));
        }

        public Builder disabled(@Nullable Boolean disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param end - (Optional) The end time for this match. Specify the value in UTC in seconds since the epoch.
         * 
         * @return builder
         * 
         */
        public Builder end(@Nullable Integer end) {
            $.end = end;
            return this;
        }

        /**
         * @param forwardSettings - (Required) Defines data used to construct a new request URL if all conditions are met. If all of the conditions you set are true, the Edge Server returns an HTTP response from the rewritten URL.
         * 
         * @return builder
         * 
         */
        public Builder forwardSettings(List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSetting> forwardSettings) {
            $.forwardSettings = forwardSettings;
            return this;
        }

        /**
         * @param forwardSettings - (Required) Defines data used to construct a new request URL if all conditions are met. If all of the conditions you set are true, the Edge Server returns an HTTP response from the rewritten URL.
         * 
         * @return builder
         * 
         */
        public Builder forwardSettings(GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSetting... forwardSettings) {
            return forwardSettings(List.of(forwardSettings));
        }

        /**
         * @param id - (Optional) An identifier for Akamai internal use only.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Integer id) {
            $.id = id;
            return this;
        }

        /**
         * @param matchUrl - (Optional) The URL that the Cloudlet uses to match the incoming request.
         * 
         * @return builder
         * 
         */
        public Builder matchUrl(@Nullable String matchUrl) {
            $.matchUrl = matchUrl;
            return this;
        }

        /**
         * @param matches - (Optional) A list of conditions to apply to a Cloudlet, including:
         * 
         * @return builder
         * 
         */
        public Builder matches(@Nullable List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatch> matches) {
            $.matches = matches;
            return this;
        }

        /**
         * @param matches - (Optional) A list of conditions to apply to a Cloudlet, including:
         * 
         * @return builder
         * 
         */
        public Builder matches(GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatch... matches) {
            return matches(List.of(matches));
        }

        /**
         * @param matchesAlways - (Optional) Whether the match supports default rules that apply to all requests.
         * 
         * @return builder
         * 
         */
        public Builder matchesAlways(@Nullable Boolean matchesAlways) {
            $.matchesAlways = matchesAlways;
            return this;
        }

        /**
         * @param name - (Optional) If you&#39;re using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param start - (Optional) The start time for this match. Specify the value in UTC in seconds since the epoch.
         * 
         * @return builder
         * 
         */
        public Builder start(@Nullable Integer start) {
            $.start = start;
            return this;
        }

        /**
         * @param type - (Required) The type of the array, either `object`, `range`, or `simple`. Use the `simple` option when adding only an array of string-based values.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            $.type = type;
            return this;
        }

        public GetCloudletsApplicationLoadBalancerMatchRuleMatchRule build() {
            $.forwardSettings = Objects.requireNonNull($.forwardSettings, "expected parameter 'forwardSettings' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
