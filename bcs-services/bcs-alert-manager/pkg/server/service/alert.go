

package service

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/pkg/proto/alertmanager"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/pkg/remote/metrics"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/pkg/server/utils"
)

const (
	// GrpcSchema grpc
	GrpcSchema = "grpc"
)

// CreateRawAlertInfo create raw alert info for httpLayer
func (am *AlertManager) CreateRawAlertInfo(ctx context.Context,
	req *alertmanager.CreateRawAlertInfoReq, resp *alertmanager.CreateRawAlertInfoResp) error {
	const (
		traceName = "CreateRawAlertInfo"
	)

	start := time.Now()

	headerInfo, err := utils.GetRequestHeaderInfo(ctx)
	if err != nil {
		blog.Errorf("get traceID failed from context: %v", err)
		return err
	}
	ctx, tracer := utils.WithTraceForContext(ctx, traceName, headerInfo.TraceID)
	tracer.DefaultRequestInEvent(headerInfo.ClientIP, GrpcSchema, http.MethodPost)
	defer tracer.DefaultRequestOutEvent()

	// console action
	am.consoleAction.CreateRawAlertInfo(ctx, req, resp)
	// metricsInfo
	metrics.ReportAlertManagerAPIMetrics("CreateRawAlertInfo", GrpcSchema, strconv.Itoa(int(resp.ErrCode)), start)
	tracer.Infof("req %v, resp %v", req, resp)

	return nil
}

// CreateBusinessAlertInfo create business alertInfo for httpLayer
func (am *AlertManager) CreateBusinessAlertInfo(ctx context.Context,
	req *alertmanager.CreateBusinessAlertInfoReq, resp *alertmanager.CreateBusinessAlertInfoResp) error {
	const (
		traceName = "CreateBusinessAlertInfo"
	)

	start := time.Now()

	headerInfo, err := utils.GetRequestHeaderInfo(ctx)
	if err != nil {
		blog.Errorf("get traceID failed from context: %v", err)
		return err
	}
	ctx, tracer := utils.WithTraceForContext(ctx, traceName, headerInfo.TraceID)
	tracer.DefaultRequestInEvent(headerInfo.ClientIP, GrpcSchema, http.MethodPost)
	defer tracer.DefaultRequestOutEvent()

	// console action
	am.consoleAction.CreateBusinessAlertInfo(ctx, req, resp)
	// metricsInfo
	metrics.ReportAlertManagerAPIMetrics("CreateBusinessAlertInfo", GrpcSchema, strconv.Itoa(int(resp.ErrCode)), start)
	tracer.Infof("req %v, resp %v", req, resp)

	return nil
}
