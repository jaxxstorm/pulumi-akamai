// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAppSecEvalRulesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecEvalRulesPlainArgs Empty = new GetAppSecEvalRulesPlainArgs();

    /**
     * . Unique identifier of the security configuration running in evaluation mode.
     * 
     */
    @Import(name="configId", required=true)
    private Integer configId;

    /**
     * @return . Unique identifier of the security configuration running in evaluation mode.
     * 
     */
    public Integer configId() {
        return this.configId;
    }

    /**
     * . Unique identifier of the evaluation rule you want to return information for. If not included, information is returned for all your evaluation rules.
     * 
     */
    @Import(name="ruleId")
    private @Nullable Integer ruleId;

    /**
     * @return . Unique identifier of the evaluation rule you want to return information for. If not included, information is returned for all your evaluation rules.
     * 
     */
    public Optional<Integer> ruleId() {
        return Optional.ofNullable(this.ruleId);
    }

    /**
     * . Unique identifier of the security policy associated with the evaluation rule.
     * 
     */
    @Import(name="securityPolicyId", required=true)
    private String securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the evaluation rule.
     * 
     */
    public String securityPolicyId() {
        return this.securityPolicyId;
    }

    private GetAppSecEvalRulesPlainArgs() {}

    private GetAppSecEvalRulesPlainArgs(GetAppSecEvalRulesPlainArgs $) {
        this.configId = $.configId;
        this.ruleId = $.ruleId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecEvalRulesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecEvalRulesPlainArgs $;

        public Builder() {
            $ = new GetAppSecEvalRulesPlainArgs();
        }

        public Builder(GetAppSecEvalRulesPlainArgs defaults) {
            $ = new GetAppSecEvalRulesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration running in evaluation mode.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param ruleId . Unique identifier of the evaluation rule you want to return information for. If not included, information is returned for all your evaluation rules.
         * 
         * @return builder
         * 
         */
        public Builder ruleId(@Nullable Integer ruleId) {
            $.ruleId = ruleId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the evaluation rule.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        public GetAppSecEvalRulesPlainArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.securityPolicyId = Objects.requireNonNull($.securityPolicyId, "expected parameter 'securityPolicyId' to be non-null");
            return $;
        }
    }

}
