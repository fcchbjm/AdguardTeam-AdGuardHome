//go:build darwin || freebsd || openbsd

package aghnet

import (
	"context"
	"log/slog"

	"github.com/fcchbjm/AdGuardHome/internal/aghos"
)

func canBindPrivilegedPorts(_ context.Context, _ *slog.Logger) (can bool, err error) {
	return aghos.HaveAdminRights()
}
