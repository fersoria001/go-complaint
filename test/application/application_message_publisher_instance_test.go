package application_test

import (
	"go-complaint/application"
	"go-complaint/dto"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationMessagePublisherInstance(t *testing.T) {
	instance := application.ApplicationMessagePublisherInstance()
	//go instance.Start()
	for n := range 5 {
		ch := make(chan application.ApplicationMessage)
		instance.Subscribe(
			&application.Subscriber{
				Id:   strconv.Itoa(n),
				Send: ch,
			},
		)
		go func() {
			for {
				m := <-ch
				assert.Equal(t, "notification", m.DataType())
				t.Log(m)
			}
		}()
	}
	for range 25 {
		min := 0
		max := 5
		random := rand.Intn(max-min+1) + min
		randomId := strconv.Itoa(random)
		instance.Publish(
			application.NewApplicationMessage(
				randomId,
				"notification",
				dto.Notification{
					Id: randomId + "asdxc",
				},
			),
		)
	}
}
