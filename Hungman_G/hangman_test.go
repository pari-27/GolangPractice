package main

// Basic imports
import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockGame struct {
	mock.Mock
}

func (m *MockGame) RenderGame(placeholder []string, entries map[string]bool, chances int) error {
	args := m.Called(placeholder, entries)
	return args.Error(0)
}

func (m *MockGame) GetInput() (str string) {
	args := m.Called()
	if args.String(0) == "word" {
		return args.String(1)
	}

	// "char"
	charset := args.String(1)
	idx := rand.Intn(len(charset))
	return string(charset[idx])
}

type HangmanTestSuite struct {
	word     string
	mockGame *MockGame
	suite.Suite
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *HangmanTestSuite) SetupTest() {
	suite.mockGame = &MockGame{}
	suite.word = "elephant"
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestHangmanTestSuite(t *testing.T) {
	suite.Run(t, new(HangmanTestSuite))
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *HangmanTestSuite) TestPlaySuccess() {
	suite.mockGame.On("RenderGame", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	suite.mockGame.On("GetInput").Return("word", "elephant")

	result := play(suite.mockGame, suite.word)
	assert.Equal(suite.T(), result, true)
}

func (suite *HangmanTestSuite) TestPlayWhenUserOutOfChances() {
	suite.mockGame.On("RenderGame", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	//suite.mockGame.On("GetInput").Return(randomString, "zxbjkimqdp", 1)
	suite.mockGame.On("GetInput").Return("char", "zxbjkimqdp")

	result := play(suite.mockGame, suite.word)
	assert.Equal(suite.T(), result, false)
}

func (suite *HangmanTestSuite) TestPlayWhenTimedOut() {
}

func (suite *HangmanTestSuite) TestPlayUserEntersCompleteWord() {
}

func (suite *HangmanTestSuite) TestPlayUserEntersWrongWord() {
	suite.mockGame.On("RenderGame", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	suite.mockGame.On("GetInput").Return("word", "wrong")

	result := play(suite.mockGame, suite.word)
	assert.Equal(suite.T(), result, false)
}
