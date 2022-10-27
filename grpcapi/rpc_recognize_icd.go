package grpcapi

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/blueai2022/appsubmission/db/sqlc"
	"github.com/blueai2022/appsubmission/healthapi"
	"github.com/blueai2022/appsubmission/pb"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
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

	icd := &pb.ICD10{}
	err = healthapi.ICD10(server.config, req.MedicalText, icd)
	if err != nil {
		if err == healthapi.ErrMedicalEntityNotFound {
			server.store.DebitApiAccountBalance(ctx, apiAcct.ID)
			return nil, status.Errorf(codes.NotFound, "medical entity not detected")
		}
		return nil, status.Errorf(codes.Internal, "failed to call health api: %s", err)
	}

	server.store.DebitApiAccountBalance(ctx, apiAcct.ID)

	rsp := &pb.RecognizeICD10Response{
		Success: true,
		Result:  icd,
	}
	//Field masking:
	//Not Shown: reduce DB/backend calls according to field masks
	//Shown below: apply mask to reduce response payload
	// Only the fields mentioned in the field mask will be copied to userDst, other fields are left intact
	rspDst := &pb.RecognizeICD10Response{} // a struct to copy to
	mask, _ := fieldmask_utils.MaskFromPaths(req.FieldMask.Paths, naming)
	fieldmask_utils.StructToStruct(mask, rsp, rspDst)

	return rspDst, nil
}
