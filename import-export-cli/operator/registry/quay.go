/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package registry

// HttpRegistry represents private HTTPS registry
// copy of HttpsRegistry
var QuayRegistry = *HttpsRegistry

func init() {
	QuayRegistry.Name = "QUAY"
	QuayRegistry.Caption = "Quay.io"
	QuayRegistry.Option = 6
	QuayRegistry.Repository.ServerUrl = "quay.io/"

	add(&QuayRegistry)
}
