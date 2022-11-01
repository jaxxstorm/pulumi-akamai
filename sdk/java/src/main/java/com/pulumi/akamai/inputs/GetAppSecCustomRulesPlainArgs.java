// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAppSecCustomRulesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecCustomRulesPlainArgs Empty = new GetAppSecCustomRulesPlainArgs();

    /**
     * . Unique identifier of the security configuration associated with the custom rules.
     * 
     */
    @Import(name="configId", required=true)
    private Integer configId;

    /**
     * @return . Unique identifier of the security configuration associated with the custom rules.
     * 
     */
    public Integer configId() {
        return this.configId;
    }

    /**
     * . Unique identifier of the custom rule you want to return information for. If not included, information is returned for all your custom rules.
     * 
     */
    @Import(name="customRuleId")
    private @Nullable Integer customRuleId;

    /**
     * @return . Unique identifier of the custom rule you want to return information for. If not included, information is returned for all your custom rules.
     * 
     */
    public Optional<Integer> customRuleId() {
        return Optional.ofNullable(this.customRuleId);
    }

    private GetAppSecCustomRulesPlainArgs() {}

    private GetAppSecCustomRulesPlainArgs(GetAppSecCustomRulesPlainArgs $) {
        this.configId = $.configId;
        this.customRuleId = $.customRuleId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecCustomRulesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecCustomRulesPlainArgs $;

        public Builder() {
            $ = new GetAppSecCustomRulesPlainArgs();
        }

        public Builder(GetAppSecCustomRulesPlainArgs defaults) {
            $ = new GetAppSecCustomRulesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the custom rules.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param customRuleId . Unique identifier of the custom rule you want to return information for. If not included, information is returned for all your custom rules.
         * 
         * @return builder
         * 
         */
        public Builder customRuleId(@Nullable Integer customRuleId) {
            $.customRuleId = customRuleId;
            return this;
        }

        public GetAppSecCustomRulesPlainArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            return $;
        }
    }

}