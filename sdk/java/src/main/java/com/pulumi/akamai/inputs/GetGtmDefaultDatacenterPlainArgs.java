// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGtmDefaultDatacenterPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGtmDefaultDatacenterPlainArgs Empty = new GetGtmDefaultDatacenterPlainArgs();

    /**
     * The default is `5400`.
     * 
     */
    @Import(name="datacenter")
    private @Nullable Integer datacenter;

    /**
     * @return The default is `5400`.
     * 
     */
    public Optional<Integer> datacenter() {
        return Optional.ofNullable(this.datacenter);
    }

    @Import(name="domain", required=true)
    private String domain;

    public String domain() {
        return this.domain;
    }

    private GetGtmDefaultDatacenterPlainArgs() {}

    private GetGtmDefaultDatacenterPlainArgs(GetGtmDefaultDatacenterPlainArgs $) {
        this.datacenter = $.datacenter;
        this.domain = $.domain;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGtmDefaultDatacenterPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGtmDefaultDatacenterPlainArgs $;

        public Builder() {
            $ = new GetGtmDefaultDatacenterPlainArgs();
        }

        public Builder(GetGtmDefaultDatacenterPlainArgs defaults) {
            $ = new GetGtmDefaultDatacenterPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datacenter The default is `5400`.
         * 
         * @return builder
         * 
         */
        public Builder datacenter(@Nullable Integer datacenter) {
            $.datacenter = datacenter;
            return this;
        }

        public Builder domain(String domain) {
            $.domain = domain;
            return this;
        }

        public GetGtmDefaultDatacenterPlainArgs build() {
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            return $;
        }
    }

}
