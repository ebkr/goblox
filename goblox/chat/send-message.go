package chat

import "time"

func (ref *Chat) SendMessage(subject, body string, userId int) (string, error) {
	data := map[string]interface{}{
		"subject":     subject,
		"body":        body,
		"recipientid": userId,
		"cacheBuster": time.Unix(time.Now().Unix(), 0).String(),
	}
	return ref.SendRequest("http://www.roblox.com/messages/send", data)
}
