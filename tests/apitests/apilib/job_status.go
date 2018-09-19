/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * OpenAPI spec version: 0.3.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package apilib

type JobStatus struct {

	// The job ID.
	Id int64 `json:"id,omitempty"`

	// The status of the job.
	Status string `json:"status,omitempty"`

	// The repository handled by the job.
	Repository string `json:"repository,omitempty"`

	// The ID of the policy that triggered this job.
	PolicyId int64 `json:"policy_id,omitempty"`

	// The operation of the job.
	Operation string `json:"operation,omitempty"`

	// The repository's used tag list.
	Tags []Tags `json:"tags,omitempty"`

	// The creation time of the job.
	CreationTime string `json:"creation_time,omitempty"`

	// The update time of the job.
	UpdateTime string `json:"update_time,omitempty"`
}
