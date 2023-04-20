package shortener

import (
	"github.com/sirupsen/logrus"
	"github.com/timickb/link-shortener/internal/repository/memory"
	"testing"
)

func TestCreateLinkSuccess(t *testing.T) {
	mockRepo := memory.New()
	service := New(logrus.New(), mockRepo)

	cases := []string{
		"https://github.com",
		"https://github.com/sirupsen/logrus",
		"http://example.org/",
		"http://example.org/test?key=val",
		"https://example.org/test?key=val&key1=val1",
		"https://sub.domain.org/",
	}

	for _, value := range cases {
		if _, err := service.CreateLink(value); err != nil {
			t.Fatal(err)
		}
	}
}

func TestCreateLinkTwice(t *testing.T) {
	mockRepo := memory.New()
	service := New(logrus.New(), mockRepo)
	url := "https://github.com"

	short, err := service.CreateLink(url)
	if err != nil {
		t.Fatal(err)
	}

	short1, err := service.CreateLink(url)
	if err != nil {
		t.Fatal(err)
	}

	if short1 != short {
		t.Fatal("shortenings don't match")
	}
}

func TestCreateLinkFail(t *testing.T) {
	mockRepo := memory.New()
	service := New(logrus.New(), mockRepo)

	cases := []string{
		"htt://github.com",
		"github.com/sirupsen/logrus",
		"ftp://example.org/",
		"http:/example.org/test?key=val",
		"https//example.org/test?key=val&key1=val1",
		"https:sub.domain.org/",
		"word",
	}

	for _, value := range cases {
		if _, err := service.CreateLink(value); err == nil {
			t.Fatalf("expected err on value %s", value)
		}
	}
}

func TestRestoreLinkSuccess(t *testing.T) {
	mockRepo := memory.New()
	service := New(logrus.New(), mockRepo)

	url := "https://github.com"
	short, _ := service.CreateLink(url)

	original, err := service.RestoreLink(short)
	if err != nil {
		t.Fatal(err)
	}

	if original != url {
		t.Fatal("expected original = url")
	}
}

func TestRestoreLinkFail(t *testing.T) {
	mockRepo := memory.New()
	service := New(logrus.New(), mockRepo)

	_, err := service.RestoreLink("blabla")
	if err == nil {
		t.Fatal("expected err")
	}
}
