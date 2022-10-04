// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CpsDvValidationArgs extends com.pulumi.resources.ResourceArgs {

    public static final CpsDvValidationArgs Empty = new CpsDvValidationArgs();

    /**
     * Unique identifier for the DV certificate enrollment.
     * 
     */
    @Import(name="enrollmentId", required=true)
    private Output<Integer> enrollmentId;

    /**
     * @return Unique identifier for the DV certificate enrollment.
     * 
     */
    public Output<Integer> enrollmentId() {
        return this.enrollmentId;
    }

    /**
     * The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
     * 
     */
    @Import(name="sans")
    private @Nullable Output<List<String>> sans;

    /**
     * @return The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
     * 
     */
    public Optional<Output<List<String>>> sans() {
        return Optional.ofNullable(this.sans);
    }

    private CpsDvValidationArgs() {}

    private CpsDvValidationArgs(CpsDvValidationArgs $) {
        this.enrollmentId = $.enrollmentId;
        this.sans = $.sans;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CpsDvValidationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CpsDvValidationArgs $;

        public Builder() {
            $ = new CpsDvValidationArgs();
        }

        public Builder(CpsDvValidationArgs defaults) {
            $ = new CpsDvValidationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enrollmentId Unique identifier for the DV certificate enrollment.
         * 
         * @return builder
         * 
         */
        public Builder enrollmentId(Output<Integer> enrollmentId) {
            $.enrollmentId = enrollmentId;
            return this;
        }

        /**
         * @param enrollmentId Unique identifier for the DV certificate enrollment.
         * 
         * @return builder
         * 
         */
        public Builder enrollmentId(Integer enrollmentId) {
            return enrollmentId(Output.of(enrollmentId));
        }

        /**
         * @param sans The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
         * 
         * @return builder
         * 
         */
        public Builder sans(@Nullable Output<List<String>> sans) {
            $.sans = sans;
            return this;
        }

        /**
         * @param sans The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
         * 
         * @return builder
         * 
         */
        public Builder sans(List<String> sans) {
            return sans(Output.of(sans));
        }

        /**
         * @param sans The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
         * 
         * @return builder
         * 
         */
        public Builder sans(String... sans) {
            return sans(List.of(sans));
        }

        public CpsDvValidationArgs build() {
            $.enrollmentId = Objects.requireNonNull($.enrollmentId, "expected parameter 'enrollmentId' to be non-null");
            return $;
        }
    }

}
