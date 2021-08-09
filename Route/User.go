package Route

import (
	"errors"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)
import "../Repo"
import "../Entity"

var getAuthorizationFiller = func(ctx *gin.Context) []func(c jwt.MapClaims) error {
	return []func(c jwt.MapClaims) error{
		func(c jwt.MapClaims) error {
			if c["sub"].(string) != ctx.Param("uuid") {
				return errors.New("not equals uuid")
			}
			return nil
		},
	}
}

func GetAllUsersProfileInGroupHandle(repo *Repo.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		groupID := ctx.Query("group")

		group, err := repo.GetGroupRepo().Get(groupID)

		if err != nil {
			fmt.Printf(fmt.Errorf("GetAllUsersProfileInGroupHandle() occurred error - %s", err.Error()).Error())
			ctx.JSONP(http.StatusInternalServerError, gin.H{
				"error": "internal_server_error",
			})
			return
		}
		ctx.JSONP(http.StatusOK, group.Users)
	}
}

func GetUserProfileHandle(repo *Repo.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := ctx.Param("uuid")
		if uuid == "" {
			ctx.JSONP(http.StatusBadRequest, gin.H{
				"error": "invalid_profile_uuid",
			})
			return
		}

		profile, err := repo.GetUserRepository().FindByID(uuid)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSONP(http.StatusNotFound, gin.H{
					"error": "not_found",
				})
			} else {
				ctx.JSONP(http.StatusInternalServerError, gin.H{
					"error": "internal_server_error",
				})
			}
			return
		}

		ctx.JSONP(http.StatusOK, profile)
	}
}

func ConvertUpdateRequestToUser(before *Entity.User, ctx *gin.Context) (*Entity.User, bool) {
	updateSource := map[string]string{}
	ctx.BindJSON(&updateSource)

	changed := false

	if picture, exist := updateSource["picture"]; exist {
		before.Picture = picture
		changed = true
	}

	if privateName, exist := updateSource["private_name"]; exist {
		before.RealName = privateName
		changed = true
	}

	if nick, exist := updateSource["nickname"]; exist {
		before.NickName = nick
		changed = true
	}

	return before, changed
}

func UpdateUserProfileHandle(repo *Repo.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		AuthFilter(getAuthorizationFiller(ctx), func() {

			before, errRepo := repo.GetUserRepository().FindByID(ctx.Param("uuid"))

			if errRepo != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "internal_server_error",
				})
			}

			updatedProfile, changed := ConvertUpdateRequestToUser(before, ctx)

			if !changed {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "ok"})
			}

			err := repo.GetUserRepository().Update(updatedProfile)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
						"error": "not_found",
					})
				} else {
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": "internal_server_error",
					})
				}
				return
			}
			ctx.JSONP(http.StatusOK, gin.H{
				"status": "successful",
			})
		})
	}
}

func ConvertAddRequestToUser(addSource map[string]interface{}) (*Entity.User, error) {
	picture, exist := addSource["picture"]
	if !exist {
		return nil, errors.New("need params \"picture\"")
	}

	mail, exist := addSource["mail"]
	if !exist {
		return nil, errors.New("need params \"mail\"")
	}

	nick, exist := addSource["nickname"]
	if !exist {
		return nil, errors.New("need params \"nick\"")
	}

	createdAt, exist := addSource["created_at"]
	if !exist {
		return nil, errors.New("need params \"created_at\"")
	}

	uid, exist := addSource["sub"]
	if !exist {
		return nil, errors.New("need params \"sub\"")
	}

	newProfile := Entity.User{
		IsCompleted: false,
		CreatedAt:   createdAt.(string),
		Picture:     picture.(string),
		NickName:    nick.(string),
		RealName:    "",
		Mail:        mail.(string),
		ID:          uid.(string),
	}
	return &newProfile, nil
}

func AddUserHandle(repo *Repo.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		newProfile, convertErr := ConvertAddRequestToUser(ExtractClaims(ctx))

		if convertErr != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": convertErr.Error(),
			})
			return
		}

		repoErr := repo.GetUserRepository().Add(newProfile)
		if repoErr != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal_server_error",
			})
			return
		}
		ctx.JSONP(http.StatusOK, gin.H{
			"status": "successful",
		})
	}
}

func DeleteUserHandle(repo *Repo.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		AuthFilter(getAuthorizationFiller(ctx), func() {
			uuid := ctx.Param("uuid")

			err := repo.GetUserRepository().Delete(uuid)
			if err != nil {
				ctx.JSONP(http.StatusInternalServerError, gin.H{
					"error": "internal_server_error",
				})
				return
			}
			ctx.JSONP(http.StatusOK, gin.H{
				"status": "successful",
			})
		})
	}
}
