package globle

type (
	ChatHistory struct {
		Id         int64  `json:"id" gorm:"primaryKey;autoIncrement"` // bigint(20) unsigned NOT NULL AUTO_INCREMENT,
		Uid        int    `json:"uid"`                                //int(11) DEFAULT NULL COMMENT '用户id',
		RoomId     int    `json:"room_id"`                            //varchar(100) DEFAULT NULL COMMENT '房间id',
		CreateTime int64  `json:"create_time"`                        //datetime DEFAULT NULL,
		Content    string `json:"content"`
	}

	SendGiftReq struct {
		Anchorid int   `json:"anchorid" validate:"required"`
		Giftid   int   `json:"giftid" validate:"required"`
		Liveid   int64 `json:"liveid"`
		Count    int   `json:"count"`
		Uid      int
	}

	RoomScoreLiveMessage struct {
		SentMessageStruct
		ScoreLive RoomScoreLive `json:"score_live"`
	}

	RoomScoreLive struct {
		Id          int   `json:"match_id"`
		Status      int   `json:"match_status"`
		HomeScore   Score `json:"home_score"`
		AwayScore   Score `json:"away_score"`
		KickOutTime int64 `json:"kick_out_time"`
	}

	Score struct {
		Score        int `json:"score"`
		HalfScore    int `json:"half_score"`
		RedCard      int `json:"red_card"`
		YellowCard   int `json:"yellow_card"`
		CornerKick   int `json:"corner_kick"`
		OTScore      int `json:"ot_score"`
		PenaltyScore int `json:"penalty_score"`
	}

	StatsLiveMessage struct {
		Id   int64 `json:"match_id"`
		Type int   `json:"type"`
		Home int   `json:"home"`
		Away int   `json:"away"`
	}

	RoomTextLiveMessage struct {
		SentMessageStruct
		Id         int64  `json:"match_id"`    // 赛事id
		Time       string `json:"time"`        // 事件时间
		Type       int8   `json:"type"`        // 事件类型
		Data       string `json:"data"`        // 事件文本
		Position   int8   `json:"position"`    // 事件發生方， 0-中立 1-主队 2-客队
		Main       int8   `json:"main"`        // 是否重要事件 0-否 1-是
		CreateTime int64  `json:"create_time"` // 創建時間
	}

	BasketballText struct {
		MatchId    int64  `json:"match_id"`    // 篮球ID
		Time       string `json:"time"`        // 时间
		EventTeam  int8   `json:"event_team"`  // 事件发生在团队（0-中立，1-主队，2-客队
		AwayScore  int    `json:"away_score"`  // 客队比分
		HomeScore  int    `json:"home_score"`  // 主队比分
		Text       string `json:"text"`        // 文字内容
		CreateTime int64  `json:"create_time"` // 創建時間
	}

	SentMessageStruct struct {
		Type    string `json:"type"`
		Message string `default:"" json:"message,omitempty"`
	}

	Lineup struct {
		MatchId       int          `json:"match_id"`       // 比賽ID
		Confirmed     int          `json:"confirmed"`      // 正式阵容，1-是、0-不是
		HomeFormation string       `json:"home_formation"` // 主队阵型 3-4-3
		AwayFormation string       `json:"away_formation"` // 客队阵型 4-4-1-1
		Home          []LineupItem `json:"home"`           // 主隊球員信息
		Away          []LineupItem `json:"away"`           // 客隊球員信息
	}
	LineupItem struct {
		LineupId     int         `json:"id"`                  // 球员id
		TeamId       int         `json:"team_id"`             // 球队id
		First        int         `json:"first"`               // 是否首发，1-是、0-否
		Captain      int         `json:"captain"`             // 是否队长，1-是、0-否
		Name         string      `json:"name"`                // 球员名称
		Logo         string      `json:"logo"`                // 球员logo
		NationalLogo string      `json:"national_logo"`       // 球员logo(国家队)
		ShirtNumber  int         `json:"shirt_number"`        // 球衣号
		Position     string      `json:"position"`            // 球员位置，F前锋、M中场、D后卫、G守门员、其他为未知
		X            int         `json:"x"`                   // 阵容x坐标，总共100
		Y            int         `json:"y"`                   // 阵容y坐标，总共100
		Rating       string      `json:"rating"`              // 评分，10为满分
		Incidents    []Incidents `json:"Incidents,omitempty"` // 球员事件列表，有事件存在，默认不存在
	}
	Incidents struct {
		Type       int    `json:"type"`        // 事件类型，详见状态码->技术统计
		Time       string `json:"time"`        // 事件发生时间（含加时时间，'A+B':A-比赛时间,B-加时时间）
		Belong     int    `json:"belong"`      // 发生方，0-中立、1-主队、2-客队
		HomeScore  int    `json:"home_score"`  // 主队比分
		AwayScore  int    `json:"away_score"`  // 客队比分
		ReasonType int    `json:"reason_type"` // 红黄牌、换人事件原因，详见状态码->事件原因（红黄牌、换人事件存在）
		Player     Player `json:"player"`      // player-相关球员
		Assist1    Player `json:"assist1"`     // assist1-助攻球员1
		Assist2    Player `json:"assist2"`     // assist1-助攻球员2
		InPlayer   Player `json:"in_player"`   // in_player-换上球员
		OutPlayer  Player `json:"out_player"`  // out_player-换下球员
	}
	Player struct {
		PlayerId int    `json:"id"`   // 球员id
		Name     string `json:"name"` // 中文名称
	}
)
