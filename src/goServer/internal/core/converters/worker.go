package converters

import (
	"goServer/internal/core/pb"

	"github.com/relvacode/iso8601"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertMapToWorkerProto(mp map[string]any) (*pb.Worker, []error) {
	var errors []error
	var result = &pb.Worker{}

	if mapId := mp["id"]; mapId != nil {
		result.Id = int64(mapId.(float64))
	}

	if mapUserId := mp["userId"]; mapUserId != nil {
		userId := int64(mapUserId.(float64))
		result.UserId = &userId
	}

	if mapCreatedById := mp["createdById"]; mapCreatedById != nil {
		createdById := int64(mapCreatedById.(float64))
		result.CreatedById = &createdById
	}

	if mapUpdatedById := mp["updatedById"]; mapUpdatedById != nil {
		updatedById := int64(mapUpdatedById.(float64))
		result.UpdatedById = &updatedById
	}

	if mapCreatedAt := mp["createdAt"]; mapCreatedAt != nil {
		createdAt, parseError := iso8601.ParseString(mapCreatedAt.(string))

		if parseError != nil {
			panic(parseError)
		}

		result.CreatedAt = timestamppb.New(createdAt)
	}

	if mapLastUpdatedAt := mp["lastUpdatedAt"]; mapLastUpdatedAt != nil {
		createdAt, parseError := iso8601.ParseString(mapLastUpdatedAt.(string))

		if parseError != nil {
			panic(parseError)
		}

		result.CreatedAt = timestamppb.New(createdAt)
	}

	if mapPictureUrl := mp["pictureUrl"]; mapPictureUrl != nil {
		pictureUrl := mapPictureUrl.(string)
		result.PictureUrl = &pictureUrl
	}

	if mapFirstName := mp["firstName"]; mapFirstName != nil {
		result.FirstName = mapFirstName.(string)
	}

	if mapLastName := mp["lastName"]; mapLastName != nil {
		result.LastName = mapLastName.(string)
	}

	if mapMiddleName := mp["middleName"]; mapMiddleName != nil {
		middleName := mapMiddleName.(string)
		result.MiddleName = &middleName
	}

	if mapBirthdate := mp["birthdate"]; mapBirthdate != nil {
		birthdate, parseError := iso8601.ParseString(mapBirthdate.(string))

		if parseError != nil {
			panic(parseError)
		}

		result.Birthdate = timestamppb.New(birthdate)
	}

	if mapUsername := mp["username"]; mapUsername != nil {
		result.Username = mapUsername.(string)
	}

	if mapEmail := mp["email"]; mapEmail != nil {
		result.Email = mapEmail.(string)
	}

	if mapMobile := mp["mobile"]; mapMobile != nil {
		mobile := mapMobile.(string)
		result.Mobile = &mobile
	}

	if mapAddresses := mp["addresses"]; mapAddresses != nil {
		addresses := mapAddresses.([]any)
		var structValues []*structpb.Struct

		for _, val := range addresses {
			address := val.(map[string]any)
			structVal, parseError := structpb.NewStruct(address)

			if parseError != nil {
				panic(parseError)
			}

			structValues = append(structValues, structVal)
		}

		result.Addresses = structValues
	}

	if mapUuid := mp["uuid"]; mapUuid != nil {
		result.Uuid = mapUuid.(string)
	}

	result.Flags = uint32(mp["flags"].(float64))

	return result, errors
}
