package processors

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/grpc/services/data"
	"github.com/tinyci/ci-agents/grpc/types"
)

// GetErrors retrieves the errors for the provided user.
func (ds *DataServer) GetErrors(ctx context.Context, name *data.Name) (*types.UserErrors, error) {
	u, err := ds.H.Model.FindUserByName(name.Name)
	if err != nil {
		return nil, err
	}

	errors := &types.UserErrors{}
	for _, err := range u.Errors {
		errors.Errors = append(errors.Errors, err.ToProto())
	}

	return errors, nil
}

// AddError adds an error to the user's error stack.
func (ds *DataServer) AddError(ctx context.Context, ue *types.UserError) (*empty.Empty, error) {
	u, err := ds.H.Model.FindUserByID(ue.UserID)
	if err != nil {
		return nil, err
	}

	u.AddError(errors.New(ue.Error))

	if err := ds.H.Model.Save(u).Error; err != nil {
		return nil, errors.New(err)
	}

	return &empty.Empty{}, nil
}

// DeleteError removes an error from errors list. The error string does not need to be provided.
func (ds *DataServer) DeleteError(ctx context.Context, ue *types.UserError) (*empty.Empty, error) {
	u, eErr := ds.H.Model.FindUserByID(ue.UserID)
	if eErr != nil {
		return nil, eErr
	}

	if err := ds.H.Model.DeleteError(u, ue.Id); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
