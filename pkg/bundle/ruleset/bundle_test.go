// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ruleset

import (
	"testing"

	bundlev1 "github.com/elastic/harp/api/gen/go/harp/bundle/v1"
	"github.com/golang/protobuf/proto"
)

func TestFromBundle(t *testing.T) {
	type args struct {
		b *bundlev1.Bundle
	}
	tests := []struct {
		name    string
		args    args
		want    *bundlev1.RuleSet
		wantErr bool
	}{
		{
			name: "nil",
			args: args{
				b: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "packages are nil",
			args: args{
				b: &bundlev1.Bundle{
					Labels: map[string]string{
						"test": "true",
					},
					Annotations: map[string]string{
						"harp.elastic.co/v1/testing#bundlePurpose": "test",
					},
					Packages: nil,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "secrets are nil",
			args: args{
				b: &bundlev1.Bundle{
					Labels: map[string]string{
						"test": "true",
					},
					Annotations: map[string]string{
						"harp.elastic.co/v1/testing#bundlePurpose": "test",
					},
					Packages: []*bundlev1.Package{
						{
							Labels: map[string]string{
								"external": "true",
							},
							Annotations: map[string]string{
								"infosec.elastic.co/v1/SecretPolicy#rotationMethod": "ci",
								"infosec.elastic.co/v1/SecretPolicy#rotationPeriod": "90d",
								"infosec.elastic.co/v1/SecretPolicy#serviceType":    "authentication",
								"infosec.elastic.co/v1/SecretPolicy#severity":       "high",
								"infra.elastic.co/v1/CI#jobName":                    "rotate-external-api-key",
								"harp.elastic.co/v1/package#encryptionKeyAlias":     "test",
							},
							Name:    "app/production/testAccount/testService/v1.0.0/internalTestComponent/authentication/api_key",
							Secrets: nil,
						},
					},
				},
			},
			want: &bundlev1.RuleSet{
				ApiVersion: "harp.elastic.co/v1",
				Kind:       "RuleSet",
				Meta: &bundlev1.RuleSetMeta{
					Description: "Generated from bundle content",
				},
				Spec: &bundlev1.RuleSetSpec{
					Rules: []*bundlev1.Rule{},
				},
			},
			wantErr: false,
		},
		{
			name: "secret data is nil",
			args: args{
				b: &bundlev1.Bundle{
					Labels: map[string]string{
						"test": "true",
					},
					Annotations: map[string]string{
						"harp.elastic.co/v1/testing#bundlePurpose": "test",
					},
					Packages: []*bundlev1.Package{
						{
							Labels: map[string]string{
								"external": "true",
							},
							Annotations: map[string]string{
								"infosec.elastic.co/v1/SecretPolicy#rotationMethod": "ci",
								"infosec.elastic.co/v1/SecretPolicy#rotationPeriod": "90d",
								"infosec.elastic.co/v1/SecretPolicy#serviceType":    "authentication",
								"infosec.elastic.co/v1/SecretPolicy#severity":       "high",
								"infra.elastic.co/v1/CI#jobName":                    "rotate-external-api-key",
								"harp.elastic.co/v1/package#encryptionKeyAlias":     "test",
							},
							Name: "app/production/testAccount/testService/v1.0.0/internalTestComponent/authentication/api_key",
							Secrets: &bundlev1.SecretChain{
								Data: nil,
							},
						},
					},
				},
			},
			want: &bundlev1.RuleSet{
				ApiVersion: "harp.elastic.co/v1",
				Kind:       "RuleSet",
				Meta: &bundlev1.RuleSetMeta{
					Description: "Generated from bundle content",
				},
				Spec: &bundlev1.RuleSetSpec{
					Rules: []*bundlev1.Rule{},
				},
			},
			wantErr: false,
		},
		{
			name: "package and secrets define with annotations and labels",
			args: args{
				b: &bundlev1.Bundle{
					Labels: map[string]string{
						"test": "true",
					},
					Annotations: map[string]string{
						"harp.elastic.co/v1/testing#bundlePurpose": "test",
					},
					Packages: []*bundlev1.Package{
						{
							Labels: map[string]string{
								"external": "true",
							},
							Annotations: map[string]string{
								"harp.elastic.co/v1/package#encryptionKeyAlias":     "test",
								"infra.elastic.co/v1/CI#jobName":                    "rotate-external-api-key",
								"infosec.elastic.co/v1/SecretPolicy#rotationMethod": "ci",
								"infosec.elastic.co/v1/SecretPolicy#rotationPeriod": "90d",
								"infosec.elastic.co/v1/SecretPolicy#serviceType":    "authentication",
								"infosec.elastic.co/v1/SecretPolicy#severity":       "high",
							},
							Name: "app/production/testAccount/testService/v1.0.0/internalTestComponent/authentication/api_key",
							Secrets: &bundlev1.SecretChain{
								Labels: map[string]string{
									"vendor": "true",
								},
								Data: []*bundlev1.KV{
									{
										Key:   "API_KEY",
										Type:  "string",
										Value: []byte("3YGVuHwUqYVkjk-c6lQgfVQwFHawPG36TgAm72sPZGE="),
									},
								},
							},
						},
					},
				},
			},
			want: &bundlev1.RuleSet{
				ApiVersion: "harp.elastic.co/v1",
				Kind:       "RuleSet",
				Meta: &bundlev1.RuleSetMeta{
					Description: "Generated from bundle content",
					Name:        "D0QMaOw378tey2m_TvEuhBPkHOZQAgG88MUV4g6XiLk616urhu5an_hUf_N-k_-PF0TqslvGPFSpUUgZcxRhpg",
				},
				Spec: &bundlev1.RuleSetSpec{
					Rules: []*bundlev1.Rule{
						&bundlev1.Rule{
							Name: "LINT-D0QMaO-1",
							Path: "app/production/testAccount/testService/v1.0.0/internalTestComponent/authentication/api_key",
							Constraints: []string{
								"p.match_label(\"external\")",
								"p.match_annotation(\"harp.elastic.co/v1/package#encryptionKeyAlias\")",
								"p.match_annotation(\"infra.elastic.co/v1/CI#jobName\")",
								"p.match_annotation(\"infosec.elastic.co/v1/SecretPolicy#rotationMethod\")",
								"p.match_annotation(\"infosec.elastic.co/v1/SecretPolicy#rotationPeriod\")",
								"p.match_annotation(\"infosec.elastic.co/v1/SecretPolicy#serviceType\")",
								"p.match_annotation(\"infosec.elastic.co/v1/SecretPolicy#severity\")",
								"p.has_secret(\"API_KEY\")",
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromBundle(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if !proto.Equal(got, tt.want) {
				t.Errorf("Ruleset not equal = %v, want %v", got, tt.want)
			}
		})
	}
}
