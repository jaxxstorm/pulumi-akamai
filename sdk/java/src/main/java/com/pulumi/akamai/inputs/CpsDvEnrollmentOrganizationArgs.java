// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CpsDvEnrollmentOrganizationArgs extends com.pulumi.resources.ResourceArgs {

    public static final CpsDvEnrollmentOrganizationArgs Empty = new CpsDvEnrollmentOrganizationArgs();

    /**
     * The address of your organization.
     * 
     */
    @Import(name="addressLineOne", required=true)
    private Output<String> addressLineOne;

    /**
     * @return The address of your organization.
     * 
     */
    public Output<String> addressLineOne() {
        return this.addressLineOne;
    }

    /**
     * The address of your organization.
     * 
     */
    @Import(name="addressLineTwo")
    private @Nullable Output<String> addressLineTwo;

    /**
     * @return The address of your organization.
     * 
     */
    public Optional<Output<String>> addressLineTwo() {
        return Optional.ofNullable(this.addressLineTwo);
    }

    /**
     * The city where your organization resides.
     * 
     */
    @Import(name="city", required=true)
    private Output<String> city;

    /**
     * @return The city where your organization resides.
     * 
     */
    public Output<String> city() {
        return this.city;
    }

    /**
     * The code for the country where your organization resides.
     * 
     */
    @Import(name="countryCode", required=true)
    private Output<String> countryCode;

    /**
     * @return The code for the country where your organization resides.
     * 
     */
    public Output<String> countryCode() {
        return this.countryCode;
    }

    /**
     * The name of your organization.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of your organization.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The phone number of the administrator who you want to use as a contact at your company.
     * 
     */
    @Import(name="phone", required=true)
    private Output<String> phone;

    /**
     * @return The phone number of the administrator who you want to use as a contact at your company.
     * 
     */
    public Output<String> phone() {
        return this.phone;
    }

    /**
     * The postal code of your organization.
     * 
     */
    @Import(name="postalCode", required=true)
    private Output<String> postalCode;

    /**
     * @return The postal code of your organization.
     * 
     */
    public Output<String> postalCode() {
        return this.postalCode;
    }

    /**
     * The region of your organization, typically a state or province.
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return The region of your organization, typically a state or province.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    private CpsDvEnrollmentOrganizationArgs() {}

    private CpsDvEnrollmentOrganizationArgs(CpsDvEnrollmentOrganizationArgs $) {
        this.addressLineOne = $.addressLineOne;
        this.addressLineTwo = $.addressLineTwo;
        this.city = $.city;
        this.countryCode = $.countryCode;
        this.name = $.name;
        this.phone = $.phone;
        this.postalCode = $.postalCode;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CpsDvEnrollmentOrganizationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CpsDvEnrollmentOrganizationArgs $;

        public Builder() {
            $ = new CpsDvEnrollmentOrganizationArgs();
        }

        public Builder(CpsDvEnrollmentOrganizationArgs defaults) {
            $ = new CpsDvEnrollmentOrganizationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressLineOne The address of your organization.
         * 
         * @return builder
         * 
         */
        public Builder addressLineOne(Output<String> addressLineOne) {
            $.addressLineOne = addressLineOne;
            return this;
        }

        /**
         * @param addressLineOne The address of your organization.
         * 
         * @return builder
         * 
         */
        public Builder addressLineOne(String addressLineOne) {
            return addressLineOne(Output.of(addressLineOne));
        }

        /**
         * @param addressLineTwo The address of your organization.
         * 
         * @return builder
         * 
         */
        public Builder addressLineTwo(@Nullable Output<String> addressLineTwo) {
            $.addressLineTwo = addressLineTwo;
            return this;
        }

        /**
         * @param addressLineTwo The address of your organization.
         * 
         * @return builder
         * 
         */
        public Builder addressLineTwo(String addressLineTwo) {
            return addressLineTwo(Output.of(addressLineTwo));
        }

        /**
         * @param city The city where your organization resides.
         * 
         * @return builder
         * 
         */
        public Builder city(Output<String> city) {
            $.city = city;
            return this;
        }

        /**
         * @param city The city where your organization resides.
         * 
         * @return builder
         * 
         */
        public Builder city(String city) {
            return city(Output.of(city));
        }

        /**
         * @param countryCode The code for the country where your organization resides.
         * 
         * @return builder
         * 
         */
        public Builder countryCode(Output<String> countryCode) {
            $.countryCode = countryCode;
            return this;
        }

        /**
         * @param countryCode The code for the country where your organization resides.
         * 
         * @return builder
         * 
         */
        public Builder countryCode(String countryCode) {
            return countryCode(Output.of(countryCode));
        }

        /**
         * @param name The name of your organization.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of your organization.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param phone The phone number of the administrator who you want to use as a contact at your company.
         * 
         * @return builder
         * 
         */
        public Builder phone(Output<String> phone) {
            $.phone = phone;
            return this;
        }

        /**
         * @param phone The phone number of the administrator who you want to use as a contact at your company.
         * 
         * @return builder
         * 
         */
        public Builder phone(String phone) {
            return phone(Output.of(phone));
        }

        /**
         * @param postalCode The postal code of your organization.
         * 
         * @return builder
         * 
         */
        public Builder postalCode(Output<String> postalCode) {
            $.postalCode = postalCode;
            return this;
        }

        /**
         * @param postalCode The postal code of your organization.
         * 
         * @return builder
         * 
         */
        public Builder postalCode(String postalCode) {
            return postalCode(Output.of(postalCode));
        }

        /**
         * @param region The region of your organization, typically a state or province.
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region of your organization, typically a state or province.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public CpsDvEnrollmentOrganizationArgs build() {
            $.addressLineOne = Objects.requireNonNull($.addressLineOne, "expected parameter 'addressLineOne' to be non-null");
            $.city = Objects.requireNonNull($.city, "expected parameter 'city' to be non-null");
            $.countryCode = Objects.requireNonNull($.countryCode, "expected parameter 'countryCode' to be non-null");
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.phone = Objects.requireNonNull($.phone, "expected parameter 'phone' to be non-null");
            $.postalCode = Objects.requireNonNull($.postalCode, "expected parameter 'postalCode' to be non-null");
            $.region = Objects.requireNonNull($.region, "expected parameter 'region' to be non-null");
            return $;
        }
    }

}
