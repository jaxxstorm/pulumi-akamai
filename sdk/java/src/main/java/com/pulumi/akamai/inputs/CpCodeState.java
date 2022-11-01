// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CpCodeState extends com.pulumi.resources.ResourceArgs {

    public static final CpCodeState Empty = new CpCodeState();

    /**
     * Replaced by `contract_id`. Maintained for legacy purposes.
     * 
     * @deprecated
     * The setting &#34;contract&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""contract"" has been deprecated. */
    @Import(name="contract")
    private @Nullable Output<String> contract;

    /**
     * @return Replaced by `contract_id`. Maintained for legacy purposes.
     * 
     * @deprecated
     * The setting &#34;contract&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""contract"" has been deprecated. */
    public Optional<Output<String>> contract() {
        return Optional.ofNullable(this.contract);
    }

    /**
     * - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
     * 
     */
    @Import(name="contractId")
    private @Nullable Output<String> contractId;

    /**
     * @return - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
     * 
     */
    public Optional<Output<String>> contractId() {
        return Optional.ofNullable(this.contractId);
    }

    /**
     * Replaced by `group_id`. Maintained for legacy purposes.
     * 
     * @deprecated
     * The setting &#34;group&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""group"" has been deprecated. */
    @Import(name="group")
    private @Nullable Output<String> group;

    /**
     * @return Replaced by `group_id`. Maintained for legacy purposes.
     * 
     * @deprecated
     * The setting &#34;group&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""group"" has been deprecated. */
    public Optional<Output<String>> group() {
        return Optional.ofNullable(this.group);
    }

    /**
     * - (Required) A group&#39;s unique ID, including the `grp_` prefix.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<String> groupId;

    /**
     * @return - (Required) A group&#39;s unique ID, including the `grp_` prefix.
     * 
     */
    public Optional<Output<String>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * - (Required) A descriptive label for the CP code. If you&#39;re creating a new CP code, the name can&#39;t include commas, underscores, quotes, or any of these special characters: ^ # %.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return - (Required) A descriptive label for the CP code. If you&#39;re creating a new CP code, the name can&#39;t include commas, underscores, quotes, or any of these special characters: ^ # %.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Replaced by `product_id`. Maintained for legacy purposes.
     * 
     * @deprecated
     * The setting &#34;product&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""product"" has been deprecated. */
    @Import(name="product")
    private @Nullable Output<String> product;

    /**
     * @return Replaced by `product_id`. Maintained for legacy purposes.
     * 
     * @deprecated
     * The setting &#34;product&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""product"" has been deprecated. */
    public Optional<Output<String>> product() {
        return Optional.ofNullable(this.product);
    }

    @Import(name="productId")
    private @Nullable Output<String> productId;

    public Optional<Output<String>> productId() {
        return Optional.ofNullable(this.productId);
    }

    private CpCodeState() {}

    private CpCodeState(CpCodeState $) {
        this.contract = $.contract;
        this.contractId = $.contractId;
        this.group = $.group;
        this.groupId = $.groupId;
        this.name = $.name;
        this.product = $.product;
        this.productId = $.productId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CpCodeState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CpCodeState $;

        public Builder() {
            $ = new CpCodeState();
        }

        public Builder(CpCodeState defaults) {
            $ = new CpCodeState(Objects.requireNonNull(defaults));
        }

        /**
         * @param contract Replaced by `contract_id`. Maintained for legacy purposes.
         * 
         * @return builder
         * 
         * @deprecated
         * The setting &#34;contract&#34; has been deprecated.
         * 
         */
        @Deprecated /* The setting ""contract"" has been deprecated. */
        public Builder contract(@Nullable Output<String> contract) {
            $.contract = contract;
            return this;
        }

        /**
         * @param contract Replaced by `contract_id`. Maintained for legacy purposes.
         * 
         * @return builder
         * 
         * @deprecated
         * The setting &#34;contract&#34; has been deprecated.
         * 
         */
        @Deprecated /* The setting ""contract"" has been deprecated. */
        public Builder contract(String contract) {
            return contract(Output.of(contract));
        }

        /**
         * @param contractId - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
         * 
         * @return builder
         * 
         */
        public Builder contractId(@Nullable Output<String> contractId) {
            $.contractId = contractId;
            return this;
        }

        /**
         * @param contractId - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
         * 
         * @return builder
         * 
         */
        public Builder contractId(String contractId) {
            return contractId(Output.of(contractId));
        }

        /**
         * @param group Replaced by `group_id`. Maintained for legacy purposes.
         * 
         * @return builder
         * 
         * @deprecated
         * The setting &#34;group&#34; has been deprecated.
         * 
         */
        @Deprecated /* The setting ""group"" has been deprecated. */
        public Builder group(@Nullable Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group Replaced by `group_id`. Maintained for legacy purposes.
         * 
         * @return builder
         * 
         * @deprecated
         * The setting &#34;group&#34; has been deprecated.
         * 
         */
        @Deprecated /* The setting ""group"" has been deprecated. */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param groupId - (Required) A group&#39;s unique ID, including the `grp_` prefix.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId - (Required) A group&#39;s unique ID, including the `grp_` prefix.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param name - (Required) A descriptive label for the CP code. If you&#39;re creating a new CP code, the name can&#39;t include commas, underscores, quotes, or any of these special characters: ^ # %.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name - (Required) A descriptive label for the CP code. If you&#39;re creating a new CP code, the name can&#39;t include commas, underscores, quotes, or any of these special characters: ^ # %.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param product Replaced by `product_id`. Maintained for legacy purposes.
         * 
         * @return builder
         * 
         * @deprecated
         * The setting &#34;product&#34; has been deprecated.
         * 
         */
        @Deprecated /* The setting ""product"" has been deprecated. */
        public Builder product(@Nullable Output<String> product) {
            $.product = product;
            return this;
        }

        /**
         * @param product Replaced by `product_id`. Maintained for legacy purposes.
         * 
         * @return builder
         * 
         * @deprecated
         * The setting &#34;product&#34; has been deprecated.
         * 
         */
        @Deprecated /* The setting ""product"" has been deprecated. */
        public Builder product(String product) {
            return product(Output.of(product));
        }

        public Builder productId(@Nullable Output<String> productId) {
            $.productId = productId;
            return this;
        }

        public Builder productId(String productId) {
            return productId(Output.of(productId));
        }

        public CpCodeState build() {
            return $;
        }
    }

}