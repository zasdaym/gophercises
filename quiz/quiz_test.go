package main

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestCheckAnswer(t *testing.T) {
	type args struct {
		answer   string
		solution string
		point    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "correct answer",
			args: args{
				answer:   "10",
				solution: "10",
				point:    0,
			},
			want: 1,
		},
		{
			name: "incorrect answer",
			args: args{
				answer:   "10",
				solution: "20",
				point:    0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkAnswer(tt.args.answer, tt.args.solution, tt.args.point); got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestGetAnswer(t *testing.T) {
	ch := make(chan string)
	s := "10"
	go getAnswer(ch, strings.NewReader(s))
	got := <-ch
	if got != s {
		t.Errorf("failed to get answer")
	}
}

func TestGetQuizzes(t *testing.T) {
	got := getQuizzes(strings.NewReader("1+1,2"))
	want := Quizzes{
		Quiz{"1+1", "2"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("failed to parse valid csv")
	}
}

func TestQuizString(t *testing.T) {
	got := quizString(0, "1+1")
	want := "Question #1: 1+1 = "
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAskQuizzes(t *testing.T) {
	quizzes := Quizzes{
		Quiz{"1+1", "2"},
	}
	timeLimit := 30 * time.Second
	var w strings.Builder
	r := strings.NewReader("2")

	got := askQuizzes(quizzes, timeLimit, &w, r)
	want := 1
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestShuffleQuizzes(t *testing.T) {
	type args struct {
		quizzes Quizzes
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "should shuffle quizzes",
			args: args{
				quizzes: Quizzes{
					Quiz{"1+1", "2"},
					Quiz{"1+3", "4"},
					Quiz{"1+4", "5"},
					Quiz{"1+5", "6"},
					Quiz{"1+6", "7"},
					Quiz{"1+7", "8"},
					Quiz{"1+8", "9"},
					Quiz{"1+9", "10"},
					Quiz{"2+1", "3"},
					Quiz{"2+2", "4"},
					Quiz{"2+3", "5"},
					Quiz{"2+4", "6"},
					Quiz{"2+5", "7"},
					Quiz{"2+6", "8"},
					Quiz{"2+7", "9"},
					Quiz{"2+7", "9"},
					Quiz{"2+8", "10"},
					Quiz{"2+9", "11"},
					Quiz{"3+1", "4"},
					Quiz{"3+2", "5"},
					Quiz{"3+3", "6"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quizzesCopy := append(tt.args.quizzes[:0:0], tt.args.quizzes...)
			shuffleQuizzes(quizzesCopy)
			if reflect.DeepEqual(quizzesCopy, tt.args.quizzes) {
				t.Errorf("quizzes not shuffled, got %v", quizzesCopy)
			}
		})
	}
}
