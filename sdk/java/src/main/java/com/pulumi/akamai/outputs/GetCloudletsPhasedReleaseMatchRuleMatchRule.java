// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.akamai.outputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettings;
import com.pulumi.akamai.outputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleMatch;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCloudletsPhasedReleaseMatchRuleMatchRule {
    /**
     * @return - (Optional) Whether to disable a rule so it is not evaluated against incoming requests.
     * 
     */
    private @Nullable Boolean disabled;
    /**
     * @return - (Optional) The end time for this match. Specify the value in UTC in seconds since the epoch.
     * 
     */
    private @Nullable Integer end;
    /**
     * @return (Required) The data used to construct a new request URL if all match conditions are met. If all conditions are met, the edge server returns an HTTP response from the rewritten URL.
     * 
     */
    private GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettings forwardSettings;
    /**
     * @return - (Optional) If you&#39;re using a URL match, this specifies the URL that the Cloudlet uses to match the incoming request.
     * 
     */
    private @Nullable String matchUrl;
    /**
     * @return - (Optional) A list of conditions to apply to a Cloudlet, including:
     * 
     */
    private @Nullable List<GetCloudletsPhasedReleaseMatchRuleMatchRuleMatch> matches;
    /**
     * @return - (Optional) Whether the match supports default rules that apply to all requests.
     * 
     */
    private @Nullable Boolean matchesAlways;
    /**
     * @return - (Optional) If you&#39;re using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
     * 
     */
    private @Nullable String name;
    /**
     * @return - (Optional) The start time for this match. Specify the value in UTC in seconds since the epoch.
     * 
     */
    private @Nullable Integer start;
    /**
     * @return - (Required) The type of the array, either `object` or `simple`. Use the `simple` option when adding only an array of string-based values.
     * 
     */
    private String type;

    private GetCloudletsPhasedReleaseMatchRuleMatchRule() {}
    /**
     * @return - (Optional) Whether to disable a rule so it is not evaluated against incoming requests.
     * 
     */
    public Optional<Boolean> disabled() {
        return Optional.ofNullable(this.disabled);
    }
    /**
     * @return - (Optional) The end time for this match. Specify the value in UTC in seconds since the epoch.
     * 
     */
    public Optional<Integer> end() {
        return Optional.ofNullable(this.end);
    }
    /**
     * @return (Required) The data used to construct a new request URL if all match conditions are met. If all conditions are met, the edge server returns an HTTP response from the rewritten URL.
     * 
     */
    public GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettings forwardSettings() {
        return this.forwardSettings;
    }
    /**
     * @return - (Optional) If you&#39;re using a URL match, this specifies the URL that the Cloudlet uses to match the incoming request.
     * 
     */
    public Optional<String> matchUrl() {
        return Optional.ofNullable(this.matchUrl);
    }
    /**
     * @return - (Optional) A list of conditions to apply to a Cloudlet, including:
     * 
     */
    public List<GetCloudletsPhasedReleaseMatchRuleMatchRuleMatch> matches() {
        return this.matches == null ? List.of() : this.matches;
    }
    /**
     * @return - (Optional) Whether the match supports default rules that apply to all requests.
     * 
     */
    public Optional<Boolean> matchesAlways() {
        return Optional.ofNullable(this.matchesAlways);
    }
    /**
     * @return - (Optional) If you&#39;re using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return - (Optional) The start time for this match. Specify the value in UTC in seconds since the epoch.
     * 
     */
    public Optional<Integer> start() {
        return Optional.ofNullable(this.start);
    }
    /**
     * @return - (Required) The type of the array, either `object` or `simple`. Use the `simple` option when adding only an array of string-based values.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudletsPhasedReleaseMatchRuleMatchRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean disabled;
        private @Nullable Integer end;
        private GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettings forwardSettings;
        private @Nullable String matchUrl;
        private @Nullable List<GetCloudletsPhasedReleaseMatchRuleMatchRuleMatch> matches;
        private @Nullable Boolean matchesAlways;
        private @Nullable String name;
        private @Nullable Integer start;
        private String type;
        public Builder() {}
        public Builder(GetCloudletsPhasedReleaseMatchRuleMatchRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.disabled = defaults.disabled;
    	      this.end = defaults.end;
    	      this.forwardSettings = defaults.forwardSettings;
    	      this.matchUrl = defaults.matchUrl;
    	      this.matches = defaults.matches;
    	      this.matchesAlways = defaults.matchesAlways;
    	      this.name = defaults.name;
    	      this.start = defaults.start;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder disabled(@Nullable Boolean disabled) {
            this.disabled = disabled;
            return this;
        }
        @CustomType.Setter
        public Builder end(@Nullable Integer end) {
            this.end = end;
            return this;
        }
        @CustomType.Setter
        public Builder forwardSettings(GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettings forwardSettings) {
            this.forwardSettings = Objects.requireNonNull(forwardSettings);
            return this;
        }
        @CustomType.Setter
        public Builder matchUrl(@Nullable String matchUrl) {
            this.matchUrl = matchUrl;
            return this;
        }
        @CustomType.Setter
        public Builder matches(@Nullable List<GetCloudletsPhasedReleaseMatchRuleMatchRuleMatch> matches) {
            this.matches = matches;
            return this;
        }
        public Builder matches(GetCloudletsPhasedReleaseMatchRuleMatchRuleMatch... matches) {
            return matches(List.of(matches));
        }
        @CustomType.Setter
        public Builder matchesAlways(@Nullable Boolean matchesAlways) {
            this.matchesAlways = matchesAlways;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder start(@Nullable Integer start) {
            this.start = start;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetCloudletsPhasedReleaseMatchRuleMatchRule build() {
            final var o = new GetCloudletsPhasedReleaseMatchRuleMatchRule();
            o.disabled = disabled;
            o.end = end;
            o.forwardSettings = forwardSettings;
            o.matchUrl = matchUrl;
            o.matches = matches;
            o.matchesAlways = matchesAlways;
            o.name = name;
            o.start = start;
            o.type = type;
            return o;
        }
    }
}