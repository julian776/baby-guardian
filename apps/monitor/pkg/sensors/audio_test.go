package sensors

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAudio_Start(t *testing.T) {
	type fields struct {
		name     string
		interval time.Duration
	}
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test receiving signal from sensor",
			fields: fields{
				name:     "audio",
				interval: time.Millisecond * 100,
			},
			args: args{
				in0: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAudio(tt.fields.name, tt.fields.interval)

			got, err := a.Start(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("Audio.Start() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			signal := <-got
			assert.Equal(t, AudioTyp.String(), signal.Type)
			assert.NotZerof(t, signal.Timestamp, "Timestamp should not be zero")
			assert.NotZerof(t, signal.Value, "Value should not be zero")
		})
	}
}
