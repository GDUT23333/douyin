package vo

/**
 * @Author: Ember
 * @Date: 2022/5/10 15:02
 * @Description: TODO
 **/
type VideoVo struct{
	//主键
	ID int64 `json:"id"`
	//作者信息
	Author UserVo `json:"author"`
	//播放地址
	PlayUrl string `json:"play_url"`
	//封面地址
	CoverUrl string `json:"cover_url"`
	//点赞数
	FavoriteCount int64 `json:"favorite_count"`
	//评论数
	CommentCount int64 `json:"comment_count"`
	//是否点赞
	IsFavorite bool `json:"is_favorite"`
}