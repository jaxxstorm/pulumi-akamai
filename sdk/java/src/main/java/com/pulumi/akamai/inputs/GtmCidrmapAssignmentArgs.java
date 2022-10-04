// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GtmCidrmapAssignmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final GtmCidrmapAssignmentArgs Empty = new GtmCidrmapAssignmentArgs();

    /**
     * Specifies an array of CIDR blocks.
     * 
     */
    @Import(name="blocks")
    private @Nullable Output<List<String>> blocks;

    /**
     * @return Specifies an array of CIDR blocks.
     * 
     */
    public Optional<Output<List<String>>> blocks() {
        return Optional.ofNullable(this.blocks);
    }

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
     * A descriptive label for the CIDR zone group, up to 256 characters.
     * 
     */
    @Import(name="nickname", required=true)
    private Output<String> nickname;

    /**
     * @return A descriptive label for the CIDR zone group, up to 256 characters.
     * 
     */
    public Output<String> nickname() {
        return this.nickname;
    }

    private GtmCidrmapAssignmentArgs() {}

    private GtmCidrmapAssignmentArgs(GtmCidrmapAssignmentArgs $) {
        this.blocks = $.blocks;
        this.datacenterId = $.datacenterId;
        this.nickname = $.nickname;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GtmCidrmapAssignmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GtmCidrmapAssignmentArgs $;

        public Builder() {
            $ = new GtmCidrmapAssignmentArgs();
        }

        public Builder(GtmCidrmapAssignmentArgs defaults) {
            $ = new GtmCidrmapAssignmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param blocks Specifies an array of CIDR blocks.
         * 
         * @return builder
         * 
         */
        public Builder blocks(@Nullable Output<List<String>> blocks) {
            $.blocks = blocks;
            return this;
        }

        /**
         * @param blocks Specifies an array of CIDR blocks.
         * 
         * @return builder
         * 
         */
        public Builder blocks(List<String> blocks) {
            return blocks(Output.of(blocks));
        }

        /**
         * @param blocks Specifies an array of CIDR blocks.
         * 
         * @return builder
         * 
         */
        public Builder blocks(String... blocks) {
            return blocks(List.of(blocks));
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
         * @param nickname A descriptive label for the CIDR zone group, up to 256 characters.
         * 
         * @return builder
         * 
         */
        public Builder nickname(Output<String> nickname) {
            $.nickname = nickname;
            return this;
        }

        /**
         * @param nickname A descriptive label for the CIDR zone group, up to 256 characters.
         * 
         * @return builder
         * 
         */
        public Builder nickname(String nickname) {
            return nickname(Output.of(nickname));
        }

        public GtmCidrmapAssignmentArgs build() {
            $.datacenterId = Objects.requireNonNull($.datacenterId, "expected parameter 'datacenterId' to be non-null");
            $.nickname = Objects.requireNonNull($.nickname, "expected parameter 'nickname' to be non-null");
            return $;
        }
    }

}
