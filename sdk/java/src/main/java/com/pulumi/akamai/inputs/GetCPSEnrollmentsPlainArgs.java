// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetCPSEnrollmentsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCPSEnrollmentsPlainArgs Empty = new GetCPSEnrollmentsPlainArgs();

    /**
     * A contract&#39;s ID, optionally with the `ctr_` prefix.
     * 
     */
    @Import(name="contractId", required=true)
    private String contractId;

    /**
     * @return A contract&#39;s ID, optionally with the `ctr_` prefix.
     * 
     */
    public String contractId() {
        return this.contractId;
    }

    private GetCPSEnrollmentsPlainArgs() {}

    private GetCPSEnrollmentsPlainArgs(GetCPSEnrollmentsPlainArgs $) {
        this.contractId = $.contractId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCPSEnrollmentsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCPSEnrollmentsPlainArgs $;

        public Builder() {
            $ = new GetCPSEnrollmentsPlainArgs();
        }

        public Builder(GetCPSEnrollmentsPlainArgs defaults) {
            $ = new GetCPSEnrollmentsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param contractId A contract&#39;s ID, optionally with the `ctr_` prefix.
         * 
         * @return builder
         * 
         */
        public Builder contractId(String contractId) {
            $.contractId = contractId;
            return this;
        }

        public GetCPSEnrollmentsPlainArgs build() {
            $.contractId = Objects.requireNonNull($.contractId, "expected parameter 'contractId' to be non-null");
            return $;
        }
    }

}