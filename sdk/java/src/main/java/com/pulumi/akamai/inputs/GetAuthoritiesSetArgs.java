// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetAuthoritiesSetArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAuthoritiesSetArgs Empty = new GetAuthoritiesSetArgs();

    /**
     * The contract ID.
     * 
     */
    @Import(name="contract", required=true)
    private Output<String> contract;

    /**
     * @return The contract ID.
     * 
     */
    public Output<String> contract() {
        return this.contract;
    }

    private GetAuthoritiesSetArgs() {}

    private GetAuthoritiesSetArgs(GetAuthoritiesSetArgs $) {
        this.contract = $.contract;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAuthoritiesSetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAuthoritiesSetArgs $;

        public Builder() {
            $ = new GetAuthoritiesSetArgs();
        }

        public Builder(GetAuthoritiesSetArgs defaults) {
            $ = new GetAuthoritiesSetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param contract The contract ID.
         * 
         * @return builder
         * 
         */
        public Builder contract(Output<String> contract) {
            $.contract = contract;
            return this;
        }

        /**
         * @param contract The contract ID.
         * 
         * @return builder
         * 
         */
        public Builder contract(String contract) {
            return contract(Output.of(contract));
        }

        public GetAuthoritiesSetArgs build() {
            $.contract = Objects.requireNonNull($.contract, "expected parameter 'contract' to be non-null");
            return $;
        }
    }

}
