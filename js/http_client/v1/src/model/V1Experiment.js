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

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/V1Dict'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./V1Dict'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.V1Experiment = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.V1Dict);
  }
}(this, function(ApiClient, V1Dict) {
  'use strict';




  /**
   * The V1Experiment model module.
   * @module model/V1Experiment
   * @version 1.14.4
   */

  /**
   * Constructs a new <code>V1Experiment</code>.
   * @alias module:model/V1Experiment
   * @class
   */
  var exports = function() {
    var _this = this;


































  };

  /**
   * Constructs a <code>V1Experiment</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V1Experiment} obj Optional instance to populate.
   * @return {module:model/V1Experiment} The populated <code>V1Experiment</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('id')) {
        obj['id'] = ApiClient.convertToType(data['id'], 'String');
      }
      if (data.hasOwnProperty('uuid')) {
        obj['uuid'] = ApiClient.convertToType(data['uuid'], 'String');
      }
      if (data.hasOwnProperty('unique_name')) {
        obj['unique_name'] = ApiClient.convertToType(data['unique_name'], 'String');
      }
      if (data.hasOwnProperty('name')) {
        obj['name'] = ApiClient.convertToType(data['name'], 'String');
      }
      if (data.hasOwnProperty('description')) {
        obj['description'] = ApiClient.convertToType(data['description'], 'String');
      }
      if (data.hasOwnProperty('tags')) {
        obj['tags'] = ApiClient.convertToType(data['tags'], ['String']);
      }
      if (data.hasOwnProperty('deleted')) {
        obj['deleted'] = ApiClient.convertToType(data['deleted'], 'Boolean');
      }
      if (data.hasOwnProperty('user')) {
        obj['user'] = ApiClient.convertToType(data['user'], 'String');
      }
      if (data.hasOwnProperty('created_at')) {
        obj['created_at'] = ApiClient.convertToType(data['created_at'], 'String');
      }
      if (data.hasOwnProperty('updated_at')) {
        obj['updated_at'] = ApiClient.convertToType(data['updated_at'], 'String');
      }
      if (data.hasOwnProperty('started_at')) {
        obj['started_at'] = ApiClient.convertToType(data['started_at'], 'String');
      }
      if (data.hasOwnProperty('finished_at')) {
        obj['finished_at'] = ApiClient.convertToType(data['finished_at'], 'String');
      }
      if (data.hasOwnProperty('project')) {
        obj['project'] = ApiClient.convertToType(data['project'], 'String');
      }
      if (data.hasOwnProperty('is_managed')) {
        obj['is_managed'] = ApiClient.convertToType(data['is_managed'], 'String');
      }
      if (data.hasOwnProperty('spec')) {
        obj['spec'] = ApiClient.convertToType(data['spec'], 'String');
      }
      if (data.hasOwnProperty('backend')) {
        obj['backend'] = ApiClient.convertToType(data['backend'], 'String');
      }
      if (data.hasOwnProperty('framework')) {
        obj['framework'] = ApiClient.convertToType(data['framework'], 'String');
      }
      if (data.hasOwnProperty('last_status')) {
        obj['last_status'] = ApiClient.convertToType(data['last_status'], 'String');
      }
      if (data.hasOwnProperty('code_reference')) {
        obj['code_reference'] = ApiClient.convertToType(data['code_reference'], 'String');
      }
      if (data.hasOwnProperty('resources')) {
        obj['resources'] = V1Dict.constructFromObject(data['resources']);
      }
      if (data.hasOwnProperty('readme')) {
        obj['readme'] = ApiClient.convertToType(data['readme'], 'String');
      }
      if (data.hasOwnProperty('bookmarked')) {
        obj['bookmarked'] = ApiClient.convertToType(data['bookmarked'], 'Boolean');
      }
      if (data.hasOwnProperty('params')) {
        obj['params'] = V1Dict.constructFromObject(data['params']);
      }
      if (data.hasOwnProperty('run_env')) {
        obj['run_env'] = V1Dict.constructFromObject(data['run_env']);
      }
      if (data.hasOwnProperty('build_job')) {
        obj['build_job'] = ApiClient.convertToType(data['build_job'], 'String');
      }
      if (data.hasOwnProperty('data_refs')) {
        obj['data_refs'] = V1Dict.constructFromObject(data['data_refs']);
      }
      if (data.hasOwnProperty('artifact_refs')) {
        obj['artifact_refs'] = V1Dict.constructFromObject(data['artifact_refs']);
      }
      if (data.hasOwnProperty('original')) {
        obj['original'] = ApiClient.convertToType(data['original'], 'String');
      }
      if (data.hasOwnProperty('cloning_strategy')) {
        obj['cloning_strategy'] = ApiClient.convertToType(data['cloning_strategy'], 'String');
      }
      if (data.hasOwnProperty('experiment_group')) {
        obj['experiment_group'] = ApiClient.convertToType(data['experiment_group'], 'String');
      }
      if (data.hasOwnProperty('num_jobs')) {
        obj['num_jobs'] = ApiClient.convertToType(data['num_jobs'], 'Number');
      }
      if (data.hasOwnProperty('has_tensorboard')) {
        obj['has_tensorboard'] = ApiClient.convertToType(data['has_tensorboard'], 'Boolean');
      }
      if (data.hasOwnProperty('last_metric')) {
        obj['last_metric'] = V1Dict.constructFromObject(data['last_metric']);
      }
    }
    return obj;
  }

  /**
   * @member {String} id
   */
  exports.prototype['id'] = undefined;
  /**
   * @member {String} uuid
   */
  exports.prototype['uuid'] = undefined;
  /**
   * @member {String} unique_name
   */
  exports.prototype['unique_name'] = undefined;
  /**
   * @member {String} name
   */
  exports.prototype['name'] = undefined;
  /**
   * @member {String} description
   */
  exports.prototype['description'] = undefined;
  /**
   * @member {Array.<String>} tags
   */
  exports.prototype['tags'] = undefined;
  /**
   * @member {Boolean} deleted
   */
  exports.prototype['deleted'] = undefined;
  /**
   * @member {String} user
   */
  exports.prototype['user'] = undefined;
  /**
   * @member {String} created_at
   */
  exports.prototype['created_at'] = undefined;
  /**
   * @member {String} updated_at
   */
  exports.prototype['updated_at'] = undefined;
  /**
   * @member {String} started_at
   */
  exports.prototype['started_at'] = undefined;
  /**
   * @member {String} finished_at
   */
  exports.prototype['finished_at'] = undefined;
  /**
   * @member {String} project
   */
  exports.prototype['project'] = undefined;
  /**
   * @member {String} is_managed
   */
  exports.prototype['is_managed'] = undefined;
  /**
   * @member {String} spec
   */
  exports.prototype['spec'] = undefined;
  /**
   * @member {String} backend
   */
  exports.prototype['backend'] = undefined;
  /**
   * @member {String} framework
   */
  exports.prototype['framework'] = undefined;
  /**
   * @member {String} last_status
   */
  exports.prototype['last_status'] = undefined;
  /**
   * @member {String} code_reference
   */
  exports.prototype['code_reference'] = undefined;
  /**
   * @member {module:model/V1Dict} resources
   */
  exports.prototype['resources'] = undefined;
  /**
   * @member {String} readme
   */
  exports.prototype['readme'] = undefined;
  /**
   * @member {Boolean} bookmarked
   */
  exports.prototype['bookmarked'] = undefined;
  /**
   * @member {module:model/V1Dict} params
   */
  exports.prototype['params'] = undefined;
  /**
   * @member {module:model/V1Dict} run_env
   */
  exports.prototype['run_env'] = undefined;
  /**
   * @member {String} build_job
   */
  exports.prototype['build_job'] = undefined;
  /**
   * @member {module:model/V1Dict} data_refs
   */
  exports.prototype['data_refs'] = undefined;
  /**
   * @member {module:model/V1Dict} artifact_refs
   */
  exports.prototype['artifact_refs'] = undefined;
  /**
   * @member {String} original
   */
  exports.prototype['original'] = undefined;
  /**
   * @member {String} cloning_strategy
   */
  exports.prototype['cloning_strategy'] = undefined;
  /**
   * @member {String} experiment_group
   */
  exports.prototype['experiment_group'] = undefined;
  /**
   * @member {Number} num_jobs
   */
  exports.prototype['num_jobs'] = undefined;
  /**
   * @member {Boolean} has_tensorboard
   */
  exports.prototype['has_tensorboard'] = undefined;
  /**
   * @member {module:model/V1Dict} last_metric
   */
  exports.prototype['last_metric'] = undefined;



  return exports;
}));

