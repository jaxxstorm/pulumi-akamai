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


public final class IamBlockedUserPropertiesState extends com.pulumi.resources.ResourceArgs {

    public static final IamBlockedUserPropertiesState Empty = new IamBlockedUserPropertiesState();

    /**
     * List of properties to block for a user. The property IDs must be an integer.
     * 
     */
    @Import(name="blockedProperties")
    private @Nullable Output<List<Integer>> blockedProperties;

    /**
     * @return List of properties to block for a user. The property IDs must be an integer.
     * 
     */
    public Optional<Output<List<Integer>>> blockedProperties() {
        return Optional.ofNullable(this.blockedProperties);
    }

    /**
     * A unique identifier for a group. Each identifier must be an integer.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<Integer> groupId;

    /**
     * @return A unique identifier for a group. Each identifier must be an integer.
     * 
     */
    public Optional<Output<Integer>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * A unique identifier that corresponds to a user&#39;s actual profile or client ID. Each identifier must be a string.
     * 
     */
    @Import(name="identityId")
    private @Nullable Output<String> identityId;

    /**
     * @return A unique identifier that corresponds to a user&#39;s actual profile or client ID. Each identifier must be a string.
     * 
     */
    public Optional<Output<String>> identityId() {
        return Optional.ofNullable(this.identityId);
    }

    private IamBlockedUserPropertiesState() {}

    private IamBlockedUserPropertiesState(IamBlockedUserPropertiesState $) {
        this.blockedProperties = $.blockedProperties;
        this.groupId = $.groupId;
        this.identityId = $.identityId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamBlockedUserPropertiesState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamBlockedUserPropertiesState $;

        public Builder() {
            $ = new IamBlockedUserPropertiesState();
        }

        public Builder(IamBlockedUserPropertiesState defaults) {
            $ = new IamBlockedUserPropertiesState(Objects.requireNonNull(defaults));
        }

        /**
         * @param blockedProperties List of properties to block for a user. The property IDs must be an integer.
         * 
         * @return builder
         * 
         */
        public Builder blockedProperties(@Nullable Output<List<Integer>> blockedProperties) {
            $.blockedProperties = blockedProperties;
            return this;
        }

        /**
         * @param blockedProperties List of properties to block for a user. The property IDs must be an integer.
         * 
         * @return builder
         * 
         */
        public Builder blockedProperties(List<Integer> blockedProperties) {
            return blockedProperties(Output.of(blockedProperties));
        }

        /**
         * @param blockedProperties List of properties to block for a user. The property IDs must be an integer.
         * 
         * @return builder
         * 
         */
        public Builder blockedProperties(Integer... blockedProperties) {
            return blockedProperties(List.of(blockedProperties));
        }

        /**
         * @param groupId A unique identifier for a group. Each identifier must be an integer.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<Integer> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId A unique identifier for a group. Each identifier must be an integer.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Integer groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param identityId A unique identifier that corresponds to a user&#39;s actual profile or client ID. Each identifier must be a string.
         * 
         * @return builder
         * 
         */
        public Builder identityId(@Nullable Output<String> identityId) {
            $.identityId = identityId;
            return this;
        }

        /**
         * @param identityId A unique identifier that corresponds to a user&#39;s actual profile or client ID. Each identifier must be a string.
         * 
         * @return builder
         * 
         */
        public Builder identityId(String identityId) {
            return identityId(Output.of(identityId));
        }

        public IamBlockedUserPropertiesState build() {
            return $;
        }
    }

}
