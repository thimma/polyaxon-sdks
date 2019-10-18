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
    define(['ApiClient', 'model/V1Auth', 'model/V1CodeReference', 'model/V1CredsBodyRequest', 'model/V1EntityResourceRequest', 'model/V1ListCodeRefsResponse', 'model/V1ListProjectsResponse', 'model/V1ListRunsResponse', 'model/V1LogHandler', 'model/V1Project', 'model/V1Run', 'model/V1Status', 'model/V1StatusCondition', 'model/V1User', 'model/V1Uuids', 'model/V1Version', 'model/V1Versions', 'api/AuthV1Api', 'api/ProjectsV1Api', 'api/RunsV1Api', 'api/UsersV1Api', 'api/VersionsV1Api'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('./ApiClient'), require('./model/V1Auth'), require('./model/V1CodeReference'), require('./model/V1CredsBodyRequest'), require('./model/V1EntityResourceRequest'), require('./model/V1ListCodeRefsResponse'), require('./model/V1ListProjectsResponse'), require('./model/V1ListRunsResponse'), require('./model/V1LogHandler'), require('./model/V1Project'), require('./model/V1Run'), require('./model/V1Status'), require('./model/V1StatusCondition'), require('./model/V1User'), require('./model/V1Uuids'), require('./model/V1Version'), require('./model/V1Versions'), require('./api/AuthV1Api'), require('./api/ProjectsV1Api'), require('./api/RunsV1Api'), require('./api/UsersV1Api'), require('./api/VersionsV1Api'));
  }
}(function(ApiClient, V1Auth, V1CodeReference, V1CredsBodyRequest, V1EntityResourceRequest, V1ListCodeRefsResponse, V1ListProjectsResponse, V1ListRunsResponse, V1LogHandler, V1Project, V1Run, V1Status, V1StatusCondition, V1User, V1Uuids, V1Version, V1Versions, AuthV1Api, ProjectsV1Api, RunsV1Api, UsersV1Api, VersionsV1Api) {
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
     * The V1Auth model constructor.
     * @property {module:model/V1Auth}
     */
    V1Auth: V1Auth,
    /**
     * The V1CodeReference model constructor.
     * @property {module:model/V1CodeReference}
     */
    V1CodeReference: V1CodeReference,
    /**
     * The V1CredsBodyRequest model constructor.
     * @property {module:model/V1CredsBodyRequest}
     */
    V1CredsBodyRequest: V1CredsBodyRequest,
    /**
     * The V1EntityResourceRequest model constructor.
     * @property {module:model/V1EntityResourceRequest}
     */
    V1EntityResourceRequest: V1EntityResourceRequest,
    /**
     * The V1ListCodeRefsResponse model constructor.
     * @property {module:model/V1ListCodeRefsResponse}
     */
    V1ListCodeRefsResponse: V1ListCodeRefsResponse,
    /**
     * The V1ListProjectsResponse model constructor.
     * @property {module:model/V1ListProjectsResponse}
     */
    V1ListProjectsResponse: V1ListProjectsResponse,
    /**
     * The V1ListRunsResponse model constructor.
     * @property {module:model/V1ListRunsResponse}
     */
    V1ListRunsResponse: V1ListRunsResponse,
    /**
     * The V1LogHandler model constructor.
     * @property {module:model/V1LogHandler}
     */
    V1LogHandler: V1LogHandler,
    /**
     * The V1Project model constructor.
     * @property {module:model/V1Project}
     */
    V1Project: V1Project,
    /**
     * The V1Run model constructor.
     * @property {module:model/V1Run}
     */
    V1Run: V1Run,
    /**
     * The V1Status model constructor.
     * @property {module:model/V1Status}
     */
    V1Status: V1Status,
    /**
     * The V1StatusCondition model constructor.
     * @property {module:model/V1StatusCondition}
     */
    V1StatusCondition: V1StatusCondition,
    /**
     * The V1User model constructor.
     * @property {module:model/V1User}
     */
    V1User: V1User,
    /**
     * The V1Uuids model constructor.
     * @property {module:model/V1Uuids}
     */
    V1Uuids: V1Uuids,
    /**
     * The V1Version model constructor.
     * @property {module:model/V1Version}
     */
    V1Version: V1Version,
    /**
     * The V1Versions model constructor.
     * @property {module:model/V1Versions}
     */
    V1Versions: V1Versions,
    /**
     * The AuthV1Api service constructor.
     * @property {module:api/AuthV1Api}
     */
    AuthV1Api: AuthV1Api,
    /**
     * The ProjectsV1Api service constructor.
     * @property {module:api/ProjectsV1Api}
     */
    ProjectsV1Api: ProjectsV1Api,
    /**
     * The RunsV1Api service constructor.
     * @property {module:api/RunsV1Api}
     */
    RunsV1Api: RunsV1Api,
    /**
     * The UsersV1Api service constructor.
     * @property {module:api/UsersV1Api}
     */
    UsersV1Api: UsersV1Api,
    /**
     * The VersionsV1Api service constructor.
     * @property {module:api/VersionsV1Api}
     */
    VersionsV1Api: VersionsV1Api
  };

  return exports;
}));
