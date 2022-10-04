// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNetworkListsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNetworkListsArgs Empty = new GetNetworkListsArgs();

    /**
     * The name of a specific network list to retrieve. If not supplied, information about all network
     * lists will be returned.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of a specific network list to retrieve. If not supplied, information about all network
     * lists will be returned.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of a specific network list to retrieve.
     * If not supplied, information about all network lists will be returned.
     * 
     */
    @Import(name="networkListId")
    private @Nullable Output<String> networkListId;

    /**
     * @return The ID of a specific network list to retrieve.
     * If not supplied, information about all network lists will be returned.
     * 
     */
    public Optional<Output<String>> networkListId() {
        return Optional.ofNullable(this.networkListId);
    }

    /**
     * The type of network lists to be retrieved; must be either &#34;IP&#34; or &#34;GEO&#34;. If not supplied,
     * information about both types will be returned.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of network lists to be retrieved; must be either &#34;IP&#34; or &#34;GEO&#34;. If not supplied,
     * information about both types will be returned.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private GetNetworkListsArgs() {}

    private GetNetworkListsArgs(GetNetworkListsArgs $) {
        this.name = $.name;
        this.networkListId = $.networkListId;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNetworkListsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNetworkListsArgs $;

        public Builder() {
            $ = new GetNetworkListsArgs();
        }

        public Builder(GetNetworkListsArgs defaults) {
            $ = new GetNetworkListsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of a specific network list to retrieve. If not supplied, information about all network
         * lists will be returned.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of a specific network list to retrieve. If not supplied, information about all network
         * lists will be returned.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param networkListId The ID of a specific network list to retrieve.
         * If not supplied, information about all network lists will be returned.
         * 
         * @return builder
         * 
         */
        public Builder networkListId(@Nullable Output<String> networkListId) {
            $.networkListId = networkListId;
            return this;
        }

        /**
         * @param networkListId The ID of a specific network list to retrieve.
         * If not supplied, information about all network lists will be returned.
         * 
         * @return builder
         * 
         */
        public Builder networkListId(String networkListId) {
            return networkListId(Output.of(networkListId));
        }

        /**
         * @param type The type of network lists to be retrieved; must be either &#34;IP&#34; or &#34;GEO&#34;. If not supplied,
         * information about both types will be returned.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of network lists to be retrieved; must be either &#34;IP&#34; or &#34;GEO&#34;. If not supplied,
         * information about both types will be returned.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public GetNetworkListsArgs build() {
            return $;
        }
    }

}
