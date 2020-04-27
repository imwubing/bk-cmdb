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

package service

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/http/rest"
	"configcenter/src/common/metadata"
)

func (s *Service) SelectObjectTopoGraphics(ctx *rest.Contexts) {
	md := new(MetaShell)
	if err := ctx.DecodeInto(md); err != nil {
		ctx.RespAutoError(err)
		return
	}

	resp, err := s.Core.GraphicsOperation().SelectObjectTopoGraphics(ctx.Kit, ctx.Request.PathParameter("scope_type"), ctx.Request.PathParameter("scope_id"), md.Metadata)
	if err != nil {
		ctx.RespAutoError(err)
		return
	}
	ctx.RespEntity(resp)
}

func (s *Service) UpdateObjectTopoGraphics(ctx *rest.Contexts) {
	requestBody := struct {
		Data []metadata.TopoGraphics `json:"data" field:"data"`
	}{}
	if err := ctx.DecodeInto(&requestBody); err != nil {
		ctx.RespAutoError(err)
		return
	}

	err := s.Engine.CoreAPI.CoreService().Txn().AutoRunTxn(ctx.Kit.Ctx, s.EnableTxn, ctx.Kit.Header, func() error {
		err := s.Core.GraphicsOperation().UpdateObjectTopoGraphics(ctx.Kit, ctx.Request.PathParameter("scope_type"), ctx.Request.PathParameter("scope_id"), requestBody.Data, nil)
		if err != nil {
			ctx.RespAutoError(err)
			return err
		}

		ctx.RespEntity(nil)
		return nil
	})

	if err != nil {
		blog.Errorf("UpdateObjectTopoGraphics failed, err: %v, rid: %s", err, ctx.Kit.Rid)
		return
	}
}

func (s *Service) UpdateObjectTopoGraphicsNew(ctx *rest.Contexts) {
	input := metadata.UpdateTopoGraphicsInput{}
	err := ctx.DecodeInto(&input)
	if nil != err {
		ctx.RespAutoError(ctx.Kit.CCError.New(common.CCErrCommParamsIsInvalid, "not set anything"))
		return
	}

	err = s.Engine.CoreAPI.CoreService().Txn().AutoRunTxn(ctx.Kit.Ctx, s.EnableTxn, ctx.Kit.Header, func() error {
		err := s.Core.GraphicsOperation().UpdateObjectTopoGraphics(ctx.Kit, ctx.Request.PathParameter("scope_type"), ctx.Request.PathParameter("scope_id"), input.Origin, input.Metadata)
		if err != nil {
			ctx.RespAutoError(err)
			return err
		}

		ctx.RespEntity(nil)
		return nil
	})

	if err != nil {
		blog.Errorf("UpdateObjectTopoGraphicsNew failed, err: %v, rid: %s", err, ctx.Kit.Rid)
		return
	}
}
