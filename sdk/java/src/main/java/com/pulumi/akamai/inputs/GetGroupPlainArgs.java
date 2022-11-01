// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGroupPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupPlainArgs Empty = new GetGroupPlainArgs();

    /**
     * Replaced by `contract_id`. Maintained for legacy purposes.
     * 
     * @deprecated
     * The setting &#34;contract&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""contract"" has been deprecated. */
    @Import(name="contract")
    private @Nullable String contract;

    /**
     * @return Replaced by `contract_id`. Maintained for legacy purposes.
     * 
     * @deprecated
     * The setting &#34;contract&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""contract"" has been deprecated. */
    public Optional<String> contract() {
        return Optional.ofNullable(this.contract);
    }

    /**
     * - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
     * 
     */
    @Import(name="contractId")
    private @Nullable String contractId;

    /**
     * @return - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
     * 
     */
    public Optional<String> contractId() {
        return Optional.ofNullable(this.contractId);
    }

    /**
     * The group name.
     * 
     */
    @Import(name="groupName")
    private @Nullable String groupName;

    /**
     * @return The group name.
     * 
     */
    public Optional<String> groupName() {
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
    private @Nullable String name;

    /**
     * @return Replaced by `group_name`. Maintained for legacy purposes.
     * 
     * @deprecated
     * The setting &#34;name&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""name"" has been deprecated. */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    private GetGroupPlainArgs() {}

    private GetGroupPlainArgs(GetGroupPlainArgs $) {
        this.contract = $.contract;
        this.contractId = $.contractId;
        this.groupName = $.groupName;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupPlainArgs $;

        public Builder() {
            $ = new GetGroupPlainArgs();
        }

        public Builder(GetGroupPlainArgs defaults) {
            $ = new GetGroupPlainArgs(Objects.requireNonNull(defaults));
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
        public Builder contract(@Nullable String contract) {
            $.contract = contract;
            return this;
        }

        /**
         * @param contractId - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
         * 
         * @return builder
         * 
         */
        public Builder contractId(@Nullable String contractId) {
            $.contractId = contractId;
            return this;
        }

        /**
         * @param groupName The group name.
         * 
         * @return builder
         * 
         */
        public Builder groupName(@Nullable String groupName) {
            $.groupName = groupName;
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
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public GetGroupPlainArgs build() {
            return $;
        }
    }

}