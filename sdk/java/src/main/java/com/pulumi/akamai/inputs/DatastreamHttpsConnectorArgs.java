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


public final class DatastreamHttpsConnectorArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatastreamHttpsConnectorArgs Empty = new DatastreamHttpsConnectorArgs();

    /**
     * Either `NONE` for no authentication, or `BASIC`. For basic authentication, provide the `user_name` and `password` you set in your custom HTTPS endpoint.
     * 
     */
    @Import(name="authenticationType", required=true)
    private Output<String> authenticationType;

    /**
     * @return Either `NONE` for no authentication, or `BASIC`. For basic authentication, provide the `user_name` and `password` you set in your custom HTTPS endpoint.
     * 
     */
    public Output<String> authenticationType() {
        return this.authenticationType;
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
     * **Secret**. Enter the password you set in your custom HTTPS endpoint for authentication.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return **Secret**. Enter the password you set in your custom HTTPS endpoint for authentication.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Enter the secure URL where you want to send and store your logs.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return Enter the secure URL where you want to send and store your logs.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     * **Secret**. Enter the valid username you set in your custom HTTPS endpoint for authentication.
     * 
     */
    @Import(name="userName")
    private @Nullable Output<String> userName;

    /**
     * @return **Secret**. Enter the valid username you set in your custom HTTPS endpoint for authentication.
     * 
     */
    public Optional<Output<String>> userName() {
        return Optional.ofNullable(this.userName);
    }

    private DatastreamHttpsConnectorArgs() {}

    private DatastreamHttpsConnectorArgs(DatastreamHttpsConnectorArgs $) {
        this.authenticationType = $.authenticationType;
        this.compressLogs = $.compressLogs;
        this.connectorId = $.connectorId;
        this.connectorName = $.connectorName;
        this.password = $.password;
        this.url = $.url;
        this.userName = $.userName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatastreamHttpsConnectorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatastreamHttpsConnectorArgs $;

        public Builder() {
            $ = new DatastreamHttpsConnectorArgs();
        }

        public Builder(DatastreamHttpsConnectorArgs defaults) {
            $ = new DatastreamHttpsConnectorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authenticationType Either `NONE` for no authentication, or `BASIC`. For basic authentication, provide the `user_name` and `password` you set in your custom HTTPS endpoint.
         * 
         * @return builder
         * 
         */
        public Builder authenticationType(Output<String> authenticationType) {
            $.authenticationType = authenticationType;
            return this;
        }

        /**
         * @param authenticationType Either `NONE` for no authentication, or `BASIC`. For basic authentication, provide the `user_name` and `password` you set in your custom HTTPS endpoint.
         * 
         * @return builder
         * 
         */
        public Builder authenticationType(String authenticationType) {
            return authenticationType(Output.of(authenticationType));
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
         * @param password **Secret**. Enter the password you set in your custom HTTPS endpoint for authentication.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password **Secret**. Enter the password you set in your custom HTTPS endpoint for authentication.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param url Enter the secure URL where you want to send and store your logs.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Enter the secure URL where you want to send and store your logs.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param userName **Secret**. Enter the valid username you set in your custom HTTPS endpoint for authentication.
         * 
         * @return builder
         * 
         */
        public Builder userName(@Nullable Output<String> userName) {
            $.userName = userName;
            return this;
        }

        /**
         * @param userName **Secret**. Enter the valid username you set in your custom HTTPS endpoint for authentication.
         * 
         * @return builder
         * 
         */
        public Builder userName(String userName) {
            return userName(Output.of(userName));
        }

        public DatastreamHttpsConnectorArgs build() {
            $.authenticationType = Objects.requireNonNull($.authenticationType, "expected parameter 'authenticationType' to be non-null");
            $.connectorName = Objects.requireNonNull($.connectorName, "expected parameter 'connectorName' to be non-null");
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}
