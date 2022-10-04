// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.akamai.inputs.GtmGeomapAssignmentArgs;
import com.pulumi.akamai.inputs.GtmGeomapDefaultDatacenterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GtmGeomapState extends com.pulumi.resources.ResourceArgs {

    public static final GtmGeomapState Empty = new GtmGeomapState();

    /**
     * Contains information about the geographic zone groupings of countries. You can have multiple `assignment` arguments. If used, requires these additional arguments:
     * 
     */
    @Import(name="assignments")
    private @Nullable Output<List<GtmGeomapAssignmentArgs>> assignments;

    /**
     * @return Contains information about the geographic zone groupings of countries. You can have multiple `assignment` arguments. If used, requires these additional arguments:
     * 
     */
    public Optional<Output<List<GtmGeomapAssignmentArgs>>> assignments() {
        return Optional.ofNullable(this.assignments);
    }

    /**
     * A placeholder for all other geographic zones. Requires these additional arguments:
     * 
     */
    @Import(name="defaultDatacenter")
    private @Nullable Output<GtmGeomapDefaultDatacenterArgs> defaultDatacenter;

    /**
     * @return A placeholder for all other geographic zones. Requires these additional arguments:
     * 
     */
    public Optional<Output<GtmGeomapDefaultDatacenterArgs>> defaultDatacenter() {
        return Optional.ofNullable(this.defaultDatacenter);
    }

    /**
     * GTM Domain name for the Geographic Map.
     * 
     */
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return GTM Domain name for the Geographic Map.
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * A descriptive label for the Geographic map.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A descriptive label for the Geographic map.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
     * 
     */
    @Import(name="waitOnComplete")
    private @Nullable Output<Boolean> waitOnComplete;

    /**
     * @return A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
     * 
     */
    public Optional<Output<Boolean>> waitOnComplete() {
        return Optional.ofNullable(this.waitOnComplete);
    }

    private GtmGeomapState() {}

    private GtmGeomapState(GtmGeomapState $) {
        this.assignments = $.assignments;
        this.defaultDatacenter = $.defaultDatacenter;
        this.domain = $.domain;
        this.name = $.name;
        this.waitOnComplete = $.waitOnComplete;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GtmGeomapState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GtmGeomapState $;

        public Builder() {
            $ = new GtmGeomapState();
        }

        public Builder(GtmGeomapState defaults) {
            $ = new GtmGeomapState(Objects.requireNonNull(defaults));
        }

        /**
         * @param assignments Contains information about the geographic zone groupings of countries. You can have multiple `assignment` arguments. If used, requires these additional arguments:
         * 
         * @return builder
         * 
         */
        public Builder assignments(@Nullable Output<List<GtmGeomapAssignmentArgs>> assignments) {
            $.assignments = assignments;
            return this;
        }

        /**
         * @param assignments Contains information about the geographic zone groupings of countries. You can have multiple `assignment` arguments. If used, requires these additional arguments:
         * 
         * @return builder
         * 
         */
        public Builder assignments(List<GtmGeomapAssignmentArgs> assignments) {
            return assignments(Output.of(assignments));
        }

        /**
         * @param assignments Contains information about the geographic zone groupings of countries. You can have multiple `assignment` arguments. If used, requires these additional arguments:
         * 
         * @return builder
         * 
         */
        public Builder assignments(GtmGeomapAssignmentArgs... assignments) {
            return assignments(List.of(assignments));
        }

        /**
         * @param defaultDatacenter A placeholder for all other geographic zones. Requires these additional arguments:
         * 
         * @return builder
         * 
         */
        public Builder defaultDatacenter(@Nullable Output<GtmGeomapDefaultDatacenterArgs> defaultDatacenter) {
            $.defaultDatacenter = defaultDatacenter;
            return this;
        }

        /**
         * @param defaultDatacenter A placeholder for all other geographic zones. Requires these additional arguments:
         * 
         * @return builder
         * 
         */
        public Builder defaultDatacenter(GtmGeomapDefaultDatacenterArgs defaultDatacenter) {
            return defaultDatacenter(Output.of(defaultDatacenter));
        }

        /**
         * @param domain GTM Domain name for the Geographic Map.
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain GTM Domain name for the Geographic Map.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param name A descriptive label for the Geographic map.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A descriptive label for the Geographic map.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param waitOnComplete A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
         * 
         * @return builder
         * 
         */
        public Builder waitOnComplete(@Nullable Output<Boolean> waitOnComplete) {
            $.waitOnComplete = waitOnComplete;
            return this;
        }

        /**
         * @param waitOnComplete A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
         * 
         * @return builder
         * 
         */
        public Builder waitOnComplete(Boolean waitOnComplete) {
            return waitOnComplete(Output.of(waitOnComplete));
        }

        public GtmGeomapState build() {
            return $;
        }
    }

}
