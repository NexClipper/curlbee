package config

import (
	"testing"
)

func TestBeePolicy_VariableMatching(t *testing.T) {
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		b       *BeePolicy
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "matching test",

			b: &BeePolicy{
				Name: "{{param.NAME}}",
				Request: BeeRequest{
					Method: "GET",
					URL:    "https://abc.com/{{param.CATEGORY}}/{{param.ZONE_ID}}",
					Headers: []BeeHeader{
						{
							Key:   "X_AUTH_KEY",
							Value: "{{param.AUTH}}",
						},
					},
				},
			},
			args: args{
				params: map[string]string{
					"NAME":     "blahblah-name",
					"ZONE_ID":  "998866",
					"AUTH":     "authenticate-key-987654321",
					"CATEGORY": "GROUP",
				},
			},
			wantErr: false,
		},
		{
			name: "matching test(empty param)",

			b: &BeePolicy{
				Name: "{{param.NAME}}",
				Request: BeeRequest{
					Method: "GET",
					URL:    "https://abc.com/{{param.CATEGORY}}/{{param.ZONE_ID}}",
					Headers: []BeeHeader{
						{
							Key:   "X_AUTH_KEY",
							Value: "{{param.AUTH}}",
						},
					},
				},
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.VariableMatching(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("BeePolicy.VariableMatching() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
