// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGroupArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupArgs Empty = new GetGroupArgs();

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
     * The group name.
     * 
     */
    @Import(name="groupName")
    private @Nullable Output<String> groupName;

    /**
     * @return The group name.
     * 
     */
    public Optional<Output<String>> groupName() {
        return Optional.ofNullable(this.groupName);
    }

    /**
     * Replaced by `group_name`. Maintained for legacy purposes.
     * 
     * @deprecated
     * The setting &#34;name&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""name"" has been deprecated. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Replaced by `group_name`. Maintained for legacy purposes.
     * 
     * @deprecated
     * The setting &#34;name&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""name"" has been deprecated. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private GetGroupArgs() {}

    private GetGroupArgs(GetGroupArgs $) {
        this.contract = $.contract;
        this.contractId = $.contractId;
        this.groupName = $.groupName;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupArgs $;

        public Builder() {
            $ = new GetGroupArgs();
        }

        public Builder(GetGroupArgs defaults) {
            $ = new GetGroupArgs(Objects.requireNonNull(defaults));
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
         * @param groupName The group name.
         * 
         * @return builder
         * 
         */
        public Builder groupName(@Nullable Output<String> groupName) {
            $.groupName = groupName;
            return this;
        }

        /**
         * @param groupName The group name.
         * 
         * @return builder
         * 
         */
        public Builder groupName(String groupName) {
            return groupName(Output.of(groupName));
        }

        /**
         * @param name Replaced by `group_name`. Maintained for legacy purposes.
         * 
         * @return builder
         * 
         * @deprecated
         * The setting &#34;name&#34; has been deprecated.
         * 
         */
        @Deprecated /* The setting ""name"" has been deprecated. */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Replaced by `group_name`. Maintained for legacy purposes.
         * 
         * @return builder
         * 
         * @deprecated
         * The setting &#34;name&#34; has been deprecated.
         * 
         */
        @Deprecated /* The setting ""name"" has been deprecated. */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetGroupArgs build() {
            return $;
        }
    }

}
