// Copyright 2022-2025 Salesforce, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datastore

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/toughtackle/slack-cli/internal/config"
	"github.com/toughtackle/slack-cli/internal/logger"
	"github.com/toughtackle/slack-cli/internal/shared"
	"github.com/toughtackle/slack-cli/internal/shared/types"
)

func Put(ctx context.Context, clients *shared.ClientFactory, log *logger.Logger, request types.AppDatastorePut) (*logger.LogEvent, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "pkg.datastore.put")
	defer span.Finish()

	// Get auth token
	var token = config.GetContextToken(ctx)

	putResult, err := clients.APIInterface().AppsDatastorePut(ctx, token, request)
	if err != nil {
		return nil, err
	}

	// Notify listeners
	log.Data["putResult"] = putResult
	log.Log("info", "on_put_result")

	return log.SuccessEvent(), nil
}
