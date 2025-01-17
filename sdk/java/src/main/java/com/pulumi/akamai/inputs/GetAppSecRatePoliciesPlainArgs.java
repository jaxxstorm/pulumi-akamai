// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAppSecRatePoliciesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecRatePoliciesPlainArgs Empty = new GetAppSecRatePoliciesPlainArgs();

    /**
     * . Unique identifier of the security configuration associated with the rate policies.
     * 
     */
    @Import(name="configId", required=true)
    private Integer configId;

    /**
     * @return . Unique identifier of the security configuration associated with the rate policies.
     * 
     */
    public Integer configId() {
        return this.configId;
    }

    /**
     * . Unique identifier of the rate policy you want to return information for. If not included, information is returned for all your rate policies.
     * 
     */
    @Import(name="ratePolicyId")
    private @Nullable Integer ratePolicyId;

    /**
     * @return . Unique identifier of the rate policy you want to return information for. If not included, information is returned for all your rate policies.
     * 
     */
    public Optional<Integer> ratePolicyId() {
        return Optional.ofNullable(this.ratePolicyId);
    }

    private GetAppSecRatePoliciesPlainArgs() {}

    private GetAppSecRatePoliciesPlainArgs(GetAppSecRatePoliciesPlainArgs $) {
        this.configId = $.configId;
        this.ratePolicyId = $.ratePolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecRatePoliciesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecRatePoliciesPlainArgs $;

        public Builder() {
            $ = new GetAppSecRatePoliciesPlainArgs();
        }

        public Builder(GetAppSecRatePoliciesPlainArgs defaults) {
            $ = new GetAppSecRatePoliciesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the rate policies.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param ratePolicyId . Unique identifier of the rate policy you want to return information for. If not included, information is returned for all your rate policies.
         * 
         * @return builder
         * 
         */
        public Builder ratePolicyId(@Nullable Integer ratePolicyId) {
            $.ratePolicyId = ratePolicyId;
            return this;
        }

        public GetAppSecRatePoliciesPlainArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            return $;
        }
    }

}
