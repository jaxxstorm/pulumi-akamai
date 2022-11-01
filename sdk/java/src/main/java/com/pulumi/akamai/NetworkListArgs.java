// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NetworkListArgs extends com.pulumi.resources.ResourceArgs {

    public static final NetworkListArgs Empty = new NetworkListArgs();

    /**
     * The contract ID of the network list. If supplied, group_id must also be supplied. The
     * contract_id value of an existing network list may not be modified.
     * 
     */
    @Import(name="contractId")
    private @Nullable Output<String> contractId;

    /**
     * @return The contract ID of the network list. If supplied, group_id must also be supplied. The
     * contract_id value of an existing network list may not be modified.
     * 
     */
    public Optional<Output<String>> contractId() {
        return Optional.ofNullable(this.contractId);
    }

    /**
     * The description to be assigned to the network list.
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return The description to be assigned to the network list.
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    /**
     * The group ID of the network list. If supplied, contract_id must also be supplied. The
     * group_id value of an existing network list may not be modified.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<Integer> groupId;

    /**
     * @return The group ID of the network list. If supplied, contract_id must also be supplied. The
     * group_id value of an existing network list may not be modified.
     * 
     */
    public Optional<Output<Integer>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
     * removed from an existing list.
     * 
     */
    @Import(name="lists")
    private @Nullable Output<List<String>> lists;

    /**
     * @return : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
     * removed from an existing list.
     * 
     */
    public Optional<Output<List<String>>> lists() {
        return Optional.ofNullable(this.lists);
    }

    /**
     * A string specifying the interpretation of the `list` parameter. Must be one of the following:
     * 
     */
    @Import(name="mode", required=true)
    private Output<String> mode;

    /**
     * @return A string specifying the interpretation of the `list` parameter. Must be one of the following:
     * 
     */
    public Output<String> mode() {
        return this.mode;
    }

    /**
     * The name to be assigned to the network list.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name to be assigned to the network list.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The type of the network list; must be either &#34;IP&#34; or &#34;GEO&#34;.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of the network list; must be either &#34;IP&#34; or &#34;GEO&#34;.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private NetworkListArgs() {}

    private NetworkListArgs(NetworkListArgs $) {
        this.contractId = $.contractId;
        this.description = $.description;
        this.groupId = $.groupId;
        this.lists = $.lists;
        this.mode = $.mode;
        this.name = $.name;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkListArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkListArgs $;

        public Builder() {
            $ = new NetworkListArgs();
        }

        public Builder(NetworkListArgs defaults) {
            $ = new NetworkListArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param contractId The contract ID of the network list. If supplied, group_id must also be supplied. The
         * contract_id value of an existing network list may not be modified.
         * 
         * @return builder
         * 
         */
        public Builder contractId(@Nullable Output<String> contractId) {
            $.contractId = contractId;
            return this;
        }

        /**
         * @param contractId The contract ID of the network list. If supplied, group_id must also be supplied. The
         * contract_id value of an existing network list may not be modified.
         * 
         * @return builder
         * 
         */
        public Builder contractId(String contractId) {
            return contractId(Output.of(contractId));
        }

        /**
         * @param description The description to be assigned to the network list.
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description to be assigned to the network list.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param groupId The group ID of the network list. If supplied, contract_id must also be supplied. The
         * group_id value of an existing network list may not be modified.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<Integer> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The group ID of the network list. If supplied, contract_id must also be supplied. The
         * group_id value of an existing network list may not be modified.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Integer groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param lists : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
         * removed from an existing list.
         * 
         * @return builder
         * 
         */
        public Builder lists(@Nullable Output<List<String>> lists) {
            $.lists = lists;
            return this;
        }

        /**
         * @param lists : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
         * removed from an existing list.
         * 
         * @return builder
         * 
         */
        public Builder lists(List<String> lists) {
            return lists(Output.of(lists));
        }

        /**
         * @param lists : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
         * removed from an existing list.
         * 
         * @return builder
         * 
         */
        public Builder lists(String... lists) {
            return lists(List.of(lists));
        }

        /**
         * @param mode A string specifying the interpretation of the `list` parameter. Must be one of the following:
         * 
         * @return builder
         * 
         */
        public Builder mode(Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode A string specifying the interpretation of the `list` parameter. Must be one of the following:
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param name The name to be assigned to the network list.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name to be assigned to the network list.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param type The type of the network list; must be either &#34;IP&#34; or &#34;GEO&#34;.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of the network list; must be either &#34;IP&#34; or &#34;GEO&#34;.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public NetworkListArgs build() {
            $.description = Objects.requireNonNull($.description, "expected parameter 'description' to be non-null");
            $.mode = Objects.requireNonNull($.mode, "expected parameter 'mode' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}