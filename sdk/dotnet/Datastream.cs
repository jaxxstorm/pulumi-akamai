// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    /// <summary>
    /// ## Import
    /// 
    /// Basic usagehcl resource "akamai_datastream" "example" {
    /// 
    /// # (resource arguments)
    /// 
    ///  } You can import your Akamai DataStream configuration using a stream version ID. For example
    /// 
    /// ```sh
    ///  $ pulumi import akamai:index/datastream:Datastream example 1234
    /// ```
    /// 
    ///  ~&gt; **IMPORTANT:** For security reasons, this command doesn't import any secrets you specify for your connector. To make sure the state file includes complete data, use this resource to manually add the arguments marked **Secret** above.
    /// </summary>
    [AkamaiResourceType("akamai:index/datastream:Datastream")]
    public partial class Datastream : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether you want to start activating the stream when applying the resource. Either `true` for activating the stream upon sending the request or `false` for leaving the stream inactive after the request.
        /// </summary>
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        /// <summary>
        /// Specify details about the Azure Storage connector configuration in a data stream. Note that currently DataStream supports only streaming data to [block objects](https://docs.microsoft.com/en-us/rest/api/storageservices/understanding-block-blobs--append-blobs--and-page-blobs). The argument includes these sub-arguments:
        /// </summary>
        [Output("azureConnector")]
        public Output<Outputs.DatastreamAzureConnector?> AzureConnector { get; private set; } = null!;

        /// <summary>
        /// Provides information about the log line configuration, log file format, names of log files sent, and file delivery. The argument includes these sub-arguments:
        /// </summary>
        [Output("config")]
        public Output<Outputs.DatastreamConfig> Config { get; private set; } = null!;

        /// <summary>
        /// Identifies the contract that has access to the product.
        /// </summary>
        [Output("contractId")]
        public Output<string> ContractId { get; private set; } = null!;

        /// <summary>
        /// The username who created the stream
        /// </summary>
        [Output("createdBy")]
        public Output<string> CreatedBy { get; private set; } = null!;

        /// <summary>
        /// The date and time when the stream was created
        /// </summary>
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        /// <summary>
        /// Specify details about the Datadog connector in a stream, including:
        /// </summary>
        [Output("datadogConnector")]
        public Output<Outputs.DatastreamDatadogConnector?> DatadogConnector { get; private set; } = null!;

        /// <summary>
        /// Identifiers of the data set fields within the template that you want to receive in logs. The order of the identifiers define how the value for these fields appears in the log lines. See [Data set parameters](https://techdocs.akamai.com/datastream2/reference/data-set-parameters-1).
        /// </summary>
        [Output("datasetFieldsIds")]
        public Output<ImmutableArray<int>> DatasetFieldsIds { get; private set; } = null!;

        /// <summary>
        /// A list of email addresses you want to notify about activations and deactivations of the stream.
        /// </summary>
        [Output("emailIds")]
        public Output<ImmutableArray<string>> EmailIds { get; private set; } = null!;

        /// <summary>
        /// Specify details about the Google Cloud Storage connector you can use in a stream. When validating this connector, DataStream uses the private access key to create an `Akamai_access_verification_&lt;timestamp&gt;.txt` object file in your GCS bucket. You can only see this file if the validation process is successful, and you have access to the Google Cloud Storage bucket where you are trying to send logs. The argument includes these sub-arguments:
        /// </summary>
        [Output("gcsConnector")]
        public Output<Outputs.DatastreamGcsConnector?> GcsConnector { get; private set; } = null!;

        /// <summary>
        /// Identifies the group that has access to the product and this stream configuration.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the user group for which the stream was created
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// Specify details about the custom HTTPS endpoint you can use as a connector for a stream, including:
        /// </summary>
        [Output("httpsConnector")]
        public Output<Outputs.DatastreamHttpsConnector?> HttpsConnector { get; private set; } = null!;

        /// <summary>
        /// The username who modified the stream
        /// </summary>
        [Output("modifiedBy")]
        public Output<string> ModifiedBy { get; private set; } = null!;

        /// <summary>
        /// The date and time when the stream was modified
        /// </summary>
        [Output("modifiedDate")]
        public Output<string> ModifiedDate { get; private set; } = null!;

        /// <summary>
        /// Specify details about the Oracle Cloud Storage connector in a stream. When validating this connector, DataStream uses the provided `access_key` and `secret_access_key` values and tries to save an `Akamai_access_verification_&lt;timestamp&gt;.txt` file in your Oracle Cloud Storage folder. You can only see this file if the validation process is successful, and you have access to the Oracle Cloud Storage bucket and folder that you’re trying to send logs to.
        /// </summary>
        [Output("oracleConnector")]
        public Output<Outputs.DatastreamOracleConnector?> OracleConnector { get; private set; } = null!;

        /// <summary>
        /// The configuration in JSON format that can be copy-pasted into PAPI configuration to enable datastream behavior
        /// </summary>
        [Output("papiJson")]
        public Output<string> PapiJson { get; private set; } = null!;

        /// <summary>
        /// The ID of the product for which the stream was created
        /// </summary>
        [Output("productId")]
        public Output<string> ProductId { get; private set; } = null!;

        /// <summary>
        /// The name of the product for which the stream was created
        /// </summary>
        [Output("productName")]
        public Output<string> ProductName { get; private set; } = null!;

        /// <summary>
        /// Identifies the properties that you want to monitor in the stream. Note that a stream can only log data for active properties.
        /// </summary>
        [Output("propertyIds")]
        public Output<ImmutableArray<string>> PropertyIds { get; private set; } = null!;

        /// <summary>
        /// Specify details about the Amazon S3 connector in a stream. When validating this connector, DataStream uses the provided `access_key` and `secret_access_key` values and saves an `akamai_write_test_2147483647.txt` file in your Amazon S3 folder. You can only see this file if validation succeeds, and you have access to the Amazon S3 bucket and folder that you’re trying to send logs to. The argument includes these sub-arguments:
        /// </summary>
        [Output("s3Connector")]
        public Output<Outputs.DatastreamS3Connector?> S3Connector { get; private set; } = null!;

        /// <summary>
        /// Specify details about the Splunk connector in your stream. Note that currently DataStream supports only endpoint URLs ending with `collector/raw`. The argument includes these sub-arguments:
        /// </summary>
        [Output("splunkConnector")]
        public Output<Outputs.DatastreamSplunkConnector?> SplunkConnector { get; private set; } = null!;

        /// <summary>
        /// The name of the stream.
        /// </summary>
        [Output("streamName")]
        public Output<string> StreamName { get; private set; } = null!;

        /// <summary>
        /// The type of stream that you want to create. Currently, `RAW_LOGS` is the only possible stream type.
        /// </summary>
        [Output("streamType")]
        public Output<string> StreamType { get; private set; } = null!;

        /// <summary>
        /// Identifies the configuration version of the stream
        /// </summary>
        [Output("streamVersionId")]
        public Output<int> StreamVersionId { get; private set; } = null!;

        /// <summary>
        /// Specify details about the Sumo Logic connector in a stream, including:
        /// </summary>
        [Output("sumologicConnector")]
        public Output<Outputs.DatastreamSumologicConnector?> SumologicConnector { get; private set; } = null!;

        /// <summary>
        /// The name of the data set template available for the product that you want to use in the stream. Currently, `EDGE_LOGS` is the only data set template available.
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;


        /// <summary>
        /// Create a Datastream resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Datastream(string name, DatastreamArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/datastream:Datastream", name, args ?? new DatastreamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Datastream(string name, Input<string> id, DatastreamState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/datastream:Datastream", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Datastream resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Datastream Get(string name, Input<string> id, DatastreamState? state = null, CustomResourceOptions? options = null)
        {
            return new Datastream(name, id, state, options);
        }
    }

    public sealed class DatastreamArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether you want to start activating the stream when applying the resource. Either `true` for activating the stream upon sending the request or `false` for leaving the stream inactive after the request.
        /// </summary>
        [Input("active", required: true)]
        public Input<bool> Active { get; set; } = null!;

        /// <summary>
        /// Specify details about the Azure Storage connector configuration in a data stream. Note that currently DataStream supports only streaming data to [block objects](https://docs.microsoft.com/en-us/rest/api/storageservices/understanding-block-blobs--append-blobs--and-page-blobs). The argument includes these sub-arguments:
        /// </summary>
        [Input("azureConnector")]
        public Input<Inputs.DatastreamAzureConnectorArgs>? AzureConnector { get; set; }

        /// <summary>
        /// Provides information about the log line configuration, log file format, names of log files sent, and file delivery. The argument includes these sub-arguments:
        /// </summary>
        [Input("config", required: true)]
        public Input<Inputs.DatastreamConfigArgs> Config { get; set; } = null!;

        /// <summary>
        /// Identifies the contract that has access to the product.
        /// </summary>
        [Input("contractId", required: true)]
        public Input<string> ContractId { get; set; } = null!;

        /// <summary>
        /// Specify details about the Datadog connector in a stream, including:
        /// </summary>
        [Input("datadogConnector")]
        public Input<Inputs.DatastreamDatadogConnectorArgs>? DatadogConnector { get; set; }

        [Input("datasetFieldsIds", required: true)]
        private InputList<int>? _datasetFieldsIds;

        /// <summary>
        /// Identifiers of the data set fields within the template that you want to receive in logs. The order of the identifiers define how the value for these fields appears in the log lines. See [Data set parameters](https://techdocs.akamai.com/datastream2/reference/data-set-parameters-1).
        /// </summary>
        public InputList<int> DatasetFieldsIds
        {
            get => _datasetFieldsIds ?? (_datasetFieldsIds = new InputList<int>());
            set => _datasetFieldsIds = value;
        }

        [Input("emailIds")]
        private InputList<string>? _emailIds;

        /// <summary>
        /// A list of email addresses you want to notify about activations and deactivations of the stream.
        /// </summary>
        public InputList<string> EmailIds
        {
            get => _emailIds ?? (_emailIds = new InputList<string>());
            set => _emailIds = value;
        }

        /// <summary>
        /// Specify details about the Google Cloud Storage connector you can use in a stream. When validating this connector, DataStream uses the private access key to create an `Akamai_access_verification_&lt;timestamp&gt;.txt` object file in your GCS bucket. You can only see this file if the validation process is successful, and you have access to the Google Cloud Storage bucket where you are trying to send logs. The argument includes these sub-arguments:
        /// </summary>
        [Input("gcsConnector")]
        public Input<Inputs.DatastreamGcsConnectorArgs>? GcsConnector { get; set; }

        /// <summary>
        /// Identifies the group that has access to the product and this stream configuration.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// Specify details about the custom HTTPS endpoint you can use as a connector for a stream, including:
        /// </summary>
        [Input("httpsConnector")]
        public Input<Inputs.DatastreamHttpsConnectorArgs>? HttpsConnector { get; set; }

        /// <summary>
        /// Specify details about the Oracle Cloud Storage connector in a stream. When validating this connector, DataStream uses the provided `access_key` and `secret_access_key` values and tries to save an `Akamai_access_verification_&lt;timestamp&gt;.txt` file in your Oracle Cloud Storage folder. You can only see this file if the validation process is successful, and you have access to the Oracle Cloud Storage bucket and folder that you’re trying to send logs to.
        /// </summary>
        [Input("oracleConnector")]
        public Input<Inputs.DatastreamOracleConnectorArgs>? OracleConnector { get; set; }

        [Input("propertyIds", required: true)]
        private InputList<string>? _propertyIds;

        /// <summary>
        /// Identifies the properties that you want to monitor in the stream. Note that a stream can only log data for active properties.
        /// </summary>
        public InputList<string> PropertyIds
        {
            get => _propertyIds ?? (_propertyIds = new InputList<string>());
            set => _propertyIds = value;
        }

        /// <summary>
        /// Specify details about the Amazon S3 connector in a stream. When validating this connector, DataStream uses the provided `access_key` and `secret_access_key` values and saves an `akamai_write_test_2147483647.txt` file in your Amazon S3 folder. You can only see this file if validation succeeds, and you have access to the Amazon S3 bucket and folder that you’re trying to send logs to. The argument includes these sub-arguments:
        /// </summary>
        [Input("s3Connector")]
        public Input<Inputs.DatastreamS3ConnectorArgs>? S3Connector { get; set; }

        /// <summary>
        /// Specify details about the Splunk connector in your stream. Note that currently DataStream supports only endpoint URLs ending with `collector/raw`. The argument includes these sub-arguments:
        /// </summary>
        [Input("splunkConnector")]
        public Input<Inputs.DatastreamSplunkConnectorArgs>? SplunkConnector { get; set; }

        /// <summary>
        /// The name of the stream.
        /// </summary>
        [Input("streamName", required: true)]
        public Input<string> StreamName { get; set; } = null!;

        /// <summary>
        /// The type of stream that you want to create. Currently, `RAW_LOGS` is the only possible stream type.
        /// </summary>
        [Input("streamType", required: true)]
        public Input<string> StreamType { get; set; } = null!;

        /// <summary>
        /// Specify details about the Sumo Logic connector in a stream, including:
        /// </summary>
        [Input("sumologicConnector")]
        public Input<Inputs.DatastreamSumologicConnectorArgs>? SumologicConnector { get; set; }

        /// <summary>
        /// The name of the data set template available for the product that you want to use in the stream. Currently, `EDGE_LOGS` is the only data set template available.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        public DatastreamArgs()
        {
        }
    }

    public sealed class DatastreamState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether you want to start activating the stream when applying the resource. Either `true` for activating the stream upon sending the request or `false` for leaving the stream inactive after the request.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Specify details about the Azure Storage connector configuration in a data stream. Note that currently DataStream supports only streaming data to [block objects](https://docs.microsoft.com/en-us/rest/api/storageservices/understanding-block-blobs--append-blobs--and-page-blobs). The argument includes these sub-arguments:
        /// </summary>
        [Input("azureConnector")]
        public Input<Inputs.DatastreamAzureConnectorGetArgs>? AzureConnector { get; set; }

        /// <summary>
        /// Provides information about the log line configuration, log file format, names of log files sent, and file delivery. The argument includes these sub-arguments:
        /// </summary>
        [Input("config")]
        public Input<Inputs.DatastreamConfigGetArgs>? Config { get; set; }

        /// <summary>
        /// Identifies the contract that has access to the product.
        /// </summary>
        [Input("contractId")]
        public Input<string>? ContractId { get; set; }

        /// <summary>
        /// The username who created the stream
        /// </summary>
        [Input("createdBy")]
        public Input<string>? CreatedBy { get; set; }

        /// <summary>
        /// The date and time when the stream was created
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// Specify details about the Datadog connector in a stream, including:
        /// </summary>
        [Input("datadogConnector")]
        public Input<Inputs.DatastreamDatadogConnectorGetArgs>? DatadogConnector { get; set; }

        [Input("datasetFieldsIds")]
        private InputList<int>? _datasetFieldsIds;

        /// <summary>
        /// Identifiers of the data set fields within the template that you want to receive in logs. The order of the identifiers define how the value for these fields appears in the log lines. See [Data set parameters](https://techdocs.akamai.com/datastream2/reference/data-set-parameters-1).
        /// </summary>
        public InputList<int> DatasetFieldsIds
        {
            get => _datasetFieldsIds ?? (_datasetFieldsIds = new InputList<int>());
            set => _datasetFieldsIds = value;
        }

        [Input("emailIds")]
        private InputList<string>? _emailIds;

        /// <summary>
        /// A list of email addresses you want to notify about activations and deactivations of the stream.
        /// </summary>
        public InputList<string> EmailIds
        {
            get => _emailIds ?? (_emailIds = new InputList<string>());
            set => _emailIds = value;
        }

        /// <summary>
        /// Specify details about the Google Cloud Storage connector you can use in a stream. When validating this connector, DataStream uses the private access key to create an `Akamai_access_verification_&lt;timestamp&gt;.txt` object file in your GCS bucket. You can only see this file if the validation process is successful, and you have access to the Google Cloud Storage bucket where you are trying to send logs. The argument includes these sub-arguments:
        /// </summary>
        [Input("gcsConnector")]
        public Input<Inputs.DatastreamGcsConnectorGetArgs>? GcsConnector { get; set; }

        /// <summary>
        /// Identifies the group that has access to the product and this stream configuration.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the user group for which the stream was created
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// Specify details about the custom HTTPS endpoint you can use as a connector for a stream, including:
        /// </summary>
        [Input("httpsConnector")]
        public Input<Inputs.DatastreamHttpsConnectorGetArgs>? HttpsConnector { get; set; }

        /// <summary>
        /// The username who modified the stream
        /// </summary>
        [Input("modifiedBy")]
        public Input<string>? ModifiedBy { get; set; }

        /// <summary>
        /// The date and time when the stream was modified
        /// </summary>
        [Input("modifiedDate")]
        public Input<string>? ModifiedDate { get; set; }

        /// <summary>
        /// Specify details about the Oracle Cloud Storage connector in a stream. When validating this connector, DataStream uses the provided `access_key` and `secret_access_key` values and tries to save an `Akamai_access_verification_&lt;timestamp&gt;.txt` file in your Oracle Cloud Storage folder. You can only see this file if the validation process is successful, and you have access to the Oracle Cloud Storage bucket and folder that you’re trying to send logs to.
        /// </summary>
        [Input("oracleConnector")]
        public Input<Inputs.DatastreamOracleConnectorGetArgs>? OracleConnector { get; set; }

        /// <summary>
        /// The configuration in JSON format that can be copy-pasted into PAPI configuration to enable datastream behavior
        /// </summary>
        [Input("papiJson")]
        public Input<string>? PapiJson { get; set; }

        /// <summary>
        /// The ID of the product for which the stream was created
        /// </summary>
        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        /// <summary>
        /// The name of the product for which the stream was created
        /// </summary>
        [Input("productName")]
        public Input<string>? ProductName { get; set; }

        [Input("propertyIds")]
        private InputList<string>? _propertyIds;

        /// <summary>
        /// Identifies the properties that you want to monitor in the stream. Note that a stream can only log data for active properties.
        /// </summary>
        public InputList<string> PropertyIds
        {
            get => _propertyIds ?? (_propertyIds = new InputList<string>());
            set => _propertyIds = value;
        }

        /// <summary>
        /// Specify details about the Amazon S3 connector in a stream. When validating this connector, DataStream uses the provided `access_key` and `secret_access_key` values and saves an `akamai_write_test_2147483647.txt` file in your Amazon S3 folder. You can only see this file if validation succeeds, and you have access to the Amazon S3 bucket and folder that you’re trying to send logs to. The argument includes these sub-arguments:
        /// </summary>
        [Input("s3Connector")]
        public Input<Inputs.DatastreamS3ConnectorGetArgs>? S3Connector { get; set; }

        /// <summary>
        /// Specify details about the Splunk connector in your stream. Note that currently DataStream supports only endpoint URLs ending with `collector/raw`. The argument includes these sub-arguments:
        /// </summary>
        [Input("splunkConnector")]
        public Input<Inputs.DatastreamSplunkConnectorGetArgs>? SplunkConnector { get; set; }

        /// <summary>
        /// The name of the stream.
        /// </summary>
        [Input("streamName")]
        public Input<string>? StreamName { get; set; }

        /// <summary>
        /// The type of stream that you want to create. Currently, `RAW_LOGS` is the only possible stream type.
        /// </summary>
        [Input("streamType")]
        public Input<string>? StreamType { get; set; }

        /// <summary>
        /// Identifies the configuration version of the stream
        /// </summary>
        [Input("streamVersionId")]
        public Input<int>? StreamVersionId { get; set; }

        /// <summary>
        /// Specify details about the Sumo Logic connector in a stream, including:
        /// </summary>
        [Input("sumologicConnector")]
        public Input<Inputs.DatastreamSumologicConnectorGetArgs>? SumologicConnector { get; set; }

        /// <summary>
        /// The name of the data set template available for the product that you want to use in the stream. Currently, `EDGE_LOGS` is the only data set template available.
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        public DatastreamState()
        {
        }
    }
}
