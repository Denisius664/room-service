package chatservice

import (
    "context"
    "errors"
    "testing"

    "github.com/Denisius664/room-service/internal/models"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
)

// mockProducer is a testify mock for chatCommandProducer
type mockProducer struct{ mock.Mock }

func (m *mockProducer) Produce(ctx context.Context, cmd *models.SendMessageCommand) error {
    args := m.Called(ctx, cmd)
    return args.Error(0)
}

type ChatServiceTestSuite struct {
    suite.Suite
    prod *mockProducer
    svc  *ChatService
}

func (s *ChatServiceTestSuite) SetupTest() {
    s.prod = &mockProducer{}
    s.svc = NewChatService(s.prod)
}

func (s *ChatServiceTestSuite) TestSend_Success() {
    ctx := context.Background()
    room := "room1"
    sender := "alice"
    content := "hello"

    s.prod.On("Produce", mock.Anything, mock.MatchedBy(func(cmd *models.SendMessageCommand) bool {
        return cmd != nil && cmd.ToRoom == room && cmd.Sender == sender && cmd.Content == content
    })).Return(nil)

    err := s.svc.Send(ctx, room, sender, content)
    s.Require().NoError(err)

    s.prod.AssertExpectations(s.T())
}

func (s *ChatServiceTestSuite) TestSend_ProducerError() {
    ctx := context.Background()
    room := "room1"
    sender := "bob"
    content := "bye"

    s.prod.On("Produce", mock.Anything, mock.Anything).Return(errors.New("kafka error"))

    err := s.svc.Send(ctx, room, sender, content)
    s.Require().Error(err)

    s.prod.AssertExpectations(s.T())
}

func TestChatServiceSuite(t *testing.T) {
    suite.Run(t, new(ChatServiceTestSuite))
}
