/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */
package user

import (
	"context"
	"net/http"

	"configcenter/src/common/metadata"
	"configcenter/src/thirdpartyclient/esbserver/esbutil"
)

// Deprecated
func (p *user) GetAllUsers(ctx context.Context, h http.Header) (resp *metadata.EsbUserListResponse, err error) {
	resp = &metadata.EsbUserListResponse{}
	subPath := "/v2/usermanage/get_all_users/"
	h.Set("Accept", "application/json")

	err = p.client.Get().
		WithContext(ctx).
		SubResource(subPath).
		WithParams(esbutil.GetEsbQueryParameters(p.config.GetConfig(), h)).
		WithHeaders(h).
		Do().
		Into(resp)

	return
}

func (p *user) ListUsers(ctx context.Context, h http.Header) (resp *metadata.EsbListUserResponse, err error) {
	// response demo
	/*
		{
		  "status": "NORMAL",
		  "domain": "default.local",
		  "code": null,
		  "telephone": "11111111111",
		  "create_time": "2019-12-11T19:45:21.518283Z",
		  "country_code": "86",
		  "iso_code": "CN",
		  "logo": null,
		  "id": 22222,
		  "display_name": "foo",
		  "role": 0,
		  "type": "",
		  "leader": [],
		  "username": "foo",
		  "update_time": "2019-12-11T19:50:07.903422Z",
		  "wx_userid": "",
		  "staff_status": "IN",
		  "position": 0,
		  "qq": "",
		  "language": "zh-cn",
		  "enabled": true,
		  "time_zone": "Asia/Shanghai",
		  "departments": [
			{
			  "order": 1,
			  "id": 22222,
			  "full_name": "总公司",
			  "name": "总公司"
			}
		  ],
		  "email": "foo@qq.com",
		  "extras": [],
		  "wx_openid": "",
		  "password_valid_days": 180,
		  "category_id": 1
		}
	*/
	resp = &metadata.EsbListUserResponse{}
	subPath := "/v2/usermanage/list_users/"
	h.Set("Accept", "application/json")

	err = p.client.Get().
		WithContext(ctx).
		WithParam("fields", "username,id").
		WithParam("no_page", "true").
		SubResource(subPath).
		WithParams(esbutil.GetEsbQueryParameters(p.config.GetConfig(), h)).
		WithHeaders(h).
		Do().
		Into(resp)

	return
}
