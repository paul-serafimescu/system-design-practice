package models

import (
	"context"
	"http-server/database"
	"time"

	"github.com/rs/zerolog/log"
)

type Channel struct {
	ID            string
	Name          string
	Description   string
	CreatedAt     time.Time
	LastUpdatedAt time.Time
	Position      int
	LastMessageId string
	communityId   string
}

func (c *Channel) GetCommunity() *Community {
	var community Community

	sql := "SELECT community_id, community_name, community_description, owner_id, created_at, updated_at, region, icon_url FROM community WHERE community_id = $1"
	err := database.Get().QueryRow(context.Background(), sql, c.communityId).Scan(
		&community.ID,
		&community.Name,
		&community.Description,
		&community.ownerId,
		&community.CreatedAt,
		&community.LastUpdatedAt,
		&community.Region,
		&community.IconUrl,
	)

	if err != nil {
		log.Error().Msgf("%s", err.Error())
		return nil
	}

	return &community
}
