// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatastreamAzureConnectorArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatastreamAzureConnectorArgs Empty = new DatastreamAzureConnectorArgs();

    /**
     * **Secret**. The access key identifier that you use to authenticate requests to your Oracle Cloud account. See [Managing user credentials in OCS](https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm).
     * 
     */
    @Import(name="accessKey", required=true)
    private Output<String> accessKey;

    /**
     * @return **Secret**. The access key identifier that you use to authenticate requests to your Oracle Cloud account. See [Managing user credentials in OCS](https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm).
     * 
     */
    public Output<String> accessKey() {
        return this.accessKey;
    }

    /**
     * Specifies the Azure Storage account name.
     * 
     */
    @Import(name="accountName", required=true)
    private Output<String> accountName;

    /**
     * @return Specifies the Azure Storage account name.
     * 
     */
    public Output<String> accountName() {
        return this.accountName;
    }

    /**
     * Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `true`.
     * 
     */
    @Import(name="compressLogs")
    private @Nullable Output<Boolean> compressLogs;

    /**
     * @return Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> compressLogs() {
        return Optional.ofNullable(this.compressLogs);
    }

    @Import(name="connectorId")
    private @Nullable Output<Integer> connectorId;

    public Optional<Output<Integer>> connectorId() {
        return Optional.ofNullable(this.connectorId);
    }

    /**
     * The name of the connector.
     * 
     */
    @Import(name="connectorName", required=true)
    private Output<String> connectorName;

    /**
     * @return The name of the connector.
     * 
     */
    public Output<String> connectorName() {
        return this.connectorName;
    }

    /**
     * Specifies the Azure Storage container name.
     * 
     */
    @Import(name="containerName", required=true)
    private Output<String> containerName;

    /**
     * @return Specifies the Azure Storage container name.
     * 
     */
    public Output<String> containerName() {
        return this.containerName;
    }

    /**
     * The path to the folder within your Oracle Cloud Storage bucket where you want to store your logs.
     * 
     */
    @Import(name="path", required=true)
    private Output<String> path;

    /**
     * @return The path to the folder within your Oracle Cloud Storage bucket where you want to store your logs.
     * 
     */
    public Output<String> path() {
        return this.path;
    }

    private DatastreamAzureConnectorArgs() {}

    private DatastreamAzureConnectorArgs(DatastreamAzureConnectorArgs $) {
        this.accessKey = $.accessKey;
        this.accountName = $.accountName;
        this.compressLogs = $.compressLogs;
        this.connectorId = $.connectorId;
        this.connectorName = $.connectorName;
        this.containerName = $.containerName;
        this.path = $.path;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatastreamAzureConnectorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatastreamAzureConnectorArgs $;

        public Builder() {
            $ = new DatastreamAzureConnectorArgs();
        }

        public Builder(DatastreamAzureConnectorArgs defaults) {
            $ = new DatastreamAzureConnectorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessKey **Secret**. The access key identifier that you use to authenticate requests to your Oracle Cloud account. See [Managing user credentials in OCS](https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm).
         * 
         * @return builder
         * 
         */
        public Builder accessKey(Output<String> accessKey) {
            $.accessKey = accessKey;
            return this;
        }

        /**
         * @param accessKey **Secret**. The access key identifier that you use to authenticate requests to your Oracle Cloud account. See [Managing user credentials in OCS](https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm).
         * 
         * @return builder
         * 
         */
        public Builder accessKey(String accessKey) {
            return accessKey(Output.of(accessKey));
        }

        /**
         * @param accountName Specifies the Azure Storage account name.
         * 
         * @return builder
         * 
         */
        public Builder accountName(Output<String> accountName) {
            $.accountName = accountName;
            return this;
        }

        /**
         * @param accountName Specifies the Azure Storage account name.
         * 
         * @return builder
         * 
         */
        public Builder accountName(String accountName) {
            return accountName(Output.of(accountName));
        }

        /**
         * @param compressLogs Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder compressLogs(@Nullable Output<Boolean> compressLogs) {
            $.compressLogs = compressLogs;
            return this;
        }

        /**
         * @param compressLogs Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder compressLogs(Boolean compressLogs) {
            return compressLogs(Output.of(compressLogs));
        }

        public Builder connectorId(@Nullable Output<Integer> connectorId) {
            $.connectorId = connectorId;
            return this;
        }

        public Builder connectorId(Integer connectorId) {
            return connectorId(Output.of(connectorId));
        }

        /**
         * @param connectorName The name of the connector.
         * 
         * @return builder
         * 
         */
        public Builder connectorName(Output<String> connectorName) {
            $.connectorName = connectorName;
            return this;
        }

        /**
         * @param connectorName The name of the connector.
         * 
         * @return builder
         * 
         */
        public Builder connectorName(String connectorName) {
            return connectorName(Output.of(connectorName));
        }

        /**
         * @param containerName Specifies the Azure Storage container name.
         * 
         * @return builder
         * 
         */
        public Builder containerName(Output<String> containerName) {
            $.containerName = containerName;
            return this;
        }

        /**
         * @param containerName Specifies the Azure Storage container name.
         * 
         * @return builder
         * 
         */
        public Builder containerName(String containerName) {
            return containerName(Output.of(containerName));
        }

        /**
         * @param path The path to the folder within your Oracle Cloud Storage bucket where you want to store your logs.
         * 
         * @return builder
         * 
         */
        public Builder path(Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path to the folder within your Oracle Cloud Storage bucket where you want to store your logs.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        public DatastreamAzureConnectorArgs build() {
            $.accessKey = Objects.requireNonNull($.accessKey, "expected parameter 'accessKey' to be non-null");
            $.accountName = Objects.requireNonNull($.accountName, "expected parameter 'accountName' to be non-null");
            $.connectorName = Objects.requireNonNull($.connectorName, "expected parameter 'connectorName' to be non-null");
            $.containerName = Objects.requireNonNull($.containerName, "expected parameter 'containerName' to be non-null");
            $.path = Objects.requireNonNull($.path, "expected parameter 'path' to be non-null");
            return $;
        }
    }

}
