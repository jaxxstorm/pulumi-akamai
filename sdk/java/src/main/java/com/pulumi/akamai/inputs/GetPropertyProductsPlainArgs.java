// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetPropertyProductsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPropertyProductsPlainArgs Empty = new GetPropertyProductsPlainArgs();

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

    private GetPropertyProductsPlainArgs() {}

    private GetPropertyProductsPlainArgs(GetPropertyProductsPlainArgs $) {
        this.contractId = $.contractId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPropertyProductsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPropertyProductsPlainArgs $;

        public Builder() {
            $ = new GetPropertyProductsPlainArgs();
        }

        public Builder(GetPropertyProductsPlainArgs defaults) {
            $ = new GetPropertyProductsPlainArgs(Objects.requireNonNull(defaults));
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

        public GetPropertyProductsPlainArgs build() {
            $.contractId = Objects.requireNonNull($.contractId, "expected parameter 'contractId' to be non-null");
            return $;
        }
    }

}
