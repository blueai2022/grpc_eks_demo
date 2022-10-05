package grpcapi

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"encoding/json"

	"github.com/blueai2022/appsubmission/http"
	"github.com/blueai2022/appsubmission/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) RecognizeICD10(ctx context.Context, req *pb.RecognizeICD10Request) (*pb.RecognizeICD10Response, error) {
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

	if backendRsp.StatusCode == 204 {
		// rsp := &pb.RecognizeICD10Response{
		// 	Success: false,
		// 	Result:  &pb.ICD10{},
		// }

		return nil, status.Errorf(codes.NotFound, "medical entity not detected")
	}

	if backendRsp.StatusCode != 200 {
		verr := fmt.Errorf("status code %s", backendRsp.Status)
		return nil, status.Errorf(codes.Internal, "received error health api response: %s", verr)
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

	rsp := &pb.RecognizeICD10Response{
		Success: true,
		Result:  icd,
	}

	return rsp, nil
}
