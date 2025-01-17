// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAppSecReputationProfilesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecReputationProfilesArgs Empty = new GetAppSecReputationProfilesArgs();

    /**
     * . Unique identifier of the security configuration associated with the reputation profiles.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the reputation profiles.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * . Unique identifier of the reputation profile you want to return information for. If not included, information is returned for all your reputation profiles.
     * 
     */
    @Import(name="reputationProfileId")
    private @Nullable Output<Integer> reputationProfileId;

    /**
     * @return . Unique identifier of the reputation profile you want to return information for. If not included, information is returned for all your reputation profiles.
     * 
     */
    public Optional<Output<Integer>> reputationProfileId() {
        return Optional.ofNullable(this.reputationProfileId);
    }

    private GetAppSecReputationProfilesArgs() {}

    private GetAppSecReputationProfilesArgs(GetAppSecReputationProfilesArgs $) {
        this.configId = $.configId;
        this.reputationProfileId = $.reputationProfileId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecReputationProfilesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecReputationProfilesArgs $;

        public Builder() {
            $ = new GetAppSecReputationProfilesArgs();
        }

        public Builder(GetAppSecReputationProfilesArgs defaults) {
            $ = new GetAppSecReputationProfilesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the reputation profiles.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the reputation profiles.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param reputationProfileId . Unique identifier of the reputation profile you want to return information for. If not included, information is returned for all your reputation profiles.
         * 
         * @return builder
         * 
         */
        public Builder reputationProfileId(@Nullable Output<Integer> reputationProfileId) {
            $.reputationProfileId = reputationProfileId;
            return this;
        }

        /**
         * @param reputationProfileId . Unique identifier of the reputation profile you want to return information for. If not included, information is returned for all your reputation profiles.
         * 
         * @return builder
         * 
         */
        public Builder reputationProfileId(Integer reputationProfileId) {
            return reputationProfileId(Output.of(reputationProfileId));
        }

        public GetAppSecReputationProfilesArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            return $;
        }
    }

}
