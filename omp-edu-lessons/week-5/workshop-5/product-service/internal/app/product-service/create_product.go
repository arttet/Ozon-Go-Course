package product_service

import (
	"context"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/week-5-workshop/product-service/internal/pkg/logger"
	product_service "github.com/ozonmp/week-5-workshop/product-service/internal/service/product"
	desc "github.com/ozonmp/week-5-workshop/product-service/pkg/product-service"
	"go.uber.org/zap/zapcore"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateProduct(ctx context.Context, req *desc.CreateProductRequest) (*desc.CreateProductResponse, error) {
	attributes := make([]product_service.ProductAttribute, len(req.GetAttributes()))
	for idx, val := range req.GetAttributes() {
		attributes[idx] = convertPbToProductAttributes(val)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		levels := md.Get("log-level")
		logger.InfoKV(ctx, "got log level", "levels", levels)
		if len(levels) > 0 {
			if parsedLevel, ok := parseLevel(levels[0]); ok {
				newLogger := logger.CloneWithLevel(ctx, parsedLevel)
				ctx = logger.AttachLogger(ctx, newLogger)
			}
		}
	}

	logger.DebugKV(ctx, "CreateProduct: some debug message")

	res, err := i.productService.CreateProduct(ctx, req.GetName(), req.GetCategoryId(), attributes)
	if err != nil {
		if err == product_service.ErrWrongCategory {
			details := &errdetails.BadRequest{
				FieldViolations: []*errdetails.BadRequest_FieldViolation{
					{
						Field:       "categoryId",
						Description: "wrong category",
					},
				},
			}

			st := status.New(codes.InvalidArgument, "wrong category")

			withDetails, stErr := st.WithDetails(details)
			if stErr != nil {
				return nil, st.Err()
			}

			return nil, withDetails.Err()
		}

		return nil, err
	}

	go func() {
		time.Sleep(time.Second)

		createProductSpan := opentracing.SpanFromContext(ctx)
		notifySpan := opentracing.StartSpan("notify", opentracing.FollowsFrom(createProductSpan.Context()))
		defer notifySpan.Finish()

		time.Sleep(time.Second)
	}()

	return &desc.CreateProductResponse{
		Result: convertProductToPb(res),
	}, nil
}

func convertProductToPb(res *product_service.Product) *desc.Product {
	resAttr := make([]*desc.ProductAttribute, len(res.Attributes))
	for idx, val := range res.Attributes {
		resAttr[idx] = convertProductAttributesToPb(val)
	}
	return &desc.Product{
		Id:         res.ID,
		Name:       res.Name,
		CategoryId: res.CategoryID,
		Attributes: resAttr,
	}
}

func convertPbToProductAttributes(pa *desc.ProductAttribute) product_service.ProductAttribute {
	return product_service.ProductAttribute{
		ID:    pa.GetId(),
		Value: pa.GetValue(),
	}
}

func convertProductAttributesToPb(pa product_service.ProductAttribute) *desc.ProductAttribute {
	return &desc.ProductAttribute{
		Id:    pa.ID,
		Value: pa.Value,
	}
}

func parseLevel(str string) (zapcore.Level, bool) {
	switch strings.ToLower(str) {
	case "debug":
		return zapcore.DebugLevel, true
	case "info":
		return zapcore.InfoLevel, true
	case "warn":
		return zapcore.WarnLevel, true
	case "error":
		return zapcore.ErrorLevel, true
	default:
		return zapcore.DebugLevel, false
	}
}
