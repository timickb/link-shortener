package shortener

import (
	"errors"
	"fmt"
	"github.com/timickb/link-shortener/internal/interfaces"
	"net/url"
	"strings"
)

type Repository interface {
	CreateShortening(shortened, original string) error
	GetOriginal(shortened string) (string, error)
}

type Service struct {
	repo Repository
	log  interfaces.Logger
}

func New(log interfaces.Logger, repo Repository) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}

func (s *Service) CreateLink(link string) (string, error) {
	if !s.validateLink(link) {
		return "", errors.New("err create link: invalid http url format")
	}

	short, err := generateShortening(link)
	if err != nil {
		return "", fmt.Errorf("err create link: %w", err)
	}

	_, err = s.repo.GetOriginal(short)
	if err == nil {
		return short, nil
	}

	if err := s.repo.CreateShortening(short, link); err != nil {
		return "", fmt.Errorf("err create link: %w", err)
	}

	return short, nil
}

func (s *Service) RestoreLink(shortened string) (string, error) {
	original, err := s.repo.GetOriginal(shortened)
	if err != nil {
		return "", fmt.Errorf("err restore link: %w", err)
	}

	return original, nil
}

func (s *Service) validateLink(link string) bool {
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		return false
	}

	parsedUrl, err := url.Parse(link)
	if err != nil {
		return false
	}

	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return false
	}

	return true
}
