// Package tictactoe solves exercise 3.10 3-8.
package tictactoe

import (
	"testing"

	"github.com/teleivo/assertive/require"
)

func TestTicTacToe(t *testing.T) {
	t.Run("WinRow", func(t *testing.T) {
		g := New(3)

		tests := []struct {
			row    uint
			col    uint
			player bool
			want   bool
		}{
			{
				row:    0,
				col:    0,
				player: false,
				want:   false,
			},
			{
				row:    0,
				col:    1,
				player: false,
				want:   false,
			},
			{
				row:    0,
				col:    2,
				player: false,
				want:   true,
			},
		}

		for _, tc := range tests {
			got := g.Move(tc.row, tc.col, tc.player)

			require.Equals(t, got, tc.want)
		}
	})

	t.Run("DenyRow", func(t *testing.T) {
		g := New(3)

		tests := []struct {
			row    uint
			col    uint
			player bool
			want   bool
		}{
			{
				row:    0,
				col:    0,
				player: false,
				want:   false,
			},
			{
				row:    0,
				col:    1,
				player: true,
				want:   false,
			},
			{
				row:    0,
				col:    2,
				player: false,
				want:   false,
			},
		}

		for _, tc := range tests {
			got := g.Move(tc.row, tc.col, tc.player)

			require.Equals(t, got, tc.want)
		}
	})

	t.Run("WinCol", func(t *testing.T) {
		g := New(3)

		tests := []struct {
			row    uint
			col    uint
			player bool
			want   bool
		}{
			{
				row:    0,
				col:    2,
				player: false,
				want:   false,
			},
			{
				row:    1,
				col:    2,
				player: false,
				want:   false,
			},
			{
				row:    2,
				col:    2,
				player: false,
				want:   true,
			},
		}

		for _, tc := range tests {
			got := g.Move(tc.row, tc.col, tc.player)

			require.Equals(t, got, tc.want)
		}
	})

	t.Run("DenyCol", func(t *testing.T) {
		g := New(3)

		tests := []struct {
			row    uint
			col    uint
			player bool
			want   bool
		}{
			{
				row:    0,
				col:    1,
				player: false,
				want:   false,
			},
			{
				row:    1,
				col:    1,
				player: true,
				want:   false,
			},
			{
				row:    2,
				col:    1,
				player: false,
				want:   false,
			},
		}

		for _, tc := range tests {
			got := g.Move(tc.row, tc.col, tc.player)

			require.Equals(t, got, tc.want)
		}
	})

	t.Run("WinDiagonalTopLeftToBottomRight", func(t *testing.T) {
		g := New(3)

		tests := []struct {
			row    uint
			col    uint
			player bool
			want   bool
		}{
			{
				row:    0,
				col:    0,
				player: false,
				want:   false,
			},
			{
				row:    1,
				col:    1,
				player: false,
				want:   false,
			},
			{
				row:    2,
				col:    2,
				player: false,
				want:   true,
			},
		}

		for _, tc := range tests {
			got := g.Move(tc.row, tc.col, tc.player)

			require.Equals(t, got, tc.want)
		}
	})

	t.Run("DenyDiagonalTopLeftToBottomRight", func(t *testing.T) {
		g := New(3)

		tests := []struct {
			row    uint
			col    uint
			player bool
			want   bool
		}{
			{
				row:    0,
				col:    0,
				player: false,
				want:   false,
			},
			{
				row:    1,
				col:    1,
				player: true,
				want:   false,
			},
			{
				row:    2,
				col:    2,
				player: false,
				want:   false,
			},
		}

		for _, tc := range tests {
			got := g.Move(tc.row, tc.col, tc.player)

			require.Equals(t, got, tc.want)
		}
	})

	t.Run("WinDiagonalBottomLeftToTopRight", func(t *testing.T) {
		g := New(3)

		tests := []struct {
			row    uint
			col    uint
			player bool
			want   bool
		}{
			{
				row:    2,
				col:    0,
				player: false,
				want:   false,
			},
			{
				row:    1,
				col:    1,
				player: false,
				want:   false,
			},
			{
				row:    0,
				col:    2,
				player: false,
				want:   true,
			},
		}

		for _, tc := range tests {
			got := g.Move(tc.row, tc.col, tc.player)

			require.Equals(t, got, tc.want)
		}
	})

	t.Run("DenyDiagonalBottomLeftToTopRight", func(t *testing.T) {
		g := New(3)

		tests := []struct {
			row    uint
			col    uint
			player bool
			want   bool
		}{
			{
				row:    2,
				col:    0,
				player: false,
				want:   false,
			},
			{
				row:    1,
				col:    1,
				player: true,
				want:   false,
			},
			{
				row:    0,
				col:    2,
				player: false,
				want:   false,
			},
		}

		for _, tc := range tests {
			got := g.Move(tc.row, tc.col, tc.player)

			require.Equals(t, got, tc.want)
		}
	})

	t.Run("SecondPlayerWins", func(t *testing.T) {
		g := New(4)

		tests := []struct {
			row    uint
			col    uint
			player bool
			want   bool
		}{
			{
				row:    0,
				col:    1,
				player: false,
				want:   false,
			},
			{
				row:    0,
				col:    0,
				player: true,
				want:   false,
			},
			{
				row:    2,
				col:    2,
				player: false,
				want:   false,
			},
			{
				row:    1,
				col:    1,
				player: true,
				want:   false,
			},
			{
				row:    2,
				col:    0,
				player: false,
				want:   false,
			},
			{
				row:    3,
				col:    0,
				player: true,
				want:   false,
			},
			{
				row:    3,
				col:    1,
				player: false,
				want:   false,
			},
			{
				row:    2,
				col:    1,
				player: true,
				want:   false,
			},
			{
				row:    3,
				col:    2,
				player: false,
				want:   false,
			},
			{
				row:    1,
				col:    2,
				player: true,
				want:   false,
			},
			{
				row:    3,
				col:    3,
				player: false,
				want:   false,
			},
			{
				row:    0,
				col:    3,
				player: true,
				want:   true,
			},
		}

		for _, tc := range tests {
			got := g.Move(tc.row, tc.col, tc.player)

			require.Equals(t, got, tc.want)
		}
	})
}
