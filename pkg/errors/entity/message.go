package entity

import "net/http"

var (
	ErrMsgISE = Message{
		StatusCode: http.StatusInternalServerError,
		EN:         `An unexpected error occurred. Please try again later or contact support if the problem persists.`,
		ID:         `Terjadi kesalahan yang tidak terduga. Silakan coba lagi nanti atau hubungi dukungan jika masalah berlanjut.`,
	}
	ErrMsgNotFound = Message{
		StatusCode: http.StatusNotFound,
		EN:         `The requested resource was not found. Please verify your input and try again.`,
		ID:         `Data yang diminta tidak ditemukan. Silakan periksa kembali input Anda dan coba lagi.`,
	}
	ErrMsgBadRequest = Message{
		StatusCode: http.StatusBadRequest,
		EN:         `Invalid request. Please check your input and try again.`,
		ID:         `Permintaan tidak valid. Silakan periksa input Anda dan coba lagi.`,
	}
	ErrMsgBadRequestCustom = Message{
		StatusCode:    http.StatusBadRequest,
		EN:            `%s`,
		ID:            `%s`,
		CustomMessage: true,
	}
	ErrMsgUnauthorized = Message{
		StatusCode: http.StatusUnauthorized,
		EN:         `Authentication required. Please log in to access this resource.`,
		ID:         `Autentikasi diperlukan. Silakan masuk untuk mengakses sumber daya ini.`,
	}
	ErrMsgUniqueConst = Message{
		StatusCode: http.StatusConflict,
		EN:         `A record with this information already exists. Please use different data or contact support.`,
		ID:         `Data dengan informasi ini sudah ada. Silakan gunakan data yang berbeda atau hubungi dukungan.`,
	}
	ErrMsgTooManyRequest = Message{
		StatusCode: http.StatusTooManyRequests,
		EN:         `Too many requests. Please wait a moment before trying again.`,
		ID:         `Terlalu banyak permintaan. Silakan tunggu sebentar sebelum mencoba lagi.`,
	}
	ErrMsgUnprocessable = Message{
		StatusCode: http.StatusUnprocessableEntity,
		EN:         `Unable to process your request. Please verify your input and try again.`,
		ID:         `Tidak dapat memproses permintaan Anda. Silakan verifikasi input Anda dan coba lagi.`,
	}
	ErrMsgForbidden = Message{
		StatusCode: http.StatusForbidden,
		EN:         `Access denied. You don't have permission to perform this action.`,
		ID:         `Akses ditolak. Anda tidak memiliki izin untuk melakukan tindakan ini.`,
	}
	ErrMsgContextCancelled = Message{
		StatusCode: 499,
		EN:         `Request was cancelled by the client.`,
		ID:         `Permintaan dibatalkan oleh klien.`,
	}
	ErrMsgContextTimeout = Message{
		StatusCode: http.StatusRequestTimeout,
		EN:         `Request timed out. Please try again.`,
		ID:         `Permintaan habis waktu. Silakan coba lagi.`,
	}
	ErrMsgConflict = Message{
		StatusCode: http.StatusConflict,
		EN:         `A record with this information already exists. Please use different data or contact support.`,
		ID:         `Data dengan informasi ini sudah ada. Silakan gunakan data yang berbeda atau hubungi dukungan.`,
	}
	ErrMsgServiceUnavailable = Message{
		StatusCode: http.StatusServiceUnavailable,
		EN:         `Service is temporarily unavailable. Please try again later.`,
		ID:         `Layanan sedang tidak tersedia sementara. Silakan coba lagi nanti.`,
	}
	ErrMsgVersionConstraint = Message{
		StatusCode: http.StatusUnprocessableEntity,
		EN:         `Your app version is outdated. Please update to the latest version to continue.`,
		ID:         `Versi aplikasi Anda sudah usang. Silakan perbarui ke versi terbaru untuk melanjutkan.`,
	}
)
