package chat

// type EventMsgPlugin struct {
// 	Platform string `json:"platform"`
// 	Channel  struct {
// 		Name     string `json:"name"`
// 		ID       string `json:"id"`
// 		UserMeta struct {
// 		} `json:"user_meta"`
// 	} `json:"channel"`
// 	Timestamp string `json:"timestamp"`
// 	Sender    struct {
// 		Name     string `json:"name"`
// 		ID       string `json:"id"`
// 		UserMeta struct {
// 		} `json:"user_meta"`
// 	} `json:"sender"`
// 	Message      string   `json:"message"`
// 	SplitMsg     []string `json:"split_msg"`
// 	PlatformMeta struct {
// 	} `json:"platform_meta"`
// 	Plugin struct {
// 		PluginID   int    `json:"plugin_id"`
// 		PluginType int    `json:"plugin_type"`
// 		Trigger    string `json:"trigger"`
// 		Data       struct {
// 		} `json:"data"`
// 	} `json:"plugin"`
// }

type Channel struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type UserMeta struct {
	isMod   bool     `json:"is_mod"`
	isSub   bool     `json:"isSub"`
	isTurbo bool     `json:"isTurbo"`
	color   string   `json:"color"`
	badges  []string `json:"user_type"`
}
type Sender struct {
	Name     string   `json:"name"`
	ID       string   `json:"id"`
	UserMeta UserMeta `json:"user_meta"`
}

type NanPlugin struct {
	Platform     string   `json:"platform"`
	Timestamp    string   `json:"timestamp"`
	Channel      Channel  `json:"channel"`
	Sender       Sender   `json:"sender"`
	Message      string   `json:"message"`
	SplitMsg     []string `json:"split_msg"`
	PlatformMeta struct {
	} `json:"platform_meta"`
	Plugin interface{} `json:"plugin"`
}
