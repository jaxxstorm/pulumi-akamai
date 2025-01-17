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


public final class DatastreamOracleConnectorArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatastreamOracleConnectorArgs Empty = new DatastreamOracleConnectorArgs();

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
     * The name of the Oracle Cloud Storage bucket. See [Working with Oracle Cloud Storage buckets](https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/managingbuckets.htm).
     * 
     */
    @Import(name="bucket", required=true)
    private Output<String> bucket;

    /**
     * @return The name of the Oracle Cloud Storage bucket. See [Working with Oracle Cloud Storage buckets](https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/managingbuckets.htm).
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
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
     * The namespace of your Oracle Cloud Storage account. See [Understanding Object Storage namespaces](https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/understandingnamespaces.htm).
     * 
     */
    @Import(name="namespace", required=true)
    private Output<String> namespace;

    /**
     * @return The namespace of your Oracle Cloud Storage account. See [Understanding Object Storage namespaces](https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/understandingnamespaces.htm).
     * 
     */
    public Output<String> namespace() {
        return this.namespace;
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

    /**
     * The Oracle Cloud Storage region where your bucket resides. See [Regions and availability domains in OCS](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm).
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return The Oracle Cloud Storage region where your bucket resides. See [Regions and availability domains in OCS](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm).
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     * **Secret**. The secret access key identifier that you use to authenticate requests to your Oracle Cloud account.
     * 
     */
    @Import(name="secretAccessKey", required=true)
    private Output<String> secretAccessKey;

    /**
     * @return **Secret**. The secret access key identifier that you use to authenticate requests to your Oracle Cloud account.
     * 
     */
    public Output<String> secretAccessKey() {
        return this.secretAccessKey;
    }

    private DatastreamOracleConnectorArgs() {}

    private DatastreamOracleConnectorArgs(DatastreamOracleConnectorArgs $) {
        this.accessKey = $.accessKey;
        this.bucket = $.bucket;
        this.compressLogs = $.compressLogs;
        this.connectorId = $.connectorId;
        this.connectorName = $.connectorName;
        this.namespace = $.namespace;
        this.path = $.path;
        this.region = $.region;
        this.secretAccessKey = $.secretAccessKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatastreamOracleConnectorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatastreamOracleConnectorArgs $;

        public Builder() {
            $ = new DatastreamOracleConnectorArgs();
        }

        public Builder(DatastreamOracleConnectorArgs defaults) {
            $ = new DatastreamOracleConnectorArgs(Objects.requireNonNull(defaults));
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
         * @param bucket The name of the Oracle Cloud Storage bucket. See [Working with Oracle Cloud Storage buckets](https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/managingbuckets.htm).
         * 
         * @return builder
         * 
         */
        public Builder bucket(Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket The name of the Oracle Cloud Storage bucket. See [Working with Oracle Cloud Storage buckets](https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/managingbuckets.htm).
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
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
         * @param namespace The namespace of your Oracle Cloud Storage account. See [Understanding Object Storage namespaces](https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/understandingnamespaces.htm).
         * 
         * @return builder
         * 
         */
        public Builder namespace(Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The namespace of your Oracle Cloud Storage account. See [Understanding Object Storage namespaces](https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/understandingnamespaces.htm).
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
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

        /**
         * @param region The Oracle Cloud Storage region where your bucket resides. See [Regions and availability domains in OCS](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm).
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The Oracle Cloud Storage region where your bucket resides. See [Regions and availability domains in OCS](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm).
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param secretAccessKey **Secret**. The secret access key identifier that you use to authenticate requests to your Oracle Cloud account.
         * 
         * @return builder
         * 
         */
        public Builder secretAccessKey(Output<String> secretAccessKey) {
            $.secretAccessKey = secretAccessKey;
            return this;
        }

        /**
         * @param secretAccessKey **Secret**. The secret access key identifier that you use to authenticate requests to your Oracle Cloud account.
         * 
         * @return builder
         * 
         */
        public Builder secretAccessKey(String secretAccessKey) {
            return secretAccessKey(Output.of(secretAccessKey));
        }

        public DatastreamOracleConnectorArgs build() {
            $.accessKey = Objects.requireNonNull($.accessKey, "expected parameter 'accessKey' to be non-null");
            $.bucket = Objects.requireNonNull($.bucket, "expected parameter 'bucket' to be non-null");
            $.connectorName = Objects.requireNonNull($.connectorName, "expected parameter 'connectorName' to be non-null");
            $.namespace = Objects.requireNonNull($.namespace, "expected parameter 'namespace' to be non-null");
            $.path = Objects.requireNonNull($.path, "expected parameter 'path' to be non-null");
            $.region = Objects.requireNonNull($.region, "expected parameter 'region' to be non-null");
            $.secretAccessKey = Objects.requireNonNull($.secretAccessKey, "expected parameter 'secretAccessKey' to be non-null");
            return $;
        }
    }

}
