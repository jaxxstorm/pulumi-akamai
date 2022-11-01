// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudletsPolicyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudletsPolicyArgs Empty = new GetCloudletsPolicyArgs();

    /**
     * - (Required) An integer identifier that is associated with all versions of a policy.
     * 
     */
    @Import(name="policyId", required=true)
    private Output<Integer> policyId;

    /**
     * @return - (Required) An integer identifier that is associated with all versions of a policy.
     * 
     */
    public Output<Integer> policyId() {
        return this.policyId;
    }

    /**
     * - (Optional) The version number of a policy.
     * 
     */
    @Import(name="version")
    private @Nullable Output<Integer> version;

    /**
     * @return - (Optional) The version number of a policy.
     * 
     */
    public Optional<Output<Integer>> version() {
        return Optional.ofNullable(this.version);
    }

    private GetCloudletsPolicyArgs() {}

    private GetCloudletsPolicyArgs(GetCloudletsPolicyArgs $) {
        this.policyId = $.policyId;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudletsPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudletsPolicyArgs $;

        public Builder() {
            $ = new GetCloudletsPolicyArgs();
        }

        public Builder(GetCloudletsPolicyArgs defaults) {
            $ = new GetCloudletsPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param policyId - (Required) An integer identifier that is associated with all versions of a policy.
         * 
         * @return builder
         * 
         */
        public Builder policyId(Output<Integer> policyId) {
            $.policyId = policyId;
            return this;
        }

        /**
         * @param policyId - (Required) An integer identifier that is associated with all versions of a policy.
         * 
         * @return builder
         * 
         */
        public Builder policyId(Integer policyId) {
            return policyId(Output.of(policyId));
        }

        /**
         * @param version - (Optional) The version number of a policy.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<Integer> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version - (Optional) The version number of a policy.
         * 
         * @return builder
         * 
         */
        public Builder version(Integer version) {
            return version(Output.of(version));
        }

        public GetCloudletsPolicyArgs build() {
            $.policyId = Objects.requireNonNull($.policyId, "expected parameter 'policyId' to be non-null");
            return $;
        }
    }

}