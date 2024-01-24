package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// User routes
	rt.router.POST("/session", rt.wrap(rt.logUser))
	rt.router.PUT("/username/:username", rt.wrap(rt.setMyUserName))
	rt.router.GET("/profile/:pUsername", rt.wrap(rt.getUserProfile))
	rt.router.GET("/stream", rt.wrap(rt.getMyStream))
	rt.router.GET("/username/:username/checkUser", rt.wrap(rt.checkUser))

	// Photo routes
	rt.router.POST("/photo", rt.wrap(rt.uploadPhoto))
	rt.router.GET("/photouser/:photoUsername", rt.wrap(rt.getUserPhotos))
	rt.router.DELETE("/photo/:photoID", rt.wrap(rt.deletePhoto))

	// Like routes
	rt.router.POST("/photo/:photoID/like", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/photo/:photoID/like", rt.wrap(rt.unlikePhoto))

	// Comment routes
	rt.router.POST("/photo/:photoID", rt.wrap(rt.commentPhoto))
	rt.router.GET("/photo/:photoID", rt.wrap(rt.getComments))
	rt.router.DELETE("/comment/:commentID", rt.wrap(rt.uncommentPhoto))

	// Follow routes
	rt.router.POST("/follow/:fUsername", rt.wrap(rt.followUser))
	rt.router.DELETE("/follow/:fUsername", rt.wrap(rt.unfollowUser))
	rt.router.GET("/followers", rt.wrap(rt.getUserFollowers))
	rt.router.GET("/followings", rt.wrap(rt.getUserFollowings))
	rt.router.GET("/follow/:fUsername", rt.wrap(rt.ifFollowing))

	// Ban routes
	rt.router.POST("/ban/:bUsername", rt.wrap(rt.banUser))
	rt.router.DELETE("/ban/:bUsername", rt.wrap(rt.unbanUser))
	rt.router.GET("/username/:username/ban/:bUsername", rt.wrap(rt.checkIfBan))

	return rt.router
}