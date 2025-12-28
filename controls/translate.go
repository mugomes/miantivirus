// Copyright (C) 2024-2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.mugomes.com.br

package controls

import "github.com/mugomes/mglang"

func LoadTranslations() {
	lang := mglang.GetLang()
	if lang == "pt" {
		
	}
}

func T(key string, args ...interface{}) string {
	return mglang.T(key, args...)
}
