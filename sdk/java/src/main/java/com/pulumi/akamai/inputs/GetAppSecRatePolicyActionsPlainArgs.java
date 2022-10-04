// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAppSecRatePolicyActionsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecRatePolicyActionsPlainArgs Empty = new GetAppSecRatePolicyActionsPlainArgs();

    /**
     * . Unique identifier of the security configuration associated with the rate policies and rate policy actions.
     * 
     */
    @Import(name="configId", required=true)
    private Integer configId;

    /**
     * @return . Unique identifier of the security configuration associated with the rate policies and rate policy actions.
     * 
     */
    public Integer configId() {
        return this.configId;
    }

    /**
     * . Unique identifier of the rate policy you want to return action information for. If not included, action information is returned for all your rate policies.
     * 
     */
    @Import(name="ratePolicyId")
    private @Nullable Integer ratePolicyId;

    /**
     * @return . Unique identifier of the rate policy you want to return action information for. If not included, action information is returned for all your rate policies.
     * 
     */
    public Optional<Integer> ratePolicyId() {
        return Optional.ofNullable(this.ratePolicyId);
    }

    /**
     * . Unique identifier of the security policy associated with the rate policies and rate policy actions.
     * 
     */
    @Import(name="securityPolicyId", required=true)
    private String securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the rate policies and rate policy actions.
     * 
     */
    public String securityPolicyId() {
        return this.securityPolicyId;
    }

    private GetAppSecRatePolicyActionsPlainArgs() {}

    private GetAppSecRatePolicyActionsPlainArgs(GetAppSecRatePolicyActionsPlainArgs $) {
        this.configId = $.configId;
        this.ratePolicyId = $.ratePolicyId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecRatePolicyActionsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecRatePolicyActionsPlainArgs $;

        public Builder() {
            $ = new GetAppSecRatePolicyActionsPlainArgs();
        }

        public Builder(GetAppSecRatePolicyActionsPlainArgs defaults) {
            $ = new GetAppSecRatePolicyActionsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the rate policies and rate policy actions.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param ratePolicyId . Unique identifier of the rate policy you want to return action information for. If not included, action information is returned for all your rate policies.
         * 
         * @return builder
         * 
         */
        public Builder ratePolicyId(@Nullable Integer ratePolicyId) {
            $.ratePolicyId = ratePolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the rate policies and rate policy actions.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        public GetAppSecRatePolicyActionsPlainArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.securityPolicyId = Objects.requireNonNull($.securityPolicyId, "expected parameter 'securityPolicyId' to be non-null");
            return $;
        }
    }

}
