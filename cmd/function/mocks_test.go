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

package function_test

import (
	"context"

	"github.com/toughtackle/slack-cli/internal/logger"
	"github.com/toughtackle/slack-cli/internal/shared"
	"github.com/toughtackle/slack-cli/internal/shared/types"
	"github.com/stretchr/testify/mock"
)

type FunctionDistributorMock struct {
	mock.Mock
}

func (m *FunctionDistributorMock) List(ctx context.Context, clients *shared.ClientFactory, fn string, log *logger.Logger) (types.Permission, []types.FunctionDistributionUser, error) {
	args := m.Called()
	return args.Get(0).(types.Permission), args.Get(1).([]types.FunctionDistributionUser), args.Error(2)
}

func (m *FunctionDistributorMock) Set(ctx context.Context, clients *shared.ClientFactory, function, distributionType, users string, log *logger.Logger) (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *FunctionDistributorMock) AddUsers(ctx context.Context, clients *shared.ClientFactory, function, users string, log *logger.Logger) (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *FunctionDistributorMock) RemoveUsers(ctx context.Context, clients *shared.ClientFactory, function, users string, log *logger.Logger) (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}
