package domain

import "errors"

var (
	// General Errors
	ErrInvalidInput = errors.New("invalid input")

	// Household Errors
	ErrHouseholdNotFound = errors.New("data household tidak ditemukan")

	// Waste Pickup Errors & Business Rules
	ErrPickupNotFound         = errors.New("data pickup tidak ditemukan")
	ErrPendingPaymentExists   = errors.New("payment terakhir masih pending, tidak bisa membuat request pickup baru")
	ErrPickupNotPending       = errors.New("pickup hanya bisa dijadwalkan ketika status nya pending")
	ErrSafetyCheckRequired    = errors.New("safety check harus true sebelum request pickup jenis elektronik")
	ErrPickupAlreadyCompleted = errors.New("pickup sudah selesai, tidak bisa dijadwalkan")
	ErrPickupAlreadyCanceled  = errors.New("pickup sudah dibatalkan, tidak bisa dijadwalkan")
	ErrCannotCancelCompleted  = errors.New("tidak bisa membatalkan pickup yang sudah selesai")

	// Payment Errors & Business Rules
	ErrPaymentNotFound    = errors.New("data payment tidak ditemukan")
	ErrPaymentAlreadyPaid = errors.New("payment sudah lunas, tidak bisa dibayar ulang")
	ErrProofFileRequired  = errors.New("bukti pembayaran wajib diisi")
	ErrInvalidFileType    = errors.New("invalid file type: only jpg, jpeg, png, and pdf are supported")
	ErrFileSizeExceeded   = errors.New("file size exceeds maximum allowed limit")
)
