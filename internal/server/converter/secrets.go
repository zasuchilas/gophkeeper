package converter

import (
	"github.com/zasuchilas/gophkeeper/internal/server/api/helper"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	desc "github.com/zasuchilas/gophkeeper/pkg/secretsv1"
)

func ToSecretFromService(item *model.Secret) *desc.Secret {
	updatedAt := helper.TimeToProto(item.UpdatedAt)
	createdAt := helper.TimeToProto(item.CreatedAt)

	var secretType desc.SecretType
	v, ok := desc.SecretType_value[item.SecretType]
	if !ok {
		secretType = desc.SecretType_UNKNOWN
	} else {
		secretType = desc.SecretType(v)
	}

	return &desc.Secret{
		Id:         item.ID,
		Name:       item.Name,
		Data:       item.Data,
		Size:       item.Size,
		SecretType: secretType,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		UserId:     item.UserID,
	}
}

func ToSecretListFromService(items []*model.Secret) []*desc.Secret {
	if len(items) == 0 {
		return nil
	}

	result := make([]*desc.Secret, len(items))
	for i, item := range items {
		result[i] = ToSecretFromService(item)
	}
	return result
}

func ToSecretFromApi(in *desc.Secret) *model.Secret {

	createdAt, _ := helper.ProtoToTime("created_api", in.CreatedAt)
	updatedAt, _ := helper.ProtoToTime("updated_api", in.UpdatedAt)

	return &model.Secret{
		ID:         in.Id,
		Name:       in.Name,
		Data:       in.Data,
		Size:       in.Size,
		SecretType: in.SecretType.String(),
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		UserID:     in.UserId,
	}
}
