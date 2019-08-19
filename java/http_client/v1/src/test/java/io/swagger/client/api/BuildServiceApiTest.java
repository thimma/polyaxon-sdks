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

/*
 * Polyaxon sdk
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.14.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package io.swagger.client.api;

import io.swagger.client.ApiException;
import io.swagger.client.model.V1Build;
import io.swagger.client.model.V1BuildBodyRequest;
import io.swagger.client.model.V1BuildStatus;
import io.swagger.client.model.V1CodeReference;
import io.swagger.client.model.V1CodeReferenceBodyRequest;
import io.swagger.client.model.V1ListBuildStatusesResponse;
import io.swagger.client.model.V1ListBuildsResponse;
import io.swagger.client.model.V1OwnedEntityIdRequest;
import io.swagger.client.model.V1ProjectBodyRequest;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for BuildServiceApi
 */
@Ignore
public class BuildServiceApiTest {

    private final BuildServiceApi api = new BuildServiceApi();

    
    /**
     * Archive build
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void archiveBuildTest() throws ApiException {
        String owner = null;
        String project = null;
        String id = null;
        Object response = api.archiveBuild(owner, project, id);

        // TODO: test validations
    }
    
    /**
     * Bookmark build
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void bookmarkBuildTest() throws ApiException {
        String owner = null;
        String project = null;
        String id = null;
        Object response = api.bookmarkBuild(owner, project, id);

        // TODO: test validations
    }
    
    /**
     * Create new build
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createBuildTest() throws ApiException {
        String owner = null;
        String project = null;
        V1BuildBodyRequest body = null;
        V1Build response = api.createBuild(owner, project, body);

        // TODO: test validations
    }
    
    /**
     * Create build code ref
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createBuildCodeRefTest() throws ApiException {
        String entityOwner = null;
        String entityProject = null;
        String entityId = null;
        V1CodeReferenceBodyRequest body = null;
        V1CodeReference response = api.createBuildCodeRef(entityOwner, entityProject, entityId, body);

        // TODO: test validations
    }
    
    /**
     * Create new build status
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createBuildStatusTest() throws ApiException {
        String owner = null;
        String project = null;
        String id = null;
        V1OwnedEntityIdRequest body = null;
        V1BuildStatus response = api.createBuildStatus(owner, project, id, body);

        // TODO: test validations
    }
    
    /**
     * Delete build
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteBuildTest() throws ApiException {
        String owner = null;
        String project = null;
        String id = null;
        Object response = api.deleteBuild(owner, project, id);

        // TODO: test validations
    }
    
    /**
     * Delete builds
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteBuildsTest() throws ApiException {
        String owner = null;
        String project = null;
        V1OwnedEntityIdRequest body = null;
        Object response = api.deleteBuilds(owner, project, body);

        // TODO: test validations
    }
    
    /**
     * Get build
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getBuildTest() throws ApiException {
        String owner = null;
        String project = null;
        String id = null;
        V1Build response = api.getBuild(owner, project, id);

        // TODO: test validations
    }
    
    /**
     * Get build code ref
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getBuildCodeRefTest() throws ApiException {
        String owner = null;
        String project = null;
        String id = null;
        V1CodeReference response = api.getBuildCodeRef(owner, project, id);

        // TODO: test validations
    }
    
    /**
     * List archived builds
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listArchivedBuildsTest() throws ApiException {
        String owner = null;
        V1ListBuildsResponse response = api.listArchivedBuilds(owner);

        // TODO: test validations
    }
    
    /**
     * List bookmarked builds
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listBookmarkedBuildsTest() throws ApiException {
        String owner = null;
        V1ListBuildsResponse response = api.listBookmarkedBuilds(owner);

        // TODO: test validations
    }
    
    /**
     * List build statuses
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listBuildStatusesTest() throws ApiException {
        String owner = null;
        String project = null;
        String id = null;
        V1ListBuildStatusesResponse response = api.listBuildStatuses(owner, project, id);

        // TODO: test validations
    }
    
    /**
     * List builds
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listBuildsTest() throws ApiException {
        String owner = null;
        String project = null;
        V1ListBuildsResponse response = api.listBuilds(owner, project);

        // TODO: test validations
    }
    
    /**
     * Restart build
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void restartBuildTest() throws ApiException {
        String owner = null;
        String project = null;
        String id = null;
        V1OwnedEntityIdRequest body = null;
        V1Build response = api.restartBuild(owner, project, id, body);

        // TODO: test validations
    }
    
    /**
     * Restore build
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void restoreBuildTest() throws ApiException {
        String owner = null;
        String project = null;
        String id = null;
        Object response = api.restoreBuild(owner, project, id);

        // TODO: test validations
    }
    
    /**
     * Stop build
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void stopBuildTest() throws ApiException {
        String owner = null;
        String project = null;
        String id = null;
        V1OwnedEntityIdRequest body = null;
        Object response = api.stopBuild(owner, project, id, body);

        // TODO: test validations
    }
    
    /**
     * Stop builds
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void stopBuildsTest() throws ApiException {
        String owner = null;
        String project = null;
        V1ProjectBodyRequest body = null;
        Object response = api.stopBuilds(owner, project, body);

        // TODO: test validations
    }
    
    /**
     * UnBookmark build
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void unBookmarkBuildTest() throws ApiException {
        String owner = null;
        String project = null;
        String id = null;
        Object response = api.unBookmarkBuild(owner, project, id);

        // TODO: test validations
    }
    
    /**
     * Update build
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateBuild2Test() throws ApiException {
        String owner = null;
        String project = null;
        String buildId = null;
        V1BuildBodyRequest body = null;
        V1Build response = api.updateBuild2(owner, project, buildId, body);

        // TODO: test validations
    }
    
}