// Copyright 2023 The Ryan SU Authors. All Rights Reserved.
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

package vben

import (
	_ "embed"
)

var (
	//go:embed api.tpl
	apiTpl string

	//go:embed model.tpl
	modelTpl string

	//go:embed data.tpl
	dataTpl string

	//go:embed drawer.tpl
	drawerTpl string

	//go:embed index.tpl
	indexTpl string

	//go:embed locale.tpl
	localeTpl string

	//go:embed statusrender.tpl
	statusRenderTpl string
)
