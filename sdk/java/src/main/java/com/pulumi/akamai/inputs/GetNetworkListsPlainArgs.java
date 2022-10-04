// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNetworkListsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNetworkListsPlainArgs Empty = new GetNetworkListsPlainArgs();

    /**
     * The name of a specific network list to retrieve. If not supplied, information about all network
     * lists will be returned.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of a specific network list to retrieve. If not supplied, information about all network
     * lists will be returned.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of a specific network list to retrieve.
     * If not supplied, information about all network lists will be returned.
     * 
     */
    @Import(name="networkListId")
    private @Nullable String networkListId;

    /**
     * @return The ID of a specific network list to retrieve.
     * If not supplied, information about all network lists will be returned.
     * 
     */
    public Optional<String> networkListId() {
        return Optional.ofNullable(this.networkListId);
    }

    /**
     * The type of network lists to be retrieved; must be either &#34;IP&#34; or &#34;GEO&#34;. If not supplied,
     * information about both types will be returned.
     * 
     */
    @Import(name="type")
    private @Nullable String type;

    /**
     * @return The type of network lists to be retrieved; must be either &#34;IP&#34; or &#34;GEO&#34;. If not supplied,
     * information about both types will be returned.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    private GetNetworkListsPlainArgs() {}

    private GetNetworkListsPlainArgs(GetNetworkListsPlainArgs $) {
        this.name = $.name;
        this.networkListId = $.networkListId;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNetworkListsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNetworkListsPlainArgs $;

        public Builder() {
            $ = new GetNetworkListsPlainArgs();
        }

        public Builder(GetNetworkListsPlainArgs defaults) {
            $ = new GetNetworkListsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of a specific network list to retrieve. If not supplied, information about all network
         * lists will be returned.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param networkListId The ID of a specific network list to retrieve.
         * If not supplied, information about all network lists will be returned.
         * 
         * @return builder
         * 
         */
        public Builder networkListId(@Nullable String networkListId) {
            $.networkListId = networkListId;
            return this;
        }

        /**
         * @param type The type of network lists to be retrieved; must be either &#34;IP&#34; or &#34;GEO&#34;. If not supplied,
         * information about both types will be returned.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        public GetNetworkListsPlainArgs build() {
            return $;
        }
    }

}
