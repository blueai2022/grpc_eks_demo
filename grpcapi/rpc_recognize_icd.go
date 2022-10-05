package grpcapi

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"encoding/json"

	db "github.com/blueai2022/appsubmission/db/sqlc"
	"github.com/blueai2022/appsubmission/http"
	"github.com/blueai2022/appsubmission/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) RecognizeICD10(ctx context.Context, req *pb.RecognizeICD10Request) (*pb.RecognizeICD10Response, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	getAccountParam := db.GetActiveApiAccountParams{
		Username:    authPayload.Username,
		ServiceType: "ICD",
	}

	apiAcct, err := server.store.GetActiveApiAccount(ctx, getAccountParam)
	if err != nil {
		if err == sql.ErrNoRows {
			verr := fmt.Errorf("active api account not found for user %s", authPayload.Username)
			return nil, unauthenticatedError(verr)
		}
		return nil, status.Errorf(codes.Internal, "failed to look up active api accounts: %s", err)
	}

	if apiAcct.CreditBalance < 1 {
		verr := fmt.Errorf("api credits exhausted for user %s", authPayload.Username)
		return nil, unauthenticatedError(verr)
	}

	reqBody := fmt.Sprintf("{\"medical_text\": \"%s\"}", req.MedicalText)
	backendRsp, err := http.PostProxy(
		server.config.ProxyTargetServer,
		"/backend/healthai/icd10",
		reqBody,
		"application/json",
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to call health api: %s", err)
	}

	if backendRsp.StatusCode != 200 && backendRsp.StatusCode != 204 {
		verr := fmt.Errorf("status code %s", backendRsp.Status)
		return nil, status.Errorf(codes.Internal, "received error health api response: %s", verr)
	}

	if backendRsp.StatusCode == 204 {
		server.store.DebitApiAccountBalance(ctx, apiAcct.ID)
		return nil, status.Errorf(codes.NotFound, "medical entity not detected")
	}

	icd := &pb.ICD10{}
	body, err := ioutil.ReadAll(backendRsp.Body)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read health api response: %s", err)
	}

	err = json.Unmarshal(body, icd)
	if err != nil {
		log.Printf("unexpected response body from health api: %s", body)
		return nil, status.Errorf(codes.Internal, "invalid health api response: %s", err)
	}

	server.store.DebitApiAccountBalance(ctx, apiAcct.ID)

	rsp := &pb.RecognizeICD10Response{
		Success: true,
		Result:  icd,
	}

	return rsp, nil
}
