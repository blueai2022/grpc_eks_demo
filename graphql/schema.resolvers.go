package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
)

// Sessions is the resolver for the sessions field.
func (r *queryResolver) Sessions(ctx context.Context, userID string) ([]*Session, error) {
	sessions, err := r.Store.GetUserSessions(ctx, userID)
	if err != nil {
		return nil, err
	}

	var res []*Session
	for _, s := range sessions {
		res = append(res, &Session{
			ID:           s.ID.String(),
			Username:     s.Username,
			RefreshToken: s.RefreshToken,
			ClientIp:     s.ClientIp,
			UserAgent:    s.UserAgent,
			ExpiresAt:    s.ExpiresAt,
		})
	}

	return res, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
