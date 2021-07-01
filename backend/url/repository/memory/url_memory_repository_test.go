package memory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/eyalch/shrtr/backend/domain"
	"github.com/eyalch/shrtr/backend/url/repository/memory"
)

type URLMemoryRepositoryTestSuite struct {
	suite.Suite

	repo domain.URLRepository
}

func (s *URLMemoryRepositoryTestSuite) SetupTest() {
	s.repo = memory.NewURLMemoryRepository()
}

func (s *URLMemoryRepositoryTestSuite) TestCreateAndGet() {
	// Act
	err := s.repo.Create("abc123", "http://example.com")
	s.Require().NoError(err)

	url, err := s.repo.Get("abc123")

	// Assert
	if s.NoError(err) {
		s.Equal("http://example.com", url)
	}
}

func (s *URLMemoryRepositoryTestSuite) TestGet_NotFound() {
	// Act
	_, err := s.repo.Get("abc123")

	// Assert
	s.ErrorIs(err, domain.ErrKeyNotFound)
}

func (s *URLMemoryRepositoryTestSuite) TestCreate_Duplicate() {
	// Arrange
	err := s.repo.Create("abc123", "http://example.com")
	s.Require().NoError(err)

	// Act
	err = s.repo.Create("abc123", "http://another-example.com")

	// Assert
	s.ErrorIs(err, domain.ErrDuplicateKey)
}

func TestURLMemoryRepository(t *testing.T) {
	suite.Run(t, new(URLMemoryRepositoryTestSuite))
}
