// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Polyaxon sdk
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.14.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.7
 *
 * Do not edit the class manually.
 *
 */

(function(factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/ProtobufAny', 'model/V1Build', 'model/V1BuildBodyRequest', 'model/V1BuildStatus', 'model/V1CodeReference', 'model/V1CodeReferenceBodyRequest', 'model/V1Dict', 'model/V1Experiment', 'model/V1ExperimentBodyRequest', 'model/V1ExperimentStatus', 'model/V1Job', 'model/V1JobBodyRequest', 'model/V1JobStatus', 'model/V1KV', 'model/V1ListBuildStatusesResponse', 'model/V1ListBuildsResponse', 'model/V1ListExperimentStatusesResponse', 'model/V1ListExperimentsResponse', 'model/V1ListJobStatusesResponse', 'model/V1ListJobsResponse', 'model/V1OwnedEntityIdRequest', 'model/V1ProjectBodyRequest', 'model/V1StatusResponse', 'api/BuildServiceApi', 'api/ExperimentServiceApi', 'api/JobServiceApi'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('./ApiClient'), require('./model/ProtobufAny'), require('./model/V1Build'), require('./model/V1BuildBodyRequest'), require('./model/V1BuildStatus'), require('./model/V1CodeReference'), require('./model/V1CodeReferenceBodyRequest'), require('./model/V1Dict'), require('./model/V1Experiment'), require('./model/V1ExperimentBodyRequest'), require('./model/V1ExperimentStatus'), require('./model/V1Job'), require('./model/V1JobBodyRequest'), require('./model/V1JobStatus'), require('./model/V1KV'), require('./model/V1ListBuildStatusesResponse'), require('./model/V1ListBuildsResponse'), require('./model/V1ListExperimentStatusesResponse'), require('./model/V1ListExperimentsResponse'), require('./model/V1ListJobStatusesResponse'), require('./model/V1ListJobsResponse'), require('./model/V1OwnedEntityIdRequest'), require('./model/V1ProjectBodyRequest'), require('./model/V1StatusResponse'), require('./api/BuildServiceApi'), require('./api/ExperimentServiceApi'), require('./api/JobServiceApi'));
  }
}(function(ApiClient, ProtobufAny, V1Build, V1BuildBodyRequest, V1BuildStatus, V1CodeReference, V1CodeReferenceBodyRequest, V1Dict, V1Experiment, V1ExperimentBodyRequest, V1ExperimentStatus, V1Job, V1JobBodyRequest, V1JobStatus, V1KV, V1ListBuildStatusesResponse, V1ListBuildsResponse, V1ListExperimentStatusesResponse, V1ListExperimentsResponse, V1ListJobStatusesResponse, V1ListJobsResponse, V1OwnedEntityIdRequest, V1ProjectBodyRequest, V1StatusResponse, BuildServiceApi, ExperimentServiceApi, JobServiceApi) {
  'use strict';

  /**
   * ERROR_UNKNOWN.<br>
   * The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
   * <p>
   * An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
   * <pre>
   * var PolyaxonSdk = require('index'); // See note below*.
   * var xxxSvc = new PolyaxonSdk.XxxApi(); // Allocate the API class we're going to use.
   * var yyyModel = new PolyaxonSdk.Yyy(); // Construct a model instance.
   * yyyModel.someProperty = 'someValue';
   * ...
   * var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
   * ...
   * </pre>
   * <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
   * and put the application logic within the callback function.</em>
   * </p>
   * <p>
   * A non-AMD browser application (discouraged) might do something like this:
   * <pre>
   * var xxxSvc = new PolyaxonSdk.XxxApi(); // Allocate the API class we're going to use.
   * var yyy = new PolyaxonSdk.Yyy(); // Construct a model instance.
   * yyyModel.someProperty = 'someValue';
   * ...
   * var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
   * ...
   * </pre>
   * </p>
   * @module index
   * @version 1.14.4
   */
  var exports = {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient: ApiClient,
    /**
     * The ProtobufAny model constructor.
     * @property {module:model/ProtobufAny}
     */
    ProtobufAny: ProtobufAny,
    /**
     * The V1Build model constructor.
     * @property {module:model/V1Build}
     */
    V1Build: V1Build,
    /**
     * The V1BuildBodyRequest model constructor.
     * @property {module:model/V1BuildBodyRequest}
     */
    V1BuildBodyRequest: V1BuildBodyRequest,
    /**
     * The V1BuildStatus model constructor.
     * @property {module:model/V1BuildStatus}
     */
    V1BuildStatus: V1BuildStatus,
    /**
     * The V1CodeReference model constructor.
     * @property {module:model/V1CodeReference}
     */
    V1CodeReference: V1CodeReference,
    /**
     * The V1CodeReferenceBodyRequest model constructor.
     * @property {module:model/V1CodeReferenceBodyRequest}
     */
    V1CodeReferenceBodyRequest: V1CodeReferenceBodyRequest,
    /**
     * The V1Dict model constructor.
     * @property {module:model/V1Dict}
     */
    V1Dict: V1Dict,
    /**
     * The V1Experiment model constructor.
     * @property {module:model/V1Experiment}
     */
    V1Experiment: V1Experiment,
    /**
     * The V1ExperimentBodyRequest model constructor.
     * @property {module:model/V1ExperimentBodyRequest}
     */
    V1ExperimentBodyRequest: V1ExperimentBodyRequest,
    /**
     * The V1ExperimentStatus model constructor.
     * @property {module:model/V1ExperimentStatus}
     */
    V1ExperimentStatus: V1ExperimentStatus,
    /**
     * The V1Job model constructor.
     * @property {module:model/V1Job}
     */
    V1Job: V1Job,
    /**
     * The V1JobBodyRequest model constructor.
     * @property {module:model/V1JobBodyRequest}
     */
    V1JobBodyRequest: V1JobBodyRequest,
    /**
     * The V1JobStatus model constructor.
     * @property {module:model/V1JobStatus}
     */
    V1JobStatus: V1JobStatus,
    /**
     * The V1KV model constructor.
     * @property {module:model/V1KV}
     */
    V1KV: V1KV,
    /**
     * The V1ListBuildStatusesResponse model constructor.
     * @property {module:model/V1ListBuildStatusesResponse}
     */
    V1ListBuildStatusesResponse: V1ListBuildStatusesResponse,
    /**
     * The V1ListBuildsResponse model constructor.
     * @property {module:model/V1ListBuildsResponse}
     */
    V1ListBuildsResponse: V1ListBuildsResponse,
    /**
     * The V1ListExperimentStatusesResponse model constructor.
     * @property {module:model/V1ListExperimentStatusesResponse}
     */
    V1ListExperimentStatusesResponse: V1ListExperimentStatusesResponse,
    /**
     * The V1ListExperimentsResponse model constructor.
     * @property {module:model/V1ListExperimentsResponse}
     */
    V1ListExperimentsResponse: V1ListExperimentsResponse,
    /**
     * The V1ListJobStatusesResponse model constructor.
     * @property {module:model/V1ListJobStatusesResponse}
     */
    V1ListJobStatusesResponse: V1ListJobStatusesResponse,
    /**
     * The V1ListJobsResponse model constructor.
     * @property {module:model/V1ListJobsResponse}
     */
    V1ListJobsResponse: V1ListJobsResponse,
    /**
     * The V1OwnedEntityIdRequest model constructor.
     * @property {module:model/V1OwnedEntityIdRequest}
     */
    V1OwnedEntityIdRequest: V1OwnedEntityIdRequest,
    /**
     * The V1ProjectBodyRequest model constructor.
     * @property {module:model/V1ProjectBodyRequest}
     */
    V1ProjectBodyRequest: V1ProjectBodyRequest,
    /**
     * The V1StatusResponse model constructor.
     * @property {module:model/V1StatusResponse}
     */
    V1StatusResponse: V1StatusResponse,
    /**
     * The BuildServiceApi service constructor.
     * @property {module:api/BuildServiceApi}
     */
    BuildServiceApi: BuildServiceApi,
    /**
     * The ExperimentServiceApi service constructor.
     * @property {module:api/ExperimentServiceApi}
     */
    ExperimentServiceApi: ExperimentServiceApi,
    /**
     * The JobServiceApi service constructor.
     * @property {module:api/JobServiceApi}
     */
    JobServiceApi: JobServiceApi
  };

  return exports;
}));