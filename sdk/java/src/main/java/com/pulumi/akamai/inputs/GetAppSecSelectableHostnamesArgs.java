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


public final class GetAppSecSelectableHostnamesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecSelectableHostnamesArgs Empty = new GetAppSecSelectableHostnamesArgs();

    @Import(name="activeInProduction")
    private @Nullable Output<Boolean> activeInProduction;

    public Optional<Output<Boolean>> activeInProduction() {
        return Optional.ofNullable(this.activeInProduction);
    }

    @Import(name="activeInStaging")
    private @Nullable Output<Boolean> activeInStaging;

    public Optional<Output<Boolean>> activeInStaging() {
        return Optional.ofNullable(this.activeInStaging);
    }

    /**
     * . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can&#39;t be used with either the `contractid` or the `groupid` arguments.
     * 
     */
    @Import(name="configId")
    private @Nullable Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can&#39;t be used with either the `contractid` or the `groupid` arguments.
     * 
     */
    public Optional<Output<Integer>> configId() {
        return Optional.ofNullable(this.configId);
    }

    /**
     * . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can&#39;t be used with the `config_id` argument.
     * 
     */
    @Import(name="contractid")
    private @Nullable Output<String> contractid;

    /**
     * @return . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can&#39;t be used with the `config_id` argument.
     * 
     */
    public Optional<Output<String>> contractid() {
        return Optional.ofNullable(this.contractid);
    }

    /**
     * . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can&#39;t be used with the `config_id` argument.
     * 
     */
    @Import(name="groupid")
    private @Nullable Output<Integer> groupid;

    /**
     * @return . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can&#39;t be used with the `config_id` argument.
     * 
     */
    public Optional<Output<Integer>> groupid() {
        return Optional.ofNullable(this.groupid);
    }

    private GetAppSecSelectableHostnamesArgs() {}

    private GetAppSecSelectableHostnamesArgs(GetAppSecSelectableHostnamesArgs $) {
        this.activeInProduction = $.activeInProduction;
        this.activeInStaging = $.activeInStaging;
        this.configId = $.configId;
        this.contractid = $.contractid;
        this.groupid = $.groupid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecSelectableHostnamesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecSelectableHostnamesArgs $;

        public Builder() {
            $ = new GetAppSecSelectableHostnamesArgs();
        }

        public Builder(GetAppSecSelectableHostnamesArgs defaults) {
            $ = new GetAppSecSelectableHostnamesArgs(Objects.requireNonNull(defaults));
        }

        public Builder activeInProduction(@Nullable Output<Boolean> activeInProduction) {
            $.activeInProduction = activeInProduction;
            return this;
        }

        public Builder activeInProduction(Boolean activeInProduction) {
            return activeInProduction(Output.of(activeInProduction));
        }

        public Builder activeInStaging(@Nullable Output<Boolean> activeInStaging) {
            $.activeInStaging = activeInStaging;
            return this;
        }

        public Builder activeInStaging(Boolean activeInStaging) {
            return activeInStaging(Output.of(activeInStaging));
        }

        /**
         * @param configId . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can&#39;t be used with either the `contractid` or the `groupid` arguments.
         * 
         * @return builder
         * 
         */
        public Builder configId(@Nullable Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can&#39;t be used with either the `contractid` or the `groupid` arguments.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param contractid . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can&#39;t be used with the `config_id` argument.
         * 
         * @return builder
         * 
         */
        public Builder contractid(@Nullable Output<String> contractid) {
            $.contractid = contractid;
            return this;
        }

        /**
         * @param contractid . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can&#39;t be used with the `config_id` argument.
         * 
         * @return builder
         * 
         */
        public Builder contractid(String contractid) {
            return contractid(Output.of(contractid));
        }

        /**
         * @param groupid . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can&#39;t be used with the `config_id` argument.
         * 
         * @return builder
         * 
         */
        public Builder groupid(@Nullable Output<Integer> groupid) {
            $.groupid = groupid;
            return this;
        }

        /**
         * @param groupid . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can&#39;t be used with the `config_id` argument.
         * 
         * @return builder
         * 
         */
        public Builder groupid(Integer groupid) {
            return groupid(Output.of(groupid));
        }

        public GetAppSecSelectableHostnamesArgs build() {
            return $;
        }
    }

}
