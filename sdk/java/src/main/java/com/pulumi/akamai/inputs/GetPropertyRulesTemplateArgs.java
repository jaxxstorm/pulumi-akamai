// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.akamai.inputs.GetPropertyRulesTemplateTemplateArgs;
import com.pulumi.akamai.inputs.GetPropertyRulesTemplateVariableArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPropertyRulesTemplateArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPropertyRulesTemplateArgs Empty = new GetPropertyRulesTemplateArgs();

    /**
     * The absolute path to your top-level JSON template file. The top-level template combines smaller, nested JSON templates to form your property rule tree. This argument conflicts with the `template` argument.
     * 
     */
    @Import(name="templateFile")
    private @Nullable Output<String> templateFile;

    /**
     * @return The absolute path to your top-level JSON template file. The top-level template combines smaller, nested JSON templates to form your property rule tree. This argument conflicts with the `template` argument.
     * 
     */
    public Optional<Output<String>> templateFile() {
        return Optional.ofNullable(this.templateFile);
    }

    /**
     * The template you use in your configuration. This argument conflicts with the `template_file` argument.
     * 
     */
    @Import(name="templates")
    private @Nullable Output<List<GetPropertyRulesTemplateTemplateArgs>> templates;

    /**
     * @return The template you use in your configuration. This argument conflicts with the `template_file` argument.
     * 
     */
    public Optional<Output<List<GetPropertyRulesTemplateTemplateArgs>>> templates() {
        return Optional.ofNullable(this.templates);
    }

    /**
     * Required when using `var_values_file`. The absolute path to the file containing variable definitions and defaults. This argument conflicts with the `variables` argument.
     * 
     */
    @Import(name="varDefinitionFile")
    private @Nullable Output<String> varDefinitionFile;

    /**
     * @return Required when using `var_values_file`. The absolute path to the file containing variable definitions and defaults. This argument conflicts with the `variables` argument.
     * 
     */
    public Optional<Output<String>> varDefinitionFile() {
        return Optional.ofNullable(this.varDefinitionFile);
    }

    /**
     * Required when using `var_definition_file`. The absolute path to the file containing variable values. This argument conflicts with the `variables` argument.
     * 
     */
    @Import(name="varValuesFile")
    private @Nullable Output<String> varValuesFile;

    /**
     * @return Required when using `var_definition_file`. The absolute path to the file containing variable values. This argument conflicts with the `variables` argument.
     * 
     */
    public Optional<Output<String>> varValuesFile() {
        return Optional.ofNullable(this.varValuesFile);
    }

    /**
     * The definition of one or more variables. This argument conflicts with the `var_definition_file` and `var_values_file` arguments. A `variables` block includes:
     * 
     */
    @Import(name="variables")
    private @Nullable Output<List<GetPropertyRulesTemplateVariableArgs>> variables;

    /**
     * @return The definition of one or more variables. This argument conflicts with the `var_definition_file` and `var_values_file` arguments. A `variables` block includes:
     * 
     */
    public Optional<Output<List<GetPropertyRulesTemplateVariableArgs>>> variables() {
        return Optional.ofNullable(this.variables);
    }

    private GetPropertyRulesTemplateArgs() {}

    private GetPropertyRulesTemplateArgs(GetPropertyRulesTemplateArgs $) {
        this.templateFile = $.templateFile;
        this.templates = $.templates;
        this.varDefinitionFile = $.varDefinitionFile;
        this.varValuesFile = $.varValuesFile;
        this.variables = $.variables;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPropertyRulesTemplateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPropertyRulesTemplateArgs $;

        public Builder() {
            $ = new GetPropertyRulesTemplateArgs();
        }

        public Builder(GetPropertyRulesTemplateArgs defaults) {
            $ = new GetPropertyRulesTemplateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param templateFile The absolute path to your top-level JSON template file. The top-level template combines smaller, nested JSON templates to form your property rule tree. This argument conflicts with the `template` argument.
         * 
         * @return builder
         * 
         */
        public Builder templateFile(@Nullable Output<String> templateFile) {
            $.templateFile = templateFile;
            return this;
        }

        /**
         * @param templateFile The absolute path to your top-level JSON template file. The top-level template combines smaller, nested JSON templates to form your property rule tree. This argument conflicts with the `template` argument.
         * 
         * @return builder
         * 
         */
        public Builder templateFile(String templateFile) {
            return templateFile(Output.of(templateFile));
        }

        /**
         * @param templates The template you use in your configuration. This argument conflicts with the `template_file` argument.
         * 
         * @return builder
         * 
         */
        public Builder templates(@Nullable Output<List<GetPropertyRulesTemplateTemplateArgs>> templates) {
            $.templates = templates;
            return this;
        }

        /**
         * @param templates The template you use in your configuration. This argument conflicts with the `template_file` argument.
         * 
         * @return builder
         * 
         */
        public Builder templates(List<GetPropertyRulesTemplateTemplateArgs> templates) {
            return templates(Output.of(templates));
        }

        /**
         * @param templates The template you use in your configuration. This argument conflicts with the `template_file` argument.
         * 
         * @return builder
         * 
         */
        public Builder templates(GetPropertyRulesTemplateTemplateArgs... templates) {
            return templates(List.of(templates));
        }

        /**
         * @param varDefinitionFile Required when using `var_values_file`. The absolute path to the file containing variable definitions and defaults. This argument conflicts with the `variables` argument.
         * 
         * @return builder
         * 
         */
        public Builder varDefinitionFile(@Nullable Output<String> varDefinitionFile) {
            $.varDefinitionFile = varDefinitionFile;
            return this;
        }

        /**
         * @param varDefinitionFile Required when using `var_values_file`. The absolute path to the file containing variable definitions and defaults. This argument conflicts with the `variables` argument.
         * 
         * @return builder
         * 
         */
        public Builder varDefinitionFile(String varDefinitionFile) {
            return varDefinitionFile(Output.of(varDefinitionFile));
        }

        /**
         * @param varValuesFile Required when using `var_definition_file`. The absolute path to the file containing variable values. This argument conflicts with the `variables` argument.
         * 
         * @return builder
         * 
         */
        public Builder varValuesFile(@Nullable Output<String> varValuesFile) {
            $.varValuesFile = varValuesFile;
            return this;
        }

        /**
         * @param varValuesFile Required when using `var_definition_file`. The absolute path to the file containing variable values. This argument conflicts with the `variables` argument.
         * 
         * @return builder
         * 
         */
        public Builder varValuesFile(String varValuesFile) {
            return varValuesFile(Output.of(varValuesFile));
        }

        /**
         * @param variables The definition of one or more variables. This argument conflicts with the `var_definition_file` and `var_values_file` arguments. A `variables` block includes:
         * 
         * @return builder
         * 
         */
        public Builder variables(@Nullable Output<List<GetPropertyRulesTemplateVariableArgs>> variables) {
            $.variables = variables;
            return this;
        }

        /**
         * @param variables The definition of one or more variables. This argument conflicts with the `var_definition_file` and `var_values_file` arguments. A `variables` block includes:
         * 
         * @return builder
         * 
         */
        public Builder variables(List<GetPropertyRulesTemplateVariableArgs> variables) {
            return variables(Output.of(variables));
        }

        /**
         * @param variables The definition of one or more variables. This argument conflicts with the `var_definition_file` and `var_values_file` arguments. A `variables` block includes:
         * 
         * @return builder
         * 
         */
        public Builder variables(GetPropertyRulesTemplateVariableArgs... variables) {
            return variables(List.of(variables));
        }

        public GetPropertyRulesTemplateArgs build() {
            return $;
        }
    }

}
