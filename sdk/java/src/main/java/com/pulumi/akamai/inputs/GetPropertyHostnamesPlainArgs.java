// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetPropertyHostnamesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPropertyHostnamesPlainArgs Empty = new GetPropertyHostnamesPlainArgs();

    /**
     * - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
     * 
     */
    @Import(name="contractId", required=true)
    private String contractId;

    /**
     * @return - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
     * 
     */
    public String contractId() {
        return this.contractId;
    }

    /**
     * - (Required) A group&#39;s unique ID, including the `grp_` prefix.
     * 
     */
    @Import(name="groupId", required=true)
    private String groupId;

    /**
     * @return - (Required) A group&#39;s unique ID, including the `grp_` prefix.
     * 
     */
    public String groupId() {
        return this.groupId;
    }

    /**
     * - (Required) A property&#39;s unique ID, including the `prp_` prefix.
     * 
     */
    @Import(name="propertyId", required=true)
    private String propertyId;

    /**
     * @return - (Required) A property&#39;s unique ID, including the `prp_` prefix.
     * 
     */
    public String propertyId() {
        return this.propertyId;
    }

    private GetPropertyHostnamesPlainArgs() {}

    private GetPropertyHostnamesPlainArgs(GetPropertyHostnamesPlainArgs $) {
        this.contractId = $.contractId;
        this.groupId = $.groupId;
        this.propertyId = $.propertyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPropertyHostnamesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPropertyHostnamesPlainArgs $;

        public Builder() {
            $ = new GetPropertyHostnamesPlainArgs();
        }

        public Builder(GetPropertyHostnamesPlainArgs defaults) {
            $ = new GetPropertyHostnamesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param contractId - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
         * 
         * @return builder
         * 
         */
        public Builder contractId(String contractId) {
            $.contractId = contractId;
            return this;
        }

        /**
         * @param groupId - (Required) A group&#39;s unique ID, including the `grp_` prefix.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param propertyId - (Required) A property&#39;s unique ID, including the `prp_` prefix.
         * 
         * @return builder
         * 
         */
        public Builder propertyId(String propertyId) {
            $.propertyId = propertyId;
            return this;
        }

        public GetPropertyHostnamesPlainArgs build() {
            $.contractId = Objects.requireNonNull($.contractId, "expected parameter 'contractId' to be non-null");
            $.groupId = Objects.requireNonNull($.groupId, "expected parameter 'groupId' to be non-null");
            $.propertyId = Objects.requireNonNull($.propertyId, "expected parameter 'propertyId' to be non-null");
            return $;
        }
    }

}
