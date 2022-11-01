// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPropertyRulesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPropertyRulesPlainArgs Empty = new GetPropertyRulesPlainArgs();

    /**
     * - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
     * 
     */
    @Import(name="contractId")
    private @Nullable String contractId;

    /**
     * @return - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
     * 
     */
    public Optional<String> contractId() {
        return Optional.ofNullable(this.contractId);
    }

    /**
     * - (Required) A group&#39;s unique ID, including the `grp_` prefix.
     * 
     */
    @Import(name="groupId")
    private @Nullable String groupId;

    /**
     * @return - (Required) A group&#39;s unique ID, including the `grp_` prefix.
     * 
     */
    public Optional<String> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * - (Required) A property&#39;s unique ID, including the `prp_` prefix.
     * 
     */
    @Import(name="propertyId", required=true)
    private String propertyId;

    /**
     * @return - (Required) A property&#39;s unique ID, including the `prp_` prefix.
     * 
     */
    public String propertyId() {
        return this.propertyId;
    }

    @Import(name="ruleFormat")
    private @Nullable String ruleFormat;

    public Optional<String> ruleFormat() {
        return Optional.ofNullable(this.ruleFormat);
    }

    /**
     * - (Optional) The version to return. Returns the latest version by default.
     * 
     */
    @Import(name="version")
    private @Nullable Integer version;

    /**
     * @return - (Optional) The version to return. Returns the latest version by default.
     * 
     */
    public Optional<Integer> version() {
        return Optional.ofNullable(this.version);
    }

    private GetPropertyRulesPlainArgs() {}

    private GetPropertyRulesPlainArgs(GetPropertyRulesPlainArgs $) {
        this.contractId = $.contractId;
        this.groupId = $.groupId;
        this.propertyId = $.propertyId;
        this.ruleFormat = $.ruleFormat;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPropertyRulesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPropertyRulesPlainArgs $;

        public Builder() {
            $ = new GetPropertyRulesPlainArgs();
        }

        public Builder(GetPropertyRulesPlainArgs defaults) {
            $ = new GetPropertyRulesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param contractId - (Required) A contract&#39;s unique ID, including the `ctr_` prefix.
         * 
         * @return builder
         * 
         */
        public Builder contractId(@Nullable String contractId) {
            $.contractId = contractId;
            return this;
        }

        /**
         * @param groupId - (Required) A group&#39;s unique ID, including the `grp_` prefix.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable String groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param propertyId - (Required) A property&#39;s unique ID, including the `prp_` prefix.
         * 
         * @return builder
         * 
         */
        public Builder propertyId(String propertyId) {
            $.propertyId = propertyId;
            return this;
        }

        public Builder ruleFormat(@Nullable String ruleFormat) {
            $.ruleFormat = ruleFormat;
            return this;
        }

        /**
         * @param version - (Optional) The version to return. Returns the latest version by default.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Integer version) {
            $.version = version;
            return this;
        }

        public GetPropertyRulesPlainArgs build() {
            $.propertyId = Objects.requireNonNull($.propertyId, "expected parameter 'propertyId' to be non-null");
            return $;
        }
    }

}