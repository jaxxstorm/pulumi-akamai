// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.trafficmanagement.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GtmASmapDefaultDatacenterArgs extends com.pulumi.resources.ResourceArgs {

    public static final GtmASmapDefaultDatacenterArgs Empty = new GtmASmapDefaultDatacenterArgs();

    /**
     * A unique identifier for an existing data center in the domain.
     * 
     */
    @Import(name="datacenterId", required=true)
    private Output<Integer> datacenterId;

    /**
     * @return A unique identifier for an existing data center in the domain.
     * 
     */
    public Output<Integer> datacenterId() {
        return this.datacenterId;
    }

    /**
     * A descriptive label for the group.
     * 
     */
    @Import(name="nickname")
    private @Nullable Output<String> nickname;

    /**
     * @return A descriptive label for the group.
     * 
     */
    public Optional<Output<String>> nickname() {
        return Optional.ofNullable(this.nickname);
    }

    private GtmASmapDefaultDatacenterArgs() {}

    private GtmASmapDefaultDatacenterArgs(GtmASmapDefaultDatacenterArgs $) {
        this.datacenterId = $.datacenterId;
        this.nickname = $.nickname;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GtmASmapDefaultDatacenterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GtmASmapDefaultDatacenterArgs $;

        public Builder() {
            $ = new GtmASmapDefaultDatacenterArgs();
        }

        public Builder(GtmASmapDefaultDatacenterArgs defaults) {
            $ = new GtmASmapDefaultDatacenterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datacenterId A unique identifier for an existing data center in the domain.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(Output<Integer> datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        /**
         * @param datacenterId A unique identifier for an existing data center in the domain.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(Integer datacenterId) {
            return datacenterId(Output.of(datacenterId));
        }

        /**
         * @param nickname A descriptive label for the group.
         * 
         * @return builder
         * 
         */
        public Builder nickname(@Nullable Output<String> nickname) {
            $.nickname = nickname;
            return this;
        }

        /**
         * @param nickname A descriptive label for the group.
         * 
         * @return builder
         * 
         */
        public Builder nickname(String nickname) {
            return nickname(Output.of(nickname));
        }

        public GtmASmapDefaultDatacenterArgs build() {
            $.datacenterId = Objects.requireNonNull($.datacenterId, "expected parameter 'datacenterId' to be non-null");
            return $;
        }
    }

}
