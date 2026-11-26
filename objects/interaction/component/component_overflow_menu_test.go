package component

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func strPtr(s string) *string { return &s }

func TestOverflowMenuMarshal(t *testing.T) {
	tests := []struct {
		name     string
		input    Component
		expected string
	}{
		{
			name: "single group without label",
			input: BuildOverflowMenu(OverflowMenu{
				Groups: []OverflowMenuGroup{{
					Components: []MenuActionItem{{
						CustomId: "close",
						Label:    "Close",
						Style:    MenuActionItemStyleDanger,
					}},
				}},
			}),
			expected: `{"type":25,"groups":[{"components":[{"type":26,"custom_id":"close","label":"Close","style":2}]}]}`,
		},
		{
			name: "link item with id, label and description",
			input: BuildOverflowMenu(OverflowMenu{
				Id: new(int),
				Groups: []OverflowMenuGroup{{
					Label: strPtr("Links"),
					Components: []MenuActionItem{{
						Label:       "Docs",
						Style:       MenuActionItemStyleLink,
						Url:         strPtr("https://docs.tickets.bot"),
						Description: strPtr("Open the docs"),
						Disabled:    true,
					}},
				}},
			}),
			expected: `{"type":25,"id":0,"groups":[{"label":"Links","components":[{"type":26,"label":"Docs","style":3,"description":"Open the docs","url":"https://docs.tickets.bot","disabled":true}]}]}`,
		},
		{
			name: "empty groups",
			input: BuildOverflowMenu(OverflowMenu{
				Groups: []OverflowMenuGroup{},
			}),
			expected: `{"type":25,"groups":[]}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, err := json.Marshal(tc.input)
			require.NoError(t, err)
			require.JSONEq(t, tc.expected, string(out))
		})
	}
}

func TestOverflowMenuUnmarshal(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(t *testing.T, c Component)
	}{
		{
			name:  "inside action row",
			input: `{"type":1,"components":[{"type":25,"groups":[{"label":"Ticket","components":[{"type":26,"custom_id":"claim","label":"Claim","style":1}]}]}]}`,
			check: func(t *testing.T, c Component) {
				row, ok := c.ComponentData.(ActionRow)
				require.True(t, ok)
				require.Len(t, row.Components, 1)
				menu, ok := row.Components[0].ComponentData.(OverflowMenu)
				require.True(t, ok)
				require.Len(t, menu.Groups, 1)
				require.Equal(t, "Ticket", *menu.Groups[0].Label)
				require.Equal(t, "claim", menu.Groups[0].Components[0].CustomId)
				require.Equal(t, MenuActionItemStyleDefault, menu.Groups[0].Components[0].Style)
			},
		},
		{
			name:  "as section accessory",
			input: `{"type":9,"components":[{"type":10,"content":"hi"}],"accessory":{"type":25,"groups":[{"components":[]}]}}`,
			check: func(t *testing.T, c Component) {
				section, ok := c.ComponentData.(Section)
				require.True(t, ok)
				require.Equal(t, ComponentOverflowMenu, section.Accessory.Type)
				_, ok = section.Accessory.ComponentData.(OverflowMenu)
				require.True(t, ok)
			},
		},
		{
			name:  "standalone menu action item",
			input: `{"type":26,"id":4,"label":"Delete","style":2,"emoji":{"name":"x"}}`,
			check: func(t *testing.T, c Component) {
				item, ok := c.ComponentData.(MenuActionItem)
				require.True(t, ok)
				require.Equal(t, 4, *item.Id)
				require.Equal(t, MenuActionItemStyleDanger, item.Style)
				require.NotNil(t, item.Emoji)
				require.False(t, item.Disabled)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var c Component
			require.NoError(t, json.Unmarshal([]byte(tc.input), &c))
			tc.check(t, c)
		})
	}
}

func TestComponentTypeValues(t *testing.T) {
	require.Equal(t, ComponentType(23), ComponentCheckbox)
	require.Equal(t, ComponentType(25), ComponentOverflowMenu)
	require.Equal(t, ComponentType(26), ComponentMenuActionItem)
}
