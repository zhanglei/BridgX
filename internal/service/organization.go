package service

import (
	"context"
	"errors"
	"time"

	"github.com/galaxy-future/BridgX/internal/constants"
	"github.com/galaxy-future/BridgX/internal/model"
	"github.com/galaxy-future/BridgX/pkg/utils"
)

func GetOrgList(ctx context.Context) (ret []model.Org, err error) {
	queryMap := map[string]interface{}{}

	err = model.QueryAll(queryMap, &ret, "id DESC")
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func GetOrgInfoById(ctx context.Context, id int64) (*model.Org, error) {
	ret := model.Org{}
	err := model.Get(id, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func CreateOrg(ctx context.Context, orgName, username, password, createBy string) error {
	now := time.Now()

	org := &model.Org{
		OrgName: orgName,
	}
	org.CreateAt = &now
	org.UpdateAt = &now
	err := model.Create(org)
	if err != nil {
		return err
	}

	user := &model.User{
		Username:   username,
		Password:   utils.Base64Md5(password),
		OrgId:      org.Id,
		UserStatus: constants.UserStatusEnable,
		UserType:   constants.UserTypeAdmin,
		CreateBy:   createBy,
	}
	user.CreateAt = &now
	user.UpdateAt = &now

	return model.Create(user)
}

func EditOrg(ctx context.Context, orgId int64, orgName string) error {
	org, err := GetOrgInfoById(ctx, orgId)
	if err != nil || org == nil {
		return errors.New("org not found")
	}
	org.OrgName = orgName
	now := time.Now()
	org.UpdateAt = &now
	return model.Save(org)
}
