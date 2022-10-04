// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAppSecSelectableHostnamesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecSelectableHostnamesPlainArgs Empty = new GetAppSecSelectableHostnamesPlainArgs();

    @Import(name="activeInProduction")
    private @Nullable Boolean activeInProduction;

    public Optional<Boolean> activeInProduction() {
        return Optional.ofNullable(this.activeInProduction);
    }

    @Import(name="activeInStaging")
    private @Nullable Boolean activeInStaging;

    public Optional<Boolean> activeInStaging() {
        return Optional.ofNullable(this.activeInStaging);
    }

    /**
     * . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can&#39;t be used with either the `contractid` or the `groupid` arguments.
     * 
     */
    @Import(name="configId")
    private @Nullable Integer configId;

    /**
     * @return . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can&#39;t be used with either the `contractid` or the `groupid` arguments.
     * 
     */
    public Optional<Integer> configId() {
        return Optional.ofNullable(this.configId);
    }

    /**
     * . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can&#39;t be used with the `config_id` argument.
     * 
     */
    @Import(name="contractid")
    private @Nullable String contractid;

    /**
     * @return . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can&#39;t be used with the `config_id` argument.
     * 
     */
    public Optional<String> contractid() {
        return Optional.ofNullable(this.contractid);
    }

    /**
     * . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can&#39;t be used with the `config_id` argument.
     * 
     */
    @Import(name="groupid")
    private @Nullable Integer groupid;

    /**
     * @return . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can&#39;t be used with the `config_id` argument.
     * 
     */
    public Optional<Integer> groupid() {
        return Optional.ofNullable(this.groupid);
    }

    private GetAppSecSelectableHostnamesPlainArgs() {}

    private GetAppSecSelectableHostnamesPlainArgs(GetAppSecSelectableHostnamesPlainArgs $) {
        this.activeInProduction = $.activeInProduction;
        this.activeInStaging = $.activeInStaging;
        this.configId = $.configId;
        this.contractid = $.contractid;
        this.groupid = $.groupid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecSelectableHostnamesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecSelectableHostnamesPlainArgs $;

        public Builder() {
            $ = new GetAppSecSelectableHostnamesPlainArgs();
        }

        public Builder(GetAppSecSelectableHostnamesPlainArgs defaults) {
            $ = new GetAppSecSelectableHostnamesPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder activeInProduction(@Nullable Boolean activeInProduction) {
            $.activeInProduction = activeInProduction;
            return this;
        }

        public Builder activeInStaging(@Nullable Boolean activeInStaging) {
            $.activeInStaging = activeInStaging;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can&#39;t be used with either the `contractid` or the `groupid` arguments.
         * 
         * @return builder
         * 
         */
        public Builder configId(@Nullable Integer configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param contractid . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can&#39;t be used with the `config_id` argument.
         * 
         * @return builder
         * 
         */
        public Builder contractid(@Nullable String contractid) {
            $.contractid = contractid;
            return this;
        }

        /**
         * @param groupid . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can&#39;t be used with the `config_id` argument.
         * 
         * @return builder
         * 
         */
        public Builder groupid(@Nullable Integer groupid) {
            $.groupid = groupid;
            return this;
        }

        public GetAppSecSelectableHostnamesPlainArgs build() {
            return $;
        }
    }

}
