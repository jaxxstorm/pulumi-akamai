// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.properties.outputs;

import com.pulumi.akamai.properties.outputs.PropertyHostnameCertStatus;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PropertyHostname {
    /**
     * @return The certificate&#39;s provisioning type, either the default `CPS_MANAGED` type for the custom certificates you provision with the [Certificate Provisioning System (CPS)](https://techdocs.akamai.com/cps/docs), or `DEFAULT` for certificates provisioned automatically.
     * 
     */
    private String certProvisioningType;
    private @Nullable List<PropertyHostnameCertStatus> certStatuses;
    /**
     * @return A string containing the original origin&#39;s hostname. For example, `&#34;example.org&#34;`.
     * 
     */
    private String cnameFrom;
    /**
     * @return A string containing the hostname for edge content. For example,  `&#34;example.org.edgesuite.net&#34;`.
     * 
     */
    private String cnameTo;
    private @Nullable String cnameType;
    private @Nullable String edgeHostnameId;

    private PropertyHostname() {}
    /**
     * @return The certificate&#39;s provisioning type, either the default `CPS_MANAGED` type for the custom certificates you provision with the [Certificate Provisioning System (CPS)](https://techdocs.akamai.com/cps/docs), or `DEFAULT` for certificates provisioned automatically.
     * 
     */
    public String certProvisioningType() {
        return this.certProvisioningType;
    }
    public List<PropertyHostnameCertStatus> certStatuses() {
        return this.certStatuses == null ? List.of() : this.certStatuses;
    }
    /**
     * @return A string containing the original origin&#39;s hostname. For example, `&#34;example.org&#34;`.
     * 
     */
    public String cnameFrom() {
        return this.cnameFrom;
    }
    /**
     * @return A string containing the hostname for edge content. For example,  `&#34;example.org.edgesuite.net&#34;`.
     * 
     */
    public String cnameTo() {
        return this.cnameTo;
    }
    public Optional<String> cnameType() {
        return Optional.ofNullable(this.cnameType);
    }
    public Optional<String> edgeHostnameId() {
        return Optional.ofNullable(this.edgeHostnameId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PropertyHostname defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String certProvisioningType;
        private @Nullable List<PropertyHostnameCertStatus> certStatuses;
        private String cnameFrom;
        private String cnameTo;
        private @Nullable String cnameType;
        private @Nullable String edgeHostnameId;
        public Builder() {}
        public Builder(PropertyHostname defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.certProvisioningType = defaults.certProvisioningType;
    	      this.certStatuses = defaults.certStatuses;
    	      this.cnameFrom = defaults.cnameFrom;
    	      this.cnameTo = defaults.cnameTo;
    	      this.cnameType = defaults.cnameType;
    	      this.edgeHostnameId = defaults.edgeHostnameId;
        }

        @CustomType.Setter
        public Builder certProvisioningType(String certProvisioningType) {
            this.certProvisioningType = Objects.requireNonNull(certProvisioningType);
            return this;
        }
        @CustomType.Setter
        public Builder certStatuses(@Nullable List<PropertyHostnameCertStatus> certStatuses) {
            this.certStatuses = certStatuses;
            return this;
        }
        public Builder certStatuses(PropertyHostnameCertStatus... certStatuses) {
            return certStatuses(List.of(certStatuses));
        }
        @CustomType.Setter
        public Builder cnameFrom(String cnameFrom) {
            this.cnameFrom = Objects.requireNonNull(cnameFrom);
            return this;
        }
        @CustomType.Setter
        public Builder cnameTo(String cnameTo) {
            this.cnameTo = Objects.requireNonNull(cnameTo);
            return this;
        }
        @CustomType.Setter
        public Builder cnameType(@Nullable String cnameType) {
            this.cnameType = cnameType;
            return this;
        }
        @CustomType.Setter
        public Builder edgeHostnameId(@Nullable String edgeHostnameId) {
            this.edgeHostnameId = edgeHostnameId;
            return this;
        }
        public PropertyHostname build() {
            final var o = new PropertyHostname();
            o.certProvisioningType = certProvisioningType;
            o.certStatuses = certStatuses;
            o.cnameFrom = cnameFrom;
            o.cnameTo = cnameTo;
            o.cnameType = cnameType;
            o.edgeHostnameId = edgeHostnameId;
            return o;
        }
    }
}