package twitter

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	twitterstream "github.com/fallenstedt/twitter-stream"
)

type StreamData struct {
	Data struct {
		Text      string    `json:"text"`
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		AuthorID  string    `json:"author_id"`
	} `json:"data"`
	Includes struct {
		Users []struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"users"`
	} `json:"includes"`
	MatchingRules []struct {
		ID  string `json:"id"`
		Tag string `json:"tag"`
	} `json:"matching_rules"`
}

type Twitter struct {
	api         *twitterstream.TwitterApi
	commentChan chan<- string
}

func NewTwitter(comment chan<- string, clientID string, clientSecret string, query string) (*Twitter, error) {
	tok, err := twitterstream.
		NewTokenGenerator().
		SetApiKeyAndSecret(clientID, clientSecret).
		RequestBearerToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get bearer token: %w", err)
	}

	api := twitterstream.NewTwitterStream(tok.AccessToken)

	res, err := api.Rules.Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get rules: %w", err)
	}

	ruleIDs := make([]int, 0, len(res.Data))
	for _, rule := range res.Data {
		log.Printf("rule:%s\n", rule)
		id, err := strconv.Atoi(rule.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to convert rule id to int: %w", err)
		}

		ruleIDs = append(ruleIDs, id)
	}

	_, err = api.Rules.Delete(twitterstream.NewRuleDelete(ruleIDs...), false)
	if err != nil {
		return nil, fmt.Errorf("failed to delete rules: %w", err)
	}

	rules := twitterstream.NewRuleBuilder().
		AddRule(query, "query").
		Build()
	_, err = api.Rules.Create(rules, false)
	if err != nil {
		return nil, fmt.Errorf("failed to create rules: %w", err)
	}
	api.Stream.SetUnmarshalHook(func(bytes []byte) (interface{}, error) {
		data := StreamData{}

		if err := json.Unmarshal(bytes, &data); err != nil {
			fmt.Printf("failed to unmarshal bytes: %v", err)
		}
		return data, err
	})

	err = api.Stream.StartStream(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start stream: %w", err)
	}

	return &Twitter{
		api:         api,
		commentChan: comment,
	}, nil
}

func (t *Twitter) Start() error {
	defer t.api.Stream.StopStream()

	for tweet := range t.api.Stream.GetMessages() {
		log.Printf("message: %+v\n", tweet)
		if tweet.Err != nil {
			continue
		}

		result, ok := tweet.Data.(StreamData)
		if !ok {
			log.Printf("failed to cast tweet.Data to StreamData\n")
			continue
		}

		t.commentChan <- result.Data.Text
	}

	return nil
}
