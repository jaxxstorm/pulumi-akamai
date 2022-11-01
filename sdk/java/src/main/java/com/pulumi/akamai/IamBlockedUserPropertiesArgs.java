// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class IamBlockedUserPropertiesArgs extends com.pulumi.resources.ResourceArgs {

    public static final IamBlockedUserPropertiesArgs Empty = new IamBlockedUserPropertiesArgs();

    /**
     * List of properties to block for a user. The property IDs must be an integer.
     * 
     */
    @Import(name="blockedProperties", required=true)
    private Output<List<Integer>> blockedProperties;

    /**
     * @return List of properties to block for a user. The property IDs must be an integer.
     * 
     */
    public Output<List<Integer>> blockedProperties() {
        return this.blockedProperties;
    }

    /**
     * A unique identifier for a group. Each identifier must be an integer.
     * 
     */
    @Import(name="groupId", required=true)
    private Output<Integer> groupId;

    /**
     * @return A unique identifier for a group. Each identifier must be an integer.
     * 
     */
    public Output<Integer> groupId() {
        return this.groupId;
    }

    /**
     * A unique identifier that corresponds to a user&#39;s actual profile or client ID. Each identifier must be a string.
     * 
     */
    @Import(name="identityId", required=true)
    private Output<String> identityId;

    /**
     * @return A unique identifier that corresponds to a user&#39;s actual profile or client ID. Each identifier must be a string.
     * 
     */
    public Output<String> identityId() {
        return this.identityId;
    }

    private IamBlockedUserPropertiesArgs() {}

    private IamBlockedUserPropertiesArgs(IamBlockedUserPropertiesArgs $) {
        this.blockedProperties = $.blockedProperties;
        this.groupId = $.groupId;
        this.identityId = $.identityId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamBlockedUserPropertiesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamBlockedUserPropertiesArgs $;

        public Builder() {
            $ = new IamBlockedUserPropertiesArgs();
        }

        public Builder(IamBlockedUserPropertiesArgs defaults) {
            $ = new IamBlockedUserPropertiesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param blockedProperties List of properties to block for a user. The property IDs must be an integer.
         * 
         * @return builder
         * 
         */
        public Builder blockedProperties(Output<List<Integer>> blockedProperties) {
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
        public Builder groupId(Output<Integer> groupId) {
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
        public Builder identityId(Output<String> identityId) {
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

        public IamBlockedUserPropertiesArgs build() {
            $.blockedProperties = Objects.requireNonNull($.blockedProperties, "expected parameter 'blockedProperties' to be non-null");
            $.groupId = Objects.requireNonNull($.groupId, "expected parameter 'groupId' to be non-null");
            $.identityId = Objects.requireNonNull($.identityId, "expected parameter 'identityId' to be non-null");
            return $;
        }
    }

}