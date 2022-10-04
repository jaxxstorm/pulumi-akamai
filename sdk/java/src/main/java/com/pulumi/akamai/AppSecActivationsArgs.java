// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppSecActivationsArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppSecActivationsArgs Empty = new AppSecActivationsArgs();

    /**
     * . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
     * 
     * @deprecated
     * The setting activate has been deprecated; &#34;terraform apply&#34; will always perform activation. (Use &#34;terraform destroy&#34; for deactivation.)
     * 
     */
    @Deprecated /* The setting activate has been deprecated; ""terraform apply"" will always perform activation. (Use ""terraform destroy"" for deactivation.) */
    @Import(name="activate")
    private @Nullable Output<Boolean> activate;

    /**
     * @return . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
     * 
     * @deprecated
     * The setting activate has been deprecated; &#34;terraform apply&#34; will always perform activation. (Use &#34;terraform destroy&#34; for deactivation.)
     * 
     */
    @Deprecated /* The setting activate has been deprecated; ""terraform apply"" will always perform activation. (Use ""terraform destroy"" for deactivation.) */
    public Optional<Output<Boolean>> activate() {
        return Optional.ofNullable(this.activate);
    }

    /**
     * . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
     * * **PRODUCTION**
     * * **STAGING**
     * 
     */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
     * * **PRODUCTION**
     * * **STAGING**
     * 
     */
    public Optional<Output<String>> network() {
        return Optional.ofNullable(this.network);
    }

    /**
     * . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That&#39;s because something must be different in order to trigger these processes. Because of that, it&#39;s recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
     * 
     */
    @Import(name="note")
    private @Nullable Output<String> note;

    /**
     * @return . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That&#39;s because something must be different in order to trigger these processes. Because of that, it&#39;s recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
     * 
     */
    public Optional<Output<String>> note() {
        return Optional.ofNullable(this.note);
    }

    /**
     * . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That&#39;s because something must be different in order to trigger one of these processes. Because of that, it&#39;s recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
     * 
     * @deprecated
     * The setting notes has been deprecated. Use &#34;note&#34; instead.
     * 
     */
    @Deprecated /* The setting notes has been deprecated. Use ""note"" instead. */
    @Import(name="notes")
    private @Nullable Output<String> notes;

    /**
     * @return . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That&#39;s because something must be different in order to trigger one of these processes. Because of that, it&#39;s recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
     * 
     * @deprecated
     * The setting notes has been deprecated. Use &#34;note&#34; instead.
     * 
     */
    @Deprecated /* The setting notes has been deprecated. Use ""note"" instead. */
    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
    }

    /**
     * . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
     * 
     */
    @Import(name="notificationEmails", required=true)
    private Output<List<String>> notificationEmails;

    /**
     * @return . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
     * 
     */
    public Output<List<String>> notificationEmails() {
        return this.notificationEmails;
    }

    /**
     * . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
     * 
     */
    @Import(name="version", required=true)
    private Output<Integer> version;

    /**
     * @return . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
     * 
     */
    public Output<Integer> version() {
        return this.version;
    }

    private AppSecActivationsArgs() {}

    private AppSecActivationsArgs(AppSecActivationsArgs $) {
        this.activate = $.activate;
        this.configId = $.configId;
        this.network = $.network;
        this.note = $.note;
        this.notes = $.notes;
        this.notificationEmails = $.notificationEmails;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecActivationsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecActivationsArgs $;

        public Builder() {
            $ = new AppSecActivationsArgs();
        }

        public Builder(AppSecActivationsArgs defaults) {
            $ = new AppSecActivationsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param activate . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
         * 
         * @return builder
         * 
         * @deprecated
         * The setting activate has been deprecated; &#34;terraform apply&#34; will always perform activation. (Use &#34;terraform destroy&#34; for deactivation.)
         * 
         */
        @Deprecated /* The setting activate has been deprecated; ""terraform apply"" will always perform activation. (Use ""terraform destroy"" for deactivation.) */
        public Builder activate(@Nullable Output<Boolean> activate) {
            $.activate = activate;
            return this;
        }

        /**
         * @param activate . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
         * 
         * @return builder
         * 
         * @deprecated
         * The setting activate has been deprecated; &#34;terraform apply&#34; will always perform activation. (Use &#34;terraform destroy&#34; for deactivation.)
         * 
         */
        @Deprecated /* The setting activate has been deprecated; ""terraform apply"" will always perform activation. (Use ""terraform destroy"" for deactivation.) */
        public Builder activate(Boolean activate) {
            return activate(Output.of(activate));
        }

        /**
         * @param configId . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param network . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
         * * **PRODUCTION**
         * * **STAGING**
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
         * * **PRODUCTION**
         * * **STAGING**
         * 
         * @return builder
         * 
         */
        public Builder network(String network) {
            return network(Output.of(network));
        }

        /**
         * @param note . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That&#39;s because something must be different in order to trigger these processes. Because of that, it&#39;s recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
         * 
         * @return builder
         * 
         */
        public Builder note(@Nullable Output<String> note) {
            $.note = note;
            return this;
        }

        /**
         * @param note . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That&#39;s because something must be different in order to trigger these processes. Because of that, it&#39;s recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
         * 
         * @return builder
         * 
         */
        public Builder note(String note) {
            return note(Output.of(note));
        }

        /**
         * @param notes . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That&#39;s because something must be different in order to trigger one of these processes. Because of that, it&#39;s recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
         * 
         * @return builder
         * 
         * @deprecated
         * The setting notes has been deprecated. Use &#34;note&#34; instead.
         * 
         */
        @Deprecated /* The setting notes has been deprecated. Use ""note"" instead. */
        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        /**
         * @param notes . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That&#39;s because something must be different in order to trigger one of these processes. Because of that, it&#39;s recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
         * 
         * @return builder
         * 
         * @deprecated
         * The setting notes has been deprecated. Use &#34;note&#34; instead.
         * 
         */
        @Deprecated /* The setting notes has been deprecated. Use ""note"" instead. */
        public Builder notes(String notes) {
            return notes(Output.of(notes));
        }

        /**
         * @param notificationEmails . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
         * 
         * @return builder
         * 
         */
        public Builder notificationEmails(Output<List<String>> notificationEmails) {
            $.notificationEmails = notificationEmails;
            return this;
        }

        /**
         * @param notificationEmails . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
         * 
         * @return builder
         * 
         */
        public Builder notificationEmails(List<String> notificationEmails) {
            return notificationEmails(Output.of(notificationEmails));
        }

        /**
         * @param notificationEmails . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
         * 
         * @return builder
         * 
         */
        public Builder notificationEmails(String... notificationEmails) {
            return notificationEmails(List.of(notificationEmails));
        }

        /**
         * @param version . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
         * 
         * @return builder
         * 
         */
        public Builder version(Output<Integer> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
         * 
         * @return builder
         * 
         */
        public Builder version(Integer version) {
            return version(Output.of(version));
        }

        public AppSecActivationsArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.notificationEmails = Objects.requireNonNull($.notificationEmails, "expected parameter 'notificationEmails' to be non-null");
            $.version = Objects.requireNonNull($.version, "expected parameter 'version' to be non-null");
            return $;
        }
    }

}
