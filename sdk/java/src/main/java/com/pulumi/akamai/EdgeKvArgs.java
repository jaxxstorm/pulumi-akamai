// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.inputs.EdgeKvInitialDataArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EdgeKvArgs extends com.pulumi.resources.ResourceArgs {

    public static final EdgeKvArgs Empty = new EdgeKvArgs();

    /**
     * Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
     * 
     */
    @Import(name="geoLocation")
    private @Nullable Output<String> geoLocation;

    /**
     * @return Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
     * 
     */
    public Optional<Output<String>> geoLocation() {
        return Optional.ofNullable(this.geoLocation);
    }

    /**
     * - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
     * 
     */
    @Import(name="groupId", required=true)
    private Output<Integer> groupId;

    /**
     * @return - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
     * 
     */
    public Output<Integer> groupId() {
        return this.groupId;
    }

    /**
     * List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
     * 
     */
    @Import(name="initialDatas")
    private @Nullable Output<List<EdgeKvInitialDataArgs>> initialDatas;

    /**
     * @return List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
     * 
     */
    public Optional<Output<List<EdgeKvInitialDataArgs>>> initialDatas() {
        return Optional.ofNullable(this.initialDatas);
    }

    /**
     * - (Required) The name of the namespace.
     * 
     */
    @Import(name="namespaceName", required=true)
    private Output<String> namespaceName;

    /**
     * @return - (Required) The name of the namespace.
     * 
     */
    public Output<String> namespaceName() {
        return this.namespaceName;
    }

    /**
     * The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
     * 
     */
    @Import(name="network", required=true)
    private Output<String> network;

    /**
     * @return The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
     * 
     */
    public Output<String> network() {
        return this.network;
    }

    /**
     * - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
     * 
     */
    @Import(name="retentionInSeconds", required=true)
    private Output<Integer> retentionInSeconds;

    /**
     * @return - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
     * 
     */
    public Output<Integer> retentionInSeconds() {
        return this.retentionInSeconds;
    }

    private EdgeKvArgs() {}

    private EdgeKvArgs(EdgeKvArgs $) {
        this.geoLocation = $.geoLocation;
        this.groupId = $.groupId;
        this.initialDatas = $.initialDatas;
        this.namespaceName = $.namespaceName;
        this.network = $.network;
        this.retentionInSeconds = $.retentionInSeconds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EdgeKvArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EdgeKvArgs $;

        public Builder() {
            $ = new EdgeKvArgs();
        }

        public Builder(EdgeKvArgs defaults) {
            $ = new EdgeKvArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param geoLocation Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
         * 
         * @return builder
         * 
         */
        public Builder geoLocation(@Nullable Output<String> geoLocation) {
            $.geoLocation = geoLocation;
            return this;
        }

        /**
         * @param geoLocation Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
         * 
         * @return builder
         * 
         */
        public Builder geoLocation(String geoLocation) {
            return geoLocation(Output.of(geoLocation));
        }

        /**
         * @param groupId - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Output<Integer> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Integer groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param initialDatas List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
         * 
         * @return builder
         * 
         */
        public Builder initialDatas(@Nullable Output<List<EdgeKvInitialDataArgs>> initialDatas) {
            $.initialDatas = initialDatas;
            return this;
        }

        /**
         * @param initialDatas List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
         * 
         * @return builder
         * 
         */
        public Builder initialDatas(List<EdgeKvInitialDataArgs> initialDatas) {
            return initialDatas(Output.of(initialDatas));
        }

        /**
         * @param initialDatas List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
         * 
         * @return builder
         * 
         */
        public Builder initialDatas(EdgeKvInitialDataArgs... initialDatas) {
            return initialDatas(List.of(initialDatas));
        }

        /**
         * @param namespaceName - (Required) The name of the namespace.
         * 
         * @return builder
         * 
         */
        public Builder namespaceName(Output<String> namespaceName) {
            $.namespaceName = namespaceName;
            return this;
        }

        /**
         * @param namespaceName - (Required) The name of the namespace.
         * 
         * @return builder
         * 
         */
        public Builder namespaceName(String namespaceName) {
            return namespaceName(Output.of(namespaceName));
        }

        /**
         * @param network The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
         * 
         * @return builder
         * 
         */
        public Builder network(Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
         * 
         * @return builder
         * 
         */
        public Builder network(String network) {
            return network(Output.of(network));
        }

        /**
         * @param retentionInSeconds - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
         * 
         * @return builder
         * 
         */
        public Builder retentionInSeconds(Output<Integer> retentionInSeconds) {
            $.retentionInSeconds = retentionInSeconds;
            return this;
        }

        /**
         * @param retentionInSeconds - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
         * 
         * @return builder
         * 
         */
        public Builder retentionInSeconds(Integer retentionInSeconds) {
            return retentionInSeconds(Output.of(retentionInSeconds));
        }

        public EdgeKvArgs build() {
            $.groupId = Objects.requireNonNull($.groupId, "expected parameter 'groupId' to be non-null");
            $.namespaceName = Objects.requireNonNull($.namespaceName, "expected parameter 'namespaceName' to be non-null");
            $.network = Objects.requireNonNull($.network, "expected parameter 'network' to be non-null");
            $.retentionInSeconds = Objects.requireNonNull($.retentionInSeconds, "expected parameter 'retentionInSeconds' to be non-null");
            return $;
        }
    }

}
